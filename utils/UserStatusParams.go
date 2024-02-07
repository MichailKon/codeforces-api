package utils

import (
	"github.com/rkennedy/optional"
	"net/url"
	"strconv"
)

type UserStatusParams struct {
	from  optional.Value[int]
	count optional.Value[int]
}

func NewUserStatusParams() *UserStatusParams {
	return &UserStatusParams{}
}

func (params *UserStatusParams) WithFrom(from int) *UserStatusParams {
	params.from = optional.New(from)
	return params
}

func (params *UserStatusParams) WithCount(count int) *UserStatusParams {
	params.count = optional.New(count)
	return params
}

func (params *UserStatusParams) FillUrlValues(values *url.Values) {
	params.from.If(func(i int) {
		values.Set("from", strconv.Itoa(i))
	})
	params.count.If(func(i int) {
		values.Set("count", strconv.Itoa(i))
	})
}
