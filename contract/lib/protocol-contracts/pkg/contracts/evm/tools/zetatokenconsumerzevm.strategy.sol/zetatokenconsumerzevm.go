// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zetatokenconsumerzevm

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ZetaTokenConsumerZEVMMetaData contains all meta data concerning the ZetaTokenConsumerZEVM contract.
var ZetaTokenConsumerZEVMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"WETH9Address_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapV2Router_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorSendingETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputCantBeZeta\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidForZEVM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWZETAAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutputCantBeZeta\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyError\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"EthExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"TokenExchangedForZeta\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForEth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"ZetaExchangedForToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH9Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getEthFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zetaTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenFromZeta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"getZetaFromEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getZetaFromToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasZetaLiquidity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200220938038062002209833981810160405281019062000037919062000164565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614806200009f5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b15620000d7576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b815250505050620001fe565b6000815190506200015e81620001e4565b92915050565b600080604083850312156200017e576200017d620001df565b5b60006200018e858286016200014d565b9250506020620001a1858286016200014d565b9150509250929050565b6000620001b882620001bf565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001ef81620001ab565b8114620001fb57600080fd5b50565b60805160601c60a05160601c611f7e6200028b600039600081816105a4015281816106fa01528181610cb50152610e2b0152600081816060015281816103060152818161038c015281816104ee01528181610689015281816109180152818161095f01528181610bdf01528181610c6b01528181610cd701528181610d6b0152610f660152611f7e6000f3fe6080604052600436106100595760003560e01c8063013b2ff8146100ea5780632405620a1461011a57806354c49a2a1461015757806380801f8414610194578063a53fb10b146101bf578063c469cf14146101fc576100e5565b366100e5577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100e3576040517f290ee5a500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b61010460048036038101906100ff919061157c565b610227565b6040516101119190611aa8565b60405180910390f35b34801561012657600080fd5b50610141600480360381019061013c91906115bc565b610412565b60405161014e9190611aa8565b60405180910390f35b34801561016357600080fd5b5061017e60048036038101906101799190611623565b610833565b60405161018b9190611aa8565b60405180910390f35b3480156101a057600080fd5b506101a9610acf565b6040516101b691906119eb565b60405180910390f35b3480156101cb57600080fd5b506101e660048036038101906101e191906115bc565b610b03565b6040516101f39190611aa8565b60405180910390f35b34801561020857600080fd5b50610211610f64565b60405161021e9190611910565b60405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561028f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003414156102ca576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81341015610304576040517fe2f844a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561036c57600080fd5b505af1158015610380573d6000803e3d6000fd5b50505050506103d083347f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610f889092919063ffffffff16565b7f87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da113434604051610401929190611ac3565b60405180910390a134905092915050565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16148061047a5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b156104b1576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008214156104ec576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610572576040517f6edfe50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61059f3330848673ffffffffffffffffffffffffffffffffffffffff1661100e909392919063ffffffff16565b6105ea7f0000000000000000000000000000000000000000000000000000000000000000838573ffffffffffffffffffffffffffffffffffffffff166110979092919063ffffffff16565b6000600267ffffffffffffffff81111561060757610606611d96565b5b6040519080825280602002602001820160405280156106355781602001602082028036833780820191505090505b509050838160008151811061064d5761064c611d67565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250507f0000000000000000000000000000000000000000000000000000000000000000816001815181106106bc576106bb611d67565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398588858b60c8426107459190611c02565b6040518663ffffffff1660e01b8152600401610765959493929190611aec565b600060405180830381600087803b15801561077f57600080fd5b505af1158015610793573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906107bc9190611676565b9050600081600184516107cf9190611c58565b815181106107e0576107df611d67565b5b602002602001015190507f017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f86868360405161081d939291906119b4565b60405180910390a1809350505050949350505050565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141561089b576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008214156108d6576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82821015610910576040517fe2f844a000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61095d3330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661100e909392919063ffffffff16565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d836040518263ffffffff1660e01b81526004016109b69190611aa8565b600060405180830381600087803b1580156109d057600080fd5b505af11580156109e4573d6000803e3d6000fd5b505050507f74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae8283604051610a19929190611ac3565b60405180910390a160008473ffffffffffffffffffffffffffffffffffffffff1683604051610a47906118fb565b60006040518083038185875af1925050503d8060008114610a84576040519150601f19603f3d011682016040523d82523d6000602084013e610a89565b606091505b5050905080610ac4576040517f3794aeaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b829150509392505050565b60006040517f0e6a82b100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161480610b6b5750600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b15610ba2576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415610bdd576040517fb813f54900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c63576040517f8c51927900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610cb03330847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661100e909392919063ffffffff16565b610d1b7f0000000000000000000000000000000000000000000000000000000000000000837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166110979092919063ffffffff16565b6000600267ffffffffffffffff811115610d3857610d37611d96565b5b604051908082528060200260200182016040528015610d665781602001602082028036833780820191505090505b5090507f000000000000000000000000000000000000000000000000000000000000000081600081518110610d9e57610d9d611d67565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508381600181518110610ded57610dec611d67565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398588858b60c842610e769190611c02565b6040518663ffffffff1660e01b8152600401610e96959493929190611aec565b600060405180830381600087803b158015610eb057600080fd5b505af1158015610ec4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610eed9190611676565b905060008160018451610f009190611c58565b81518110610f1157610f10611d67565b5b602002602001015190507f0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b868683604051610f4e939291906119b4565b60405180910390a1809350505050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6110098363a9059cbb60e01b8484604051602401610fa792919061198b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506111f5565b505050565b611091846323b872dd60e01b85858560405160240161102f93929190611954565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506111f5565b50505050565b6000811480611130575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b81526004016110de92919061192b565b60206040518083038186803b1580156110f657600080fd5b505afa15801561110a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061112e91906116ec565b145b61116f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161116690611a88565b60405180910390fd5b6111f08363095ea7b360e01b848460405160240161118e92919061198b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506111f5565b505050565b6000611257826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166112bc9092919063ffffffff16565b90506000815111156112b7578080602001905181019061127791906116bf565b6112b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ad90611a68565b60405180910390fd5b5b505050565b60606112cb84846000856112d4565b90509392505050565b606082471015611319576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161131090611a28565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161134291906118e4565b60006040518083038185875af1925050503d806000811461137f576040519150601f19603f3d011682016040523d82523d6000602084013e611384565b606091505b5091509150611395878383876113a1565b92505050949350505050565b60608315611404576000835114156113fc576113bc85611417565b6113fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113f290611a48565b60405180910390fd5b5b82905061140f565b61140e838361143a565b5b949350505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60008251111561144d5781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114819190611a06565b60405180910390fd5b600061149d61149884611b6b565b611b46565b905080838252602082019050828560208602820111156114c0576114bf611dca565b5b60005b858110156114f057816114d68882611567565b8452602084019350602083019250506001810190506114c3565b5050509392505050565b60008135905061150981611f03565b92915050565b600082601f83011261152457611523611dc5565b5b815161153484826020860161148a565b91505092915050565b60008151905061154c81611f1a565b92915050565b60008135905061156181611f31565b92915050565b60008151905061157681611f31565b92915050565b6000806040838503121561159357611592611dd4565b5b60006115a1858286016114fa565b92505060206115b285828601611552565b9150509250929050565b600080600080608085870312156115d6576115d5611dd4565b5b60006115e4878288016114fa565b94505060206115f587828801611552565b9350506040611606878288016114fa565b925050606061161787828801611552565b91505092959194509250565b60008060006060848603121561163c5761163b611dd4565b5b600061164a868287016114fa565b935050602061165b86828701611552565b925050604061166c86828701611552565b9150509250925092565b60006020828403121561168c5761168b611dd4565b5b600082015167ffffffffffffffff8111156116aa576116a9611dcf565b5b6116b68482850161150f565b91505092915050565b6000602082840312156116d5576116d4611dd4565b5b60006116e38482850161153d565b91505092915050565b60006020828403121561170257611701611dd4565b5b600061171084828501611567565b91505092915050565b60006117258383611731565b60208301905092915050565b61173a81611c8c565b82525050565b61174981611c8c565b82525050565b600061175a82611ba7565b6117648185611bd5565b935061176f83611b97565b8060005b838110156117a05781516117878882611719565b975061179283611bc8565b925050600181019050611773565b5085935050505092915050565b6117b681611c9e565b82525050565b60006117c782611bb2565b6117d18185611be6565b93506117e1818560208601611cd4565b80840191505092915050565b60006117f882611bbd565b6118028185611bf1565b9350611812818560208601611cd4565b61181b81611dd9565b840191505092915050565b6000611833602683611bf1565b915061183e82611dea565b604082019050919050565b6000611856600083611be6565b915061186182611e39565b600082019050919050565b6000611879601d83611bf1565b915061188482611e3c565b602082019050919050565b600061189c602a83611bf1565b91506118a782611e65565b604082019050919050565b60006118bf603683611bf1565b91506118ca82611eb4565b604082019050919050565b6118de81611cca565b82525050565b60006118f082846117bc565b915081905092915050565b600061190682611849565b9150819050919050565b60006020820190506119256000830184611740565b92915050565b60006040820190506119406000830185611740565b61194d6020830184611740565b9392505050565b60006060820190506119696000830186611740565b6119766020830185611740565b61198360408301846118d5565b949350505050565b60006040820190506119a06000830185611740565b6119ad60208301846118d5565b9392505050565b60006060820190506119c96000830186611740565b6119d660208301856118d5565b6119e360408301846118d5565b949350505050565b6000602082019050611a0060008301846117ad565b92915050565b60006020820190508181036000830152611a2081846117ed565b905092915050565b60006020820190508181036000830152611a4181611826565b9050919050565b60006020820190508181036000830152611a618161186c565b9050919050565b60006020820190508181036000830152611a818161188f565b9050919050565b60006020820190508181036000830152611aa1816118b2565b9050919050565b6000602082019050611abd60008301846118d5565b92915050565b6000604082019050611ad860008301856118d5565b611ae560208301846118d5565b9392505050565b600060a082019050611b0160008301886118d5565b611b0e60208301876118d5565b8181036040830152611b20818661174f565b9050611b2f6060830185611740565b611b3c60808301846118d5565b9695505050505050565b6000611b50611b61565b9050611b5c8282611d07565b919050565b6000604051905090565b600067ffffffffffffffff821115611b8657611b85611d96565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611c0d82611cca565b9150611c1883611cca565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611c4d57611c4c611d38565b5b828201905092915050565b6000611c6382611cca565b9150611c6e83611cca565b925082821015611c8157611c80611d38565b5b828203905092915050565b6000611c9782611caa565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015611cf2578082015181840152602081019050611cd7565b83811115611d01576000848401525b50505050565b611d1082611dd9565b810181811067ffffffffffffffff82111715611d2f57611d2e611d96565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f416464726573733a20696e73756666696369656e742062616c616e636520666f60008201527f722063616c6c0000000000000000000000000000000000000000000000000000602082015250565b50565b7f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000600082015250565b7f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008201527f6f74207375636365656400000000000000000000000000000000000000000000602082015250565b7f5361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f60008201527f20746f206e6f6e2d7a65726f20616c6c6f77616e636500000000000000000000602082015250565b611f0c81611c8c565b8114611f1757600080fd5b50565b611f2381611c9e565b8114611f2e57600080fd5b50565b611f3a81611cca565b8114611f4557600080fd5b5056fea26469706673582212208d9034aea3c8399eb52d47a1eab88ad2e74879d4a2d03d208a0237686f63877a64736f6c63430008070033",
}

// ZetaTokenConsumerZEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use ZetaTokenConsumerZEVMMetaData.ABI instead.
var ZetaTokenConsumerZEVMABI = ZetaTokenConsumerZEVMMetaData.ABI

// ZetaTokenConsumerZEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZetaTokenConsumerZEVMMetaData.Bin instead.
var ZetaTokenConsumerZEVMBin = ZetaTokenConsumerZEVMMetaData.Bin

// DeployZetaTokenConsumerZEVM deploys a new Ethereum contract, binding an instance of ZetaTokenConsumerZEVM to it.
func DeployZetaTokenConsumerZEVM(auth *bind.TransactOpts, backend bind.ContractBackend, WETH9Address_ common.Address, uniswapV2Router_ common.Address) (common.Address, *types.Transaction, *ZetaTokenConsumerZEVM, error) {
	parsed, err := ZetaTokenConsumerZEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZetaTokenConsumerZEVMBin), backend, WETH9Address_, uniswapV2Router_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZetaTokenConsumerZEVM{ZetaTokenConsumerZEVMCaller: ZetaTokenConsumerZEVMCaller{contract: contract}, ZetaTokenConsumerZEVMTransactor: ZetaTokenConsumerZEVMTransactor{contract: contract}, ZetaTokenConsumerZEVMFilterer: ZetaTokenConsumerZEVMFilterer{contract: contract}}, nil
}

