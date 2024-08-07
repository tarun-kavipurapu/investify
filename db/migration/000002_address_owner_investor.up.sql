CREATE TABLE
    "bk_address" (
        "address_id" BIGSERIAL PRIMARY KEY,
        "address_street" VARCHAR NOT NULL,
        "address_city" VARCHAR NOT NULL,
        "address_state" VARCHAR NOT NULL,
        "address_country" VARCHAR,
        "address_zipcode" VARCHAR NOT NULL
    );

CREATE TABLE
    "bk_owner" (
        "owner_id" BIGSERIAL PRIMARY KEY,
        "owner_name" VARCHAR(255),
        "owner_user_id" BIGINT NOT NULL UNIQUE REFERENCES "bk_users" ("user_id"),
        "owner_address_id" BIGINT NOT NULL REFERENCES "bk_address" ("address_id"),
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP
    );

CREATE TABLE
    "bk_investor" (
        "investor_id" BIGSERIAL PRIMARY KEY,
        "investor_name" VARCHAR(255),
        "investor_user_id" BIGINT NOT NULL UNIQUE REFERENCES "bk_users" ("user_id"),
        "investor_address_id" BIGINT NOT NULL REFERENCES "bk_address" ("address_id"),
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP
    );