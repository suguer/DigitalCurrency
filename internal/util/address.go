package util

import (
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func Address2HexString(address string) (string, error) {
	base, err := DecodeCheck(address)
	if err != nil {
		return "", err
	}
	str := BytesToHexString(base)
	if len(str) == 44 {
		str = str[:2] + str[4:]
	}
	return str, nil
}

func HexString2Address(hex string) (string, error) {
	if len(hex) == 40 {
		hex = "0x41" + hex
	} else if len(hex) == 42 && hex[:2] == "0x" {
		hex = hex[:2] + "41" + hex[2:]
	}
	b, err := HexStringToBytes(hex)
	if err != nil {
		return "", err
	}
	a := EncodeCheck(b)
	return a, nil
}

func Generate() (address string, key string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("无法生成私钥: %v", err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	walletAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	address = walletAddress.Hex()
	key = hex.EncodeToString(privateKeyBytes)
	return address, key, nil
}
