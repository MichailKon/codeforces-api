package utils

import (
	"net/url"
	"testing"
)

func checkStringArrays(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func checkMaps(a, b map[string][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		val1, ok := b[key]
		if !ok || !checkStringArrays(val1, value) {
			return false
		}
	}
	return true
}

func TestContestStandingsParams_FillUrlValues(t *testing.T) {
	params := NewContestStandingsParams().
		WithShowUnofficial(true).
		WithCount(5).
		WithHandles("a;b;c")
	out := url.Values{}
	params.FillUrlValues(&out)
	expected := make(map[string][]string)
	expected["from"] = append(expected["from"], "5")
	expected["handles"] = append(expected["handles"], "a;b;c")
	expected["showUnofficial"] = append(expected["showUnofficial"], "true")

	if !checkMaps(expected, out) {
		t.Errorf("Expected %v; got %v", expected, out)
	}

	params = NewContestStandingsParams()
	out = url.Values{}
	params.FillUrlValues(&out)
	if len(out) != 0 {
		t.Errorf("Expected empty params, got %v", out)
	}
}

func TestContestStandingsParams_WithAsManager(t *testing.T) {
	params := NewContestStandingsParams().WithAsManager(true)
	if val, err := params.asManager.Get(); err != nil || !val {
		t.Errorf("Expected true in asManager; got error/false: %v; %v", val, err)
	}
}

func TestContestStandingsParams_WithCount(t *testing.T) {
	params := NewContestStandingsParams().WithCount(5)
	if val, err := params.count.Get(); err != nil || val != 5 {
		t.Errorf("Expected 5 in count; got error/wrong value: %v; %v", val, err)
	}
}

func TestContestStandingsParams_WithFrom(t *testing.T) {
	params := NewContestStandingsParams().WithFrom(5)
	if val, err := params.from.Get(); err != nil || val != 5 {
		t.Errorf("Expected 5 in from; got error/wrong value: %v; %v", val, err)
	}
}

func TestContestStandingsParams_WithHandles(t *testing.T) {
	params := NewContestStandingsParams().WithHandles("a;b;c")
	if val, err := params.handles.Get(); err != nil || val != "a;b;c" {
		t.Errorf("Expected \"a;b;c\" in handles; got error/wrong value: %v; %v", val, err)
	}
}

func TestContestStandingsParams_WithHandlesList(t *testing.T) {
	params := NewContestStandingsParams().WithHandlesList([]string{"a", "b", "c"})
	if val, err := params.handles.Get(); err != nil || val != "a;b;c" {
		t.Errorf("Expected \"a;b;c\" in handles; got error/wrong value: %v; %v", val, err)
	}
}

func TestContestStandingsParams_WithShowUnofficial(t *testing.T) {
	params := NewContestStandingsParams().WithShowUnofficial(true)
	if val, err := params.showUnofficial.Get(); err != nil || !val {
		t.Errorf("Expected true in showUnofficial; got error/false: %v; %v", val, err)
	}
}

func TestContestStatusParams_FillUrlValues(t *testing.T) {
	params := NewContestStatusParams().
		WithAsManager(true).
		WithFrom(1)
	out := url.Values{}
	params.FillUrlValues(&out)
	expected := make(map[string][]string)
	expected["asManager"] = append(expected["asManager"], "true")
	expected["from"] = append(expected["from"], "1")

	if !checkMaps(expected, out) {
		t.Errorf("Expected %v; got %v", expected, out)
	}

	params = NewContestStatusParams()
	out = url.Values{}
	params.FillUrlValues(&out)
	if len(out) != 0 {
		t.Errorf("Expected empty params, got %v", out)
	}
}

func TestContestStatusParams_WithAsManager(t *testing.T) {
	params := NewContestStatusParams().WithAsManager(true)
	if val, err := params.asManager.Get(); err != nil || !val {
		t.Errorf("Expected true in asManager; got error/false: %v; %v", val, err)
	}
}

func TestContestStatusParams_WithCount(t *testing.T) {
	params := NewContestStatusParams().WithCount(5)
	if val, err := params.count.Get(); err != nil || val != 5 {
		t.Errorf("Expected 5 in count; got error/wrong value: %v; %v", val, err)
	}
}

func TestContestStatusParams_WithFrom(t *testing.T) {
	params := NewContestStatusParams().WithFrom(5)
	if val, err := params.from.Get(); err != nil || val != 5 {
		t.Errorf("Expected 5 in from; got error/wrong value: %v; %v", val, err)
	}
}

func TestContestStatusParams_WithHandle(t *testing.T) {
	params := NewContestStatusParams().WithHandle("a")
	if val, err := params.handle.Get(); err != nil || val != "a" {
		t.Errorf("Expected \"a\" in handle; got error/wrong value: %v; %v", val, err)
	}
}
