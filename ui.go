package subscan_plugin

type UiConfig struct {
	Type string `json:"type"`
	Body struct {
		Type string `json:"type"`
		Api  struct {
			Method         string `json:"method"`
			Url            string `json:"url"`
			RequestAdaptor string `json:"requestAdaptor"`
			Adaptor        string `json:"adaptor"`
		} `json:"api"`
		SyncLocation  bool     `json:"syncLocation"`
		HeaderToolbar []string `json:"headerToolbar"`
		Columns       []struct {
			Name  string `json:"name"`
			Label string `json:"label"`
		} `json:"columns"`
	} `json:"body"`
}

// https://baidu.gitee.io/amis/docs/concepts/schema
func (u *UiConfig) Init() {
	u.Type = "page"
	u.Body.Type = "crud"
	u.Body.SyncLocation = false
	u.Body.HeaderToolbar = []string{}
	u.Body.Api.RequestAdaptor = "return {...api, data: {...api.data, page: api.data.page - 1, row: api.data.perPage,} }"
	u.Body.Api.Adaptor = "return {...payload, status: payload.code, data: {items: payload.data.%s, count: payload.data.count}, msg: payload.message };"
}
