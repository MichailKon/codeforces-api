package objects

import "github.com/rkennedy/optional"

type HackVerdict string

const (
	HACK_SUCCESSFUL             HackVerdict = "HACK_SUCCESSFUL"
	HACK_UNSUCCESSFUL           HackVerdict = "HACK_UNSUCCESSFUL"
	HACK_INVALID_INPUT          HackVerdict = "INVALID_INPUT"
	HACK_GENERATOR_INCOMPILABLE HackVerdict = "GENERATOR_INCOMPILABLE"
	HACK_GENERATOR_CRASHED      HackVerdict = "GENERATOR_CRASHED"
	HACK_IGNORED                HackVerdict = "IGNORED"
	HACK_TESTING                HackVerdict = "TESTING"
	HACK_OTHER                  HackVerdict = "OTHER"
)

type Hack struct {
	Id                  int                               `json:"id"`
	CreationTimeSeconds int                               `json:"creationTimeSeconds"`
	Hacker              Party                             `json:"hacker"`
	Defender            Party                             `json:"defender"`
	Verdict             optional.Value[HackVerdict]       `json:"verdict"`
	Problem             Problem                           `json:"problem"`
	Test                optional.Value[string]            `json:"test"`
	JudgeProtocol       optional.Value[HackJudgeProtocol] `json:"judgeProtocol"`
}
