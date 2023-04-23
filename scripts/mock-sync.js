import http, { mock } from 'k6/x/mock'
import { check } from 'k6'
import { test } from 'k6/execution'

mock({ sync: true }, 'https://example.com', app => {
  app.get('/', (req, res) => {
    res.json({ greeting: 'Hello World!' })
  })
})

export default function () {
  const res = http.get('https://example.com')
  const ok = check(res, {
    'response code was 200': res => res.status == 200,
    '"greeting" was "Hello World!"': res =>
      res.json('greeting') == 'Hello World!'
  })

  if (!ok) {
    test.abort('unexpected response')
  }
}
