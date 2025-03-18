CREATE TABLE "response" (
        "id" SERIAL PRIMARY KEY,
        "task_id" INTEGER NOT NULL,
        "user_id" INTEGER NOT NULL,
        "status_id" INTEGER DEFAULT 1
);