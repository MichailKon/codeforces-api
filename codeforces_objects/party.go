package codeforces_objects

type PartyParticipantType string

const (
	CONTESTANT         PartyParticipantType = "CONTESTANT"
	PRACTICE           PartyParticipantType = "PRACTICE"
	VIRTUAL            PartyParticipantType = "VIRTUAL"
	MANAGER            PartyParticipantType = "MANAGER"
	OUT_OF_COMPETITION PartyParticipantType = "OUT_OF_COMPETITION"
)

type Party struct {
	ContestId        *int                 `json:"contestId"`
	Members          []*Member            `json:"members"`
	ParticipantType  PartyParticipantType `json:"participantType"`
	TeamId           *int                 `json:"teamId"`
	TeamName         *string              `json:"teamName"`
	Ghost            bool                 `json:"ghost"`
	Room             *int                 `json:"room"`
	StartTimeSeconds *int                 `json:"startTimeSeconds"`
}
