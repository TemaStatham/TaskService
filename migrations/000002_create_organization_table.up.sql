CREATE TABLE "organization" (
    "id" SERIAL PRIMARY KEY,
    "email" VARCHAR(150) UNIQUE NOT NULL,
    "phone" VARCHAR(18) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "inn" VARCHAR(12) UNIQUE NOT NULL,
    "legal_address" VARCHAR(255) NOT NULL,
    "actual_address" VARCHAR(255) NOT NULL,
    "status_id" INTEGER NOT NULL,
    "full_name_owner" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);