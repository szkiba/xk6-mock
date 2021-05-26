# xk6-mock

A [k6](https://go.k6.io/k6) extension for mocking HTTP(s) servers during test development.
The design of the library was inspired by [Express](https://expressjs.com/). If you already know Express framework, using this library should be very simple.

## Features

- Start mock HTTP server inside of a k6 process
- Familiar, Express like mock route definitions
- Almost transparent for test scripts: change import from `k6/http` to `k6/x/mock/http`

## http module

The `k6/x/mock/http` module is a drop-in replacement for `k6/http` module. During the k6 test development simply replace `k6/http` imports with `k6/x/mock/http`. Each URL's host and port part will be automatically replaced with the mock server's host and port value. Other parts of the request are remain untouched.

## mock module

The `k6/x/mock` module's default export is an Express like default mock server with the usual HTTP method functions (get, post, ..) and use function for defining middlewares.

You can add route definitions both in [Init and VU stages](https://k6.io/docs/using-k6/test-life-cycle/#init-and-vu-stages).

## Table of contents

### Classes

- [Module](classes/module.md)
- [Server](classes/server.md)

### Interfaces

- [Request](interfaces/request.md)
- [Response](interfaces/response.md)

### Type aliases

- [CallbackFunc](README.md#callbackfunc)
- [NextFunc](README.md#nextfunc)

### Variables

- [default](README.md#default)

## Type aliases

### CallbackFunc

Ƭ **CallbackFunc**: (`req`: [*Request*](interfaces/request.md), `res`: [*Response*](interfaces/response.md), `next`: [*NextFunc*](README.md#nextfunc)) => *void*

CallbackFunc defines middleware and request handler callback function.

**`param`** the request object

**`param`** the response object

**`param`** indicating the next middleware function

#### Type declaration

▸ (`req`: [*Request*](interfaces/request.md), `res`: [*Response*](interfaces/response.md), `next`: [*NextFunc*](README.md#nextfunc)): *void*

#### Parameters

| Name | Type |
| :------ | :------ |
| `req` | [*Request*](interfaces/request.md) |
| `res` | [*Response*](interfaces/response.md) |
| `next` | [*NextFunc*](README.md#nextfunc) |

**Returns:** *void*

___

### NextFunc

Ƭ **NextFunc**: () => *void*

NextFunc defines callback function's `next` parameter function.
Calling from middleware enables processing next middleware.

#### Type declaration

▸ (): *void*

**Returns:** *void*

## Variables

### default

• `Const` **default**: [*Module*](classes/module.md)

default export
