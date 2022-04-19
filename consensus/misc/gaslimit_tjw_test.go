package misc

import (
	//"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/params"
	//"math/big"
	"testing"
)

func TestBlockGasLimits1(t *testing.T) {
	parentG := uint64(4000000)
	// TJW：gaslimit数值在parent和now之间不能差距太大，但是这个修改是哪里完成的？
	// TJW：修改位置找到，core/block_validator.go:110--CalcGasLimit()函数，每次修改一定限度，以靠近desired gaslimit
	nowG := uint64(4500000)
	err := VerifyGaslimit(parentG, nowG)
	if err != nil {
		//println("err们: ", err)
		t.Errorf("%s",err)

	}

}
