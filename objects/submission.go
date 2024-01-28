package objects

import "github.com/rkennedy/optional"

type SubmissionVerdict string

const (
	FAILED                    SubmissionVerdict = "FAILED"
	OK                        SubmissionVerdict = "OK"
	PARTIAL                   SubmissionVerdict = "PARTIAL"
	COMPILATION_ERROR         SubmissionVerdict = "COMPILATION_ERROR"
	RUNTIME_ERROR             SubmissionVerdict = "RUNTIME_ERROR"
	WRONG_ANSWER              SubmissionVerdict = "WRONG_ANSWER"
	PRESENTATION_ERROR        SubmissionVerdict = "PRESENTATION_ERROR"
	TIME_LIMIT_EXCEEDED       SubmissionVerdict = "TIME_LIMIT_EXCEEDED"
	MEMORY_LIMIT_EXCEEDED     SubmissionVerdict = "MEMORY_LIMIT_EXCEEDED"
	IDLENESS_LIMIT_EXCEEDED   SubmissionVerdict = "IDLENESS_LIMIT_EXCEEDED"
	SECURITY_VIOLATED         SubmissionVerdict = "SECURITY_VIOLATED"
	CRASHED                   SubmissionVerdict = "CRASHED"
	INPUT_PREPARATION_CRASHED SubmissionVerdict = "INPUT_PREPARATION_CRASHED"
	CHALLENGED                SubmissionVerdict = "CHALLENGED"
	SKIPPED                   SubmissionVerdict = "SKIPPED"
	TESTING                   SubmissionVerdict = "TESTING"
	REJECTED                  SubmissionVerdict = "REJECTED"
)

type SubmissionTestset string

const (
	SAMPLES    SubmissionTestset = "SAMPLES"
	PRETESTS   SubmissionTestset = "PRETESTS"
	TESTS      SubmissionTestset = "TESTS"
	CHALLENGES SubmissionTestset = "CHALLENGES"
	TESTS1     SubmissionTestset = "TESTS1"
	TESTS2     SubmissionTestset = "TESTS2"
	TESTS3     SubmissionTestset = "TESTS3"
	TESTS4     SubmissionTestset = "TESTS4"
	TESTS5     SubmissionTestset = "TESTS5"
	TESTS6     SubmissionTestset = "TESTS6"
	TESTS7     SubmissionTestset = "TESTS7"
	TESTS8     SubmissionTestset = "TESTS8"
	TESTS9     SubmissionTestset = "TESTS9"
	TESTS10    SubmissionTestset = "TESTS10"
)

type Submission struct {
	Id                  int                               `json:"id"`
	ContestId           optional.Value[int]               `json:"contestId"`
	CreationTimeSeconds int                               `json:"creationTimeSeconds"`
	RelativeTimeSeconds int                               `json:"relativeTimeSeconds"`
	Problem             Problem                           `json:"problem"`
	Author              Party                             `json:"author"`
	ProgrammingLanguage string                            `json:"programmingLanguage"`
	Verdict             optional.Value[SubmissionVerdict] `json:"verdict"`
	Testset             SubmissionTestset                 `json:"testset"`
	PassedTestCount     int                               `json:"passedTestCount"`
	TimeConsumedMillis  int                               `json:"timeConsumedMillis"`
	MemoryConsumedBytes int                               `json:"memoryConsumedBytes"`
	Points              optional.Value[float64]           `json:"points"`
}
