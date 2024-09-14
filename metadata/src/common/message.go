package common

import "errors"

var MsgFlags = map[int]string{
	SUCCESS:                    "success",
	INVALID_PARAM:              "invalid params",
	CANNOT_ACCESS_TO_RESOURCES: "cannot access to this resources",
	BAD_CREDENTIALS:            "bad credentials for user's signature",
	AUTH_TOKEN_ERROR:           "auth token error",
	INTERNAL_ERROR:             "internal server error",
	EXCEED_FREQUENCY_LIMITS:    "api request exceed frequency limits",
	INSUFFICIENT_PRESALE_QUOTA: "insufficient presale quota to mint",
	INVALID_MINT_TYPE:          "invalid mint type",
	NO_UNUSED_METADATA:         "no unused metadata",
	NOT_HAVE_WHITELIST:         "this address doesn't have any whitelist now",
	TRANSACTION_NOT_FOUND:      "transaction not found",
	ORDER_NOT_FOUND:            "order not found",
	NOT_SUPPORTED_CHAIN_ID:     "not supported chain id",
	API_NOT_IN_SUPPORT_NOW:     "this api is not in support now",
}

// GetMsg get error information based on Code
func GetMsg(code int) (string, error) {
	msg, ok := MsgFlags[code]
	if ok {
		return msg, nil
	}
	return msg, errors.New("no such code found")
}
