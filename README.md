# Subscan-Plugin


The subscan plugin is an interface lib to [subscan-essentials](https://github.com/itering/subscan-essentials) plugin, and [Gen](https://github.com/itering/subscan-plugin/tree/master/tools) tool can automatically generate plugin templates

## How it works

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

## Usage

### gen plugin template

```
cd tools/gen-plugin && go build -o subscan-plugin
./subscan-plugin staking
```

## Install

```
go get github.com/itering/subscan-plugin
```

## LICENSE

GPL-3.0

## Resource
 
- [subscan-essentials] https://github.com/itering/subscan-essentials
