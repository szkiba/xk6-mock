// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package mock

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/js/modulestest"
	"go.k6.io/k6/lib"
)

func TestNewLogger(t *testing.T) {
	t.Parallel()

	vu := modulestest.NewRuntime(t).VU // nolint:varnamelen

	assert.NotNil(t, vu.InitEnv())
	assert.Nil(t, vu.State())

	assert.NotNil(t, newLogger(vu))

	vu.InitEnvField.Logger = nil

	assert.NotNil(t, newLogger(vu))

	vu.InitEnvField = nil
	vu.StateField = &lib.State{} // nolint:exhaustruct

	assert.NotNil(t, vu.State())
	assert.NotNil(t, newLogger(vu))

	vu.StateField.Logger = logrus.StandardLogger()

	assert.NotNil(t, newLogger(vu))

	vu.StateField = nil

	assert.NotNil(t, newLogger(vu))
}

func TestGetopts(t *testing.T) {
	t.Parallel()

	assert.NotNil(t, getopts(nil))

	opts := getopts(nil)

	assert.False(t, opts.sync)
	assert.False(t, opts.skip)

	opts = getopts(goja.Undefined())

	assert.False(t, opts.sync)
	assert.False(t, opts.skip)

	opts = getopts(goja.Null())

	assert.False(t, opts.sync)
	assert.False(t, opts.skip)

	runtime := goja.New()
	obj := runtime.NewObject()

	opts = getopts(obj)

	assert.False(t, opts.sync)
	assert.False(t, opts.skip)

	assert.NoError(t, obj.Set("sync", goja.Null()))
	assert.NoError(t, obj.Set("skip", goja.Undefined()))

	opts = getopts(obj)

	assert.False(t, opts.sync)
	assert.False(t, opts.skip)

	assert.NoError(t, obj.Set("sync", runtime.ToValue(3)))
	assert.NoError(t, obj.Set("skip", true))

	opts = getopts(obj)

	assert.True(t, opts.sync)
	assert.True(t, opts.skip)
}

func TestNewRunner(t *testing.T) {
	t.Parallel()

	vu := modulestest.NewRuntime(t).VU // nolint:varnamelen

	registerCalled := false
	enqueueCalled := false

	vu.RegisterCallbackField = func() func(func() error) {
		registerCalled = true

		return func(f func() error) {
			enqueueCalled = true

			f() // nolint:errcheck
		}
	}

	runner := newRunner(vu)

	assert.False(t, registerCalled)
	assert.False(t, enqueueCalled)

	jobCalled := false

	runner(func() error {
		jobCalled = true

		return nil
	})

	assert.True(t, registerCalled)
	assert.True(t, enqueueCalled)
	assert.True(t, jobCalled)
}

func TestApplicationCtor(t *testing.T) {
	t.Parallel()

	module, vu := newModule(t) // nolint:varnamelen

	ctor := module.applicationCtor()

	call := goja.ConstructorCall{
		This:      vu.Runtime().GlobalObject(),
		NewTarget: vu.Runtime().NewObject(),
		Arguments: []goja.Value{},
	}

	appCtorCalled := false

	module.appCtor = func(cc goja.ConstructorCall) *goja.Object {
		appCtorCalled = true

		return nil
	}

	module.appCtorSync = nil

	ctor(call)

	assert.True(t, appCtorCalled)

	appCtorSyncCalled := false

	appCtorCalled = false

	module.appCtorSync = func(cc goja.ConstructorCall) *goja.Object {
		appCtorSyncCalled = true

		return nil
	}

	opts := vu.Runtime().NewObject()

	assert.NoError(t, opts.Set("sync", true))

	call.Arguments = append(call.Arguments, vu.Runtime().ToValue(opts))

	ctor(call)

	assert.False(t, appCtorCalled)
	assert.True(t, appCtorSyncCalled)
}

func newModule(t *testing.T) (*Module, modules.VU) {
	t.Helper()

	vu := modulestest.NewRuntime(t).VU // nolint:varnamelen

	root := New()

	var module *Module

	assert.NotPanics(t, func() { module = root.NewModuleInstance(vu).(*Module) }) // nolint:forcetypeassert
	assert.NotNil(t, module)

	return module, vu
}
