CREATE TABLE "category" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) UNIQUE NOT NULL
);

INSERT INTO "category" (name) VALUES
                                     ('Категория 1'),
                                     ('Категория 2'),
                                     ('Категория 3');
