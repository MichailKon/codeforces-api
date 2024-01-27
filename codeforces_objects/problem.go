package codeforces_objects

type ProblemType string

const (
	PROGRAMMING ProblemType = "PROGRAMMING"
	QUESTION    ProblemType = "QUESTION"
)

type Problem struct {
	ContestId      *int        `json:"contestId"`
	ProblemsetName *string     `json:"problemsetName"`
	Index          string      `json:"index"`
	Name           string      `json:"name"`
	Type           ProblemType `json:"type"`
	Points         *float64    `json:"points"`
	Rating         *int        `json:"rating"`
	Tags           []string    `json:"tags"`
}
