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

func TestBase58(t *testing.T) {
	address := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	hex, _ := DecodeCheck(address)
	hexString := BytesToHexString(hex)
	fmt.Printf("res: %v\n", hexString)
	if hexString != "41a614f803b6fd780986a42c78ec9c7f77e6ded13c" {
		t.Fatalf("res: %v\n", hexString)
	}
	hexBytes, _ := HexStringToBytes(hexString)
	address = EncodeCheck(hexBytes)
	fmt.Printf("address: %v\n", address)
}
