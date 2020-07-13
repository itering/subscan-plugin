package main

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/itering/subscan-plugin/storage"
	"github.com/itering/subscan-plugin/system/model"
	"github.com/itering/subscan-plugin/system/service"
	"github.com/itering/subscan-plugin/tools"
	"github.com/shopspring/decimal"
)

var srv *service.Service

type System struct {
	d storage.Dao
	e *bm.Engine
}

func New() *System {
	return &System{}
}

func (a *System) InitDao(d storage.Dao) {
	srv = service.New(d)
	a.d = d
	a.Migrate()
}
func (a *System) InitHttp(e *bm.Engine) {
	a.e = e
}

func (a *System) ProcessExtrinsic(block *storage.Block, extrinsic *storage.Extrinsic, events []storage.Event) error {
	return nil
}

func (a *System) ProcessEvent(block *storage.Block, event *storage.Event, fee decimal.Decimal) error {
	var paramEvent []storage.EventParam
	tools.UnmarshalToAnything(&paramEvent, event.Params)
	switch event.EventId {
	case "ExtrinsicFailed":
		srv.ExtrinsicFailed(block.SpecVersion, block.BlockTimestamp, block.Hash, event, paramEvent)
	}
	return nil
}

func (a *System) Migrate() {
	db := a.d.DB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.ExtrinsicError{},
	)
	db.Model(model.ExtrinsicError{}).AddUniqueIndex("extrinsic_hash", "extrinsic_hash")
}