// ZetaTokenConsumerZEVM is an auto generated Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVM struct {
	ZetaTokenConsumerZEVMCaller     // Read-only binding to the contract
	ZetaTokenConsumerZEVMTransactor // Write-only binding to the contract
	ZetaTokenConsumerZEVMFilterer   // Log filterer for contract events
}

// ZetaTokenConsumerZEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerZEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerZEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZetaTokenConsumerZEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZetaTokenConsumerZEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZetaTokenConsumerZEVMSession struct {
	Contract     *ZetaTokenConsumerZEVM // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerZEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZetaTokenConsumerZEVMCallerSession struct {
	Contract *ZetaTokenConsumerZEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ZetaTokenConsumerZEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZetaTokenConsumerZEVMTransactorSession struct {
	Contract     *ZetaTokenConsumerZEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ZetaTokenConsumerZEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMRaw struct {
	Contract *ZetaTokenConsumerZEVM // Generic contract binding to access the raw methods on
}

// ZetaTokenConsumerZEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMCallerRaw struct {
	Contract *ZetaTokenConsumerZEVMCaller // Generic read-only contract binding to access the raw methods on
}

// ZetaTokenConsumerZEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZetaTokenConsumerZEVMTransactorRaw struct {
	Contract *ZetaTokenConsumerZEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZetaTokenConsumerZEVM creates a new instance of ZetaTokenConsumerZEVM, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVM(address common.Address, backend bind.ContractBackend) (*ZetaTokenConsumerZEVM, error) {
	contract, err := bindZetaTokenConsumerZEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVM{ZetaTokenConsumerZEVMCaller: ZetaTokenConsumerZEVMCaller{contract: contract}, ZetaTokenConsumerZEVMTransactor: ZetaTokenConsumerZEVMTransactor{contract: contract}, ZetaTokenConsumerZEVMFilterer: ZetaTokenConsumerZEVMFilterer{contract: contract}}, nil
}

