package codeforces_objects

type ProblemResultType string

const (
	PRELIMINARY ProblemResultType = "PRELIMINARY"
	FINAL       ProblemResultType = "FINAL"
)

type ProblemResult struct {
	Points                    float64           `json:"points"`
	Penalty                   *int              `json:"penalty"`
	RejectedAttemptCount      int               `json:"rejectedAttemptCount"`
	Type                      ProblemResultType `json:"type"`
	BestSubmissionTimeSeconds *int              `json:"bestSubmissionTimeSeconds"`
}
