CREATE TABLE "approve_task" (
    "id" SERIAL PRIMARY KEY,
    "task_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "status_id" INTEGER DEFAULT 1,
    "score" INTEGER DEFAULT 0,
    "approved" INTEGER,
    "created_at" TIMESTAMP DEFAULT (NOW())
);