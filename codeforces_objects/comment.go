package codeforces_objects

type Comment struct {
	Id                  int    `json:"id"`
	CreationTimeSeconds int    `json:"creationTimeSeconds"`
	CommentatorHandle   string `json:"commentatorHandle"`
	Locale              string `json:"locale"`
	Text                string `json:"text"`
	ParentCommentId     *int   `json:"parentCommentId"`
	Rating              int    `json:"rating"`
}
