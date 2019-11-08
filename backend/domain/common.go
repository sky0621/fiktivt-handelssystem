package domain

type PageInfoModel struct {
	StartCursor string
	EndCursor   string
	HasNextPage bool
	HasPrevPage bool
}
