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
		t.Errorf("invalid transaction id")
		t.Fail()
	}
}

