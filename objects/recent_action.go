package objects

import "github.com/rkennedy/optional"

type RecentAction struct {
	TimeSeconds int                       `json:"timeSeconds"`
	BlogEntry   optional.Value[BlogEntry] `json:"blogEntry"`
	Comment     optional.Value[Comment]   `json:"comment"`
}
