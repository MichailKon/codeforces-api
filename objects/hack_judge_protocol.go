package objects

type HackJudgeProtocol struct {
	Manual   string `json:"manual"` // bruh it's true/false
	Protocol string `json:"protocol"`
	Verdict  string `json:"verdict"`
}
