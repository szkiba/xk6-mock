# xk6-mock

A [k6](https://go.k6.io/k6) extension for mocking HTTP(S) servers during test development. The design of the library was inspired by [Express](https://expressjs.com/). If you already know Express framework, using this library should be very simple.

Built for [k6](https://go.k6.io/k6) using [xk6](https://github.com/k6io/xk6).

## Features

- Start mock HTTP server inside of a k6 process
- Familiar, Express like mock route definitions
- Almost transparent for test scripts: change import from `k6/http` to `k6/x/mock/http`
- Helps testing k6 tests with mock server

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
**Table of Contents**

- [Usage](#usage)
  - [http module](#http-module)
  - [mock module](#mock-module)
- [API](#api)
- [Example](#example)
- [Best practice](#best-practice)
- [Build](#build)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Usage

This section contains an annotated, runnable k6 usage example script. You can extract the script with [codedown](https://www.npmjs.com/package/codedown) and run with k6 using the following command line:

```bash
cat README.md | codedown javascript | k6 run -
```

Or you can simply download [test/readme-usage.test.js](test/readme-usage.test.js)

### http module

The `k6/x/mock/http` module is a drop-in replacement for `k6/http` module. During the k6 test development simply replace `k6/http` imports with `k6/x/mock/http`. Each URL's host and port part will be automatically replaced with the mock server's host and port value. Other parts of the request are remain untouched.

```javascript
import http from "k6/x/mock/http";
```

### mock module

The `k6/x/mock` module's default export is an Express like default mock server with the usual HTTP method functions ([get](docs/classes/module.md#get), [head](docs/classes/module.md#head), [post](docs/classes/module.md#post), [put](docs/classes/module.md#put), ..) and [use](docs/classes/module.md#use), function for defining middlewares.

You can add route definitions both in [Init and VU stages](https://k6.io/docs/using-k6/test-life-cycle/#init-and-vu-stages).

```javascript
import mock from "k6/x/mock";

mock.get("/question", (req, res) => {
  res.json({ question: "How many?" });
});

function defaultMockExample() {
  mock.get("/answer", (req, res) => {
    res.json({ answer: 42 });
  });

  const question = http.get("https://question-api.server.url/question");
  const answer = http.get("https://answer-api.server.url/answer");
  
  console.log(question.body); // {"question":"How many?"}
  console.log(answer.body); // {"answer":42}
}
```

With using the Server constructor you can create other mock server instances as well. Custom mock servers should start/stop manually. The server address can be get from `addr()` function.

```javascript
import { Server } from "k6/x/mock";

function customMockExample() {
  const api = new Server()
    .get("/custom", (req, res) => {
      res.json({ foo: "bar" });
    })
    .start();

  const res = http.get(`http://${api.addr()}/custom`);
  
  console.log(res.body); // {"foo":"bar"}

  api.stop();
}

export default function () {
  defaultMockExample();
  customMockExample();
}
```

## API

API documentation can be found in [docs/README.md](docs/README.md).

## Example

This section contains an annotated, runnable k6 example script. You can extract the script with [codedown](https://www.npmjs.com/package/codedown) and run with k6 using the following command line:

```bash
cat README.md | codedown js | k6 run -
```

Or you can simply download [test/readme-example.test.js](test/readme-example.test.js)

 1. Import mock and mock/http modules

    ```js
    import mock from "k6/x/mock";
    import http from "k6/x/mock/http";
    ```

 2. Import and setup [Kahwah](https://github.com/szkiba/kahwah) test runner

    ```js
    import { describe, it } from "https://cdn.jsdelivr.net/npm/kahwah";
    export { options, default } from "https://cdn.jsdelivr.net/npm/kahwah";
    ```

 3. Import `expect` from [Chai Assertion Library](https://www.chaijs.com/)

    ```js
    import { expect } from "cdnjs.com/libraries/chai";
    ```

 4. Define mock routes

    ```js
    const users = {
      alice: {
        name: "alice",
        email: "alice@example.com",
      },
    };

    mock.get("/user/{name}", (req, res) => {
      const user = users[req.params.name];
      if (user) {
        res.json(user);
      } else {
        res.status(404);
      }
    });

    mock.post("/user", (req, res) => {
      const user = req.body;
      if (user.name in users) {
        res.status(409);
      } else {
        users[user.name] = user;
        res.json(user);
      }
    });
    ```

 5. Create tests

    ```js
    const base = "https://user-service.example.com";

    describe("Get user", () => {
      it("existing user", () => {
        const res = http.get(`${base}/user/alice`);
        expect(res.status).equal(200);
        expect(JSON.parse(res.body).name).equal("alice");
      });

      it("missing user", () => {
        const res = http.get(`${base}/user/anonymous`);
        expect(res.status).equal(404);
      });
    });

    describe("Create user", () => {
      const options = { headers: { "Content-Type": "application/json" } };

      it("new user", () => {
        const res = http.post(`${base}/user`, JSON.stringify({ name: "bob" }), options);
        expect(res.status).equal(200);
        expect(JSON.parse(res.body).name).equal("bob");
      });

      it("existing user", () => {
        const res = http.post(`${base}/user`, JSON.stringify({ name: "alice" }), options);
        expect(res.status).equal(409);
      });
    });
    ```

    The output of the test script will be something like this:

    ```plain
        █ Get user

          ✓ existing user
          ✓ missing user

        █ Create user

          ✓ new user
          ✓ existing user
    ```

## Best practice

This section contains an annotated, runnable k6 example script to show how to organize mocks into separated file and make it easy to swith between mock server and real server version.

You can download the mock server version of the test from [test/readme-best-practice-mock.test.js](test/readme-best-practice-mock.test.js) and the real server version from [test/readme-best-practice.test.js](test/readme-best-practice.test.js). As you see, the only difference is the first import line.

 1. Create separated `mock.js` module for mocking (download: [test/mock.js](test/mock.js))

    ```Javascript
    import mock from "k6/x/mock";

    mock.get("/user", (req, res) => {
      res.json({ name: "alice", email: "alice@example.com" });
    });
    ```

 2. Re-export `k6/x/mock/http` module content

    ```Javascript
    export { default } from "k6/x/mock/http";
    export * from "k6/x/mock/http";
    ```

 3. In test script, import `http` from `mock.js` instead of `k6/http`
    ```JavaScript
    import http from "./mock.js";
    ```

    Switching from mock to real implementation as easy as replacing the line above with real `k6/http` module import 

    ```javascripT
    import http from "k6/http"
    ```

 4. The other part of the test script is independent from mocking

    ```JavaScript
    import { describe, it } from "https://cdn.jsdelivr.net/npm/kahwah";
    export { options, default } from "https://cdn.jsdelivr.net/npm/kahwah";
    import { expect } from "cdnjs.com/libraries/chai";

    describe("Get random user", () => {
      const res = http.get("https://phantauth.net/user");

      it("200 OK", () => {
        expect(res.status).equal(200);
      });

      const user = JSON.parse(res.body);

      it("name", () => {
        expect(user).to.have.property("name");
      });

      it("email", () => {
        expect(user).to.have.property("email");
      });
    });
    ```

    The output of the test script will be something like this:

    ```plain
        █ Get random user

          ✓ 200 OK
          ✓ name
          ✓ email
    ```

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```bash
  $ go install github.com/k6io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```bash
  $ xk6 build --with github.com/szkiba/xk6-mock@latest
  ```
