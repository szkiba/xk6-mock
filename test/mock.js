import mock from "k6/x/mock";

mock.get("/user", (req, res) => {
  res.json({ name: "alice", email: "alice@example.com" });
});

export { default } from "k6/x/mock/http";
export * from "k6/x/mock/http";
