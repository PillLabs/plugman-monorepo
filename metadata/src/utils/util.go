package utils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
	"strings"
	"time"
)

var jwtSecret = []byte("Nftfanatic")

func WeiToEth(wei *big.Int) *big.Float {
	weiFloat := new(big.Float).SetInt(wei)
	ethValue := new(big.Float).Quo(weiFloat, big.NewFloat(1e18))
	return ethValue
}

func EthToWei(eth *big.Float) *big.Int {
	ethToWeiMultiplier := new(big.Float).SetFloat64(1e18)

	weiFloat := new(big.Float).Mul(eth, ethToWeiMultiplier)

	weiBigInt := new(big.Int)
	weiFloat.Int(weiBigInt)

	return weiBigInt
}

func GeneratePseudoHash() common2.Hash {
	data := make([]byte, 64)

	_, err := rand.Read(data)
	if err != nil {
		log.Fatalf("Failed to generate random data: %v", err)
	}

	var hash common2.Hash
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	copy(hash[:], hasher.Sum(nil))

	return hash
}

func PreCheckSignature(_signature string) ([]byte, bool, error) {
	_signature = Strip0x(_signature)
	signature, err := hex.DecodeString(_signature)
	if err != nil {
		return []byte{}, false, err
	}
	if len(signature) != 65 {
		return []byte{}, false, fmt.Errorf("invalid signature length: %d", len(signature))
	}
	// making sure the signature's recovery ID (the last byte) is set to 27 or 28
	if signature[64] != 27 && signature[64] != 28 {
		return []byte{}, false, fmt.Errorf("invalid recovery id: %d", signature[64])
	}
	// subtract 27 from the recovery ID to convert it to a 0 or 1 , for Ecrecover function.
	signature[64] -= 27
	return signature, true, nil
}

// EcRecoverForPubKey derive a public key from the provided signature
func EcRecoverForPubKey(signature, challengeHash []byte) (*ecdsa.PublicKey, error) {
	pubKeyRaw, err := crypto.Ecrecover(challengeHash, signature)
	if err != nil {
		return &ecdsa.PublicKey{}, fmt.Errorf("invalid signature: %s", err.Error())
	}
	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return &ecdsa.PublicKey{}, err
	}
	return pubKey, nil
}

func RecoveryAddressCheck(pubkey *ecdsa.PublicKey, signatureAddress string) (bool, error) {
	ok := common2.IsHexAddress(signatureAddress)
	if !ok {
		return false, fmt.Errorf("invalid input hex address: %s", signatureAddress)
	}
	signatureAddress_ := common2.HexToAddress(signatureAddress)
	recoveredAddress := crypto.PubkeyToAddress(*pubkey)
	return bytes.Equal(signatureAddress_.Bytes(), recoveredAddress.Bytes()), nil
}

func IsTimeOut(now, checkTime time.Time) bool {
	if now.After(checkTime) {
		return true
	} else {
		return false
	}
}

func Strip0x(s string) string {
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}
	return s
}

func MustConvertStringToByte32(s string) [32]byte {
	var bytes32 [32]byte
	bs := []byte(s)
	copy(bytes32[:], bs[:len(s)])
	return bytes32
}

func HashByteSlice(slice [][32]byte) []byte {
	hasher := sha3.NewLegacyKeccak256()
	for _, trait := range slice {
		hasher.Write(trait[:])
	}
	return hasher.Sum(nil)
}

type Claims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// GenerateToken generate tokens used for auth
func GenerateToken(userId string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Nftfanatic",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
