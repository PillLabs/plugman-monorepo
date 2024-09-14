package common

import "errors"

var (
	ErrInsufficientPresaleQuota = errors.New("insufficient presale quota to mint")
	ErrInvalidMintType          = errors.New("invalid mint type")
	ErrNoUnusedMetadata         = errors.New("no unused metadata")
	ErrReplaceUnderPriced       = errors.New("replacement transaction underpriced")
)
