package codeforces_objects

type ContestType string

const (
	CF   ContestType = "CF"
	IOI  ContestType = "IOI"
	ICPC ContestType = "ICPC"
)

type ContestPhase string

const (
	BEFORE              ContestPhase = "BEFORE"
	CODING              ContestPhase = "CODING"
	PENDING_SYSTEM_TEST ContestPhase = "PENDING_SYSTEM_TEST"
	SYSTEM_TEST         ContestPhase = "SYSTEM_TEST"
	FINISHED            ContestPhase = "FINISHED"
)

type Contest struct {
	Id                  int          `json:"id"`
	Name                string       `json:"name"`
	Type                ContestType  `json:"type"`
	Phase               ContestPhase `json:"phase"`
	Frozen              bool         `json:"frozen"`
	DurationSeconds     int          `json:"durationSeconds"`
	StartTimeSeconds    int          `json:"startTimeSeconds"`
	RelativeTimeSeconds int          `json:"relativeTimeSeconds"`
	PreparedBy          *string      `json:"preparedBy,omitempty"`
	WebsiteUrl          *string      `json:"websiteUrl,omitempty"`
	Description         *string      `json:"description,omitempty"`
	Difficulty          *int         `json:"difficulty,omitempty"`
	Kind                *string      `json:"kind,omitempty"`
	IcpcRegion          *string      `json:"icpcRegion,omitempty"`
	Country             *string      `json:"country,omitempty"`
	City                *string      `json:"city,omitempty"`
	Season              *string      `json:"season,omitempty"`
}
