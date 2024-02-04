package codeforces_api

import (
	"net/url"
	"testing"
)

func TestCodeforcesSession_generateApiSig(t *testing.T) {
	session := GetGlobalSession()
	expectedSig := "0000005f1f3de78163894a44cb12d9f1b6582fe134443c55e" +
		"67555db27327b161ece5cb57be1bcf449f40ccb9dd86973d9ebf193d2f52178d14062f9502dbd3a881f70"
	if got := session.generateApiSig("contest.hacks", make(url.Values)); got != expectedSig {
		t.Errorf("Expected ApiSig %v, got %v", expectedSig, got)
	}

	params := url.Values{}
	params["bruh"] = append(params["bruh"], "aboba")
	params["bruh"] = append(params["bruh"], "bruh")
	params["123"] = append(params["bruh"], "456")
	expectedSig = "0000001f08193c237cf90122542c4e6542480601b764dff72c2b" +
		"ce3fc691c2f5a8a77ab83f984995722af18ce6e5a6d6d4a4bdc05d69666297099f110676db224134de"
	if got := session.generateApiSig("method", params); got != expectedSig {
		t.Errorf("Expected ApiSig %v, got %v", expectedSig, got)
	}

	expectedSig = "0000007b9a1f22a5d3b0f0c58a10684e06f71afc27d517a26bcb" +
		"14a26b85865919e67927f21c492b05f83564c401d0937f64c9f95e767541aebcaace428eab59580663"
	if got := session.generateApiSig("", make(url.Values)); got != expectedSig {
		t.Errorf("Expected ApiSig %v, got %v", expectedSig, got)
	}
}
