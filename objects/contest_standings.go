package objects

type ContestStandings struct {
	Contest  Contest       `json:"contest"`
	Problems []Problem     `json:"problems"`
	Rows     []RanklistRow `json:"rows"`
}
