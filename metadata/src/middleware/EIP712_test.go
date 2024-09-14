package middleware

import (
	"github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	"github.com/PillLabs/plugman-monorepo/metadata/src/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"os"
	"testing"
)

func TestPlugmanMintSignAndVerify(t *testing.T) {
	mintTo := "0x0d228F021B1F07e02642737Fbd4d46475A603eab"
	address := "0x0d228F021B1F07e02642737Fbd4d46475A603eab"
	privateKey := "fefe107d88fea4820a739c4f0f228351df11e3886d2faf0b9cfe4f011d2126d8"
	var amount uint = 1
	var nonce uint = 0
	var mintType uint = 1
	var traitsBytes [][32]byte
	traits := []string{"meetman", "neon", "yellow", "pipboy2000", "lightsaber", "kumo"}
	for _, trait := range traits {
		traitBytes := utils.MustConvertStringToByte32(trait)
		traitsBytes = append(traitsBytes, traitBytes)
	}
	traitHash := utils.HashByteSlice(traitsBytes)
	name := conf.GConfig.GetString("erc712Domain.name")
	version := conf.GConfig.GetString("erc712Domain.version")
	chainId := conf.GConfig.GetInt64("erc712Domain.chainId")
	verifyingContract := conf.GConfig.GetString("erc712Domain.verifyingContract")
	t.Log(name, version, chainId, verifyingContract)
	t.Log(hexutil.Encode(traitHash))
	var timestamp uint64 = 1718117397
	signature, err := PlugmanMintSign(mintTo, privateKey, amount, nonce, mintType, traitsBytes, timestamp)
	if err != nil {
		t.Errorf("EIP712SignatureSign error: %s", err)
	}
	t.Log("signature:", signature)
	ok, err := PlugmanMintVerify(mintTo, address, signature, amount, nonce, mintType, traitsBytes, timestamp)
	if err != nil {
		t.Errorf("EIP712SignatureVerify error: %s", err)
	}
	if !ok {
		t.Errorf("EIP712 Signature Verify failed")
	} else {
		t.Log("okay")
	}
}

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}
