// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package mock

import (
	"testing"

	"github.com/grafana/sobek"
	"github.com/imroc/req/v3"
	"github.com/stretchr/testify/suite"
)

type scriptSuite struct {
	suiteBase
}

func TestScript(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(scriptSuite))
}

func (suite *scriptSuite) TestApplication() {
	host := suite.js(`
	// js
	const app = new Application({sync:true})
	app.get('/', (req, res) => {
	  	res.text("Hello World!")
  })
	app.listen()
	app.host
	// !js
	`)

	resp, err := req.R().Get("http://" + host.String())

	suite.NoError(err)

	body, err := resp.ToString()

	suite.NoError(err)
	suite.Equal(200, resp.GetStatusCode())
	suite.Equal("Hello World!", body)
}

func (suite *scriptSuite) TestScriptMockUnmock() {
	suite.js(`
// js
function handler(app) {
	app.get('/', (req, res) => {
		res.text("Hello World!")
  })
}

mock("https://example.com", handler, {sync:true})
// !js
`)

	suite.Run("mock", func() {
		suite.Equal(1, len(suite.module.apps))
		suite.Equal(1, len(suite.module.lookup))

		args := []sobek.Value{suite.vu.Runtime().ToValue("https://example.com")}

		suite.module.rewrite(args, 0)
		suite.Equal(suite.module.lookup["https://example.com"], args[0].String())

		res, err := req.Get(args[0].String())

		suite.NoError(err)

		body, err := res.ToString()

		suite.NoError(err)
		suite.Equal("Hello World!", body)
	})

	suite.Run("unmock", func() {
		suite.js(`unmock("https://example.com")`)

		suite.Empty(suite.module.lookup)
		suite.Empty(suite.module.apps)

		args := []sobek.Value{suite.vu.Runtime().ToValue("https://example.com")}

		suite.module.rewrite(args, 0)

		suite.Equal("https://example.com", args[0].String())
	})
}

func (suite *scriptSuite) TestScriptRewriteLocalhost() {
	suite.js(`
// js
function handler(app) {
	app.get('/', (req, res) => {
		res.text("Hello World!")
  })
}

mock("http://localhost", handler, {sync:true})
mock("http://127.0.0.1", handler, {sync:true})
// !js
`)

	suite.Equal(2, len(suite.module.apps))
	suite.Equal(2, len(suite.module.lookup))

	args := []sobek.Value{suite.vu.Runtime().ToValue("http://localhost")}

	suite.module.rewrite(args, 0)
	suite.Equal("http://localhost", args[0].String())

	args = []sobek.Value{suite.vu.Runtime().ToValue("http://127.0.0.1")}

	suite.module.rewrite(args, 0)
	suite.Equal("http://127.0.0.1", args[0].String())
}

func (suite *scriptSuite) TestScriptMockCallbackError() {
	script := `
// js
function handler(app) {
	app.get('/', (req, res) => {
		res.text("Hello World!")
  })

	throw Error()
}

mock("htts://example.com", handler, {sync:true})
// !js
`
	_, err := suite.run(script)

	suite.Error(err)
}

func (suite *scriptSuite) TestScriptMockSkip() {
	suite.js(`
// js
function handler(app) {
	app.get('/', (req, res) => {
		res.text("Hello World!")
  })
}

mock.skip("https://func.example.com", handler)
mock({ skip:true }, "https://before.example.com", handler)
mock("https://after.example.com", handler, { skip:true })
// !js
`)

	suite.Empty(suite.module.apps)
}
