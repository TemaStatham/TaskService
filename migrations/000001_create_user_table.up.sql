CREATE TABLE "user" (
    "id" SERIAL PRIMARY KEY,
    "email" VARCHAR(150) UNIQUE NOT NULL,
    "phone" VARCHAR(18) NOT NULL,
    "surname" VARCHAR(255),
    "name" VARCHAR(255) NOT NULL,
    "patronymic" VARCHAR(255),
    "date_of_birthday" DATE,
    "address" VARCHAR(255),
    "password_hash" TEXT NOT NULL,
    "is_admin" BOOLEAN DEFAULT false,
    "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
    "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);