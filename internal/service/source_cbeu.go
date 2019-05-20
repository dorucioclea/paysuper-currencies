package service

import (
    "encoding/xml"
    "errors"
    "fmt"
    "github.com/paysuper/paysuper-currencies/pkg/proto/currencies"
    "github.com/satori/go.uuid"
    "go.uber.org/zap"
    "net/http"
)

const (
    errorCbeuUrlValidationFailed   = "CBEU Rates url validation failed"
    errorCbeuRequestFailed         = "CBEU Rates request failed"
    errorCbeuResponseParsingFailed = "CBEU Rates response parsing failed"
    errorCbeuSaveRatesFailed       = "CBEU Rates save data failed"
    errorCbeuNoResults             = "CBEU Rates no results"

    cbeuTo          = "EUR"
    cbeuSource      = "CBEU"
    cbeuUrlTemplate = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml?%s"
)

type cbeuResponse struct {
    XMLName xml.Name          `xml:"http://www.gesmes.org/xml/2002-08-01 Envelope"`
    Data    cbeuResponseCube1 `xml:"Cube"`
}

type cbeuResponseCube1 struct {
    XMLName xml.Name          `xml:"Cube"`
    Rates   cbeuResponseCube2 `xml:"Cube"`
}

type cbeuResponseCube2 struct {
    XMLName xml.Name            `xml:"Cube"`
    Rates   []cbeuResponseCube3 `xml:"Cube"`
}

type cbeuResponseCube3 struct {
    XMLName      xml.Name `xml:"Cube"`
    CurrencyCode string   `xml:"currency,attr"`
    Value        float64  `xml:"rate,attr"`
}

func (s *Service) RequestRatesCbeu(c chan error) {
    zap.S().Info("Requesting rates from CBEU")

    headers := map[string]string{
        HeaderContentType: MIMEApplicationXML,
        HeaderAccept:      MIMEApplicationXML,
    }

    reqUrl, err := s.validateUrl(fmt.Sprintf(cbeuUrlTemplate, uuid.NewV4().String()))

    if err != nil {
        zap.S().Errorw(errorCbeuUrlValidationFailed, "error", err)
        s.sendCentrifugoMessage(errorCbeuUrlValidationFailed, err)
        c <- err
        return
    }

    zap.S().Info("Sending request to url: ", reqUrl.String())

    resp, err := s.request(http.MethodGet, reqUrl.String(), nil, headers)

    if err != nil {
        zap.S().Errorw(errorCbeuRequestFailed, "error", err)
        s.sendCentrifugoMessage(errorCbeuRequestFailed, err)
        c <- err
        return
    }

    res := &cbeuResponse{}
    err = s.decodeXml(resp, res)

    if err != nil {
        zap.S().Errorw(errorCbeuResponseParsingFailed, "error", err)
        s.sendCentrifugoMessage(errorCbeuResponseParsingFailed, err)
        c <- err
        return
    }

    err = s.processRatesCbeu(res)

    if err != nil {
        zap.S().Errorw(errorCbeuSaveRatesFailed, "error", err)
        s.sendCentrifugoMessage(errorCbeuSaveRatesFailed, err)
        c <- err
        return
    }

    zap.S().Info("Rates from CBEU updated")
}

func (s *Service) processRatesCbeu(res *cbeuResponse) error {

    if len(res.Data.Rates.Rates) == 0 {
        return errors.New(errorCbeuNoResults)
    }

    var rates []interface{}

    ln := len(s.cfg.SettlementCurrencies)
    if s.contains(s.cfg.SettlementCurrenciesParsed, cbeuTo) {
        ln--
    }
    c := 0

    for _, rateItem := range res.Data.Rates.Rates {

        if !s.contains(s.cfg.SettlementCurrenciesParsed, rateItem.CurrencyCode) {
            continue
        }

        if rateItem.CurrencyCode == cbeuTo {
            continue
        }

        // direct pair
        rates = append(rates, &currencies.RateData{
            Pair:   rateItem.CurrencyCode + cbeuTo,
            Rate:   s.toPrecise(rateItem.Value),
            Source: cbeuSource,
        })

        // inverse pair
        rates = append(rates, &currencies.RateData{
            Pair:   cbeuTo + rateItem.CurrencyCode,
            Rate:   s.toPrecise(1 / rateItem.Value),
            Source: cbeuSource,
        })

        c++
        if c == ln {
            break
        }
    }

    err := s.saveRates(collectionRatesNameSuffixCentralbanks, rates)
    if err != nil {
        return err
    }
    return nil
}
