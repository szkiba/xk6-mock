// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package mock

import (
	"reflect"
	"strings"

	"github.com/grafana/sobek"
)

func (mod *Module) skipMock() bool {
	if v := mod.runtime().Get("__VU"); v == nil || v.ToInteger() == 0 {
		return true
	}

	return false
}

type mockArgs struct {
	target   string
	callback sobek.Callable
	options  *options
}

func (mod *Module) newMockArgs(call sobek.FunctionCall) *mockArgs {
	args := new(mockArgs)

	for idx := 0; idx < len(call.Arguments); idx++ {
		if c, isFunc := sobek.AssertFunction(call.Argument(idx)); isFunc {
			args.callback = c

			continue
		}

		if obj, isObj := call.Argument(idx).(*sobek.Object); isObj {
			args.options = getopts(obj)

			continue
		}

		if call.Argument(idx).ExportType().Kind() == reflect.String {
			args.target = call.Argument(idx).String()
		}
	}

	if args.callback == nil {
		mod.throwf("missingr callback function", errInvalidArg)
	}

	if len(args.target) == 0 {
		mod.throwf("missing or empty mock target", errInvalidArg)
	}

	if args.options == nil {
		args.options = new(options)
	}

	return args
}

func (mod *Module) mock(call sobek.FunctionCall) sobek.Value {
	if mod.skipMock() {
		return sobek.Undefined()
	}

	args := mod.newMockArgs(call)

	if args.options.skip {
		return sobek.Undefined()
	}

	app, listen := mod.newApplication(args.options.sync)

	_, err := args.callback(mod.runtime().GlobalObject(), app)
	if err != nil {
		mod.throw(err)
	}

	_, err = listen(app)
	if err != nil {
		mod.throw(err)
	}

	addr := app.Get("host")
	if addr == nil || len(addr.String()) == 0 {
		mod.logger.WithField("target", args.target).Warn("mock server not started")

		return sobek.Undefined()
	}

	mod.apps[args.target] = app
	mod.lookup[args.target] = "http://" + addr.String()

	return sobek.Undefined()
}

func (mod *Module) mockWithSkip() sobek.Value {
	function := mod.runtime().ToValue(mod.mock).(*sobek.Object) // nolint:forcetypeassert

	function.Set("skip", func(_ sobek.FunctionCall) sobek.Value { return sobek.Undefined() }) // nolint:errcheck

	return function
}

func (mod *Module) unmock(target sobek.Value) {
	if mod.skipMock() {
		return
	}

	key := target.String()

	app, ok := mod.apps[key]
	if !ok {
		return
	}

	delete(mod.apps, key)
	delete(mod.lookup, key)

	shutdown, _ := sobek.AssertFunction(app.Get("shutdown"))

	if _, err := shutdown(app); err != nil {
		mod.throw(err)
	}
}

func (mod *Module) rewrite(args []sobek.Value, index int) {
	loc := args[index].String()

	if strings.HasPrefix(loc, "http://localhost") || strings.HasPrefix(loc, "http://127.") {
		return
	}

	for k, v := range mod.lookup {
		if strings.HasPrefix(loc, k) {
			args[index] = mod.runtime().ToValue(strings.Replace(loc, k, v, 1))

			break
		}
	}
}
