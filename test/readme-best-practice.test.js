import http from "k6/http";

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
