// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: business_images.sql

package db

import (
	"context"
)

const deleteByImageId = `-- name: DeleteByImageId :exec
DELETE FROM
    bk_business_images
WHERE
    image_id = $1
`

func (q *Queries) DeleteByImageId(ctx context.Context, imageID int64) error {
	_, err := q.db.Exec(ctx, deleteByImageId, imageID)
	return err
}

const getImageByBusinessId = `-- name: GetImageByBusinessId :many
SELECT
    image_id, business_id, image_url, created_at, updated_at
FROM
    bk_business_images
WHERE
    business_id = $1
`

func (q *Queries) GetImageByBusinessId(ctx context.Context, businessID int64) ([]BkBusinessImage, error) {
	rows, err := q.db.Query(ctx, getImageByBusinessId, businessID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BkBusinessImage{}
	for rows.Next() {
		var i BkBusinessImage
		if err := rows.Scan(
			&i.ImageID,
			&i.BusinessID,
			&i.ImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getImageByImageId = `-- name: GetImageByImageId :one
SELECT
    image_id, business_id, image_url, created_at, updated_at
FROM
    bk_business_images
WHERE
    image_id = $1
`

func (q *Queries) GetImageByImageId(ctx context.Context, imageID int64) (BkBusinessImage, error) {
	row := q.db.QueryRow(ctx, getImageByImageId, imageID)
	var i BkBusinessImage
	err := row.Scan(
		&i.ImageID,
		&i.BusinessID,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const uploadBusinessImage = `-- name: UploadBusinessImage :one
INSERT INTO
    bk_business_images (
        business_id,
        image_url
    )
VALUES
    ($1, $2) RETURNING image_id, business_id, image_url, created_at, updated_at
`

type UploadBusinessImageParams struct {
	BusinessID int64  `json:"business_id"`
	ImageUrl   string `json:"image_url"`
}

func (q *Queries) UploadBusinessImage(ctx context.Context, arg UploadBusinessImageParams) (BkBusinessImage, error) {
	row := q.db.QueryRow(ctx, uploadBusinessImage, arg.BusinessID, arg.ImageUrl)
	var i BkBusinessImage
	err := row.Scan(
		&i.ImageID,
		&i.BusinessID,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
