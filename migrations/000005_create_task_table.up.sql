CREATE TABLE "task" (
        "id" SERIAL PRIMARY KEY,
        "organization_id" INTEGER NOT NULL,
        "name" VARCHAR(255) NOT NULL,
        "type_id" INTEGER NOT NULL,
        "description" TEXT NOT NULL,
        "location" VARCHAR(255) NOT NULL,
        "task_date" TIMESTAMP NOT NULL,
        "participants_count" INTEGER,
        "max_score" INTEGER,
        "status_id" INTEGER DEFAULT 1,
        "created_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
        "updated_at" TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
);