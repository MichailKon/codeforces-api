package objects

import "github.com/rkennedy/optional"

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
	Id                  int                    `json:"id"`
	Name                string                 `json:"name"`
	Type                ContestType            `json:"type"`
	Phase               ContestPhase           `json:"phase"`
	Frozen              bool                   `json:"frozen"`
	DurationSeconds     int                    `json:"durationSeconds"`
	StartTimeSeconds    int                    `json:"startTimeSeconds"`
	RelativeTimeSeconds int                    `json:"relativeTimeSeconds"`
	PreparedBy          optional.Value[string] `json:"preparedBy,omitempty"`
	WebsiteUrl          optional.Value[string] `json:"websiteUrl,omitempty"`
	Description         optional.Value[string] `json:"description,omitempty"`
	Difficulty          optional.Value[int]    `json:"difficulty,omitempty"`
	Kind                optional.Value[string] `json:"kind,omitempty"`
	IcpcRegion          optional.Value[string] `json:"icpcRegion,omitempty"`
	Country             optional.Value[string] `json:"country,omitempty"`
	City                optional.Value[string] `json:"city,omitempty"`
	Season              optional.Value[string] `json:"season,omitempty"`
}
