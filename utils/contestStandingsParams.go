package utils

import (
	"github.com/rkennedy/optional"
	"net/url"
	"strconv"
	"strings"
)

type ContestStandingsParams struct {
	asManager      optional.Value[bool]
	from           optional.Value[int]
	count          optional.Value[int]
	handles        optional.Value[string]
	room           optional.Value[int]
	showUnofficial optional.Value[bool]
}

func NewContestStandingsParams() *ContestStandingsParams {
	return &ContestStandingsParams{}
}

func (params *ContestStandingsParams) WithAsManager(asManager bool) *ContestStandingsParams {
	params.asManager = optional.New(asManager)
	return params
}

func (params *ContestStandingsParams) WithFrom(from int) *ContestStandingsParams {
	params.from = optional.New(from)
	return params
}

func (params *ContestStandingsParams) WithCount(count int) *ContestStandingsParams {
	params.count = optional.New(count)
	return params
}

func (params *ContestStandingsParams) WithHandles(handles string) *ContestStandingsParams {
	params.handles = optional.New(handles)
	return params
}

func (params *ContestStandingsParams) WithHandlesList(handles []string) *ContestStandingsParams {
	params.handles = optional.New(strings.Join(handles, ";"))
	return params
}

func (params *ContestStandingsParams) WithShowUnofficial(showUnofficial bool) *ContestStandingsParams {
	params.showUnofficial = optional.New(showUnofficial)
	return params
}

func (params *ContestStandingsParams) FillUrlValues(values *url.Values) {
	params.asManager.If(func(b bool) {
		values.Set("asManager", strconv.FormatBool(b))
	})
	params.from.If(func(i int) {
		values.Set("from", strconv.Itoa(i))
	})
	params.count.If(func(i int) {
		values.Set("from", strconv.Itoa(i))
	})
	params.handles.If(func(s string) {
		values.Set("handles", s)
	})
	params.room.If(func(i int) {
		values.Set("room", strconv.Itoa(i))
	})
	params.showUnofficial.If(func(b bool) {
		values.Set("showUnofficial", strconv.FormatBool(b))
	})
}
