package web3

import (
	"github.com/Tinddd28/GoPTL/internal/models"
)

const (
	lenTRC20 = 34
	lenTON   = 64
	lenSOL   = 44
	lenERC20 = 42
	lenBEP20 = 42
)

func CheckAddressFormat(address string, network models.Network) bool {
	networkCode := network.NetworkCode
	switch networkCode {
	case "TRC20":
		return CheckTRC20(address)
	case "TON":
		return CheckTON(address)
	case "ERC20":
		return CheckERC20(address)
	case "BEP20":
		return CheckBEP20(address)
	case "SOL":
		return CheckSOL(address)
	}
	return false
}

// 'TRC20', 'TON', 'ERC20', 'BEP20', 'SOL'
func CheckTRC20(address string) bool {
	// check TRC20 address
	if len(address) != lenTRC20 || address[0] != 'T' {
		return false
	}
	return true
}

func CheckTON(address string) bool {
	// check TON address
	if len(address) != lenTON || address[0] != '0' {
		return false
	}
	return true
}

func CheckERC20(address string) bool {
	// check ERC20 address
	if len(address) != lenERC20 || !(address[0] == '0' && address[1] == 'x') {
		return false
	}
	return true
}

func CheckBEP20(address string) bool {
	// check BEP20 address
	if len(address) != lenBEP20 || !(address[0] == '0' && address[1] == 'x') {
		return false
	}
	return true
}

func CheckSOL(address string) bool {
	// check SOL address
	return len(address) == lenSOL
}
