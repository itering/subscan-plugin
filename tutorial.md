## tutorial

## description

The subscan plugin gives more possibilities to subscan-essentials. Subscan sends the Extrinsic or Event received by the fetch block to the plugin, which can be customized and processed into formatted data, stored in the database, and presented through http api.
You can write ui conf directly to automatically display the ui.

## Quick start

### Install gen tool

```shell script
$ go get github.com/itering/subscan-plugin/tools/gen-plugin
```

### Gen plugin

```shell script
# assume need generate balance module
$ gen-plugin balance
```

The directory structure is as follows:
```
├── dao
│   └── dao.go
├── http
│   └── http.go
├── main.go
├── model
│   └── model.go
└── service
    └── service.go

4 directories, 5 files
```

The logic can refer to https://github.com/itering/subscan-essentials/tree/master/plugins/balance

### Register plugin

```go
func init() {
	registerNative(balance.New())
	registerNative(system.New())
}
```

### UI
The UI part is rely on [amis](https://github.com/baidu/amis/blob/master/README-en.md), a Low-Code frontend UI Framework, which allows us to develop various page via different components by only using JSON configuration.

First, plugin list is fetched via /api/scan/plugins. After that, user can click different plugin in navbar and will be redirected to the plugin page. When the page is mounted, we init amis using the plugin info in JSON format fetched via /api/scan/plugins/ui to generate corresponding view.

An example page using basic table view component is as follow.
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
```

refer to [amis docs](https://baidu.gitee.io/amis/docs/index) for further config detail.

