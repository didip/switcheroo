package libstring

import (
	"testing"
)

func TestMatch(t *testing.T) {
	isMatched, params, err := Match("/a/b/{name}", "/a/b/didip")
	if err != nil {
		t.Fatalf("Matcher should not failed. Error: %#v", err)
	}
	if !isMatched {
		t.Fatalf("Matcher should have found the params.")
	}
	if len(params) != 1 {
		t.Fatalf("Matcher should have found the params.")
	}
	if params["name"] != "didip" {
		t.Fatalf("Matcher should have found the params. Got: %v", params["name"])
	}

	isMatched, params, err = Match("yolo", "/a/b/didip")
	if err != nil {
		t.Fatalf("Matcher should not failed. Error: %#v", err)
	}
	if isMatched {
		t.Fatalf("Matcher should have not matched.")
	}

	isMatched, params, err = Match("/a/b", "/a/b/didip")
	if err != nil {
		t.Fatalf("Matcher should not failed. Error: %#v", err)
	}
	if isMatched {
		t.Fatalf("Matcher should have not matched.")
	}

	isMatched, params, err = Match("/a/b", "/a/b")
	if err != nil {
		t.Fatalf("Matcher should not failed. Error: %#v", err)
	}
	if !isMatched {
		t.Fatalf("Matcher should have match exact string.")
	}
}

func TestMatchWithStar(t *testing.T) {
	isMatched, params, err := Match("/a/*/{name}", "/a/b/didip")
	if err != nil {
		t.Fatalf("Matcher should not failed. Error: %#v", err)
	}
	if !isMatched {
		t.Fatalf("Matcher should have found the params.")
	}
	if len(params) != 1 {
		t.Fatalf("Matcher should have found the params.")
	}
	if params["name"] != "didip" {
		t.Fatalf("Matcher should have found the params. Got: %v", params["name"])
	}

	// The star should only match 1 level deep.
	isMatched, params, err = Match("/a/*/{name}", "/a/b/c/d/e/f/didip")
	if err != nil {
		t.Fatalf("Matcher should not failed. Error: %#v", err)
	}
	if !isMatched {
		t.Fatalf("Matcher should have found the params.")
	}
	if len(params) != 1 {
		t.Fatalf("Matcher should have found the params.")
	}
	if params["name"] != "c/d/e/f/didip" {
		t.Fatalf("Matcher should have found the params. Got: %v", params["name"])
	}
}

func TestNonMatchWithStar(t *testing.T) {
	isMatched, _, err := Match("/b/*/{name}", "/a/b/didip")
	if err != nil {
		t.Fatalf("Matcher should not failed. Error: %#v", err)
	}
	if isMatched {
		t.Fatalf("Matcher should have miss.")
	}
}
