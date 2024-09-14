package common

// @name ReturnCodes
// @description API response codes indicating success or type of error.
// The following list enumerates all possible response codes:
// - 20000: Success
// - 20001: Invalid Parameters
// - 20002: Cannot Access Resources
// - 20003: Bad Credentials
// - 20004: Auth Token Error
// - 20015: Internal Error
// - 20020: Exceed Frequency Limits
// - 29999: API Currently Not Supported
const (
	SUCCESS                    = 20000
	INVALID_PARAM              = 20001
	CANNOT_ACCESS_TO_RESOURCES = 20002
	BAD_CREDENTIALS            = 20003
	AUTH_TOKEN_ERROR           = 20004
	INTERNAL_ERROR             = 20015
	EXCEED_FREQUENCY_LIMITS    = 20020
	INSUFFICIENT_PRESALE_QUOTA = 20021
	INVALID_MINT_TYPE          = 20022
	NO_UNUSED_METADATA         = 20023
	NOT_HAVE_WHITELIST         = 20024
	TRANSACTION_NOT_FOUND      = 20025
	ORDER_NOT_FOUND            = 20026
	NOT_SUPPORTED_CHAIN_ID     = 20027
	API_NOT_IN_SUPPORT_NOW     = 29999
)
