package controller

import (
	"github.com/sky0621/fiktivt-handelssystem/adapter/controller/model"
	"github.com/sky0621/fiktivt-handelssystem/domain"
)

func ToSearchWordConditionModel(input *model.SearchWordCondition) *domain.SearchWordConditionModel {
	if input == nil {
		return nil
	}
	pm := domain.ExactMatch
	if input.PatternMatch != nil {
		switch *input.PatternMatch {
		case model.PatternMatchPartialMatch:
			pm = domain.PartialMatch
		case model.PatternMatchForwardMatch:
			pm = domain.ForwardMatch
		case model.PatternMatchBackwardMatch:
			pm = domain.BackwardMatch
		}
	}
	return &domain.SearchWordConditionModel{
		SearchWord:   input.SearchWord,
		PatternMatch: pm,
	}
}

func ToSortConditionModel(input *model.SortCondition) *domain.SortConditionModel {
	if input == nil {
		return nil
	}
	odr := domain.Desc
	if input.SortOrder == model.SortOrderAsc {
		odr = domain.Asc
	}
	return &domain.SortConditionModel{
		SortKey:   input.SortKey,
		SortOrder: odr,
	}
}

func ToSearchDirectionType(input model.SearchDirection) domain.SearchDirection {
	switch input {
	case model.SearchDirectionPrev:
		return domain.Prev
	case model.SearchDirectionNext:
		return domain.Next
	}
	return domain.None
}
