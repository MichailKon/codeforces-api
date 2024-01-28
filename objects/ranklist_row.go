package objects

import "github.com/rkennedy/optional"

type RanklistRow struct {
	Party                     Party                   `json:"party"`
	Rank                      int                     `json:"rank"`
	Points                    float64                 `json:"points"`
	Penalty                   int                     `json:"penalty"`
	SuccessfulHackCount       int                     `json:"successfulHackCount"`
	UnsuccessfulHackCount     int                     `json:"unsuccessfulHackCount"`
	ProblemResults            []ProblemResult         `json:"problemResults"`
	LastSubmissionTimeSeconds optional.Value[float64] `json:"lastSubmissionTimeSeconds"`
}
