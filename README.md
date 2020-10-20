# Subscan-Plugin


The subscan plugin is an interface lib to [subscan-essentials](https://github.com/itering/subscan-essentials) plugin, and [Gen](https://github.com/itering/subscan-plugin/tree/master/tools) tool can automatically generate plugin templates

## Usage

### gen plugin template

```
cd tools/gen-plugin && go build -o subscan-plugin
./subscan-plugin staking
```

### Example

```
// plugin balance ui config in json format
{
  "type": "page",
  "body": {
    // component type, 'crud' is a basic table view component
    "type": "crud", 
    "api": {
      "method": "post",
      "url": "api/plugin/balance/accounts",
      //requestAdaptor and adaptor is amis API config to customize request and response
      "requestAdaptor": "return {...api, data: {...api.data, page: api.data.page - 1, row: api.data.perPage,} }",
      "adaptor": "return {...payload, status: payload.code, data: {items: payload.data.list, count: payload.data.count}, msg: payload.message };"
    },
    "syncLocation": false,
    "headerToolbar": [],
    "columns": [
      {
        "name": "address",
        "label": "address"
      },
      {
        "name": "nonce",
        "label": "nonce"
      },
      {
        "name": "balance",
        "label": "balance"
      },
      {
        "name": "lock",
        "label": "lock"
      }
    ]
  }
}

// plugin balance UiConf
func (a *Balance) UiConf() *plugin.UiConfig {
	conf := new(plugin.UiConfig)
	conf.Init()
	conf.Body.Api.Method = "post"
	conf.Body.Api.Url = "api/plugin/balance/accounts"
	conf.Body.Api.Adaptor = fmt.Sprintf(conf.Body.Api.Adaptor, "list")
	conf.Body.Columns = []plugin.UiColumns{
		{Name: "address", Label: "address"},
		{Name: "nonce", Label: "nonce"},
		{Name: "balance", Label: "balance"},
		{Name: "lock", Label: "lock"},
	}
	return conf
}
```
refer to [amis docs](https://baidu.gitee.io/amis/docs/index) for further config detail.


## Install

```
go get github.com/itering/subscan-plugin
```

## LICENSE

GPL-3.0

## Resource
 
- [subscan-essentials] https://github.com/itering/subscan-essentials
