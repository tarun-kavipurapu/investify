// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: investor.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createInvestor = `-- name: CreateInvestor :one
INSERT INTO
    bk_investor (
        investor_name,
        investor_user_id,
        investor_address_id
    )
VALUES
    ($1, $2, $3) RETURNING investor_id, investor_name, investor_user_id, investor_address_id, created_at, updated_at, deleted_at
`

type CreateInvestorParams struct {
	InvestorName      pgtype.Text `json:"investor_name"`
	InvestorUserID    int64       `json:"investor_user_id"`
	InvestorAddressID int64       `json:"investor_address_id"`
}

func (q *Queries) CreateInvestor(ctx context.Context, arg CreateInvestorParams) (BkInvestor, error) {
	row := q.db.QueryRow(ctx, createInvestor, arg.InvestorName, arg.InvestorUserID, arg.InvestorAddressID)
	var i BkInvestor
	err := row.Scan(
		&i.InvestorID,
		&i.InvestorName,
		&i.InvestorUserID,
		&i.InvestorAddressID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getInvestorById = `-- name: GetInvestorById :one
SELECT
    investor_id, investor_name, investor_user_id, investor_address_id, created_at, updated_at, deleted_at FROM bk_investor where investor_id = $1
`

func (q *Queries) GetInvestorById(ctx context.Context, investorID int64) (BkInvestor, error) {
	row := q.db.QueryRow(ctx, getInvestorById, investorID)
	var i BkInvestor
	err := row.Scan(
		&i.InvestorID,
		&i.InvestorName,
		&i.InvestorUserID,
		&i.InvestorAddressID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getInvestorByUserId = `-- name: GetInvestorByUserId :one
SELECT
    investor_id, investor_name, investor_user_id, investor_address_id, created_at, updated_at, deleted_at FROM bk_investor where investor_user_id = $1
`

func (q *Queries) GetInvestorByUserId(ctx context.Context, investorUserID int64) (BkInvestor, error) {
	row := q.db.QueryRow(ctx, getInvestorByUserId, investorUserID)
	var i BkInvestor
	err := row.Scan(
		&i.InvestorID,
		&i.InvestorName,
		&i.InvestorUserID,
		&i.InvestorAddressID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getInvestorFeed = `-- name: GetInvestorFeed :many
SELECT
    investor_id, investor_name, investor_user_id, investor_address_id, created_at, updated_at, deleted_at FROM bk_investor LIMIT 10
`

func (q *Queries) GetInvestorFeed(ctx context.Context) ([]BkInvestor, error) {
	rows, err := q.db.Query(ctx, getInvestorFeed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BkInvestor{}
	for rows.Next() {
		var i BkInvestor
		if err := rows.Scan(
			&i.InvestorID,
			&i.InvestorName,
			&i.InvestorUserID,
			&i.InvestorAddressID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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