package argus

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

type ArgusService struct {
	rpc        string
	privateKey *ecdsa.PrivateKey
	contract   *Argus
	chainId    *big.Int
}

func NewArgusService(rpc, argusAddress, operatorPk string) (*ArgusService, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	_privateKey, err := crypto.HexToECDSA(operatorPk)
	if err != nil {
		return nil, err
	}
	argusContract, err := NewArgus(common.HexToAddress(argusAddress), client)
	if err != nil {
		return nil, err
	}
	_chainId, _ := client.ChainID(context.Background())
	argus := &ArgusService{
		rpc:        rpc,
		chainId:    _chainId,
		privateKey: _privateKey,
		contract:   argusContract,
	}
	return argus, nil
}

func MakeTxDataByAbi(abiJson, method string, args ...interface{}) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return nil, err
	}
	pack, err2 := parsed.Pack(method, args)
	if err2 != nil {
		return nil, err
	}
	return pack, nil
}

func MakeCallData(to common.Address, value *big.Int, data []byte) CallData {
	cd := CallData{
		Flag:  big.NewInt(0),
		To:    to,
		Value: value,
		Data:  data,
		Hint:  common.Hex2Bytes("0x"),
		Extra: common.Hex2Bytes("0x"),
	}
	return cd
}

func (a *ArgusService) GetTransactOpts() (*bind.TransactOpts, error) {
	o, err := bind.NewKeyedTransactorWithChainID(a.privateKey, a.chainId)
	return o, err
}

func (a *ArgusService) ExecTransaction(opts *bind.TransactOpts, callData CallData) (*types.Transaction, error) {
	if opts == nil {
		o, _ := bind.NewKeyedTransactorWithChainID(a.privateKey, a.chainId)
		opts = o
	}
	tx, err := a.contract.ExecTransaction(opts, callData)
	return tx, err
}
