import http, { mock } from 'k6/x/mock'
import { check } from 'k6'
import { test } from 'k6/execution'

mock('https://httpbin.test.k6.io/get', app => {
  app.get('/', (req, res) => {
    res.json({ url: 'https://httpbin.test.k6.io/get' })
  })
})

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
