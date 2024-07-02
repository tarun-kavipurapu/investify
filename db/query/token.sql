-- name: CreateToken :one
INSERT INTO
    bk_tokens (
        token_value,
        token_user_id,
        token_expiry_date
    )
VALUES
    ($1, $2, $3) RETURNING *;