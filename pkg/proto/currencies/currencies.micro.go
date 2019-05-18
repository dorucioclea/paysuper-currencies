// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: currencies.proto

/*
Package currencies is a generated protocol buffer package.

It is generated from these files:
	currencies.proto

It has these top-level messages:
	GetRateCurrentCommonRequest
	GetRateByDateCommonRequest
	GetRateCurrentForMerchantRequest
	GetRateByDateForMerchantRequest
	RateData
	EmptyResponse
	EmptyRequest
	CorrectionCorridor
	CorrectionRule
	CommonCorrectionRule
	CommonCorrectionRuleRequest
	MerchantCorrectionRuleRequest
	ExchangeCurrencyCurrentCommonRequest
	ExchangeCurrencyCurrentForMerchantRequest
	ExchangeCurrencyByDateCommonRequest
	ExchangeCurrencyByDateForMerchantRequest
	ExchangeCurrencyResponse
	CurrenciesList
*/
package currencies

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CurrencyratesService service

type CurrencyratesService interface {
	GetRateCurrentCommon(ctx context.Context, in *GetRateCurrentCommonRequest, opts ...client.CallOption) (*RateData, error)
	GetRateByDateCommon(ctx context.Context, in *GetRateByDateCommonRequest, opts ...client.CallOption) (*RateData, error)
	GetRateCurrentForMerchant(ctx context.Context, in *GetRateCurrentForMerchantRequest, opts ...client.CallOption) (*RateData, error)
	GetRateByDateForMerchant(ctx context.Context, in *GetRateByDateForMerchantRequest, opts ...client.CallOption) (*RateData, error)
	ExchangeCurrencyCurrentCommon(ctx context.Context, in *ExchangeCurrencyCurrentCommonRequest, opts ...client.CallOption) (*ExchangeCurrencyResponse, error)
	ExchangeCurrencyCurrentForMerchant(ctx context.Context, in *ExchangeCurrencyCurrentForMerchantRequest, opts ...client.CallOption) (*ExchangeCurrencyResponse, error)
	ExchangeCurrencyByDateCommon(ctx context.Context, in *ExchangeCurrencyByDateCommonRequest, opts ...client.CallOption) (*ExchangeCurrencyResponse, error)
	ExchangeCurrencyByDateForMerchant(ctx context.Context, in *ExchangeCurrencyByDateForMerchantRequest, opts ...client.CallOption) (*ExchangeCurrencyResponse, error)
	GetCommonRateCorrectionRule(ctx context.Context, in *CommonCorrectionRuleRequest, opts ...client.CallOption) (*CorrectionRule, error)
	GetMerchantRateCorrectionRule(ctx context.Context, in *MerchantCorrectionRuleRequest, opts ...client.CallOption) (*CorrectionRule, error)
	AddCommonRateCorrectionRule(ctx context.Context, in *CommonCorrectionRule, opts ...client.CallOption) (*EmptyResponse, error)
	AddMerchantRateCorrectionRule(ctx context.Context, in *CorrectionRule, opts ...client.CallOption) (*EmptyResponse, error)
	SetPaysuperCorrectionCorridor(ctx context.Context, in *CorrectionCorridor, opts ...client.CallOption) (*EmptyResponse, error)
	GetSupportedCurrencies(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CurrenciesList, error)
	GetSettlementCurrencies(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CurrenciesList, error)
	GetPriceCurrencies(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CurrenciesList, error)
	GetVatCurrencies(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CurrenciesList, error)
	GetAccountingCurrencies(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CurrenciesList, error)
}

type currencyratesService struct {
	c    client.Client
	name string
}

func NewCurrencyratesService(name string, c client.Client) CurrencyratesService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "currencies"
	}
	return &currencyratesService{
		c:    c,
		name: name,
	}
}

