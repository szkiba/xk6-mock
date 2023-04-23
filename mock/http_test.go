// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package mock

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
)

func TestModuleWrap(t *testing.T) {
	t.Parallel()

	helper := newHelper(t)

	target := helper.vu.Runtime().NewObject()

	assert.NoError(t, target.Set("not_a_func", 1))

	assert.Panics(t, func() { helper.module.wrap(target, "not_a_func", 1) })

	var actual string

	method := func(loc string) {
		actual = loc
	}

	assert.NoError(t, target.Set("method", method))

	helper.module.lookup["https://example.com"] = "https://example.net"

	helper.module.wrap(target, "method", 0)

	callable, ok := goja.AssertFunction(target.Get("method"))

	assert.True(t, ok)

	_, err := callable(goja.Undefined(), helper.vu.Runtime().ToValue("https://example.com"))

	assert.NoError(t, err)
	assert.Equal(t, "https://example.net", actual)
}
