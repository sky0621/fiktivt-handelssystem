package domain

import (
	"fmt"
	"testing"
)

func TestSearchWordConditionModel_SearchWord(t *testing.T) {
	actual := &SearchWordConditionModel{
		SearchWord: "dummy",
	}
	if actual.SearchWord != "dummy" {
		t.Fatalf("SearchWord[%s] is not 'dummy'", actual.SearchWord)
	}
}

func TestSearchWordConditionModel_PatternMatch(t *testing.T) {
	factors := []PatternMatch{
		ExactMatch,
		PartialMatch,
		ForwardMatch,
		BackwardMatch,
	}
	for _, f := range factors {
		actual := &SearchWordConditionModel{
			PatternMatch: f,
		}
		fmt.Printf("%#v\n", actual)
		if f != actual.PatternMatch {
			t.Fatalf("PatternMatch[%v] is not '%v'", actual.PatternMatch, f)
		}
	}
}

func TestSortConditionModel_SortKey(t *testing.T) {
	actual := &SortConditionModel{
		SortKey: "dummy",
	}
	if actual.SortKey != "dummy" {
		t.Fatalf("SortKey[%s] is not 'dummy'", actual.SortKey)
	}
}

func TestSortConditionModel_SortOrder(t *testing.T) {
	factors := []SortOrder{
		Asc,
		Desc,
	}
	for _, f := range factors {
		actual := &SortConditionModel{
			SortOrder: f,
		}
		fmt.Printf("%#v\n", actual)
		if f != actual.SortOrder {
			t.Fatalf("SortOrder[%v] is not '%v'", actual.SortOrder, f)
		}
	}
}
