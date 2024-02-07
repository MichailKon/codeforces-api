package utils

import (
	"github.com/rkennedy/optional"
	"net/url"
	"strconv"
)

type UserRatedListParams struct {
	activeOnly     optional.Value[bool]
	includeRetired optional.Value[bool]
	contestId      optional.Value[int]
}

func NewUserRatedListParams() *UserRatedListParams {
	return &UserRatedListParams{}
}

func (r *UserRatedListParams) WithActiveOnly(b bool) *UserRatedListParams {
	r.activeOnly = optional.New(b)
	return r
}

func (r *UserRatedListParams) WithIncludeRetired(b bool) *UserRatedListParams {
	r.includeRetired = optional.New(b)
	return r
}

func (r *UserRatedListParams) WithContestId(id int) *UserRatedListParams {
	r.contestId = optional.New(id)
	return r
}

func (r *UserRatedListParams) FillUrlValues(values *url.Values) {
	r.activeOnly.If(func(b bool) {
		values.Set("activeOnly", strconv.FormatBool(b))
	})
	r.includeRetired.If(func(b bool) {
		values.Set("includeRetired", strconv.FormatBool(b))
	})
	r.contestId.If(func(i int) {
		values.Set("contestId", strconv.Itoa(i))
	})
}
