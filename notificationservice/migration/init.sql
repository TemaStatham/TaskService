CREATE TABLE "notification" (
                                "id" SERIAL PRIMARY KEY,
                                "user_id" INTEGER NOT NULL,
                                "message" TEXT NOT NULL,
                                "is_read" BOOLEAN DEFAULT false,
                                "created_at" TIMESTAMP DEFAULT (NOW())
);