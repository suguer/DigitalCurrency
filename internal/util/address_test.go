package util

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	address, key, err := Generate()
	tronlinkAddress, _ := HexString2Address(address)
	fmt.Printf("address: %v\n", address)
	fmt.Printf("tronlinkAddress: %v\n", tronlinkAddress)
	fmt.Printf("key: %v\n", key)
	fmt.Printf("err: %v\n", err)
}
