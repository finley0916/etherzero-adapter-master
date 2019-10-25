package openwtester

import (
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
	"github.com/finley0916/etherzero-adapter-master/tree/master/etz"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(etz.Symbol, etz.NewWalletManager())
}
