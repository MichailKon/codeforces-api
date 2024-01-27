package codeforces_objects

type User struct {
	Handle                  string  `json:"handle"`
	Email                   string  `json:"email"`
	VkId                    string  `json:"vkId"`
	OpenId                  string  `json:"openId"`
	FirstName               *string `json:"firstName"`
	LastName                *string `json:"lastName"`
	Country                 *string `json:"country"`
	City                    *string `json:"city"`
	Organization            *string `json:"organization"`
	Contribution            int     `json:"contribution"`
	Rank                    string  `json:"rank"`
	Rating                  int     `json:"rating"`
	MaxRank                 string  `json:"maxRank"`
	MaxRating               int     `json:"maxRating"`
	LastOnlineTimeSeconds   int     `json:"lastOnlineTimeSeconds"`
	RegistrationTimeSeconds int     `json:"registrationTimeSeconds"`
	FriendOfCount           int     `json:"friendOfCount"`
	Avatar                  string  `json:"avatar"`
	TitlePhoto              string  `json:"titlePhoto"`
}
