package subscan_plugin

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestUiConfig_Init(t *testing.T) {
	ui := new(UiConfig)
	ui.Init()
	ui.Body.Api.Method = "post"
	ui.Body.Api.Url = "api/accounts"
	ui.Body.Api.Adaptor = fmt.Sprintf(ui.Body.Api.Adaptor, "list")
	ui.Body.Columns = []UiColumns{{Name: "address", Label: "address"}}
	bu, _ := json.Marshal(ui)
	assert.Equal(t,
		`{"type":"page","body":{"type":"crud","api":{"method":"post","url":"api/accounts","requestAdaptor":"return {...api, data: {...api.data, page: api.data.page - 1, row: api.data.perPage,} }","adaptor":"return {...payload, status: payload.code, data: {items: payload.data.list, count: payload.data.count}, msg: payload.message };"},"syncLocation":false,"headerToolbar":[],"columns":[{"name":"address","label":"address"}]}}`,
		string(bu),
	)
}
