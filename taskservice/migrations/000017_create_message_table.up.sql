CREATE TABLE "message" (
    "id" SERIAL PRIMARY KEY,
    "chat_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "message" TEXT NOT NULL,
    "is_read" BOOLEAN DEFAULT false,
    "created_at" TIMESTAMP DEFAULT (NOW())
);