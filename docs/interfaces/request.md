# Interface: Request

The `req` object represents the HTTP request and has properties for the request query string, parameters, body, HTTP headers, and so on.

In this documentation and by convention, the object is always referred to as `req` (and the HTTP response is `res`) but its actual name is determined by the parameters to the callback function in which you’re working.

For example:

```js
mock.get("/user/{id}", function (req, res) {
  res.send("user " + req.params.id);
});
```

But you could just as well have:

```js
mock.get("/user/{id}", function (request, response) {
  response.send("user " + request.params.id);
});
```

The `req` object contains a `Request` field which is a golang's [http.Request](https://golang.org/pkg/net/http/#Request) instance.

## Table of contents

### Properties

- [body](request.md#body)
- [cookies](request.md#cookies)
- [get](request.md#get)
- [header](request.md#header)
- [method](request.md#method)
- [params](request.md#params)
- [path](request.md#path)
- [protocol](request.md#protocol)
- [query](request.md#query)

## Properties

### body

• **body**: *undefined* \| *Record*<string, any\>

Contains key-value pairs of data submitted in the request body.
By default, it is undefined, and is populated when the request
Content-Type is `application/json`.

___

### cookies

• **cookies**: *Record*<string, string\>

This property is an object that contains cookies sent by the request.

___

### get

• **get**: (`field`: *string*) => *string*

Returns the specified HTTP request header field (case-insensitive match).

**`param`** the header field name

**`returns`** the header field value.

#### Type declaration

▸ (`field`: *string*): *string*

#### Parameters

| Name | Type |
| :------ | :------ |
| `field` | *string* |

**Returns:** *string*

___

### header

• **header**: (`field`: *string*) => *string*

Returns the specified HTTP request header field (case-insensitive match).

**`param`** the header field name

**`returns`** the header field value.

#### Type declaration

▸ (`field`: *string*): *string*

#### Parameters

| Name | Type |
| :------ | :------ |
| `field` | *string* |

**Returns:** *string*

___

### method

• **method**: *string*

Contains a string corresponding to the HTTP method of the request: GET, POST, PUT, and so on.

___

### params

• **params**: *Record*<string, string\>

This property is an object containing properties mapped to the named route parameters.
For example, if you have the route /user/{name}, then the “name” property is available as req.params.name.
This object defaults to {}.

___

### path

• **path**: *string*

Contains the path part of the request URL.

___

### protocol

• **protocol**: *string*

Contains the request protocol string: either http or (for TLS requests) https.

___

### query

• **query**: *Record*<string, any\>

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
