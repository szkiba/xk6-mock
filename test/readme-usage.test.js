import http from "k6/x/mock/http";

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