func (c *currencyratesService) GetRateCurrentCommon(ctx context.Context, in *GetRateCurrentCommonRequest, opts ...client.CallOption) (*RateData, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetRateCurrentCommon", in)
	out := new(RateData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) GetRateByDateCommon(ctx context.Context, in *GetRateByDateCommonRequest, opts ...client.CallOption) (*RateData, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetRateByDateCommon", in)
	out := new(RateData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) GetRateCurrentForMerchant(ctx context.Context, in *GetRateCurrentForMerchantRequest, opts ...client.CallOption) (*RateData, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetRateCurrentForMerchant", in)
	out := new(RateData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) GetRateByDateForMerchant(ctx context.Context, in *GetRateByDateForMerchantRequest, opts ...client.CallOption) (*RateData, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetRateByDateForMerchant", in)
	out := new(RateData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) ExchangeCurrencyCurrentCommon(ctx context.Context, in *ExchangeCurrencyCurrentCommonRequest, opts ...client.CallOption) (*ExchangeCurrencyResponse, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.ExchangeCurrencyCurrentCommon", in)
	out := new(ExchangeCurrencyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) ExchangeCurrencyCurrentForMerchant(ctx context.Context, in *ExchangeCurrencyCurrentForMerchantRequest, opts ...client.CallOption) (*ExchangeCurrencyResponse, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.ExchangeCurrencyCurrentForMerchant", in)
	out := new(ExchangeCurrencyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) ExchangeCurrencyByDateCommon(ctx context.Context, in *ExchangeCurrencyByDateCommonRequest, opts ...client.CallOption) (*ExchangeCurrencyResponse, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.ExchangeCurrencyByDateCommon", in)
	out := new(ExchangeCurrencyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) ExchangeCurrencyByDateForMerchant(ctx context.Context, in *ExchangeCurrencyByDateForMerchantRequest, opts ...client.CallOption) (*ExchangeCurrencyResponse, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.ExchangeCurrencyByDateForMerchant", in)
	out := new(ExchangeCurrencyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) GetCommonRateCorrectionRule(ctx context.Context, in *CommonCorrectionRuleRequest, opts ...client.CallOption) (*CorrectionRule, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetCommonRateCorrectionRule", in)
	out := new(CorrectionRule)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) GetMerchantRateCorrectionRule(ctx context.Context, in *MerchantCorrectionRuleRequest, opts ...client.CallOption) (*CorrectionRule, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetMerchantRateCorrectionRule", in)
	out := new(CorrectionRule)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) AddCommonRateCorrectionRule(ctx context.Context, in *CommonCorrectionRule, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.AddCommonRateCorrectionRule", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) AddMerchantRateCorrectionRule(ctx context.Context, in *CorrectionRule, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.AddMerchantRateCorrectionRule", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) SetPaysuperCorrectionCorridor(ctx context.Context, in *CorrectionCorridor, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.SetPaysuperCorrectionCorridor", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) GetSupportedCurrencies(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CurrenciesList, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetSupportedCurrencies", in)
	out := new(CurrenciesList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) GetSettlementCurrencies(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CurrenciesList, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetSettlementCurrencies", in)
	out := new(CurrenciesList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) GetPriceCurrencies(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CurrenciesList, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetPriceCurrencies", in)
	out := new(CurrenciesList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) GetVatCurrencies(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CurrenciesList, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetVatCurrencies", in)
	out := new(CurrenciesList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyratesService) GetAccountingCurrencies(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*CurrenciesList, error) {
	req := c.c.NewRequest(c.name, "CurrencyratesService.GetAccountingCurrencies", in)
	out := new(CurrenciesList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CurrencyratesService service

type CurrencyratesServiceHandler interface {
	GetRateCurrentCommon(context.Context, *GetRateCurrentCommonRequest, *RateData) error
	GetRateByDateCommon(context.Context, *GetRateByDateCommonRequest, *RateData) error
	GetRateCurrentForMerchant(context.Context, *GetRateCurrentForMerchantRequest, *RateData) error
	GetRateByDateForMerchant(context.Context, *GetRateByDateForMerchantRequest, *RateData) error
	ExchangeCurrencyCurrentCommon(context.Context, *ExchangeCurrencyCurrentCommonRequest, *ExchangeCurrencyResponse) error
	ExchangeCurrencyCurrentForMerchant(context.Context, *ExchangeCurrencyCurrentForMerchantRequest, *ExchangeCurrencyResponse) error
	ExchangeCurrencyByDateCommon(context.Context, *ExchangeCurrencyByDateCommonRequest, *ExchangeCurrencyResponse) error
	ExchangeCurrencyByDateForMerchant(context.Context, *ExchangeCurrencyByDateForMerchantRequest, *ExchangeCurrencyResponse) error
	GetCommonRateCorrectionRule(context.Context, *CommonCorrectionRuleRequest, *CorrectionRule) error
	GetMerchantRateCorrectionRule(context.Context, *MerchantCorrectionRuleRequest, *CorrectionRule) error
	AddCommonRateCorrectionRule(context.Context, *CommonCorrectionRule, *EmptyResponse) error
	AddMerchantRateCorrectionRule(context.Context, *CorrectionRule, *EmptyResponse) error
	SetPaysuperCorrectionCorridor(context.Context, *CorrectionCorridor, *EmptyResponse) error
	GetSupportedCurrencies(context.Context, *EmptyRequest, *CurrenciesList) error
	GetSettlementCurrencies(context.Context, *EmptyRequest, *CurrenciesList) error
	GetPriceCurrencies(context.Context, *EmptyRequest, *CurrenciesList) error
	GetVatCurrencies(context.Context, *EmptyRequest, *CurrenciesList) error
	GetAccountingCurrencies(context.Context, *EmptyRequest, *CurrenciesList) error
}

func RegisterCurrencyratesServiceHandler(s server.Server, hdlr CurrencyratesServiceHandler, opts ...server.HandlerOption) error {
	type currencyratesService interface {
		GetRateCurrentCommon(ctx context.Context, in *GetRateCurrentCommonRequest, out *RateData) error
		GetRateByDateCommon(ctx context.Context, in *GetRateByDateCommonRequest, out *RateData) error
		GetRateCurrentForMerchant(ctx context.Context, in *GetRateCurrentForMerchantRequest, out *RateData) error
		GetRateByDateForMerchant(ctx context.Context, in *GetRateByDateForMerchantRequest, out *RateData) error
		ExchangeCurrencyCurrentCommon(ctx context.Context, in *ExchangeCurrencyCurrentCommonRequest, out *ExchangeCurrencyResponse) error
		ExchangeCurrencyCurrentForMerchant(ctx context.Context, in *ExchangeCurrencyCurrentForMerchantRequest, out *ExchangeCurrencyResponse) error
		ExchangeCurrencyByDateCommon(ctx context.Context, in *ExchangeCurrencyByDateCommonRequest, out *ExchangeCurrencyResponse) error
		ExchangeCurrencyByDateForMerchant(ctx context.Context, in *ExchangeCurrencyByDateForMerchantRequest, out *ExchangeCurrencyResponse) error
		GetCommonRateCorrectionRule(ctx context.Context, in *CommonCorrectionRuleRequest, out *CorrectionRule) error
		GetMerchantRateCorrectionRule(ctx context.Context, in *MerchantCorrectionRuleRequest, out *CorrectionRule) error
		AddCommonRateCorrectionRule(ctx context.Context, in *CommonCorrectionRule, out *EmptyResponse) error
		AddMerchantRateCorrectionRule(ctx context.Context, in *CorrectionRule, out *EmptyResponse) error
		SetPaysuperCorrectionCorridor(ctx context.Context, in *CorrectionCorridor, out *EmptyResponse) error
		GetSupportedCurrencies(ctx context.Context, in *EmptyRequest, out *CurrenciesList) error
		GetSettlementCurrencies(ctx context.Context, in *EmptyRequest, out *CurrenciesList) error
		GetPriceCurrencies(ctx context.Context, in *EmptyRequest, out *CurrenciesList) error
		GetVatCurrencies(ctx context.Context, in *EmptyRequest, out *CurrenciesList) error
		GetAccountingCurrencies(ctx context.Context, in *EmptyRequest, out *CurrenciesList) error
	}
	type CurrencyratesService struct {
		currencyratesService
	}
	h := &currencyratesServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CurrencyratesService{h}, opts...))
}

type currencyratesServiceHandler struct {
	CurrencyratesServiceHandler
}

func (h *currencyratesServiceHandler) GetRateCurrentCommon(ctx context.Context, in *GetRateCurrentCommonRequest, out *RateData) error {
	return h.CurrencyratesServiceHandler.GetRateCurrentCommon(ctx, in, out)
}

func (h *currencyratesServiceHandler) GetRateByDateCommon(ctx context.Context, in *GetRateByDateCommonRequest, out *RateData) error {
	return h.CurrencyratesServiceHandler.GetRateByDateCommon(ctx, in, out)
}

func (h *currencyratesServiceHandler) GetRateCurrentForMerchant(ctx context.Context, in *GetRateCurrentForMerchantRequest, out *RateData) error {
	return h.CurrencyratesServiceHandler.GetRateCurrentForMerchant(ctx, in, out)
}

func (h *currencyratesServiceHandler) GetRateByDateForMerchant(ctx context.Context, in *GetRateByDateForMerchantRequest, out *RateData) error {
	return h.CurrencyratesServiceHandler.GetRateByDateForMerchant(ctx, in, out)
}

func (h *currencyratesServiceHandler) ExchangeCurrencyCurrentCommon(ctx context.Context, in *ExchangeCurrencyCurrentCommonRequest, out *ExchangeCurrencyResponse) error {
	return h.CurrencyratesServiceHandler.ExchangeCurrencyCurrentCommon(ctx, in, out)
}

func (h *currencyratesServiceHandler) ExchangeCurrencyCurrentForMerchant(ctx context.Context, in *ExchangeCurrencyCurrentForMerchantRequest, out *ExchangeCurrencyResponse) error {
	return h.CurrencyratesServiceHandler.ExchangeCurrencyCurrentForMerchant(ctx, in, out)
}

func (h *currencyratesServiceHandler) ExchangeCurrencyByDateCommon(ctx context.Context, in *ExchangeCurrencyByDateCommonRequest, out *ExchangeCurrencyResponse) error {
	return h.CurrencyratesServiceHandler.ExchangeCurrencyByDateCommon(ctx, in, out)
}

func (h *currencyratesServiceHandler) ExchangeCurrencyByDateForMerchant(ctx context.Context, in *ExchangeCurrencyByDateForMerchantRequest, out *ExchangeCurrencyResponse) error {
	return h.CurrencyratesServiceHandler.ExchangeCurrencyByDateForMerchant(ctx, in, out)
}

func (h *currencyratesServiceHandler) GetCommonRateCorrectionRule(ctx context.Context, in *CommonCorrectionRuleRequest, out *CorrectionRule) error {
	return h.CurrencyratesServiceHandler.GetCommonRateCorrectionRule(ctx, in, out)
}

func (h *currencyratesServiceHandler) GetMerchantRateCorrectionRule(ctx context.Context, in *MerchantCorrectionRuleRequest, out *CorrectionRule) error {
	return h.CurrencyratesServiceHandler.GetMerchantRateCorrectionRule(ctx, in, out)
}

func (h *currencyratesServiceHandler) AddCommonRateCorrectionRule(ctx context.Context, in *CommonCorrectionRule, out *EmptyResponse) error {
	return h.CurrencyratesServiceHandler.AddCommonRateCorrectionRule(ctx, in, out)
}

func (h *currencyratesServiceHandler) AddMerchantRateCorrectionRule(ctx context.Context, in *CorrectionRule, out *EmptyResponse) error {
	return h.CurrencyratesServiceHandler.AddMerchantRateCorrectionRule(ctx, in, out)
}

func (h *currencyratesServiceHandler) SetPaysuperCorrectionCorridor(ctx context.Context, in *CorrectionCorridor, out *EmptyResponse) error {
	return h.CurrencyratesServiceHandler.SetPaysuperCorrectionCorridor(ctx, in, out)
}

func (h *currencyratesServiceHandler) GetSupportedCurrencies(ctx context.Context, in *EmptyRequest, out *CurrenciesList) error {
	return h.CurrencyratesServiceHandler.GetSupportedCurrencies(ctx, in, out)
}

func (h *currencyratesServiceHandler) GetSettlementCurrencies(ctx context.Context, in *EmptyRequest, out *CurrenciesList) error {
	return h.CurrencyratesServiceHandler.GetSettlementCurrencies(ctx, in, out)
}

func (h *currencyratesServiceHandler) GetPriceCurrencies(ctx context.Context, in *EmptyRequest, out *CurrenciesList) error {
	return h.CurrencyratesServiceHandler.GetPriceCurrencies(ctx, in, out)
}

func (h *currencyratesServiceHandler) GetVatCurrencies(ctx context.Context, in *EmptyRequest, out *CurrenciesList) error {
	return h.CurrencyratesServiceHandler.GetVatCurrencies(ctx, in, out)
}

func (h *currencyratesServiceHandler) GetAccountingCurrencies(ctx context.Context, in *EmptyRequest, out *CurrenciesList) error {
	return h.CurrencyratesServiceHandler.GetAccountingCurrencies(ctx, in, out)
}
