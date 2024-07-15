-- name: CreateBusiness :one
INSERT INTO
    bk_business (
        business_owner_id,
        business_owner_firstname,
        business_owner_lastname,
        business_domain_code,
        business_state_code,
        business_email,
        business_contact,
        business_name,
        business_address_id,
        business_ratings,
        business_investment_amount
    )
VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10,$11)
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


-- name: GetFilteredBusinesses :many
SELECT business_id, business_owner_id, business_domain_code, business_state_code, business_owner_firstname, business_owner_lastname, business_email, business_contact, business_name, business_address_id, business_ratings, business_investment_amount, created_at, updated_at, deleted_at
FROM bk_business
WHERE
    (business_domain_code = $1 OR $1 IS NULL)
    AND (business_state_code = $2 OR $2 IS NULL)
    AND (business_investment_amount >= $3 OR $3 IS NULL)
    AND (business_investment_amount <= $4 OR $4 IS NULL);
