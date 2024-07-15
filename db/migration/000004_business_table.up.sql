CREATE TABLE
    "bk_business" (
        "business_id" BIGSERIAL PRIMARY KEY,
        "business_owner_id" BIGINT NOT NULL REFERENCES "bk_owner" ("owner_id"),
        "business_domain_code" VARCHAR(50) NOT NULL REFERENCES "bk_quick_codes" ("code"),
        "business_state_code" VARCHAR(50) NOT NULL REFERENCES "bk_quick_codes" ("code"),
        "business_owner_firstname" VARCHAR NOT NULL,
        "business_owner_lastname" VARCHAR NOT NULL,
        "business_email" VARCHAR NOT NULL,
        "business_contact" VARCHAR NOT NULL,
        "business_name" VARCHAR NOT NULL,
        "business_address_id" BIGINT NOT NULL,
        "business_ratings" NUMERIC,
        "business_investment_amount" NUMERIC,
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP
    );