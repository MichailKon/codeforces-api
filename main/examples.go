package main

import (
	codeforcesapi "codeforces-api"
	"codeforces-api/utils"
	"fmt"
)

func main() {
	comments, err := codeforcesapi.GetGlobalSession().BlogEntryComments(1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%+v", comments[0])
	}

	users, err := codeforcesapi.GetGlobalSession().UserInfo("Kon567889")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("User %v with rating %d\n", users[0].Handle, users[0].Rating)
	}
	//
	session := codeforcesapi.NewCodeforcesSession("xxx", "yyy") // replace with your keys
	hacks, err := session.ContestHacks(566)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Number of hacks (manager view): %v\n", len(hacks))
	}

	hacks, err = session.ContestHacksAsManager(566)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Number of hacks (NOT manager view): %v\n", len(hacks))
	}

	standings, err := codeforcesapi.GetGlobalSession().ContestStandings(
		566,
		utils.
			NewContestStandingsParams().
			WithShowUnofficial(true),
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Contest \"%v\" contains %v rows and %v problems",
			standings.Contest.Name,
			len(standings.Rows),
			len(standings.Problems),
		)
	}
}
