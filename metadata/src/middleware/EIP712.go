package middleware

import (
	"encoding/hex"
	"fmt"
	"github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	"github.com/PillLabs/plugman-monorepo/metadata/src/utils"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// 'PlugmanMint(address mintTo,uint256 amount,uint256 nonce,uint256 timestamp,bytes32 traitHash,uint8 mintType)'
func newPlugmanMint(mintTo string, amount, nonce, mintType uint, traitHash []byte, timestamp uint64) apitypes.TypedData {
	return apitypes.TypedData{
		Types: apitypes.Types{
			"PlugmanMint": []apitypes.Type{
				{Name: "mintTo", Type: "address"},
				{Name: "amount", Type: "uint256"},
				{Name: "nonce", Type: "uint256"},
				{Name: "timestamp", Type: "uint256"},
				{Name: "traitHash", Type: "bytes32"},
				{Name: "mintType", Type: "uint8"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "PlugmanMint",
		Domain: apitypes.TypedDataDomain{
			Name:              conf.GConfig.GetString("erc712Domain.name"),
			Version:           conf.GConfig.GetString("erc712Domain.version"),
			ChainId:           math.NewHexOrDecimal256(conf.GConfig.GetInt64("erc712Domain.chainId")),
			VerifyingContract: conf.GConfig.GetString("erc712Domain.verifyingContract"),
		},
		Message: apitypes.TypedDataMessage{
			"mintTo":    mintTo,
			"amount":    math.NewHexOrDecimal256(int64(amount)),
			"nonce":     math.NewHexOrDecimal256(int64(nonce)),
			"timestamp": math.NewHexOrDecimal256(int64(timestamp)),
			"traitHash": traitHash,
			"mintType":  math.NewHexOrDecimal256(int64(mintType)),
		},
	}
}

func PlugmanMintSign(mintTo, privateKey string, amount, nonce, mintType uint, traitValue [][32]byte, timestamp uint64) (string, error) {
	traitHash := utils.HashByteSlice(traitValue)
	signerData := newPlugmanMint(mintTo, amount, nonce, mintType, traitHash, timestamp)
	return EIP712SignatureSign(signerData, privateKey)
}

func PlugmanMintVerify(mintTo, address, signature string, amount, nonce, mintType uint, traitValue [][32]byte, timestamp uint64) (bool, error) {
	traitHash := utils.HashByteSlice(traitValue)
	signerData := newPlugmanMint(mintTo, amount, nonce, mintType, traitHash, timestamp)
	return EIP712SignatureVerify(address, signature, signerData)
}

func EIP712SignatureSign(signerData apitypes.TypedData, privateKey string) (string, error) {
	typedDataHash, err := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	if err != nil {
		return "", fmt.Errorf("hash struct message error: %s", err.Error())
	}
	domainSeparator, err := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err != nil {
		return "", fmt.Errorf("hash struct EIP712 domain error: %s", err.Error())
	}
	// format them into an EIP-712 compliant byte string.
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-712.md#specification
	rawData := append([]byte("\x19\x01"), append(domainSeparator, typedDataHash...)...)
	challengeHash := crypto.Keccak256Hash(rawData)
	key, err := crypto.HexToECDSA(utils.Strip0x(privateKey))
	if err != nil {
		return "", err
	}
	if sig, err := crypto.Sign(challengeHash.Bytes(), key); err != nil {
		return "", err
	} else {
		//https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
		if sig[64] == 0 || sig[64] == 1 {
			sig[64] += 27
		}
		return "0x" + hex.EncodeToString(sig), nil
	}
}

func EIP712SignatureVerify(address, signature string, signerData apitypes.TypedData) (bool, error) {
	typedDataHash, err := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	if err != nil {
		return false, fmt.Errorf("hash struct message error: %s", err.Error())
	}
	domainSeparator, err := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err != nil {
		return false, fmt.Errorf("hash struct EIP712 domain error: %s", err.Error())
	}
	// format them into an EIP-712 compliant byte string.
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-712.md#specification
	rawData := append([]byte("\x19\x01"), append(domainSeparator, typedDataHash...)...)
	challengeHash := crypto.Keccak256Hash(rawData)
	signatureBytes, ok, err := utils.PreCheckSignature(signature)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	pubKey, err := utils.EcRecoverForPubKey(signatureBytes, challengeHash.Bytes())
	if err != nil {
		return false, err
	}
	ok, err = utils.RecoveryAddressCheck(pubKey, address)
	if err != nil {
		return false, err
	}
	return ok, nil
}
