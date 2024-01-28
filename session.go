package codeforces_api

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/MichailKon/codeforces-api/objects"
	"github.com/MichailKon/codeforces-api/utils"
	"github.com/go-resty/resty/v2"
	"github.com/rkennedy/optional"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	baseUrl  string = "https://codeforces.com/api"
	alphabet string = "abcdefghijklmnopqrstuvwxyz"
)

type CodeforcesApiError struct {
	Body       string
	StatusCode int
}

func (e CodeforcesApiError) Error() string {
	return e.Body
}

type CodeforcesSession struct {
	key    string
	secret string
	Client *resty.Client
}

var globalCodeforcesSession CodeforcesSession
var initGlobal sync.Once

func GetGlobalSession() *CodeforcesSession {
	initGlobal.Do(func() {
		globalCodeforcesSession = *NewCodeforcesSession("", "")
	})
	return &globalCodeforcesSession
}

func NewCodeforcesSession(key string, secret string) *CodeforcesSession {
	return &CodeforcesSession{
		key:    key,
		secret: secret,
		Client: resty.New(),
	}
}

func (session *CodeforcesSession) generateApiSig(methodName string, parameters url.Values) string {
	randval := strings.Builder{}
	for i := 0; i < 6; i++ {
		randval.WriteByte(alphabet[rand.IntN(len(alphabet))])
	}

	s := fmt.Sprintf("%v/%v?%v#%v", randval.String(), methodName, parameters.Encode(), session.secret)
	return fmt.Sprintf("%v%x", randval.String(), sha512.Sum512([]byte(s)))
}

func (session *CodeforcesSession) makeQuery(methodName string, parameters url.Values, out any) error {
	if session.key != "" {
		parameters.Set("apiKey", session.key)
	}
	parameters.Set("time", strconv.FormatInt(time.Now().Unix(), 10))
	if session.key != "" {
		parameters.Set("apiSig", session.generateApiSig(methodName, parameters))
	}

	query := session.Client.R()
	query.SetQueryParamsFromValues(parameters)

	response, err := query.Get(fmt.Sprintf("%v/%v", baseUrl, methodName))
	if err != nil {
		return err
	}
	if response.StatusCode() == http.StatusNoContent {
		return nil
	}
	if response.StatusCode() != http.StatusOK {
		type FailResponse struct {
			Status  string `json:"status"`
			Comment string `json:"comment"`
		}
		var res FailResponse
		if err := json.Unmarshal(response.Body(), &res); err != nil {
			return CodeforcesApiError{
				Body:       fmt.Sprintf("unmarshal error: %v", err),
				StatusCode: response.StatusCode(),
			}
		}
		return CodeforcesApiError{
			Body:       res.Comment,
			StatusCode: response.StatusCode(),
		}
	}

	// maybe this can be made in a better way
	var cur map[string]any
	err = json.Unmarshal(response.Body(), &cur)
	if err != nil {
		return err
	}
	if _, ok := cur["result"]; !ok {
		return fmt.Errorf("no \"result\" field in response")
	}
	jsonBody, err := json.Marshal(cur["result"])
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonBody, &out)
}

