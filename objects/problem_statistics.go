package objects

import "github.com/rkennedy/optional"

type ProblemStatistics struct {
	ContestId   optional.Value[int] `json:"contestId"`
	Index       string              `json:"index"`
	SolvedCount int                 `json:"solvedCount"`
}
