CREATE TABLE "approve_file" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "file_id" INTEGER NOT NULL,
    "approve_task_id" INTEGER NOT NULL
);