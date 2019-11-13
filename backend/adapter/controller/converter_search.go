package controller

import (
	"github.com/sky0621/fiktivt-handelssystem/domain"
)

func ToSearchWordConditionModel(input *SearchWordCondition) *domain.SearchWordConditionModel {
	if input == nil {
		return nil
	}
	pm := domain.ExactMatch
	if input.PatternMatch != nil {
		switch *input.PatternMatch {
		case PatternMatchPartialMatch:
			pm = domain.PartialMatch
		case PatternMatchForwardMatch:
			pm = domain.ForwardMatch
		case PatternMatchBackwardMatch:
			pm = domain.BackwardMatch
		}
	}
	return &domain.SearchWordConditionModel{
		SearchWord:   input.SearchWord,
		PatternMatch: pm,
	}
}

func ToSortConditionModel(input *SortCondition) *domain.SortConditionModel {
	if input == nil {
		return nil
	}
	odr := domain.Desc
	if input.SortOrder == SortOrderAsc {
		odr = domain.Asc
	}
	return &domain.SortConditionModel{
		SortKey:   input.SortKey,
		SortOrder: odr,
	}
}

func ToSearchDirectionType(input SearchDirection) domain.SearchDirection {
	switch input {
	case SearchDirectionPrev:
		return domain.Prev
	case SearchDirectionNext:
		return domain.Next
	}
	return domain.None
}
