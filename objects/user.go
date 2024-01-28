package objects

import "github.com/rkennedy/optional"

type User struct {
	Handle                  string                 `json:"handle"`
	Email                   string                 `json:"email"`
	VkId                    string                 `json:"vkId"`
	OpenId                  string                 `json:"openId"`
	FirstName               optional.Value[string] `json:"firstName"`
	LastName                optional.Value[string] `json:"lastName"`
	Country                 optional.Value[string] `json:"country"`
	City                    optional.Value[string] `json:"city"`
	Organization            optional.Value[string] `json:"organization"`
	Contribution            int                    `json:"contribution"`
	Rank                    string                 `json:"rank"`
	Rating                  int                    `json:"rating"`
	MaxRank                 string                 `json:"maxRank"`
	MaxRating               int                    `json:"maxRating"`
	LastOnlineTimeSeconds   int                    `json:"lastOnlineTimeSeconds"`
	RegistrationTimeSeconds int                    `json:"registrationTimeSeconds"`
	FriendOfCount           int                    `json:"friendOfCount"`
	Avatar                  string                 `json:"avatar"`
	TitlePhoto              string                 `json:"titlePhoto"`
}
