package codeforces_objects

type RecentAction struct {
	TimeSeconds int        `json:"timeSeconds"`
	BlogEntry   *BlogEntry `json:"blogEntry"`
	Comment     *Comment   `json:"comment"`
}
