import http, { mock } from 'k6/x/mock'
import { check } from 'k6'
import { test } from 'k6/execution'

mock('https://example.com/ASYNC', app => {
  app.get('/', (req, res) => {
    res.json({ greeting: 'Hello World!' })
  })
})

mock({ sync: true }, 'https://example.com/SYNC', app => {
  app.get('/', (req, res) => {
    res.json({ greeting: 'Hello World!' })
  })
})

function assert (res) {
  const ok = check(res, {
    'response code was 200': res => res.status == 200,
    '"greeting" was "Hello World!"': res =>
      res.json('greeting') == 'Hello World!'
  })

  if (!ok) {
    test.abort('unexpected response')
  }
}

export default async function () {
  assert(await http.asyncRequest('GET', 'https://example.com/ASYNC'))
  assert(http.get('https://example.com/SYNC'))
}
