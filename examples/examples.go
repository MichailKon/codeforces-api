package examples

import (
	codeforcesapi "codeforces-api"
	"fmt"
)

func main() {
	users, err := codeforcesapi.GlobalCodeforcesSession.UserInfo("Kon567889")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("User %v with rating %d\n", users[0].Handle, users[0].Rating)
	}

	session := codeforcesapi.NewCodeforcesSession("xxx", "yyy") // replace with your keys
	hacks, err := session.ContestHacks(566, &[]bool{true}[0])   // sorry for that
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Number of hacks (manager view): %v\n", len(hacks))
	}

	hacks, err = session.ContestHacks(566, &[]bool{false}[0]) // sorry for that
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Number of hacks (NOT manager view): %v\n", len(hacks))
	}
}
