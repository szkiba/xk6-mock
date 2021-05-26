# Class: Module

A module object represents module default object.

## Table of contents

### Constructors

- [constructor](module.md#constructor)

### Methods

- [all](module.md#all)
- [delete](module.md#delete)
- [get](module.md#get)
- [head](module.md#head)
- [options](module.md#options)
- [patch](module.md#patch)
- [post](module.md#post)
- [put](module.md#put)
- [resolve](module.md#resolve)
- [use](module.md#use)

## Constructors

### constructor

\+ **new Module**(): [*Module*](module.md)

**Returns:** [*Module*](module.md)

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

### resolve

▸ **resolve**(`url`: *string*): *string*

Resolve URL to mock URL, replace host and port with mock server's host and por value.

Only external URLs are resolved, URLs are already pointing to mock server will remain untouched.

#### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `url` | *string* | original URL string |

**Returns:** *string*

resolved URL, points to mock server

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
