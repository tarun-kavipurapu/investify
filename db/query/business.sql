-- name: CreateBusiness :one
INSERT INTO
    bk_business (
        business_owner_id,
        business_owner_firstname,
        business_owner_lastname,
        business_email,
        business_contact,
        business_name,
        business_address_id,
        business_ratings,
        business_minamount
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetBusinessByOwnerId :many
SELECT
    * FROM bk_business where business_owner_id = $1;

-- name: GetBusinessById :one
SELECT
    * FROM bk_business where business_id = $1;

-- name: GetBusinessFeed :many
SELECT
    * FROM bk_business LIMIT 10;