func (session *CodeforcesSession) BlogEntryComments(blogEntryId int) ([]*objects.Comment, error) {
	parameters := url.Values{}
	parameters.Set("blogEntryId", strconv.Itoa(blogEntryId))
	var res []*objects.Comment
	if err := session.makeQuery("blogEntry.comments", parameters, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (session *CodeforcesSession) BlogEntryView(blogEntryId int) (*objects.BlogEntry, error) {
	parameters := url.Values{}
	parameters.Set("blogEntryId", strconv.Itoa(blogEntryId))
	var result *objects.BlogEntry
	if err := session.makeQuery("blogEntry.view", parameters, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) generalContestHacks(contestId int, asManager bool) ([]*objects.Hack, error) {
	p := url.Values{}
	p.Set("contestId", strconv.Itoa(contestId))
	p.Set("asManager", strconv.FormatBool(asManager))
	var result []*objects.Hack
	if err := session.makeQuery("contest.hacks", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) ContestHacksAsManager(contestId int) ([]*objects.Hack, error) {
	return session.generalContestHacks(contestId, true)
}

func (session *CodeforcesSession) ContestHacks(contestId int) ([]*objects.Hack, error) {
	return session.generalContestHacks(contestId, false)
}

func (session *CodeforcesSession) generalContestList(gym bool) ([]*objects.Contest, error) {
	p := url.Values{}
	p.Set("gym", strconv.FormatBool(gym))
	var result []*objects.Contest
	if err := session.makeQuery("contest.list", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) ContestListGym() ([]*objects.Contest, error) {
	return session.generalContestList(true)
}

func (session *CodeforcesSession) ContestList() ([]*objects.Contest, error) {
	return session.generalContestList(false)
}

func (session *CodeforcesSession) ContestRatingChanges(contestId int) ([]*objects.RatingChange, error) {
	p := url.Values{}
	p.Set("contestId", strconv.Itoa(contestId))
	var result []*objects.RatingChange
	if err := session.makeQuery("contest.ratingChanges", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) ContestStandings(
	contestId int,
	params *utils.ContestStandingsParams,
) (*objects.ContestStandings, error) {
	p := url.Values{}
	p.Set("contestId", strconv.Itoa(contestId))
	params.FillUrlValues(&p)
	var result *objects.ContestStandings
	if err := session.makeQuery("contest.standings", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) ContestStatus(
	contestId int,
	params *utils.ContestStatusParams,
) ([]*objects.Submission, error) {
	p := url.Values{}
	p.Set("contestId", strconv.Itoa(contestId))
	params.FillUrlValues(&p)
	var result []*objects.Submission
	if err := session.makeQuery("contest.standings", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) ProblemsetRecentStatusWithProblemset(
	count int,
	problemsetName string,
) ([]*objects.Submission, error) {
	p := url.Values{}
	p.Set("count", strconv.Itoa(count))
	p.Add("problemsetName", problemsetName)
	var result []*objects.Submission
	if err := session.makeQuery("problemset.recentStatus", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) ProblemsetRecentStatus(
	count int,
) ([]*objects.Submission, error) {
	return session.ProblemsetRecentStatusWithProblemset(count, "")
}

func (session *CodeforcesSession) RecentActions(maxCount int) ([]*objects.RecentAction, error) {
	p := url.Values{}
	p.Set("maxCount", strconv.Itoa(maxCount))
	var result []*objects.RecentAction
	if err := session.makeQuery("recentActions", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) UserBlogEntries(handle string) ([]*objects.BlogEntry, error) {
	p := url.Values{}
	p.Set("handle", handle)
	var result []*objects.BlogEntry
	if err := session.makeQuery("recentActions", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) generalUserFriends(onlyOnline bool) ([]string, error) {
	p := url.Values{}
	p.Set("onlyOnline", strconv.FormatBool(onlyOnline))
	var result []string
	if err := session.makeQuery("user.friends", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) UserFriendsOnline() ([]string, error) {
	return session.generalUserFriends(true)
}

func (session *CodeforcesSession) UserFriends() ([]string, error) {
	return session.generalUserFriends(false)
}

func (session *CodeforcesSession) UserInfo(handles string) ([]*objects.User, error) {
	p := url.Values{}
	p.Set("handles", handles)
	var result []*objects.User
	if err := session.makeQuery("user.info", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// UserRatedList implements *user.ratedList*
// TODO
// activeOnly, includeRetires, contestId are optional
func (session *CodeforcesSession) UserRatedList(
	activeOnly optional.Value[bool],
	includeRetired optional.Value[bool],
	contestId optional.Value[int],
) ([]*objects.User, error) {
	p := url.Values{}
	activeOnly.If(func(b bool) {
		p.Set("activeOnly", strconv.FormatBool(b))
	})
	includeRetired.If(func(b bool) {
		p.Set("activeOnly", strconv.FormatBool(b))
	})
	contestId.If(func(i int) {
		p.Set("contestId", strconv.Itoa(i))
	})
	var result []*objects.User
	if err := session.makeQuery("user.ratedList", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (session *CodeforcesSession) UserRating(handle string) ([]*objects.RatingChange, error) {
	p := url.Values{}
	p.Set("handle", handle)
	var result []*objects.RatingChange
	if err := session.makeQuery("user.ratedList", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// UserStatus implements *user.status*
// TODO
// from, count are optional
func (session *CodeforcesSession) UserStatus(
	handle string,
	from optional.Value[int],
	count optional.Value[int],
) ([]*objects.Submission, error) {
	p := url.Values{}
	p.Set("handle", handle)
	from.If(func(i int) {
		p.Set("from", strconv.Itoa(i))
	})
	count.If(func(i int) {
		p.Set("count", strconv.Itoa(i))
	})
	var result []*objects.Submission
	if err := session.makeQuery("user.status", p, &result); err != nil {
		return nil, err
	}
	return result, nil
}
