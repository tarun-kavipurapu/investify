CREATE TABLE
    "bk_business_images" (
        "image_id" BIGSERIAL PRIMARY KEY,
        "business_id" BIGINT NOT NULL REFERENCES "bk_business" ("business_id"),
        "image_url" TEXT NOT NULL,
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );