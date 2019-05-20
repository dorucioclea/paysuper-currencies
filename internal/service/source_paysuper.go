package service

import (
    "github.com/globalsign/mgo/bson"
    "github.com/paysuper/paysuper-currencies/pkg/proto/currencies"
    "go.uber.org/zap"
)

const (
    paysuperSource = "PS"

    errorPaysuperRateCalc = "paysuper prediction rate calculation error"
    errorPaysuperRateSave = "paysuper prediction rates save error"
)

func (s *Service) SetRatesPaysuper(c chan error) {

    zap.S().Info("Start calculation of prediction rates for Paysuper")

    var (
        cFrom string
        cTo   string
        rates []interface{}
    )

    for _, cFrom = range s.cfg.SettlementCurrencies {
        for _, cTo = range s.cfg.RatesRequestCurrencies {

            if cFrom == cTo {
                continue
            }

            rd, err := s.getRatePaysuper(cFrom, cTo)
            if err != nil {
                zap.S().Errorw(errorPaysuperRateCalc, "error", err)
                s.sendCentrifugoMessage(errorPaysuperRateCalc, err)
                c <- err
                return
            }
            rates = append(rates, rd)

            rd, err = s.getRatePaysuper(cTo, cFrom)
            if err != nil {
                zap.S().Errorw(errorPaysuperRateCalc, "error", err)
                s.sendCentrifugoMessage(errorPaysuperRateCalc, err)
                c <- err
                return
            }
            rates = append(rates, rd)
        }
    }

    err := s.saveRates(collectionRatesNameSuffixPaysuper, rates)
    if err != nil {
        zap.S().Errorw(errorPaysuperRateSave, "error", err)
        s.sendCentrifugoMessage(errorPaysuperRateSave, err)
        c <- err
        return
    }

    zap.S().Info("Prediction rates for Paysuper updated")
}

func (s *Service) getRatePaysuper(cFrom string, cTo string) (*currencies.RateData, error) {
    res := &currencies.RateData{}

    err := s.getRate(collectionRatesNameSuffixOxr, cFrom, cTo, bson.M{}, res)
    if err != nil {
        return nil, err
    }

    correction, err := s.GetPaysuperCorrection(cFrom + cTo)
    if err != nil {
        return nil, err
    }

    rd := &currencies.RateData{
        Pair:   cFrom + cTo,
        Rate:   s.toPrecise(res.Rate + correction),
        Source: paysuperSource,
    }

    return rd, nil

}
