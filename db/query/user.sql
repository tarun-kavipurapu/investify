-- name: CreateUser :one
INSERT INTO
    bk_users (
        user_email,
        user_password,
        user_phone_number,
        users_role_id,
        users_photo_link
    )
VALUES
    ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUserById :one
SELECT
    * FROM bk_users where user_id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT
    * FROM bk_users where user_email = $1 LIMIT 1;

-- name: UpdateUserPhotoURL :exec
UPDATE bk_users
SET users_photo_link = $2
WHERE user_id = $1;

-- name: RemoveUserPhotoLink :exec
UPDATE bk_users
SET users_photo_link = NULL
WHERE user_id = $1;
