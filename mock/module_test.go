// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package mock

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
	"go.k6.io/k6/js/modules"
)

func TestNewModuleInstance(t *testing.T) {
	t.Parallel()

	tc := newHelper(t) // nolint:varnamelen
	mod := tc.module
	vu := tc.vu // nolint:varnamelen

	exports := mod.Exports()
	defaults := exports.Default.(*goja.Object) // nolint:forcetypeassert

	assert.NotEmpty(t, exports.Default)
	assert.Empty(t, exports.Named)

	appCtor := defaults.Get("Application")
	assert.NotNil(t, appCtor, "missing Application constructor")

	_, ok := goja.AssertConstructor(appCtor)

	assert.True(t, ok, "not a constructor")

	assert.NoError(t, vu.Runtime().Set("Application", appCtor))

	val, err := vu.Runtime().RunString("new Application()")

	assert.NoError(t, err)

	assert.IsType(t, vu.Runtime().GlobalObject(), val)

	testModuleExports(t, exports)
}

var exported = []string{"get", "head", "post", "put", "patch", "options", "del", "asyncRequest", "mock", "unmock", "Application"}

func testModuleExports(t *testing.T, exports modules.Exports) {
	t.Helper()

	defaults := exports.Default.(*goja.Object) // nolint:forcetypeassert

	for _, name := range exported {
		assert.NotNil(t, defaults.Get(name))
	}
}
