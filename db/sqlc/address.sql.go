// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: address.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAddress = `-- name: CreateAddress :one
INSERT INTO
    bk_address (
        address_street,
        address_city,
        address_state,
        address_country,
        address_zipcode 
    )
VALUES
    ($1, $2, $3, $4, $5) RETURNING address_id, address_street, address_city, address_state, address_country, address_zipcode
`

type CreateAddressParams struct {
	AddressStreet  string      `json:"address_street"`
	AddressCity    string      `json:"address_city"`
	AddressState   string      `json:"address_state"`
	AddressCountry pgtype.Text `json:"address_country"`
	AddressZipcode string      `json:"address_zipcode"`
}

func (q *Queries) CreateAddress(ctx context.Context, arg CreateAddressParams) (BkAddress, error) {
	row := q.db.QueryRow(ctx, createAddress,
		arg.AddressStreet,
		arg.AddressCity,
		arg.AddressState,
		arg.AddressCountry,
		arg.AddressZipcode,
	)
	var i BkAddress
	err := row.Scan(
		&i.AddressID,
		&i.AddressStreet,
		&i.AddressCity,
		&i.AddressState,
		&i.AddressCountry,
		&i.AddressZipcode,
	)
	return i, err
}

const getAddressById = `-- name: GetAddressById :one
SELECT
    address_id, address_street, address_city, address_state, address_country, address_zipcode FROM bk_address where address_id  = $1
`

func (q *Queries) GetAddressById(ctx context.Context, addressID int64) (BkAddress, error) {
	row := q.db.QueryRow(ctx, getAddressById, addressID)
	var i BkAddress
	err := row.Scan(
		&i.AddressID,
		&i.AddressStreet,
		&i.AddressCity,
		&i.AddressState,
		&i.AddressCountry,
		&i.AddressZipcode,
	)
	return i, err
}