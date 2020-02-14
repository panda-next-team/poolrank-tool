package pkg

import (
	"testing"
)

func TestConvertUintToBigInt(t *testing.T) {
	unitBigInt, err :=  ConvertUintToBigInt(100, "T")
	if err != nil {
		t.Fatal(err.Error())
	}

	if unitBigInt.String() != "100000000000000" {
		t.Errorf("invalid result")
		t.Fail()
	}
}

func TestConvertFloatToUnitString(t *testing.T) {
	floatNum := 1.546609893555465e+13
	unitString, err := ConvertFloatToUnitString(floatNum, "%.8f", "E")
	if err != nil {
		t.Fatal(err.Error())
	}

	if unitString != "15.47T" {
		t.Errorf("invalid result")
		t.Fail()
	}
}

