CREATE TABLE "task_type" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255)
);

INSERT INTO "task_type" (name) VALUES
                                 ('Открытый'),
                                 ('Закрытый');
