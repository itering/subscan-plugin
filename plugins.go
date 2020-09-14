package subscan_plugin

import (
	"github.com/itering/subscan-plugin/router"
	"github.com/itering/subscan-plugin/storage"
	"github.com/shopspring/decimal"
)

type Plugin interface {
	// init storage interface
	InitDao(d storage.Dao)

	// init http router
	InitHttp() []router.Http

	// receive Extrinsic data when subscribe extrinsic dispatch
	ProcessExtrinsic(*storage.Block, *storage.Extrinsic, []storage.Event) error

	// receive Extrinsic data when subscribe extrinsic dispatch
	ProcessEvent(*storage.Block, *storage.Event, decimal.Decimal) error

	// mysql tables schema auto migrate
	Migrate()

	// Subscribe Extrinsic with special module
	SubscribeExtrinsic() []string

	// Subscribe Events with special module
	SubscribeEvent() []string

	// plugins version
	Version() string
}
