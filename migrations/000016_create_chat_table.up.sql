CREATE TABLE "chat" (
    "id" SERIAL PRIMARY KEY,
    "user1_id" INTEGER NOT NULL,
    "user2_id" INTEGER NOT NULL,
    "updated_at" TIMESTAMP DEFAULT (NOW()),
    "created_at" TIMESTAMP DEFAULT (NOW())
);