// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package mock

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.k6.io/k6/js/modulestest"
)

type testHelper struct {
	runtime *modulestest.Runtime
	vu      *modulestest.VU
	module  *Module
}

func newHelper(t *testing.T) *testHelper {
	t.Helper()

	runtime := modulestest.NewRuntime(t)
	vu := runtime.VU // nolint:varnamelen

	assert.NoError(t, vu.Runtime().Set("__VU", 1))

	root := New()

	var module *Module

	assert.NotPanics(t, func() { module = root.NewModuleInstance(vu).(*Module) }) // nolint:forcetypeassert
	assert.NotNil(t, module)

	exports := module.Exports()
	obj := exports.Default.(*goja.Object) // nolint:forcetypeassert

	assert.NoError(t, vu.Runtime().Set("mock", obj.Get("mock")))
	assert.NoError(t, vu.Runtime().Set("unmock", obj.Get("unmock")))
	assert.NoError(t, vu.Runtime().Set("Application", obj.Get("Application")))

	return &testHelper{
		runtime: runtime,
		vu:      vu,
		module:  module,
	}
}

type suiteBase struct {
	suite.Suite
	*testHelper
}

func (suite *suiteBase) SetupSuite() {
	suite.testHelper = newHelper(suite.T())
}

func (suite *suiteBase) run(script string) (goja.Value, error) {
	return suite.vu.Runtime().RunString(script)
}

func (suite *suiteBase) js(script string) goja.Value {
	value, err := suite.run(script)

	suite.NoError(err)

	return value
}
