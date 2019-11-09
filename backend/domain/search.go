package domain

type PageInfoModel struct {
	StartCursor string
	EndCursor   string
	HasNextPage bool
	HasPrevPage bool
}

type SearchWordConditionModel struct {
	SearchWord   string
	PatternMatch PatternMatch
}

type PatternMatch int

const (
	ExactMatch PatternMatch = iota + 1
	PartialMatch
	ForwardMatch
	BackwardMatch
)

type SortConditionModel struct {
	SortKey   string
	SortOrder SortOrder
}

type SortOrder int

const (
	Asc SortOrder = iota + 1
	Desc
)

type SearchDirection int

const (
	None SearchDirection = iota + 1
	Prev
	Next
)
