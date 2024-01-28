package utils

import (
	"github.com/rkennedy/optional"
	"net/url"
	"strconv"
)

type ContestStatusParams struct {
	asManager optional.Value[bool]
	from      optional.Value[int]
	count     optional.Value[int]
	handle    optional.Value[string]
}

func (params *ContestStatusParams) WithAsManager(asManager bool) *ContestStatusParams {
	params.asManager = optional.New(asManager)
	return params
}

func (params *ContestStatusParams) WithFrom(from int) *ContestStatusParams {
	params.from = optional.New(from)
	return params
}

func (params *ContestStatusParams) WithCount(count int) *ContestStatusParams {
	params.count = optional.New(count)
	return params
}

func (params *ContestStatusParams) WithHandle(handles string) *ContestStatusParams {
	params.handle = optional.New(handles)
	return params
}

func (params *ContestStatusParams) FillUrlValues(values *url.Values) {
	params.asManager.If(func(b bool) {
		values.Set("asManager", strconv.FormatBool(b))
	})
	params.from.If(func(i int) {
		values.Set("from", strconv.Itoa(i))
	})
	params.count.If(func(i int) {
		values.Set("from", strconv.Itoa(i))
	})
	params.handle.If(func(s string) {
		values.Set("handles", s)
	})
}
