-- name: CreateAddress :one
INSERT INTO
    bk_address (
        address_street,
        address_city,
        address_state,
        address_country,
        address_zipcode 
    )
VALUES
    ($1, $2, $3, $4, $5) RETURNING *;


-- name: GetAddressById :one
SELECT
    * FROM bk_address where address_id  = $1;