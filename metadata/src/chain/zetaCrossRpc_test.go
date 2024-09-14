package chain

import (
	"fmt"
	"testing"
)

func TestCheckCrossChainTransactionStatus(t *testing.T) {
	txHash := "0xb69745eeca8cb8f6c6007ed85659088b0e198bbc2c0f6bdc05c1400217c210d0"
	status, err := CheckCrossChainTransactionStatus(txHash)
	if err != nil {
		t.Error(err)
	}
	t.Log(fmt.Sprintf("Transaction %s status is %d", txHash, status))
}
