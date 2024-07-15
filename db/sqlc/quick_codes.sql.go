// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: quick_codes.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAllDomains = `-- name: GetAllDomains :many
SELECT DISTINCT name, code, description
FROM bk_quick_codes
WHERE type = 'DOMAIN'
`

type GetAllDomainsRow struct {
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Description pgtype.Text `json:"description"`
}

func (q *Queries) GetAllDomains(ctx context.Context) ([]GetAllDomainsRow, error) {
	rows, err := q.db.Query(ctx, getAllDomains)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllDomainsRow{}
	for rows.Next() {
		var i GetAllDomainsRow
		if err := rows.Scan(&i.Name, &i.Code, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllStates = `-- name: GetAllStates :many
SELECT DISTINCT name, code, description
FROM bk_quick_codes
WHERE type = 'STATE'
`

type GetAllStatesRow struct {
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Description pgtype.Text `json:"description"`
}

func (q *Queries) GetAllStates(ctx context.Context) ([]GetAllStatesRow, error) {
	rows, err := q.db.Query(ctx, getAllStates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllStatesRow{}
	for rows.Next() {
		var i GetAllStatesRow
		if err := rows.Scan(&i.Name, &i.Code, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getQuickCodesByType = `-- name: GetQuickCodesByType :many
SELECT quick_code_id, type, name, code, description, created_at, updated_at
FROM bk_quick_codes
WHERE type = $1
`

func (q *Queries) GetQuickCodesByType(ctx context.Context, type_ string) ([]BkQuickCode, error) {
	rows, err := q.db.Query(ctx, getQuickCodesByType, type_)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BkQuickCode{}
	for rows.Next() {
		var i BkQuickCode
		if err := rows.Scan(
			&i.QuickCodeID,
			&i.Type,
			&i.Name,
			&i.Code,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}