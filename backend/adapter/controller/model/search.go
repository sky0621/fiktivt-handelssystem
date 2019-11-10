package model

import (
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
)

type PageInfo struct {
	StartCursor *string `json:"startCursor"`
	EndCursor   *string `json:"endCursor"`
	HasPrevPage bool    `json:"hasPrevPage"`
	HasNextPage bool    `json:"hasNextPage"`
}

type BaseCondition struct {
	SearchWordCondition *SearchWordCondition `json:"searchWordCondition"`
	SortCondition       *SortCondition       `json:"sortCondition"`
	SearchDirection     SearchDirection      `json:"searchDirection"`
	Limit               *int                 `json:"limit"`
	StartCursor         *string              `json:"startCursor"`
	EndCursor           *string              `json:"endCursor"`
}

type SearchWordCondition struct {
	SearchWord   string        `json:"searchWord"`
	PatternMatch *PatternMatch `json:"patternMatch"`
}

type SortCondition struct {
	SortKey   string    `json:"sortKey"`
	SortOrder SortOrder `json:"sortOrder"`
}

type PatternMatch string

const (
	PatternMatchExactMatch    PatternMatch = "EXACT_MATCH"
	PatternMatchPartialMatch  PatternMatch = "PARTIAL_MATCH"
	PatternMatchForwardMatch  PatternMatch = "FORWARD_MATCH"
	PatternMatchBackwardMatch PatternMatch = "BACKWARD_MATCH"
)

var AllPatternMatch = []PatternMatch{
	PatternMatchExactMatch,
	PatternMatchPartialMatch,
	PatternMatchForwardMatch,
	PatternMatchBackwardMatch,
}

func (e PatternMatch) IsValid() bool {
	switch e {
	case PatternMatchExactMatch, PatternMatchPartialMatch, PatternMatchForwardMatch, PatternMatchBackwardMatch:
		return true
	}
	return false
}

func (e PatternMatch) String() string {
	return string(e)
}

func (e *PatternMatch) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PatternMatch(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PatternMatch", str)
	}
	return nil
}

func (e PatternMatch) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortOrder string

const (
	SortOrderAsc  SortOrder = "ASC"
	SortOrderDesc SortOrder = "DESC"
)

var AllSortOrder = []SortOrder{
	SortOrderAsc,
	SortOrderDesc,
}

func (e SortOrder) IsValid() bool {
	switch e {
	case SortOrderAsc, SortOrderDesc:
		return true
	}
	return false
}

func (e SortOrder) String() string {
	return string(e)
}

func (e *SortOrder) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortOrder(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortOrder", str)
	}
	return nil
}

func (e SortOrder) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SearchDirection string

const (
	SearchDirectionNone SearchDirection = "NONE"
	SearchDirectionPrev SearchDirection = "PREV"
	SearchDirectionNext SearchDirection = "NEXT"
)

var AllSearchDirection = []SearchDirection{
	SearchDirectionNone,
	SearchDirectionPrev,
	SearchDirectionNext,
}

func (e SearchDirection) IsValid() bool {
	switch e {
	case SearchDirectionNone, SearchDirectionPrev, SearchDirectionNext:
		return true
	}
	return false
}

func (e SearchDirection) String() string {
	return string(e)
}

func (e *SearchDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SearchDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SearchDirection", str)
	}
	return nil
}

func (e SearchDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func EncodeCursor(key string, val interface{}) *string {
	if key == "" {
		return nil
	}
	if val == nil {
		return nil
	}
	cursor := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s=%v", key, val)))
	return &cursor
}
