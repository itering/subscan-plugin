package subscan_plugin

import (
	"github.com/itering/subscan-plugin/router"
	"github.com/itering/subscan-plugin/storage"
	"github.com/shopspring/decimal"
)

type Plugin interface {
	// Init storage interface
	InitDao(d storage.Dao)

	// Init http router
	InitHttp() []router.Http

	// Receive Extrinsic data when subscribe extrinsic dispatch
	ProcessExtrinsic(*storage.Block, *storage.Extrinsic, []storage.Event) error

	// Receive Extrinsic data when subscribe extrinsic dispatch
	ProcessEvent(*storage.Block, *storage.Event, decimal.Decimal) error

	// Mysql tables schema auto migrate
	Migrate()

	// Subscribe Extrinsic with special module
	SubscribeExtrinsic() []string

	// Subscribe Events with special module
	SubscribeEvent() []string

	// Plugins version
	Version() string
}
