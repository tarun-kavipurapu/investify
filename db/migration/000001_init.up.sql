CREATE TABLE
    "bk_role" (
        "role_id" SERIAL PRIMARY KEY,
        "role_name" VARCHAR NOT NULL
    );

-- Either phone_number or emaail
-- CHECK (user_email IS NOT NULL OR user_phone_number IS NOT NULL)
CREATE TABLE
    "bk_users" (
        "user_id" BIGSERIAL PRIMARY KEY,
        "user_email" VARCHAR NOT NULL UNIQUE,
        "user_phone_number" VARCHAR UNIQUE,
        "user_password" VARCHAR NOT NULL,
        "users_role_id" INTEGER NOT NULL REFERENCES "bk_role" ("role_id"),
        "users_photo_link" TEXT,
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMP
    );

CREATE TABLE
    "bk_tokens" (
        "token_id" BIGSERIAL PRIMARY KEY,
        "token_value" UUID NOT NULL,
        "token_user_id" BIGINT NOT NULL REFERENCES "bk_users" ("user_id"),
        "token_expiry_date" TIMESTAMPTZ
    );

CREATE UNIQUE INDEX "user_email_index" ON "bk_users" ("user_email");

CREATE UNIQUE INDEX "token_value_index" ON "bk_tokens" ("token_value");

CREATE UNIQUE INDEX "user_phone_number_index" ON "bk_users" ("user_phone_number");

INSERT INTO
    bk_role (role_name)
VALUES
    ('owner'),
    ('investor'),
    ('admin');