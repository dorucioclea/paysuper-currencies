package internal

import (
    "context"
    "github.com/globalsign/mgo/bson"
    "github.com/golang/protobuf/ptypes"
    currencyrates "github.com/paysuper/paysuper-currencies-rates/proto"
    "go.uber.org/zap"
    "time"
)

func (s *Service) GetOxrRate(
    ctx context.Context,
    req *currencyrates.GetRateRequest,
    res *currencyrates.RateData,
) error {
    err := s.getRate(collectionSuffixOxr, req.From, req.To, bson.M{}, res)
    if err != nil {
        zap.S().Errorw(errorCurrentRateRequest, "error", err, "req", req)
        return err
    }
    return nil
}

func (s *Service) GetCentralBankRateForDate(
    ctx context.Context,
    req *currencyrates.GetCentralBankRateRequest,
    res *currencyrates.RateData,
) error {

    dt, err := ptypes.Timestamp(req.Datetime)
    if err != nil {
        zap.S().Errorw(errorDatetimeConversion, "error", err, "req", req)
        return err
    }

    query := bson.M{"created_at": bson.M{"$lte": s.Eod(dt)}}
    err = s.getRate(collectionSuffixCb, req.From, req.To, query, res)
    if err != nil {
        zap.S().Errorw(errorCentralBankRateRequest, "error", err, "req", req)
        return err
    }

    return nil
}

func (s *Service) GetPaysuperRate(
    ctx context.Context,
    req *currencyrates.GetRateRequest,
    res *currencyrates.RateData,
) error {
    err := s.getRate(collectionSuffixPaysuper, req.From, req.To, bson.M{}, res)
    if err != nil {
        zap.S().Errorw(errorCurrentRateRequest, "error", err, "req", req)
        return err
    }
    return nil
}

func (s *Service) GetStockRate(
    ctx context.Context,
    req *currencyrates.GetRateRequest,
    res *currencyrates.RateData,
) error {
    err := s.getRate(collectionSuffixStock, req.From, req.To, bson.M{}, res)
    if err != nil {
        zap.S().Errorw(errorCurrentRateRequest, "error", err, "req", req)
        return err
    }
    return nil
}

func (s *Service) GetCardpayRate(
    ctx context.Context,
    req *currencyrates.GetRateRequest,
    res *currencyrates.RateData,
) error {
    err := s.getRate(collectionSuffixCardpay, req.From, req.To, bson.M{}, res)
    if err != nil {
        zap.S().Errorw(errorCurrentRateRequest, "error", err, "req", req)
        return err
    }
    return nil
}

func (s *Service) SetPaysuperCorrectionCorridor(
    ctx context.Context,
    req *currencyrates.CorrectionCorridor,
    res *currencyrates.EmptyResponse,
) error {

    corridor := PaysuperCorridor{
        Value:     req.Value,
        CreatedAt: time.Now(),
    }

    err := s.db.Collection(collectionNamePaysuperCorridors).Insert(corridor)
    if err != nil {
        zap.S().Errorw(errorDbInsertFailed, "error", err, "data", corridor)
        return err
    }

    return nil
}
