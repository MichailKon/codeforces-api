package objects

import "github.com/rkennedy/optional"

type ProblemType string

const (
	PROGRAMMING ProblemType = "PROGRAMMING"
	QUESTION    ProblemType = "QUESTION"
)

type Problem struct {
	ContestId      optional.Value[int]     `json:"contestId"`
	ProblemsetName optional.Value[string]  `json:"problemsetName"`
	Index          string                  `json:"index"`
	Name           string                  `json:"name"`
	Type           ProblemType             `json:"type"`
	Points         optional.Value[float64] `json:"points"`
	Rating         optional.Value[int]     `json:"rating"`
	Tags           []string                `json:"tags"`
}
