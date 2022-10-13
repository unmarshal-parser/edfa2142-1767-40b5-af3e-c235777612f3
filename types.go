// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"math/big"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = datatypes.JSON{}
	_ = time.Time{}
)

func GetEthWithdrawnEventHash() string {
	return "0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b"
}

type EthWithdrawnEvent struct {
	Target string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:bfd83f78-b797-4225-8411-f595bc56851d,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:bfd83f78-b797-4225-8411-f595bc56851d,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:bfd83f78-b797-4225-8411-f595bc56851d,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSwapTargetAddedEventHash() string {
	return "0xb907822409611d127ab6a64611591b98e03a6a85ade4f258bae26b7c1efdfeaf"
}

type SwapTargetAddedEvent struct {
	Target string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:3c413372-f999-48db-8b08-c2be373d6bfb,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:3c413372-f999-48db-8b08-c2be373d6bfb,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:3c413372-f999-48db-8b08-c2be373d6bfb,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSwapTargetRemovedEventHash() string {
	return "0x393b8be3e26787f19285ecd039dfd80bc6507828750f4d50367e6efe2524695c"
}

type SwapTargetRemovedEvent struct {
	Target string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:67fb4d71-5ac0-41c6-af46-2c1e644d2052,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:67fb4d71-5ac0-41c6-af46-2c1e644d2052,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:67fb4d71-5ac0-41c6-af46-2c1e644d2052,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTokenWithdrawnEventHash() string {
	return "0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620"
}

type TokenWithdrawnEvent struct {
	Token  string
	Target string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:ea9c7cfd-14f0-438a-b968-1ba5ab4382b8,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:ea9c7cfd-14f0-438a-b968-1ba5ab4382b8,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:ea9c7cfd-14f0-438a-b968-1ba5ab4382b8,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTransferOwnershipMethodHash() string {
	return "f2fde38b"
}

type TransferOwnershipMethod struct {
	NewOwner string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:21be39a0-f7be-46b2-b9eb-85a48b38c713,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:21be39a0-f7be-46b2-b9eb-85a48b38c713,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteEthToTokenMethodHash() string {
	return "3c2b9a7d"
}

type FillQuoteEthToTokenMethod struct {
	BuyTokenAddress string
	Target          string
	SwapCallData    []byte
	FeeAmount       decimal.Decimal `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:31bc0c01-de4d-4385-8100-0d8f1c7f088b,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:31bc0c01-de4d-4385-8100-0d8f1c7f088b,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToEthWithPermitMethodHash() string {
	return "b3093838"
}

type FillQuoteTokenToEthWithPermitMethod struct {
	SellTokenAddress         string
	Target                   string
	SwapCallData             []byte
	SellAmount               decimal.Decimal `gorm:"type:numeric"`
	FeePercentageBasisPoints decimal.Decimal `gorm:"type:numeric"`
	PermitData               datatypes.JSON

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeePercentageBasisPoints      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeePercentageBasisPoints float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:61b31a5f-be5c-4f3e-a534-c1c72a79e1e2,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:61b31a5f-be5c-4f3e-a534-c1c72a79e1e2,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToTokenWithPermitMethodHash() string {
	return "b0480bbd"
}

type FillQuoteTokenToTokenWithPermitMethod struct {
	SellTokenAddress string
	BuyTokenAddress  string
	Target           string
	SwapCallData     []byte
	SellAmount       decimal.Decimal `gorm:"type:numeric"`
	FeeAmount        decimal.Decimal `gorm:"type:numeric"`
	PermitData       datatypes.JSON

	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`
	TokenPriceSellAmount      float64 `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:b6ccdb15-9b10-4bac-8b4f-3f6329656d5a,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:b6ccdb15-9b10-4bac-8b4f-3f6329656d5a,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetUpdateSwapTargetsMethodHash() string {
	return "97bbda0e"
}

type UpdateSwapTargetsMethod struct {
	Target string
	Add    bool

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:ed54650d-d529-4d33-8d3c-46853df10c92,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:ed54650d-d529-4d33-8d3c-46853df10c92,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetWithdrawEthMethodHash() string {
	return "1b9a91a4"
}

type WithdrawEthMethod struct {
	To     string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:02862680-4454-495b-85f6-63501a59a520,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:02862680-4454-495b-85f6-63501a59a520,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetWithdrawTokenMethodHash() string {
	return "01e33667"
}

type WithdrawTokenMethod struct {
	Token  string
	To     string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:112f0048-3fe5-4ceb-b7b3-11da96f720d1,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:112f0048-3fe5-4ceb-b7b3-11da96f720d1,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToEthMethodHash() string {
	return "999b6464"
}

type FillQuoteTokenToEthMethod struct {
	SellTokenAddress         string
	Target                   string
	SwapCallData             []byte
	SellAmount               decimal.Decimal `gorm:"type:numeric"`
	FeePercentageBasisPoints decimal.Decimal `gorm:"type:numeric"`

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeePercentageBasisPoints      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeePercentageBasisPoints float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:2172e65b-6245-4b9c-b4ef-4c81dcaeef9a,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:2172e65b-6245-4b9c-b4ef-4c81dcaeef9a,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToTokenMethodHash() string {
	return "55e4b7be"
}

type FillQuoteTokenToTokenMethod struct {
	SellTokenAddress string
	BuyTokenAddress  string
	Target           string
	SwapCallData     []byte
	SellAmount       decimal.Decimal `gorm:"type:numeric"`
	FeeAmount        decimal.Decimal `gorm:"type:numeric"`

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`
	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:38defc0b-af1a-4699-9cbc-a3fc23749c67,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:38defc0b-af1a-4699-9cbc-a3fc23749c67,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

type LastSyncedBlock struct {
	Contract    string `gorm:"primaryKey"`
	ChainID     string `gorm:"primaryKey"`
	SyncType    string `gorm:"primaryKey"`
	BlockNumber uint64
}

// Plugin Models

type TokenDetails struct {
	ID      int
	Address string `gorm:"uniqueIndex:address_and_chain"`
	Symbol  string
	ChainID string `gorm:"uniqueIndex:address_and_chain"`
	Decimal int
}

// Config

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
	TablePrefix      string `mapstructure:"table_prefix"`
	CreateBatchSize  int    `mapstructure:"create_batch_size"`
}

type IndexerConfig struct {
	EthEndpoint       string `mapstructure:"eth_endpoint"`
	ContractAddress   string `mapstructure:"contract_address"`
	StartBlock        int    `mapstructure:"start_block"`
	ApiKey            string `mapstructure:"api_key"`
	PostgresConfig    `mapstructure:"postgres_config"`
	LagToHighestBlock int `mapstructure:"lag_to_highest_block"`
	StepSize          int `mapstructure:"step_size"`
	ParallelCalls     int `mapstructure:"parallel_calls_for_logs"`
}

func (i *IndexerConfig) AssignDefaults() {
	if i.PostgresConfig.CreateBatchSize == 0 {
		i.PostgresConfig.CreateBatchSize = 100
	}
	if i.StepSize == 0 {
		i.StepSize = 50
	}
	if i.LagToHighestBlock == 0 {
		i.LagToHighestBlock = 10
	}
}
