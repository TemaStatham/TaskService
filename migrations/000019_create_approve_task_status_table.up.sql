CREATE TABLE "approve_task_status" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255)
);

INSERT INTO "approve_task_status" ("name") VALUES
                                           ('На рассмотрении'),
                                           ('Принято'),
                                           ('Отказано');