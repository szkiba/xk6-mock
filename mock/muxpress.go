// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package mock

import (
	"github.com/grafana/sobek"
	"github.com/sirupsen/logrus"
	"github.com/szkiba/muxpress"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

func newRunner(vu modules.VU) muxpress.RunnerFunc { // nolint:varnamelen
	return func(fn func() error) {
		vu.RegisterCallback()(fn)
	}
}

func newLogger(vu modules.VU) logrus.FieldLogger { // nolint:varnamelen
	var logger logrus.FieldLogger

	if vu.InitEnv() != nil && vu.InitEnv().Logger != nil {
		logger = vu.InitEnv().Logger
	} else if vu.State() != nil && vu.State().Logger != nil {
		logger = vu.State().Logger
	}

	if logger == nil {
		logger = logrus.StandardLogger()
	}

	return logger.WithField("module", "mock")
}

func newApplicationCtor(vu modules.VU, sync bool) func(sobek.ConstructorCall) *sobek.Object { // nolint:varnamelen
	opts := []muxpress.Option{muxpress.WithLogger(newLogger(vu))}

	if !sync {
		opts = append(opts, muxpress.WithRunner(newRunner(vu)))
	}

	ctor, err := muxpress.NewApplicationConstructor(vu.Runtime(), opts...)
	if err != nil {
		common.Throw(vu.Runtime(), err)
	}

	return ctor
}

func (mod *Module) applicationCtor() func(sobek.ConstructorCall) *sobek.Object {
	return func(call sobek.ConstructorCall) *sobek.Object {
		if len(call.Arguments) == 0 || !getopts(call.Argument(0)).sync {
			return mod.appCtor(call)
		}

		return mod.appCtorSync(call)
	}
}

func (mod *Module) newApplication(sync bool) (*sobek.Object, sobek.Callable) {
	from := mod.appCtor

	if sync {
		from = mod.appCtorSync
	}

	ctor, assertOK := sobek.AssertConstructor(mod.runtime().ToValue(from))
	if !assertOK {
		mod.throwf("invalid constructor", errInvalidArg)
	}

	app, err := ctor(mod.runtime().NewObject())
	if err != nil {
		mod.throw(err)
	}

	listen, assertOK := sobek.AssertFunction(app.Get("listen"))
	if !assertOK {
		mod.throwf("missing listen method", errInvalidArg)
	}

	return app, listen
}
