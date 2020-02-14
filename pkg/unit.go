package pkg

import (
	"errors"
	"fmt"
	"math/big"
)

var (
	kInt, _ = new(big.Int).SetString("1000", 10)
	mInt, _ = new(big.Int).SetString("1000000", 10)
	gInt, _ = new(big.Int).SetString("1000000000", 10)
	tInt, _ = new(big.Int).SetString("1000000000000", 0)
	pInt, _ = new(big.Int).SetString("1000000000000000", 0)
	eInt, _ = new(big.Int).SetString("1000000000000000000", 0)
	zInt, _ = new(big.Int).SetString("1000000000000000000000", 0)
	yInt, _ = new(big.Int).SetString("1000000000000000000000000", 0)

	unitMaps map[string]*big.Int = map[string]*big.Int{
		"K": kInt, "M": mInt,
		"G": gInt, "T": tInt,
		"P": pInt, "E": eInt,
		"Z": zInt, "Y": yInt,
	}

)


func ConvertUintToBigInt(num int, unit string) (*big.Int, error) {
	if val, ok := unitMaps[unit]; !ok {
		return nil, errors.New("invalid unit")
	} else {
		return new(big.Int).Mul(big.NewInt(int64(num)), val), nil
	}
}

func ConvertFloatToUnitString(num float64, format string, unit string) (string, error) {
	var result string
	if val, ok := unitMaps[unit]; !ok {
		return "", errors.New("invalid unit")
	} else {
		numFloat64 := new(big.Float).SetFloat64(num)
		quo, _ := new(big.Float).Quo(numFloat64, new(big.Float).SetInt(val)).Float64()
		result = fmt.Sprintf(format, quo)
	}
	return fmt.Sprintf("%s%s", result, unit), nil
}
