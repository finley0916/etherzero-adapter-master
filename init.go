package openwtester

import (
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
	"github.com/cyberskycat/ethereum-classic-adapter/etc"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(etc.Symbol, etc.NewWalletManager())
}
