package storage

import (
	"github.com/itering/substrate-api-rpc/websocket"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type Dao interface {
	DB() *gorm.DB
	SpecialMetadata(int) string
	RPCPool() *websocket.Pool
}

type Block struct {
	BlockNum       int    `json:"block_num"`
	BlockTimestamp int    `json:"block_timestamp"`
	Hash           string `json:"hash"`
	SpecVersion    int    `json:"spec_version"`
	Validator      string `json:"validator"`
	Finalized      bool   `json:"finalized"`
}

type Extrinsic struct {
	ExtrinsicIndex     string          `json:"extrinsic_index" `
	CallCode           string          `json:"call_code"`
	CallModuleFunction string          `json:"call_module_function" `
	CallModule         string          `json:"call_module"`
	Params             []byte          `json:"params"`
	AccountId          string          `json:"account_id"`
	Signature          string          `json:"signature"`
	Nonce              int             `json:"nonce"`
	Era                string          `json:"era"`
	ExtrinsicHash      string          `json:"extrinsic_hash"`
	Success            bool            `json:"success"`
	Fee                decimal.Decimal `json:"fee"`
}

type Event struct {
	BlockNum      int    `json:"block_num"`
	ExtrinsicIdx  int    `json:"extrinsic_idx"`
	ModuleId      string `json:"module_id" `
	EventId       string `json:"event_id" `
	Params        []byte `json:"params" sql:"type:text;" `
	ExtrinsicHash string `json:"extrinsic_hash" sql:"default: null" `
	EventIdx      int    `json:"event_idx"`
}

type ExtrinsicParam struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	ValueRaw string      `json:"valueRaw"`
}

type EventParam struct {
	Type     string      `json:"type"`
	Value    interface{} `json:"value"`
	ValueRaw string      `json:"valueRaw"`
}
