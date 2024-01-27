package codeforces_objects

type RanklistRow struct {
	Party                     Party            `json:"party"`
	Rank                      int              `json:"rank"`
	Points                    float64          `json:"points"`
	Penalty                   int              `json:"penalty"`
	SuccessfulHackCount       int              `json:"successfulHackCount"`
	UnsuccessfulHackCount     int              `json:"unsuccessfulHackCount"`
	ProblemResults            []*ProblemResult `json:"problemResults"`
	LastSubmissionTimeSeconds *int             `json:"lastSubmissionTimeSeconds"`
}
