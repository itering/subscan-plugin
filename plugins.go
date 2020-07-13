package subscan_plugin

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/itering/subscan-plugin/storage"
	"github.com/shopspring/decimal"
)

type Plugin interface {
	InitDao(d storage.Dao)

	InitHttp(e *bm.Engine)

	ProcessExtrinsic(*storage.Block, *storage.Extrinsic, []storage.Event) error

	ProcessEvent(*storage.Block, *storage.Event, decimal.Decimal) error

	Migrate()
}
