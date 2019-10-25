
package etz

import (
	"github.com/blocktree/ethereum-adapter/ethereum"
	"github.com/blocktree/openwallet/log"
)

const (
	Symbol    = "ETZ"
)

type WalletManager struct {
	*ethereum.WalletManager
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = ethereum.NewWalletManager()
	wm.Config = ethereum.NewConfig(Symbol)
	wm.Log = log.NewOWLogger(wm.Symbol())
	return &wm
}

//FullName 币种全名
func (wm *WalletManager) FullName() string {
	return "Ethereum Classic"
}

