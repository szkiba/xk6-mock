// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package mock

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/stretchr/testify/assert"
)

func TestNewMockArgs(t *testing.T) {
	t.Parallel()

	helper := newHelper(t)

	call := goja.FunctionCall{
		This:      goja.Undefined(),
		Arguments: []goja.Value{},
	}

	assert.Panics(t, func() { helper.module.newMockArgs(call) })

	callback := helper.vu.Runtime().ToValue(func(goja.FunctionCall) goja.Value { return nil })

	call.Arguments = append(call.Arguments, callback)

	assert.Panics(t, func() { helper.module.newMockArgs(call) })

	call.Arguments = append(call.Arguments, helper.vu.Runtime().ToValue("https://example.com"))

	args := helper.module.newMockArgs(call)

	assert.NotNil(t, args.callback)
	assert.NotEmpty(t, args.target)
	assert.NotNil(t, args.options)
	assert.False(t, args.options.skip)
	assert.False(t, args.options.sync)
}

func TestSkipMock(t *testing.T) {
	t.Parallel()

	helper := newHelper(t)

	assert.False(t, helper.module.skipMock())

	assert.NoError(t, helper.vu.Runtime().Set("__VU", 0))
	assert.True(t, helper.module.skipMock())

	assert.NoError(t, helper.vu.Runtime().Set("__VU", nil))
	assert.True(t, helper.module.skipMock())
}

func TestMockUnmok_skip(t *testing.T) {
	t.Parallel()

	helper := newHelper(t)

	call := goja.FunctionCall{This: nil, Arguments: nil}

	assert.Panics(t, func() { helper.module.mock(call) })

	assert.NoError(t, helper.vu.Runtime().Set("__VU", 0))

	assert.NotPanics(t, func() { helper.module.mock(call) })

	assert.NotPanics(t, func() { helper.module.unmock(helper.vu.Runtime().ToValue("https://example.com")) })
}
