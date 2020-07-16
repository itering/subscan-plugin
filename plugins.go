package subscan_plugin

import (
	"github.com/itering/subscan-plugin/router"
	"github.com/itering/subscan-plugin/storage"
	"github.com/shopspring/decimal"
)

type Plugin interface {
	InitDao(d storage.Dao)

	InitHttp() []router.Http

	ProcessExtrinsic(*storage.Block, *storage.Extrinsic, []storage.Event) error

	ProcessEvent(*storage.Block, *storage.Event, decimal.Decimal) error

	Migrate()

	Version() string
}
