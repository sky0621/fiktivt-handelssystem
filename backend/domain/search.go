package domain

type PageInfoModel struct {
	StartCursor string
	EndCursor   string
	HasNextPage bool
	HasPrevPage bool
}

type SearchWordConditionModel struct {
	SearchWord   *string
	PatternMatch *PatternMatch
}

type PatternMatch int

const (
	ExactMatch PatternMatch = iota + 1
	PartialMatch
	ForwardMatch
	BackwardMatch
)
