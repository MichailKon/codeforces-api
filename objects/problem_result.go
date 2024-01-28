package objects

import "github.com/rkennedy/optional"

type ProblemResultType string

const (
	PRELIMINARY ProblemResultType = "PRELIMINARY"
	FINAL       ProblemResultType = "FINAL"
)

type ProblemResult struct {
	Points                    float64             `json:"points"`
	Penalty                   optional.Value[int] `json:"penalty"`
	RejectedAttemptCount      int                 `json:"rejectedAttemptCount"`
	Type                      ProblemResultType   `json:"type"`
	BestSubmissionTimeSeconds optional.Value[int] `json:"bestSubmissionTimeSeconds"`
}
