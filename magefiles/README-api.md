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
