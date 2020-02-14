package pkg

import (
	"errors"
	"math/big"
)

func ConvertUintToBigInt(num int, unit string) (*big.Int, error) {
	kInt, _ := new(big.Int).SetString("1000", 10)
	mInt, _ := new(big.Int).SetString("1000000", 10)
	gInt, _ := new(big.Int).SetString("1000000000", 10)
	tInt, _ := new(big.Int).SetString("1000000000000", 0)
	pInt, _ := new(big.Int).SetString("1000000000000000", 0)
	eInt, _ := new(big.Int).SetString("1000000000000000000", 0)
	zInt, _ := new(big.Int).SetString("1000000000000000000000", 0)
	yInt, _ := new(big.Int).SetString("1000000000000000000000000", 0)

	var unitMaps map[string]*big.Int = map[string]*big.Int{
		"K": kInt, "M": mInt,
		"G": gInt, "T": tInt,
		"P": pInt, "E": eInt,
		"Z": zInt, "Y": yInt,
	}

	if val, ok := unitMaps[unit]; !ok {
		return nil, errors.New("invalid unit")
	} else {
		return new (big.Int).Mul(big.NewInt(int64(num)), val), nil
	}
}