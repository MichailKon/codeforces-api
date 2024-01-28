package objects

import "github.com/rkennedy/optional"

type PartyParticipantType string

const (
	CONTESTANT         PartyParticipantType = "CONTESTANT"
	PRACTICE           PartyParticipantType = "PRACTICE"
	VIRTUAL            PartyParticipantType = "VIRTUAL"
	MANAGER            PartyParticipantType = "MANAGER"
	OUT_OF_COMPETITION PartyParticipantType = "OUT_OF_COMPETITION"
)

type Party struct {
	ContestId        optional.Value[int]    `json:"contestId"`
	Members          []Member               `json:"members"`
	ParticipantType  PartyParticipantType   `json:"participantType"`
	TeamId           optional.Value[int]    `json:"teamId"`
	TeamName         optional.Value[string] `json:"teamName"`
	Ghost            bool                   `json:"ghost"`
	Room             optional.Value[int]    `json:"room"`
	StartTimeSeconds optional.Value[int]    `json:"startTimeSeconds"`
}
