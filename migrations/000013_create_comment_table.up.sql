CREATE TABLE "comment" (
    "id" SERIAL PRIMARY KEY,
    "task_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "comment" TEXT NOT NULL,
    "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);