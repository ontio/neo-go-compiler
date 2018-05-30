package util

import (
	"fmt"
	"math/big"
	"testing"
)

func TestConvertBigIntegerToBytes(t *testing.T) {
	i := big.NewInt(int64(15))

	b := ConvertBigIntegerToBytes(i)
	fmt.Println(b)
}
