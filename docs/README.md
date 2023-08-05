[![Go Report Card](https://goreportcard.com/badge/github.com/szkiba/xk6-mock)](https://goreportcard.com/report/github.com/szkiba/xk6-mock)
[![GitHub Actions](https://github.com/szkiba/xk6-mock/workflows/Test/badge.svg)](https://github.com/szkiba/xk6-mock/actions?query=workflow%3ATest+branch%3Amaster)
[![codecov](https://codecov.io/gh/szkiba/xk6-mock/branch/master/graph/badge.svg?token=GK1JNCPH8U)](https://codecov.io/gh/szkiba/xk6-mock)
[![Documentation](https://img.shields.io/badge/docs-reference-blue?logo=readme&logoColor=lightgray)](https://ivan.szkiba.hu/xk6-mock)

# xk6-mock API

A [k6](https://go.k6.io/k6) extension enables mocking HTTP(S) servers during test development.
The design of the extension was inspired by [Express.js](https://expressjs.com/).

If you have already known Express.js framework, using this extension should be very simple.

```js
import http, { mock } from 'k6/x/mock'

mock('https://example.com', app => {
  app.get('/', (req, res) => {
    res.json({ greeting: 'Hello World!' })
  })
})

export default async function () {
  const res = await http.asyncRequest('GET', 'https://example.com')
  console.log(res.json())
}
```

## Features

- Starts mock HTTP server(s) inside the k6 process
- Familiar, Express like mock route definitions
- Almost transparent for test scripts: just change import statement from `k6/http` to `k6/x/mock`
- Helps testing k6 tests with mock server
- Supports sync and async `k6/http` API

<a name="readmemd"></a>

## Synchronous mode

Mixing synchronous and asynchronous calls can block k6 test execution. Both synchronous (`http.get`, `http.post`, ...) and asynchronous (`http.asyncRequest`) API mocking is supported. By default the new asynchronous mode will be used.

To switch to synchronous mode you can pass an options object to `mock` function with `sync` property set to `true`:

```js
mock({ sync:true }, "https://example.com", app => {

})
```

The `Application` class constructor also accepts similar options parameter:

```js
const app = new Application({ sync: true })
```

## `k6/http`

The `k6/x/mock` module contains a thin wrapper around the `k6/http` module. This allow to use `k6/x/mock` module as drop-in replacement for `k6/http` for seamless mocking.

```JavaScript
import http, { mock } from "k6/x/mock"
```

## Usage tips

 1. Create separated `mock.js` module for mocking

 2. Re-export `k6/x/mock` module content

    ```js
    export { default } from "k6/x/mock"
    export * from "k6/x/mock"
    ```

  3. Put mock definitions in `mock.js`

     ```js
     import { mock } from "k6/x/mock"

     mock('https://httpbin.test.k6.io/get', app => {
       app.get('/', (req, res) => {
         res.json({ url: 'https://httpbin.test.k6.io/get' })
       })
     })
     ```

 3. In test script, import `http` from `mock.js` instead of `k6/http`
    ```js
    import http from "./mock.js";
    ```

    Switching from mock to real implementation is as easy as replacing the line above with real `k6/http` module import 

    ```js
    import http from "k6/http"
    ```

 4. The other part of the test script is independent from mocking

    ```js
    import http from "./mock.js";
    import { check } from 'k6'
    import { test } from 'k6/execution'

    export default async function () {
      const res = await http.asyncRequest('GET', 'https://httpbin.test.k6.io/get')
      const ok = check(res, {
        'response code was 200': res => res.status == 200,
        '"url" was "https://httpbin.test.k6.io/get"': res =>
          res.json('url') == 'https://httpbin.test.k6.io/get'
      })

      if (!ok) {
        test.abort('unexpected response')
      }
    }
    ```

## Disabling mock

You can disable the given mock definition quickly by passing options parameter with `skip` set to true.
```JavaScript
mock(
  'https://example.com',
  app => {
    app.get('/', (req, res) => {
      res.json({ greeting: 'Hello World!' })
    })
  },
  { skip: true }
)
```

Alternatively you can put `.skip` after function name:
```JavaScript
mock.skip('https://example.com', app => {
  app.get('/', (req, res) => {
    res.json({ greeting: 'Hello World!' })
  })
})
```

---

The API documentation bellow was generated from [index.d.ts](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts) file.


<a name="classesapplicationmd"></a>

## Class: Application

An application object represents a web application.

The following example starts a server and listens for connections on port 3000.
The application responds with a JSON object for requests to the root URL.
All other routes are answered with a 404 not found message.

In this example, the name of the constructor is `Application`, but the name you use is up to you.

**`Example`**

```ts
const app = new Application()

app.get('/', (req, res) => {
  res.json({message:"Hello World!"})
})

app.listen(3000)
```

### Constructors

#### constructor

• **new Application**()

Creates a new application instance.

##### Defined in

[index.d.ts:111](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L111)

### Methods

#### delete

▸ **delete**(`path`, `...middleware`): `void`

Routes HTTP `DELETE` requests to the specified path with the specified middleware functions.

You can provide multiple middleware functions.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | `string` | The path for which the middleware function is invoked (string or path pattern) |
| `...middleware` | [`Middleware`](#middleware)[] | Middleware functions |

##### Returns

`void`

##### Defined in

[index.d.ts:171](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L171)

___

#### get

▸ **get**(`path`, `...middleware`): `void`

Routes HTTP GET requests to the specified path with the specified middleware functions.

You can provide multiple middleware functions.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | `string` | The path for which the middleware function is invoked (string or path pattern) |
| `...middleware` | [`Middleware`](#middleware)[] | Middleware functions |

##### Returns

`void`

##### Defined in

[index.d.ts:121](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L121)

___

#### head

▸ **head**(`path`, `...middleware`): `void`

Routes HTTP HEAD requests to the specified path with the specified middleware functions.

You can provide multiple middleware functions.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | `string` | The path for which the middleware function is invoked (string or path pattern) |
| `...middleware` | [`Middleware`](#middleware)[] | Middleware functions |

##### Returns

`void`

##### Defined in

[index.d.ts:131](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L131)

___

#### listen

▸ **listen**(`addr?`, `callback?`): `void`

Starts the server.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `addr?` | `string` | - |
| `callback?` | () => `void` | host name or IP address for listening on, default 127.0.0.1 |

##### Returns

`void`

The instance for fluent/chaining API

##### Defined in

[index.d.ts:206](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L206)

___

#### options

▸ **options**(`path`, `...middleware`): `void`

Routes HTTP OPTIONS requests to the specified path with the specified middleware functions.

You can provide multiple middleware functions.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | `string` | The path for which the middleware function is invoked (string or path pattern) |
| `...middleware` | [`Middleware`](#middleware)[] | Middleware functions |

##### Returns

`void`

##### Defined in

[index.d.ts:181](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L181)

___

#### patch

▸ **patch**(`path`, `...middleware`): `void`

Routes HTTP PATCH requests to the specified path with the specified middleware functions.

You can provide multiple middleware functions.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | `string` | The path for which the middleware function is invoked (string or path pattern) |
| `...middleware` | [`Middleware`](#middleware)[] | Middleware functions |

##### Returns

`void`

##### Defined in

[index.d.ts:161](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L161)

___

#### post

▸ **post**(`path`, `...middleware`): `void`

Routes HTTP POST requests to the specified path with the specified middleware functions.

You can provide multiple middleware functions.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | `string` | The path for which the middleware function is invoked (string or path pattern) |
| `...middleware` | [`Middleware`](#middleware)[] | Middleware functions |

##### Returns

`void`

##### Defined in

[index.d.ts:141](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L141)

___

#### put

▸ **put**(`path`, `...middleware`): `void`

Routes HTTP PUT requests to the specified path with the specified middleware functions.

You can provide multiple middleware functions.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | `string` | The path for which the middleware function is invoked (string or path pattern) |
| `...middleware` | [`Middleware`](#middleware)[] | Middleware functions |

##### Returns

`void`

##### Defined in

[index.d.ts:151](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L151)

___

#### static

▸ **static**(`path`, `docroot`): `void`

Mount static web content from given source directory.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | `string` | The path where the source will be mounted on |
| `docroot` | `string` | The source directory path |

##### Returns

`void`

##### Defined in

[index.d.ts:197](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L197)

___

#### use

▸ **use**(`path`, `...middleware`): `void`

Uses the specified middleware function or functions.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `path` | `string` | The path for which the middleware function is invoked (string or path pattern) |
| `...middleware` | [`Middleware`](#middleware)[] | Middleware functions |

##### Returns

`void`

##### Defined in

[index.d.ts:189](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L189)


<a name="interfacesmockoptionsmd"></a>

## Interface: MockOptions

Optional flags for `mock` function.

Passing this options object as first (or last) parameter to `mock` function (or to `Application` constructor) may
change its behavior.

**`Example`**

```ts
mock("https://example.com", callback, { sync:true });
```

### Properties

#### skip

• **skip**: `boolean`

True value indicaes that given mock definition should be ignored.

##### Defined in

[index.d.ts:29](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L29)

___

#### sync

• **sync**: `boolean`

True value indicates synchronous mode operation. You should use it for synchronous k6 http API.

##### Defined in

[index.d.ts:24](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L24)


<a name="interfacesrequestmd"></a>

## Interface: Request

The `req` object represents the HTTP request and has properties for the request query string, parameters, body, HTTP headers, and so on.

In this documentation and by convention, the object is always referred to as `req` (and the HTTP response is `res`) but its actual name is determined by the parameters to the callback function in which you’re working.

**`Example`**

```ts
app.get("/user/:id", function (req, res) {
  res.send("user " + req.params.id);
});
```

### Properties

#### body

• **body**: `Record`<`string`, `any`\>

Contains key-value pairs of data submitted in the request body.
By default, it is undefined, and is populated when the request
Content-Type is `application/json`.

##### Defined in

[index.d.ts:226](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L226)

___

#### cookies

• **cookies**: `Record`<`string`, `string`\>

This property is an object that contains cookies sent by the request.

##### Defined in

[index.d.ts:231](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L231)

___

#### get

• **get**: (`field`: `string`) => `string`

##### Type declaration

▸ (`field`): `string`

Returns the specified HTTP request header field (case-insensitive match).

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `field` | `string` | the header field name |

###### Returns

`string`

the header field value.

##### Defined in

[index.d.ts:278](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L278)

___

#### header

• **header**: (`field`: `string`) => `string`

##### Type declaration

▸ (`field`): `string`

Returns the specified HTTP request header field (case-insensitive match).

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `field` | `string` | the header field name |

###### Returns

`string`

the header field value.

##### Defined in

[index.d.ts:286](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L286)

___

#### method

• **method**: `string`

Contains a string corresponding to the HTTP method of the request: GET, POST, PUT, and so on.

##### Defined in

[index.d.ts:236](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L236)

___

#### params

• **params**: `Record`<`string`, `string`\>

This property is an object containing properties mapped to the named route parameters.
For example, if you have the route /user/:name, then the “name” property is available as req.params.name.
This object defaults to empty.

##### Defined in

[index.d.ts:243](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L243)

___

#### path

• **path**: `string`

Contains the path part of the request URL.

##### Defined in

[index.d.ts:248](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L248)

___

#### protocol

• **protocol**: `string`

Contains the request protocol string: either http or (for TLS requests) https.

##### Defined in

[index.d.ts:253](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L253)

___

#### query

• **query**: `Record`<`string`, `any`\>

This property is an object containing a property for each query string parameter in the route.

For example:

```js
// GET /search?q=tobi+ferret
console.dir(req.query.q);
// => 'tobi ferret'

// GET /shoes?color=blue&color=black&color=red
console.dir(req.query.color);
// => ['blue', 'black', 'red']
```

##### Defined in

[index.d.ts:270](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L270)


<a name="interfacesresponsemd"></a>

## Interface: Response

The `res` object represents the HTTP response that a server sends when it gets an HTTP request.

In this documentation and by convention, the object is always referred to as `res` (and the HTTP request is `req`) but its actual name is determined by the parameters to the callback function in which you’re working.

**`Example`**

```ts
app.get("/user/:id", function (req, res) {
  res.send("user " + req.params.id);
});
```

### Properties

#### append

• **append**: (`field`: `string`, `value`: `string`) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`field`, `value`): [`Response`](#interfacesresponsemd)

Appends the specified value to the HTTP response header field. If the header is not already set, it creates the header with the specified value.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `field` | `string` | the header field name |
| `value` | `string` | the value to append |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:306](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L306)

___

#### binary

• **binary**: (`body`: `string` \| `number`[] \| `ArrayBuffer`) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`body`): [`Response`](#interfacesresponsemd)

Sends a binray response. This method sends a response (with the "application/octet-stream" content-type) that is the body paramter.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `body` | `string` \| `number`[] \| `ArrayBuffer` | the data to send |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:334](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L334)

___

#### html

• **html**: (`body`: `string`) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`body`): [`Response`](#interfacesresponsemd)

Sends a HTML text response. This method sends a response (with the correct content-type) that is the body string paramter.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `body` | `string` | the string to send |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:327](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L327)

___

#### json

• **json**: (`body`: `Record`<`string`, `any`\>) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`body`): [`Response`](#interfacesresponsemd)

Sends a JSON response. This method sends a response (with the correct content-type) that is the parameter converted to a JSON string.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `body` | `Record`<`string`, `any`\> | the object to send |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:313](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L313)

___

#### redirect

• **redirect**: (`code`: `number`, `loc`: `string`) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`code`, `loc`): [`Response`](#interfacesresponsemd)

Redirects to the URL, with specified status, a positive integer that corresponds to an HTTP status code.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `code` | `number` | the HTTP status code (301, 302, ...) |
| `loc` | `string` | the location to redirect |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:382](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L382)

___

#### send

• **send**: (`body`: `string` \| `number`[] \| `ArrayBuffer`) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`body`): [`Response`](#interfacesresponsemd)

Sends the HTTP response.

When the parameter is a ArrayBuffer or number[], the method sets the Content-Type response header field to “application/octet-stream”.
When the parameter is a String, the method sets the Content-Type to “text/html”.
Otherwise the method sets the Content-Type to "application/json" and convert paramter to JSON representation before sending.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `body` | `string` \| `number`[] \| `ArrayBuffer` | the data to send |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:345](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L345)

___

#### set

• **set**: (`field`: `string`, `value`: `string`) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`field`, `value`): [`Response`](#interfacesresponsemd)

Sets the response’s HTTP header field to value.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `field` | `string` | the header field name |
| `value` | `string` | the value to set |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:374](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L374)

___

#### status

• **status**: (`code`: `number`) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`code`): [`Response`](#interfacesresponsemd)

Sets the HTTP status for the response.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `code` | `number` | the satus code value |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:352](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L352)

___

#### text

• **text**: (`format`: `string`, `v?`: `any`[]) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`format`, `v?`): [`Response`](#interfacesresponsemd)

Sends a plain text response. This method sends a response (with the correct content-type) that is the string formatting result.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `format` | `string` | go format string |
| `v?` | `any`[] | format values (if any) |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:321](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L321)

___

#### type

• **type**: (`mime`: `string`) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`mime`): [`Response`](#interfacesresponsemd)

Sets the Content-Type HTTP header to the MIME type as from mime parameter.

**`Params`**

mime the content type

###### Parameters

| Name | Type |
| :------ | :------ |
| `mime` | `string` |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:359](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L359)

___

#### vary

• **vary**: (`header`: `string`) => [`Response`](#interfacesresponsemd)

##### Type declaration

▸ (`header`): [`Response`](#interfacesresponsemd)

Adds the header field to the Vary response header.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `header` | `string` | the header filed name |

###### Returns

[`Response`](#interfacesresponsemd)

##### Defined in

[index.d.ts:366](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L366)


<a name="modulesmd"></a>

## Auxiliary

### Classes

- [Application](#classesapplicationmd)

### Interfaces

- [MockOptions](#interfacesmockoptionsmd)
- [Request](#interfacesrequestmd)
- [Response](#interfacesresponsemd)

### Type Aliases

#### Middleware

Ƭ **Middleware**: (`req`: [`Request`](#interfacesrequestmd), `res`: [`Response`](#interfacesresponsemd), `next`: () => `void`) => `void`

##### Type declaration

▸ (`req`, `res`, `next`): `void`

Middleware defines middleware and request handler callback function.

###### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `req` | [`Request`](#interfacesrequestmd) | the request object |
| `res` | [`Response`](#interfacesresponsemd) | the response object |
| `next` | () => `void` | calling from middleware enables processing next middleware |

###### Returns

`void`

##### Defined in

[index.d.ts:87](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L87)

### Functions

#### mock

▸ **mock**(`target`, `callback`, `options?`): `void`

Create URL mock definition.

This function will create a new Express.js like web application and pass it to the provided
callback function for programming. After mock programming is done, new mock HTTP server will be start.
Whan you use http API from mock module, all matching URLs will be directed to this mock server.

You can create as many mock definitions (server) as you want.

You can disable the given mock definition quickly by passing options parameter with `skip` set to true.
```JavaScript
mock(
  'https://example.com',
  app => {
    app.get('/', (req, res) => {
      res.json({ greeting: 'Hello World!' })
    })
  },
  { skip: true }
)
```

Alternatively you can put `.skip` after function name:
```JavaScript
mock.skip('https://example.com', app => {
 app.get('/', (req, res) => {
   res.json({ greeting: 'Hello World!' })
 })
})
```

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `target` | `String` | the URL or URL prefix to be mocked |
| `callback` | (`app`: [`Application`](#classesapplicationmd)) => `void` | function to for defining route definitions for mock server |
| `options?` | [`MockOptions`](#interfacesmockoptionsmd) | optional flags (`sync`, `skip`) |

##### Returns

`void`

##### Defined in

[index.d.ts:67](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L67)

___

#### unmock

▸ **unmock**(`target`): `void`

Deactivate URL mocking.

This function will remove mock definition associated to given URL and stop the related HTTP server.

##### Parameters

| Name | Type | Description |
| :------ | :------ | :------ |
| `target` | `String` | the URL or URL prefix of mock definition to be remove |

##### Returns

`void`

##### Defined in

[index.d.ts:76](https://github.com/szkiba/xk6-mock/blob/master/api/index.d.ts#L76)