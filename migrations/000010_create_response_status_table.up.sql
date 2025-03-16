CREATE TABLE "response_status" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255)
);

INSERT INTO "response_status" ("name") VALUES
    ('На рассмотрении'),
    ('Принято'),
    ('Отказано');
