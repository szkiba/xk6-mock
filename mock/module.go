// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package mock

import (
	"errors"
	"fmt"

	"github.com/grafana/sobek"
	"github.com/sirupsen/logrus"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/js/modules/k6/http"
)

type RootModule struct {
	*http.RootModule
}

func New() modules.Module {
	return &RootModule{RootModule: http.New()}
}

func (root *RootModule) NewModuleInstance(vu modules.VU) modules.Instance { // nolint:varnamelen
	return &Module{
		ModuleInstance: root.RootModule.NewModuleInstance(vu).(*http.ModuleInstance), // nolint:forcetypeassert
		vu:             vu,
		appCtor:        newApplicationCtor(vu, false),
		appCtorSync:    newApplicationCtor(vu, true),
		logger:         newLogger(vu),
		apps:           make(map[string]*sobek.Object),
		lookup:         make(map[string]string),
	}
}

type Module struct {
	*http.ModuleInstance
	vu          modules.VU
	appCtor     func(sobek.ConstructorCall) *sobek.Object
	appCtorSync func(sobek.ConstructorCall) *sobek.Object
	apps        map[string]*sobek.Object
	lookup      map[string]string
	logger      logrus.FieldLogger
}

var (
	_ modules.Module   = (*RootModule)(nil)
	_ modules.Instance = (*Module)(nil)
)

func (mod *Module) runtime() *sobek.Runtime {
	return mod.vu.Runtime()
}

func (mod *Module) throw(err error) {
	common.Throw(mod.runtime(), err)
}

func (mod *Module) throwf(format string, err error, args ...interface{}) {
	mod.throw(fmt.Errorf("%w: "+format, append(append([]interface{}{}, err), args...)...)) // nolint:goerr113
}

func (mod *Module) Exports() modules.Exports {
	exports := mod.ModuleInstance.Exports()
	defaults := exports.Default.(*sobek.Object) // nolint:forcetypeassert

	mod.wrapHTTPExports(defaults)

	mustSet := func(name string, value interface{}) {
		if err := defaults.Set(name, value); err != nil {
			common.Throw(mod.runtime(), err)
		}
	}

	mustSet("unmock", mod.unmock)
	mustSet("Application", mod.applicationCtor())
	mustSet("mock", mod.mockWithSkip())

	return exports
}

type options struct {
	sync bool
	skip bool
}

func getopts(value sobek.Value) *options {
	opts := new(options)

	if obj, ok := value.(*sobek.Object); ok {
		flag := func(name string) bool {
			v := obj.Get(name)

			return v != nil && v.ToBoolean()
		}

		opts.sync = flag("sync")
		opts.skip = flag("skip")
	}

	return opts
}

var errInvalidArg = errors.New("invalid argument")
