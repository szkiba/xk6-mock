
# Synchronous mode

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

# `k6/http`

The `k6/x/mock` module contains a thin wrapper around the `k6/http` module. This allow to use `k6/x/mock` module as drop-in replacement for `k6/http` for seamless mocking.

```JavaScript
import http, { mock } from "k6/x/mock"
```

# Usage tips

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

# Disabling mock

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

