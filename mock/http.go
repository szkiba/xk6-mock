// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package mock

import (
	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
)

// XXX: add batch function support

var (
	urlFirstMethods  = []string{"get", "head", "post", "put", "patch", "options", "del"}
	urlSecondMethods = []string{"request", "asyncRequest"}
)

func (mod *Module) wrapHTTPExports(defaults *goja.Object) {
	for _, method := range urlFirstMethods {
		mod.wrap(defaults, method, 0)
	}

	for _, method := range urlSecondMethods {
		mod.wrap(defaults, method, 1)
	}
}

func (mod *Module) wrap(this *goja.Object, method string, index int) {
	v := this.Get(method)

	callable, ok := goja.AssertFunction(v)
	if !ok {
		mod.throwf("%s must be callable", errInvalidArg, method)
	}

	wrapper := func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) > index {
			mod.rewrite(call.Arguments, index)
		}

		v, err := callable(mod.runtime().GlobalObject(), call.Arguments...)
		if err != nil {
			common.Throw(mod.runtime(), err)
		}

		return v
	}

	err := this.Set(method, mod.runtime().ToValue(wrapper))
	if err != nil {
		common.Throw(mod.runtime(), err)
	}
}
