package mysql

import (
	"context"
	"database/sql"
)

type paginQuery struct {
	DB         *sql.DB
	LimitCount int64
	PageCount  int64
}

type IPaginQuery interface {
	Limit(limit int64) IPaginQuery
	Page(page int64) IPaginQuery

	QueryPaginate(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

func New(query *sql.DB) IPaginQuery {
	return &paginQuery{
		DB: query,
	}
}

// Limit is to add limit for pagination
func (paging *paginQuery) Limit(limit int64) IPaginQuery {
	if limit < 1 {
		paging.LimitCount = 10
	} else {
		paging.LimitCount = limit
	}
	return paging
}

// Page is to specify which page to serve in mongo paginated result
func (paging *paginQuery) Page(page int64) IPaginQuery {
	if page < 1 {
		paging.PageCount = 1
	} else {
		paging.PageCount = page
	}
	return paging
}

func (paging *paginQuery) QueryPaginate(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := paging.DB.QueryContext(ctx, query, args)
	return rows, err
}
