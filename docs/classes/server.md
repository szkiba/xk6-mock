# Class: Server

A server object represents a web server.

## Table of contents

### Constructors

- [constructor](server.md#constructor)

### Methods

- [all](server.md#all)
- [delete](server.md#delete)
- [get](server.md#get)
- [head](server.md#head)
- [options](server.md#options)
- [patch](server.md#patch)
- [post](server.md#post)
- [put](server.md#put)
- [start](server.md#start)
- [use](server.md#use)

## Constructors

### constructor

\+ **new Server**(): [*Server*](server.md)

Creates a new server instance.

**Returns:** [*Server*](server.md)

## Methods

### all

▸ **all**(`path`: *string*, ...`callback`: [*CallbackFunc*](../README.md#callbackfunc)[]): [*Server*](server.md)

This method is just like the get,head,post,... methods, except that it matches all HTTP methods (verbs).

You can provide multiple callback functions that behave just like middleware.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | *string* | The path for which the callback function is invoked (string or path pattern) |
| `...callback` | [*CallbackFunc*](../README.md#callbackfunc)[] | Callback functions (middlewares or request handler) |

**Returns:** [*Server*](server.md)

The instance for fluent/chaining API

___

### delete

▸ **delete**(`path`: *string*, ...`callback`: [*CallbackFunc*](../README.md#callbackfunc)[]): [*Server*](server.md)

Routes HTTP `DELETE` requests to the specified path with the specified callback functions.

You can provide multiple callback functions that behave just like middleware.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | *string* | The path for which the callback function is invoked (string or path pattern) |
| `...callback` | [*CallbackFunc*](../README.md#callbackfunc)[] | Callback functions (middlewares or request handler) |

**Returns:** [*Server*](server.md)

The instance for fluent/chaining API

___

### get

▸ **get**(`path`: *string*, ...`callback`: [*CallbackFunc*](../README.md#callbackfunc)[]): [*Server*](server.md)

Routes HTTP GET requests to the specified path with the specified callback functions.

You can provide multiple callback functions that behave just like middleware.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | *string* | The path for which the callback function is invoked (string or path pattern) |
| `...callback` | [*CallbackFunc*](../README.md#callbackfunc)[] | Callback functions (middlewares or request handler) |

**Returns:** [*Server*](server.md)

The instance for fluent/chaining API

___

### head

▸ **head**(`path`: *string*, ...`callback`: [*CallbackFunc*](../README.md#callbackfunc)[]): [*Server*](server.md)

Routes HTTP HEAD requests to the specified path with the specified callback functions.

You can provide multiple callback functions that behave just like middleware.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | *string* | The path for which the callback function is invoked (string or path pattern) |
| `...callback` | [*CallbackFunc*](../README.md#callbackfunc)[] | Callback functions (middlewares or request handler) |

**Returns:** [*Server*](server.md)

The instance for fluent/chaining API

___

### options

▸ **options**(`path`: *string*, ...`callback`: [*CallbackFunc*](../README.md#callbackfunc)[]): [*Server*](server.md)

Routes HTTP OPTIONS requests to the specified path with the specified callback functions.

You can provide multiple callback functions that behave just like middleware.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | *string* | The path for which the callback function is invoked (string or path pattern) |
| `...callback` | [*CallbackFunc*](../README.md#callbackfunc)[] | Callback functions (middlewares or request handler) |

**Returns:** [*Server*](server.md)

The instance for fluent/chaining API

___

### patch

▸ **patch**(`path`: *string*, ...`callback`: [*CallbackFunc*](../README.md#callbackfunc)[]): [*Server*](server.md)

Routes HTTP PATCH requests to the specified path with the specified callback functions.

You can provide multiple callback functions that behave just like middleware.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | *string* | The path for which the callback function is invoked (string or path pattern) |
| `...callback` | [*CallbackFunc*](../README.md#callbackfunc)[] | Callback functions (middlewares or request handler) |

**Returns:** [*Server*](server.md)

The instance for fluent/chaining API

___

### post

▸ **post**(`path`: *string*, ...`callback`: [*CallbackFunc*](../README.md#callbackfunc)[]): [*Server*](server.md)

Routes HTTP POST requests to the specified path with the specified callback functions.

You can provide multiple callback functions that behave just like middleware.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | *string* | The path for which the callback function is invoked (string or path pattern) |
| `...callback` | [*CallbackFunc*](../README.md#callbackfunc)[] | Callback functions (middlewares or request handler) |

**Returns:** [*Server*](server.md)

The instance for fluent/chaining API

___

### put

▸ **put**(`path`: *string*, ...`callback`: [*CallbackFunc*](../README.md#callbackfunc)[]): [*Server*](server.md)

Routes HTTP PUT requests to the specified path with the specified callback functions.

You can provide multiple callback functions that behave just like middleware.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | *string* | The path for which the callback function is invoked (string or path pattern) |
| `...callback` | [*CallbackFunc*](../README.md#callbackfunc)[] | Callback functions (middlewares or request handler) |

**Returns:** [*Server*](server.md)

The instance for fluent/chaining API

___

### start

▸ **start**(`port?`: *number*, `host?`: *string*): [*Server*](server.md)

Starts the server.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `port?` | *number* | TCP port number, if 0 or missing then random unused port will be allocated |
| `host?` | *string* | host name or IP address for listening on, default 127.0.0.1 |

**Returns:** [*Server*](server.md)

The instance for fluent/chaining API

___

### use

▸ **use**(`path`: *string*, ...`callback`: [*CallbackFunc*](../README.md#callbackfunc)[]): [*Server*](server.md)

Uses the specified middleware function or functions.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | *string* | The path for which the callback function is invoked (string or path pattern) |
| `...callback` | [*CallbackFunc*](../README.md#callbackfunc)[] | Callback functions (middlewares or request handler) |

**Returns:** [*Server*](server.md)

The instance for fluent/chaining API
