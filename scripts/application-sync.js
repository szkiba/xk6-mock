import http, { Application } from 'k6/x/mock'
import { check } from 'k6'
import { test } from 'k6/execution'

const app = new Application({ sync: true })

app.get('/', (req, res) => {
  res.json({ greeting: 'Hello World!' })
})

app.listen()

export default function () {
  const res = http.get('http://' + app.host)
  const ok = check(res, {
    'response code was 200': res => res.status == 200,
    '"greeting" was "Hello World!"': res =>
      res.json('greeting') == 'Hello World!'
  })

  if (!ok) {
    test.abort('unexpected response')
  }
}
