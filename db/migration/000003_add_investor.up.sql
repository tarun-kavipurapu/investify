CREATE TABLE
    "bk_investor" (
        "investor_id" BIGSERIAL PRIMARY KEY,
        "investor_name" VARCHAR(255),
        "investor_user_id" BIGINT NOT NULL REFERENCES "bk_users" ("user_id"),
        "investor_address_id" BIGINT NOT NULL REFERENCES "bk_address" ("address_id"),
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP
    );