package decoder

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func DecodeAddress(addr string, defaultNet *chaincfg.Params) (btcutil.Address, error) {
	return btcutil.DecodeAddress(addr, defaultNet)
}

