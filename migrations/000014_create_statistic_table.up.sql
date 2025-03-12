CREATE TABLE "statistic" (
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "points" INTEGER NOT NULL DEFAULT 0
);