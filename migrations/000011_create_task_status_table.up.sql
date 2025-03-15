CREATE TABLE "task_status" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255)
);

INSERT INTO "task_status" (name) VALUES
                                 ('Выполнено'),
                                 ('Не выполнено'),
                                 ('В работе');
