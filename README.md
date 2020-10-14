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



## Install

```
go get github.com/itering/subscan-plugin
```

## Resource
 
- [subscan-essentials] https://github.com/itering/subscan-essentials
