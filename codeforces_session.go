package codeforces_api

import (
	"bytes"
	"codeforces-api/codeforces_objects"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	baseUrl  string = "https://codeforces.com/api"
	alphabet string = "abcdefghijklmnopqrstuvwxyz"
)

// CodeforcesApiError contains response body and status code. Body is a comment if response contains this field
type CodeforcesApiError struct {
	Body       string
	StatusCode int
}

func (e CodeforcesApiError) Error() string {
	return e.Body
}

// CodeforcesSession manage the access for the Codeforces API
type CodeforcesSession struct {
	key    string
	secret string
	client *resty.Client
}

// GlobalCodeforcesSession can be used for unauthorized access
var GlobalCodeforcesSession CodeforcesSession

func init() {
	GlobalCodeforcesSession = *NewCodeforcesSession("", "")
}

// NewCodeforcesSession creates new *CodeforcesSession with given key and secret
func NewCodeforcesSession(key string, secret string) *CodeforcesSession {
	return &CodeforcesSession{
		key:    key,
		secret: secret,
		client: resty.New(),
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

func (session *CodeforcesSession) makeQuery(methodName string, parameters url.Values) (map[string]any, error) {
	if session.key != "" {
		parameters.Set("apiKey", session.key)
	}
	parameters.Set("time", strconv.FormatInt(time.Now().Unix(), 10))
	if session.key != "" {
		parameters.Set("apiSig", session.generateApiSig(methodName, parameters))
	}

	query := session.client.R()
	query.SetQueryParamsFromValues(parameters)

	response, err := query.Get(fmt.Sprintf("%v/%v", baseUrl, methodName))
	if err != nil {
		return nil, err
	}

	var result any
	decoder := json.NewDecoder(bytes.NewReader(response.Body()))
	if err := decoder.Decode(&result); err != nil {
		if response.StatusCode() == http.StatusNotFound {
			// wrong request
			return nil, CodeforcesApiError{
				Body:       fmt.Sprintf("Can't find method %v", methodName),
				StatusCode: response.StatusCode(),
			}
		}
		return nil, err
	}
	if response.StatusCode() != http.StatusOK {
		if cur, ok := result.(map[string]any); ok {
			return nil, CodeforcesApiError{
				Body:       cur["comment"].(string), // for better comment
				StatusCode: response.StatusCode(),
			}
		}
		return nil, CodeforcesApiError{
			Body:       bytes.NewBuffer(response.Body()).String(),
			StatusCode: response.StatusCode(),
		}
	}
	return result.(map[string]any), nil
}

// BlogEntryComments implements *blogEntry.comments*
func (session *CodeforcesSession) BlogEntryComments(blogEntryId int) ([]*codeforces_objects.Comment, error) {
	parameters := url.Values{}
	parameters.Set("blogEntryId", strconv.Itoa(blogEntryId))
	res, err := session.makeQuery("blogEntry.comments", parameters)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.Comment
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// BlogEntryView implements *blogEntry.view*
func (session *CodeforcesSession) BlogEntryView(blogEntryId int) (*codeforces_objects.BlogEntry, error) {
	parameters := url.Values{}
	parameters.Set("blogEntryId", strconv.Itoa(blogEntryId))
	res, err := session.makeQuery("blogEntry.view", parameters)
	if err != nil {
		return nil, err
	}
	var result *codeforces_objects.BlogEntry
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ContestHacks implements *contest.hacks*
//
// asManager is optional
func (session *CodeforcesSession) ContestHacks(contestId int, asManager *bool) ([]*codeforces_objects.Hack, error) {
	p := url.Values{}
	p.Set("contestId", strconv.Itoa(contestId))
	if asManager != nil {
		p.Set("asManager", strconv.FormatBool(*asManager))
	}
	res, err := session.makeQuery("contest.hacks", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.Hack
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ContestList implements *contest.list*
//
// gym is optional
func (session *CodeforcesSession) ContestList(gym *bool) ([]*codeforces_objects.Contest, error) {
	p := url.Values{}
	if gym != nil {
		p.Set("gym", strconv.FormatBool(*gym))
	}
	res, err := session.makeQuery("contest.list", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.Contest
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ContestRatingChanges implements *contest.ratingChanges*
func (session *CodeforcesSession) ContestRatingChanges(contestId int) ([]*codeforces_objects.RatingChange, error) {
	p := url.Values{}
	p.Set("contestId", strconv.Itoa(contestId))
	res, err := session.makeQuery("contest.ratingChanges", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.RatingChange
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ContestStandings implements *contest.standings*
//
// asManager, from, count, handles, room, showUnofficial are optional
func (session *CodeforcesSession) ContestStandings(
	contestId int,
	asManager *bool,
	from *int,
	count *int,
	handles *string,
	room *int,
	showUnofficial *bool) (*codeforces_objects.ContestStandings, error) {
	p := url.Values{}
	p.Set("contestId", strconv.Itoa(contestId))
	if asManager != nil {
		p.Set("asManager", strconv.FormatBool(*asManager))
	}
	if from != nil {
		p.Set("from", strconv.Itoa(*from))
	}
	if count != nil {
		p.Set("count", strconv.Itoa(*count))
	}
	if handles != nil {
		p.Set("handles", *handles)
	}
	if room != nil {
		p.Set("room", strconv.Itoa(*room))
	}
	if showUnofficial != nil {
		p.Set("showUnofficial", strconv.FormatBool(*showUnofficial))
	}
	res, err := session.makeQuery("contest.standings", p)
	if err != nil {
		return nil, err
	}
	var result *codeforces_objects.ContestStandings
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ContestStatus implements *contest.status*
//
// asManager, handle, from, count are optional
func (session *CodeforcesSession) ContestStatus(
	contestId int,
	asManager *bool,
	handle *string,
	from *int,
	count *int) ([]*codeforces_objects.Submission, error) {

	p := url.Values{}
	p.Set("contestId", strconv.Itoa(contestId))
	if asManager != nil {
		p.Set("asManager", strconv.FormatBool(*asManager))
	}
	if handle != nil {
		p.Add("handle", *handle)
	}
	if from != nil {
		p.Set("from", strconv.Itoa(*from))
	}
	if count != nil {
		p.Set("count", strconv.Itoa(*count))
	}
	res, err := session.makeQuery("contest.standings", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.Submission
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ProblemsetProblems implements *problemset.problems*
//
// tags, problemsetName are optional
func (session *CodeforcesSession) ProblemsetProblems(
	tags *string,
	problemsetName *string,
) ([]*codeforces_objects.Problem, []*codeforces_objects.ProblemStatistics, error) {
	p := url.Values{}
	if tags != nil {
		p.Set("tags", *tags)
	}
	if problemsetName != nil {
		p.Add("problemsetName", *problemsetName)
	}
	res, err := session.makeQuery("problemset.problems", p)
	if err != nil {
		return nil, nil, err
	}
	var problems []*codeforces_objects.Problem
	err = mapstructure.Decode(res["result"].(map[string]any)["problems"], &problems)
	if err != nil {
		return nil, nil, err
	}
	var stats []*codeforces_objects.ProblemStatistics
	err = mapstructure.Decode(res["result"].(map[string]any)["problemStatistics"], &stats)
	if err != nil {
		return nil, nil, err
	}
	return problems, stats, nil
}

// ProblemsetRecentStatus implements *problemset.recentStatus*
//
// problemsetName is optional
func (session *CodeforcesSession) ProblemsetRecentStatus(
	count int,
	problemsetName *string,
) ([]*codeforces_objects.Submission, error) {
	p := url.Values{}
	p.Set("count", strconv.Itoa(count))
	if problemsetName != nil {
		p.Add("problemsetName", *problemsetName)
	}
	res, err := session.makeQuery("problemset.recentStatus", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.Submission
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RecentActions implements *recentActions*
func (session *CodeforcesSession) RecentActions(maxCount int) ([]*codeforces_objects.RecentAction, error) {
	p := url.Values{}
	p.Set("maxCount", strconv.Itoa(maxCount))
	res, err := session.makeQuery("recentActions", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.RecentAction
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserBlogEntries implements *user.blogEntries*
func (session *CodeforcesSession) UserBlogEntries(handle string) ([]*codeforces_objects.BlogEntry, error) {
	p := url.Values{}
	p.Set("handle", handle)
	res, err := session.makeQuery("recentActions", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.BlogEntry
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserFriends implements *user.friends*
//
// onlyOnline is optional
func (session *CodeforcesSession) UserFriends(onlyOnline *bool) ([]string, error) {
	p := url.Values{}
	if onlyOnline != nil {
		p.Set("onlyOnline", strconv.FormatBool(*onlyOnline))
	}
	res, err := session.makeQuery("user.friends", p)
	if err != nil {
		return nil, err
	}
	var result []string
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserInfo implements *user.info*
func (session *CodeforcesSession) UserInfo(handles string) ([]*codeforces_objects.User, error) {
	p := url.Values{}
	p.Set("handles", handles)
	res, err := session.makeQuery("user.info", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.User
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserRatedList implements *user.ratedList*
//
// activeOnly, includeRetires, contestId are optional
func (session *CodeforcesSession) UserRatedList(
	activeOnly *bool,
	includeRetired *bool,
	contestId *int,
) ([]*codeforces_objects.User, error) {

	p := url.Values{}
	if activeOnly != nil {
		p.Set("activeOnly", strconv.FormatBool(*activeOnly))
	}
	if includeRetired != nil {
		p.Set("activeOnly", strconv.FormatBool(*includeRetired))
	}
	if contestId != nil {
		p.Set("activeOnly", strconv.Itoa(*contestId))
	}
	res, err := session.makeQuery("user.ratedList", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.User
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserRating implements *user.rating*
func (session *CodeforcesSession) UserRating(handle string) ([]*codeforces_objects.RatingChange, error) {
	p := url.Values{}
	p.Set("handle", handle)
	res, err := session.makeQuery("user.ratedList", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.RatingChange
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UserStatus implements *user.status*
//
// from, count are optional
func (session *CodeforcesSession) UserStatus(
	handle string,
	from *int,
	count *int,
) ([]*codeforces_objects.Submission, error) {
	p := url.Values{}
	p.Set("handle", handle)
	if from != nil {
		p.Set("from", strconv.Itoa(*from))
	}
	if count != nil {
		p.Set("count", strconv.Itoa(*count))
	}
	res, err := session.makeQuery("user.status", p)
	if err != nil {
		return nil, err
	}
	var result []*codeforces_objects.Submission
	err = mapstructure.Decode(res["result"], &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
