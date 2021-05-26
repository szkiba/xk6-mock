import mock from "k6/x/mock";
import http from "k6/x/mock/http";

import { describe, it } from "https://cdn.jsdelivr.net/npm/kahwah";
export { options, default } from "https://cdn.jsdelivr.net/npm/kahwah";

import { expect } from "cdnjs.com/libraries/chai";

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