// NewZetaTokenConsumerZEVMCaller creates a new read-only instance of ZetaTokenConsumerZEVM, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVMCaller(address common.Address, caller bind.ContractCaller) (*ZetaTokenConsumerZEVMCaller, error) {
	contract, err := bindZetaTokenConsumerZEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMCaller{contract: contract}, nil
}

// NewZetaTokenConsumerZEVMTransactor creates a new write-only instance of ZetaTokenConsumerZEVM, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*ZetaTokenConsumerZEVMTransactor, error) {
	contract, err := bindZetaTokenConsumerZEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMTransactor{contract: contract}, nil
}

// NewZetaTokenConsumerZEVMFilterer creates a new log filterer instance of ZetaTokenConsumerZEVM, bound to a specific deployed contract.
func NewZetaTokenConsumerZEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*ZetaTokenConsumerZEVMFilterer, error) {
	contract, err := bindZetaTokenConsumerZEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMFilterer{contract: contract}, nil
}

// bindZetaTokenConsumerZEVM binds a generic wrapper to an already deployed contract.
func bindZetaTokenConsumerZEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZetaTokenConsumerZEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerZEVM.Contract.ZetaTokenConsumerZEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.ZetaTokenConsumerZEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.ZetaTokenConsumerZEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZetaTokenConsumerZEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.contract.Transact(opts, method, params...)
}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMCaller) WETH9Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZetaTokenConsumerZEVM.contract.Call(opts, &out, "WETH9Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) WETH9Address() (common.Address, error) {
	return _ZetaTokenConsumerZEVM.Contract.WETH9Address(&_ZetaTokenConsumerZEVM.CallOpts)
}

