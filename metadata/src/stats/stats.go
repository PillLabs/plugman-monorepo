package stats

import (
	"github.com/prometheus/client_golang/prometheus"
)

// stats models for prometheus monitoring

var (
	EthRpcClientError = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "plugman_eth_rpc_client_errors",
		Help: "Eth rpc client call error occurs",
	}, []string{"chain_name"})
	EthScanningBlockNumber = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "plugman_eth_scanning_block_number",
		Help: "Eth scanning block number",
	}, []string{"chain_id"})
	DbError = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "plugman_db_saving_error",
		Help: "Db saving function error occurs",
	})
	ThirdPartyClientError = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "plugman_third_party_client_errors",
		Help: "Third party client call error occurs",
	}, []string{"service_name"})
)

func init() {
	prometheus.MustRegister(EthRpcClientError)
	prometheus.MustRegister(EthScanningBlockNumber)
	prometheus.MustRegister(DbError)
	prometheus.MustRegister(ThirdPartyClientError)
}
