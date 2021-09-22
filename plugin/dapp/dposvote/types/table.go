package types

import (
	"fmt"

	"github.com/33cn/chain33/common/db"

	"github.com/33cn/chain33/common/db/table"
	"github.com/33cn/chain33/types"
)

/*
table  struct
data:  voter
index: FromAddr,Pubkey,Votes,Index,Time
*/

var optDposVoter = &table.Option{
	Prefix:  "LODB-dpos",
	Name:    "voter",
	Primary: "index",
	Index:   []string{"addr", "pubkey"},
}

func NewDposVoteTable(kvdb db.KV) *table.Table {
	rowmeta := NewDposVoterRow()
	table, err := table.NewTable(rowmeta, kvdb, optDposVoter)
	if err != nil {
		panic(err)
	}
	return table
}

type DposVoterRow struct {
	*DposVoter
}

func NewDposVoterRow() *DposVoterRow {
	return &DposVoterRow{DposVoter: &DposVoter{}}
}

func (tx *DposVoterRow) CreateRow() *table.Row {
	return &table.Row{Data: &DposVoter{}}
}

func (tx *DposVoterRow) SetPayload(data types.Message) error {
	if txdata, ok := data.(*DposVoter); ok {
		tx.DposVoter = txdata
		return nil
	}
	return types.ErrTypeAsset
}

func (tx *DposVoterRow) Get(key string) ([]byte, error) {
	if key == "index" {
		return []byte(fmt.Sprintf("%018d", tx.Index)), nil
	} else if key == "addr" {
		return []byte(tx.FromAddr), nil
	} else if key == "pubkey" {
		return tx.Pubkey, nil
	}

	return nil, types.ErrNotFound
}

var optDposCandidator = &table.Option{
	Prefix:  "LODB-dpos",
	Name:    "candidator",
	Primary: "pubkey",
	Index:   []string{"status"},
}

func NewDposCandidatorTable(kvdb db.KV) *table.Table {
	rowmeta := NewDposCandidatorRow()
	table, err := table.NewTable(rowmeta, kvdb, optDposCandidator)
	if err != nil {
		panic(err)
	}
	return table
}

type DposCandidatorRow struct {
	*CandidatorInfo
}

func NewDposCandidatorRow() *DposCandidatorRow {
	return &DposCandidatorRow{CandidatorInfo: &CandidatorInfo{}}
}

func (tx *DposCandidatorRow) CreateRow() *table.Row {
	return &table.Row{Data: &CandidatorInfo{}}
}

func (tx *DposCandidatorRow) SetPayload(data types.Message) error {
	if txdata, ok := data.(*CandidatorInfo); ok {
		tx.CandidatorInfo = txdata
		return nil
	}
	return types.ErrTypeAsset
}

func (tx *DposCandidatorRow) Get(key string) ([]byte, error) {
	if key == "pubkey" {
		return tx.Pubkey, nil
	} else if key == "status" {
		return []byte(fmt.Sprintf("%2d", tx.Status)), nil
	}

	return nil, types.ErrNotFound
}

var optDposVrfm = &table.Option{
	Prefix:  "LODB-dpos",
	Name:    "vrfm",
	Primary: "index",
	Index:   []string{"pubkey_cycle", "cycle"},
}

func NewDposVrfMTable(kvdb db.KV) *table.Table {
	rowmeta := NewDposVrfMRow()
	table, err := table.NewTable(rowmeta, kvdb, optDposVrfm)
	if err != nil {
		panic(err)
	}
	return table
}

type DposVrfMRow struct {
	*DposVrfM
}

func NewDposVrfMRow() *DposVrfMRow {
	return &DposVrfMRow{DposVrfM: &DposVrfM{}}
}

func (tx *DposVrfMRow) CreateRow() *table.Row {
	return &table.Row{Data: &DposVrfM{}}
}

func (tx *DposVrfMRow) SetPayload(data types.Message) error {
	if txdata, ok := data.(*DposVrfM); ok {
		tx.DposVrfM = txdata
		return nil
	}
	return types.ErrTypeAsset
}

func (tx *DposVrfMRow) Get(key string) ([]byte, error) {
	if key == "index" {
		return []byte(fmt.Sprintf("%018d", tx.Index)), nil
	} else if key == "pubkey_cycle" {
		return []byte(fmt.Sprintf("%X:%018d", tx.Pubkey, tx.Cycle)), nil
	} else if key == "cycle" {
		return []byte(fmt.Sprintf("%018d", tx.Cycle)), nil
	}

	return nil, types.ErrNotFound
}

var optDposVrfrp = &table.Option{
	Prefix:  "LODB-dpos",
	Name:    "vrfrp",
	Primary: "index",
	Index:   []string{"pubkey_cycle", "cycle"},
}

func NewDposVrfRPTable(kvdb db.KV) *table.Table {
	rowmeta := NewDposVrfRPRow()
	table, err := table.NewTable(rowmeta, kvdb, optDposVrfrp)
	if err != nil {
		panic(err)
	}
	return table
}

type DposVrfRPRow struct {
	*DposVrfRP
}

func NewDposVrfRPRow() *DposVrfRPRow {
	return &DposVrfRPRow{DposVrfRP: &DposVrfRP{}}
}

func (tx *DposVrfRPRow) CreateRow() *table.Row {
	return &table.Row{Data: &DposVrfRP{}}
}

func (tx *DposVrfRPRow) SetPayload(data types.Message) error {
	if txdata, ok := data.(*DposVrfRP); ok {
		tx.DposVrfRP = txdata
		return nil
	}
	return types.ErrTypeAsset
}

func (tx *DposVrfRPRow) Get(key string) ([]byte, error) {
	if key == "index" {
		return []byte(fmt.Sprintf("%018d", tx.Index)), nil
	} else if key == "pubkey_cycle" {
		return []byte(fmt.Sprintf("%X:%018d", tx.Pubkey, tx.Cycle)), nil
	} else if key == "cycle" {
		return []byte(fmt.Sprintf("%018d", tx.Cycle)), nil
	}

	return nil, types.ErrNotFound
}

var optDposCb = &table.Option{
	Prefix:  "LODB-dpos",
	Name:    "cb",
	Primary: "cycle",
	Index:   []string{"height", "hash"},
}

func NewDposCBTable(kvdb db.KV) *table.Table {
	rowmeta := NewDposCBRow()
	table, err := table.NewTable(rowmeta, kvdb, optDposCb)
	if err != nil {
		panic(err)
	}
	return table
}

type DposCBRow struct {
	*DposCycleBoundaryInfo
}

func NewDposCBRow() *DposCBRow {
	return &DposCBRow{DposCycleBoundaryInfo: &DposCycleBoundaryInfo{}}
}

func (tx *DposCBRow) CreateRow() *table.Row {
	return &table.Row{Data: &DposCycleBoundaryInfo{}}
}

func (tx *DposCBRow) SetPayload(data types.Message) error {
	if txdata, ok := data.(*DposCycleBoundaryInfo); ok {
		tx.DposCycleBoundaryInfo = txdata
		return nil
	}
	return types.ErrTypeAsset
}

func (tx *DposCBRow) Get(key string) ([]byte, error) {
	if key == "cycle" {
		return []byte(fmt.Sprintf("%018d", tx.Cycle)), nil
	} else if key == "height" {
		return []byte(fmt.Sprintf("%018d", tx.StopHeight)), nil
	} else if key == "hash" {
		return tx.StopHash, nil
	}

	return nil, types.ErrNotFound
}
