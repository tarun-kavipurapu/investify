-- name: CreateInvestor :one
INSERT INTO
    bk_investor (
        investor_name,
        investor_user_id,
        investor_address_id
    )
VALUES
    ($1, $2, $3) RETURNING *;


-- name: GetInvestorByUserId :one
SELECT
    * FROM bk_investor where investor_user_id = $1;

-- name: GetInvestorById :one
SELECT
    * FROM bk_investor where investor_id = $1;

-- name: GetInvestorFeed :many
SELECT
    * FROM bk_investor LIMIT 10;