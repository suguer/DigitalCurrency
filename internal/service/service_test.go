package service

import (
	"DigitalCurrency/internal/config"
	"context"
	"testing"
)

func TestTronStart(t *testing.T) {
	conf := &config.EthConfig{
		Chain:  "tron",
		Node:   "https://api.trongrid.io",
		ApiKey: "6bf9c088-5578-4ffb-aa66-d41e40594f76",
		Contract: []config.ContractConfig{
			{Token: "USDT", Address: "41a614f803b6fd780986a42c78ec9c7f77e6ded13c", Decimal: 6},
		},
	}
	service := NewTronService(context.Background(), conf)
	service.Start()
}
