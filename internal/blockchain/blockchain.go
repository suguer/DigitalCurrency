package blockchain

import (
	"DigitalCurrency/internal/config"
	"context"
)

type Blockchain struct {
	Chain  string
	Config *config.EthConfig
	ctx    context.Context
}
