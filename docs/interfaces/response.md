# Interface: Response

The `res` object represents the HTTP response that a server sends when it gets an HTTP request.

In this documentation and by convention, the object is always referred to as `res` (and the HTTP request is `req`) but its actual name is determined by the parameters to the callback function in which you’re working.

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

The `res` object contains a `Writer` field which is a golang's [http.ResponseWriter](https://golang.org/pkg/net/http/#ResponseWriter) instance.

## Table of contents

### Properties

- [append](response.md#append)
- [binary](response.md#binary)
- [html](response.md#html)
- [json](response.md#json)
- [redirect](response.md#redirect)
- [send](response.md#send)
- [set](response.md#set)
- [status](response.md#status)
- [text](response.md#text)
- [type](response.md#type)
- [vary](response.md#vary)

## Properties

### append

• **append**: (`field`: *string*, `value`: *string*) => [*Response*](response.md)

Appends the specified value to the HTTP response header field. If the header is not already set, it creates the header with the specified value.

**`param`** the header field name

**`param`** the value to append

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`field`: *string*, `value`: *string*): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `field` | *string* |
| `value` | *string* |

**Returns:** [*Response*](response.md)

___

### binary

• **binary**: (`body`: *string* \| *number*[] \| ArrayBuffer) => [*Response*](response.md)

Sends a binray response. This method sends a response (with the "application/octet-stream" content-type) that is the body paramter.

**`param`** the data to send

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`body`: *string* \| *number*[] \| ArrayBuffer): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `body` | *string* \| *number*[] \| ArrayBuffer |

**Returns:** [*Response*](response.md)

___

### html

• **html**: (`body`: *string*) => [*Response*](response.md)

Sends a HTML text response. This method sends a response (with the correct content-type) that is the body string paramter.

**`param`** the string to send

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`body`: *string*): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `body` | *string* |

**Returns:** [*Response*](response.md)

___

### json

• **json**: (`body`: *Record*<string, any\>) => [*Response*](response.md)

Sends a JSON response. This method sends a response (with the correct content-type) that is the parameter converted to a JSON string.

**`param`** the object to send

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`body`: *Record*<string, any\>): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `body` | *Record*<string, any\> |

**Returns:** [*Response*](response.md)

___

### redirect

• **redirect**: (`code`: *number*, `loc`: *string*) => [*Response*](response.md)

Redirects to the URL, with specified status, a positive integer that corresponds to an HTTP status code.

**`param`** the HTTP status code (301, 302, ...)

**`param`** the location to redirect

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`code`: *number*, `loc`: *string*): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `code` | *number* |
| `loc` | *string* |

**Returns:** [*Response*](response.md)

___

### send

• **send**: (`body`: *string* \| *number*[] \| ArrayBuffer) => [*Response*](response.md)

Sends the HTTP response.

When the parameter is a ArrayBuffer or number[], the method sets the Content-Type response header field to “application/octet-stream”.
When the parameter is a String, the method sets the Content-Type to “text/html”.
Otherwise the method sets the Content-Type to "application/json" and convert paramter to JSON representation before sending.

**`param`** the data to send

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`body`: *string* \| *number*[] \| ArrayBuffer): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `body` | *string* \| *number*[] \| ArrayBuffer |

**Returns:** [*Response*](response.md)

___

### set

• **set**: (`field`: *string*, `value`: *string*) => [*Response*](response.md)

Sets the response’s HTTP header field to value.

**`param`** the header field name

**`param`** the value to set

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`field`: *string*, `value`: *string*): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `field` | *string* |
| `value` | *string* |

**Returns:** [*Response*](response.md)

___

### status

• **status**: (`code`: *number*) => [*Response*](response.md)

Sets the HTTP status for the response.

**`param`** the satus code value

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`code`: *number*): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `code` | *number* |

**Returns:** [*Response*](response.md)

___

### text

• **text**: (`format`: *string*, `v?`: *any*[]) => [*Response*](response.md)

Sends a plain text response. This method sends a response (with the correct content-type) that is the string formatting result.

**`param`** go format string

**`param`** format values (if any)

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`format`: *string*, `v?`: *any*[]): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `format` | *string* |
| `v?` | *any*[] |

**Returns:** [*Response*](response.md)

___

### type

• **type**: (`mime`: *string*) => [*Response*](response.md)

Sets the Content-Type HTTP header to the MIME type as from mime parameter.

**`params`** mime the content type

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`mime`: *string*): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `mime` | *string* |

**Returns:** [*Response*](response.md)

___

### vary

• **vary**: (`header`: *string*) => [*Response*](response.md)

Adds the header field to the Vary response header.

**`param`** the header filed name

**`returns`** The instance for fluent/chaining API

#### Type declaration

▸ (`header`: *string*): [*Response*](response.md)

#### Parameters

| Name | Type |
| :------ | :------ |
| `header` | *string* |

**Returns:** [*Response*](response.md)
