package chain

import (
	"encoding/json"
	"fmt"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	"io"
	"net/http"
)

const (
	REVERTED       = "Reverted"
	PENDING_REVERT = "PendingRevert"
	OUTBOUND_MINED = "OutboundMined"
)

type CrossChainResponse struct {
	CrossChainTxs []struct {
		Index      string `json:"index"`
		ZetaFees   string `json:"zeta_fees"`
		CctxStatus struct {
			Status              string `json:"status"`
			StatusMessage       string `json:"status_message"`
			LastUpdateTimestamp string `json:"lastUpdate_timestamp"`
			IsAbortRefunded     bool   `json:"isAbortRefunded"`
		} `json:"cctx_status"`
		InboundParams struct {
			Sender               string `json:"sender"`
			SenderChainId        string `json:"sender_chain_id"`
			Amount               string `json:"amount"`
			ObservedHash         string `json:"observed_hash"`
			TxFinalizationStatus string `json:"tx_finalization_status"`
		} `json:"inbound_params"`
	} `json:"CrossChainTxs"`
}

func CheckCrossChainTransactionStatus(txHash string) (int, error) {
	crossRpc := conf.GConfig.GetString("chain.zetaCrossChainRpc")
	fullUri := fmt.Sprintf("%s/zeta-chain/crosschain/inboundHashToCctxData/%s", crossRpc, txHash)

	resp, err := http.Get(fullUri)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return common.ORDER_INIT, nil
		}
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %v", err)
	}

	var crossChainResp CrossChainResponse
	err = json.Unmarshal(body, &crossChainResp)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	if len(crossChainResp.CrossChainTxs) == 0 {
		return 0, fmt.Errorf("no cross chain transaction found")
	}

	if crossChainResp.CrossChainTxs[0].CctxStatus.Status == OUTBOUND_MINED {
		return common.ORDER_SUCCESS, nil
	} else if crossChainResp.CrossChainTxs[0].CctxStatus.Status == REVERTED ||
		crossChainResp.CrossChainTxs[0].CctxStatus.Status == PENDING_REVERT {
		return common.ORDER_FAILED, nil
	} else {
		return common.ORDER_INIT, nil
	}
}
