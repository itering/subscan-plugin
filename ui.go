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
	u.Body.Api.RequestAdaptor = "return {\n ...api,\n data: {\n ...api.data, \n page: api.data.page - 1, \n row: api.data.perPage,}\n }"
	u.Body.Api.Adaptor = "return {\n...payload,\n status: payload.code,\n data: {\n items: payload.data.%s,\n count: payload.data.count \n },\n msg: payload.message\n };"
}
