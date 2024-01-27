package codeforces_objects

type ProblemStatistics struct {
	ContestId   *int   `json:"contestId"`
	Index       string `json:"index"`
	SolvedCount int    `json:"solvedCount"`
}