// WETH9Address is a free data retrieval call binding the contract method 0xc469cf14.
//
// Solidity: function WETH9Address() view returns(address)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMCallerSession) WETH9Address() (common.Address, error) {
	return _ZetaTokenConsumerZEVM.Contract.WETH9Address(&_ZetaTokenConsumerZEVM.CallOpts)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMCaller) HasZetaLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZetaTokenConsumerZEVM.contract.Call(opts, &out, "hasZetaLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumerZEVM.Contract.HasZetaLiquidity(&_ZetaTokenConsumerZEVM.CallOpts)
}

// HasZetaLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasZetaLiquidity() view returns(bool)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMCallerSession) HasZetaLiquidity() (bool, error) {
	return _ZetaTokenConsumerZEVM.Contract.HasZetaLiquidity(&_ZetaTokenConsumerZEVM.CallOpts)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactor) GetEthFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.contract.Transact(opts, "getEthFromZeta", destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetEthFromZeta(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetEthFromZeta is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromZeta(address destinationAddress, uint256 minAmountOut, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorSession) GetEthFromZeta(destinationAddress common.Address, minAmountOut *big.Int, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetEthFromZeta(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactor) GetTokenFromZeta(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.contract.Transact(opts, "getTokenFromZeta", destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetTokenFromZeta(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetTokenFromZeta is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromZeta(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 zetaTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorSession) GetTokenFromZeta(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, zetaTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetTokenFromZeta(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, outputToken, zetaTokenAmount)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactor) GetZetaFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.contract.Transact(opts, "getZetaFromEth", destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetZetaFromEth(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getZetaFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorSession) GetZetaFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetZetaFromEth(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactor) GetZetaFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.contract.Transact(opts, "getZetaFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetZetaFromToken(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetZetaFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getZetaFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorSession) GetZetaFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.GetZetaFromToken(&_ZetaTokenConsumerZEVM.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMSession) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.Receive(&_ZetaTokenConsumerZEVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMTransactorSession) Receive() (*types.Transaction, error) {
	return _ZetaTokenConsumerZEVM.Contract.Receive(&_ZetaTokenConsumerZEVM.TransactOpts)
}

// ZetaTokenConsumerZEVMEthExchangedForZetaIterator is returned from FilterEthExchangedForZeta and is used to iterate over the raw logs and unpacked data for EthExchangedForZeta events raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMEthExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerZEVMEthExchangedForZeta // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerZEVMEthExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerZEVMEthExchangedForZeta)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerZEVMEthExchangedForZeta)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerZEVMEthExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerZEVMEthExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerZEVMEthExchangedForZeta represents a EthExchangedForZeta event raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMEthExchangedForZeta struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForZeta is a free log retrieval operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) FilterEthExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerZEVMEthExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.FilterLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMEthExchangedForZetaIterator{contract: _ZetaTokenConsumerZEVM.contract, event: "EthExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForZeta is a free log subscription operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) WatchEthExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerZEVMEthExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.WatchLogs(opts, "EthExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerZEVMEthExchangedForZeta)
				if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEthExchangedForZeta is a log parse operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForZeta(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) ParseEthExchangedForZeta(log types.Log) (*ZetaTokenConsumerZEVMEthExchangedForZeta, error) {
	event := new(ZetaTokenConsumerZEVMEthExchangedForZeta)
	if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "EthExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerZEVMTokenExchangedForZetaIterator is returned from FilterTokenExchangedForZeta and is used to iterate over the raw logs and unpacked data for TokenExchangedForZeta events raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMTokenExchangedForZetaIterator struct {
	Event *ZetaTokenConsumerZEVMTokenExchangedForZeta // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerZEVMTokenExchangedForZetaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerZEVMTokenExchangedForZeta)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerZEVMTokenExchangedForZeta)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerZEVMTokenExchangedForZetaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerZEVMTokenExchangedForZetaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerZEVMTokenExchangedForZeta represents a TokenExchangedForZeta event raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMTokenExchangedForZeta struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForZeta is a free log retrieval operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) FilterTokenExchangedForZeta(opts *bind.FilterOpts) (*ZetaTokenConsumerZEVMTokenExchangedForZetaIterator, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.FilterLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMTokenExchangedForZetaIterator{contract: _ZetaTokenConsumerZEVM.contract, event: "TokenExchangedForZeta", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForZeta is a free log subscription operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) WatchTokenExchangedForZeta(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerZEVMTokenExchangedForZeta) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.WatchLogs(opts, "TokenExchangedForZeta")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerZEVMTokenExchangedForZeta)
				if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenExchangedForZeta is a log parse operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForZeta(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) ParseTokenExchangedForZeta(log types.Log) (*ZetaTokenConsumerZEVMTokenExchangedForZeta, error) {
	event := new(ZetaTokenConsumerZEVMTokenExchangedForZeta)
	if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "TokenExchangedForZeta", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerZEVMZetaExchangedForEthIterator is returned from FilterZetaExchangedForEth and is used to iterate over the raw logs and unpacked data for ZetaExchangedForEth events raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMZetaExchangedForEthIterator struct {
	Event *ZetaTokenConsumerZEVMZetaExchangedForEth // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerZEVMZetaExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerZEVMZetaExchangedForEth)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerZEVMZetaExchangedForEth)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerZEVMZetaExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerZEVMZetaExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerZEVMZetaExchangedForEth represents a ZetaExchangedForEth event raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMZetaExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForEth is a free log retrieval operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) FilterZetaExchangedForEth(opts *bind.FilterOpts) (*ZetaTokenConsumerZEVMZetaExchangedForEthIterator, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.FilterLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMZetaExchangedForEthIterator{contract: _ZetaTokenConsumerZEVM.contract, event: "ZetaExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForEth is a free log subscription operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) WatchZetaExchangedForEth(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerZEVMZetaExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.WatchLogs(opts, "ZetaExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerZEVMZetaExchangedForEth)
				if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZetaExchangedForEth is a log parse operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event ZetaExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) ParseZetaExchangedForEth(log types.Log) (*ZetaTokenConsumerZEVMZetaExchangedForEth, error) {
	event := new(ZetaTokenConsumerZEVMZetaExchangedForEth)
	if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "ZetaExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZetaTokenConsumerZEVMZetaExchangedForTokenIterator is returned from FilterZetaExchangedForToken and is used to iterate over the raw logs and unpacked data for ZetaExchangedForToken events raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMZetaExchangedForTokenIterator struct {
	Event *ZetaTokenConsumerZEVMZetaExchangedForToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZetaTokenConsumerZEVMZetaExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZetaTokenConsumerZEVMZetaExchangedForToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZetaTokenConsumerZEVMZetaExchangedForToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZetaTokenConsumerZEVMZetaExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZetaTokenConsumerZEVMZetaExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZetaTokenConsumerZEVMZetaExchangedForToken represents a ZetaExchangedForToken event raised by the ZetaTokenConsumerZEVM contract.
type ZetaTokenConsumerZEVMZetaExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterZetaExchangedForToken is a free log retrieval operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) FilterZetaExchangedForToken(opts *bind.FilterOpts) (*ZetaTokenConsumerZEVMZetaExchangedForTokenIterator, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.FilterLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &ZetaTokenConsumerZEVMZetaExchangedForTokenIterator{contract: _ZetaTokenConsumerZEVM.contract, event: "ZetaExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchZetaExchangedForToken is a free log subscription operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) WatchZetaExchangedForToken(opts *bind.WatchOpts, sink chan<- *ZetaTokenConsumerZEVMZetaExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _ZetaTokenConsumerZEVM.contract.WatchLogs(opts, "ZetaExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZetaTokenConsumerZEVMZetaExchangedForToken)
				if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseZetaExchangedForToken is a log parse operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event ZetaExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_ZetaTokenConsumerZEVM *ZetaTokenConsumerZEVMFilterer) ParseZetaExchangedForToken(log types.Log) (*ZetaTokenConsumerZEVMZetaExchangedForToken, error) {
	event := new(ZetaTokenConsumerZEVMZetaExchangedForToken)
	if err := _ZetaTokenConsumerZEVM.contract.UnpackLog(event, "ZetaExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
