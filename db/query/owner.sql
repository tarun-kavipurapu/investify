-- name: CreateOwner :one
INSERT INTO
    bk_owner (
        owner_name,
        owner_user_id,
        owner_address_id
    )
VALUES
    ($1, $2, $3) RETURNING *;

-- name: GetOwnerByUserId :one
SELECT
    * FROM bk_owner where owner_user_id = $1 LIMIT 1;