/**
 * MIT License
 *
 * Copyright (c) 2021 Iván Szkiba
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

import { describe, it } from "https://cdn.jsdelivr.net/npm/kahwah";
export { options, default } from "https://cdn.jsdelivr.net/npm/kahwah";
import { expect } from "cdnjs.com/libraries/chai";

import http from "./script.mock.js";
import { Server } from "k6/x/mock";

describe("User", () => {
  const res = http.get("http://phantauth.net/user/joe");

  it("200 OK", () => {
    expect(res.status).equal(200);
  });

  console.log(res.body);
});

describe("Shutdown", () => {
  const api = new Server()
    .get("/custom", (req, res) => {
      res.json({ foo: "bar" });
    })
    .start();

  console.log(`http://${api.addr()}/custom`);

  const res = http.get(`http://${api.addr()}/custom`);

  console.log(res.body); // {"foo":"bar"}

  it("200 OK", () => {
    expect(res.status).equal(200);
  });

  api.stop();
});
