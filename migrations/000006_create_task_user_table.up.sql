CREATE TABLE "task_user" (
        "id" SERIAL PRIMARY KEY,
        "task_id" INTEGER NOT NULL,
        "user_id" INTEGER NOT NULL,
        "is_coordinator" BOOLEAN DEFAULT false
);