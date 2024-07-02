-- name: GetImageByBusinessId :many
SELECT
    *
FROM
    bk_business_images
WHERE
    business_id = $1;

-- name: UploadBusinessImage :one
INSERT INTO
    bk_business_images (
        business_id,
        image_url
    )
VALUES
    ($1, $2) RETURNING *;


-- name: DeleteByImageId :exec
DELETE FROM
    bk_business_images
WHERE
    image_id = $1;

-- name: GetImageByImageId :one
SELECT
    *
FROM
    bk_business_images
WHERE
    image_id = $1;