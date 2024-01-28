package objects

import "github.com/rkennedy/optional"

type Member struct {
	Handle string                 `json:"handle"`
	Name   optional.Value[string] `json:"name"`
}
