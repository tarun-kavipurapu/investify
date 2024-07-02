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
        "owner_user_id" BIGINT NOT NULL REFERENCES "bk_users" ("user_id"),
        "owner_address_id" BIGINT NOT NULL REFERENCES "bk_address" ("address_id"),
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP
    );

CREATE TABLE
    "bk_business" (
        "business_id" BIGSERIAL PRIMARY KEY,
        "business_owner_id" BIGINT NOT NULL REFERENCES "bk_owner" ("owner_id"),
        "business_owner_firstname" VARCHAR NOT NULL,
        "business_owner_lastname" VARCHAR NOT NULL,
        "business_email" VARCHAR NOT NULL,
        "business_contact" VARCHAR NOT NULL,
        "business_name" VARCHAR NOT NULL,
        "business_address_id" BIGINT NOT NULL,
        "business_ratings" NUMERIC,
        "business_minamount" NUMERIC,
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP
    );