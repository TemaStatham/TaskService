CREATE UNIQUE INDEX ON "user_organization" ("user_id", "organization_id");

CREATE INDEX "idx_user_organization_user" ON "user_organization" ("user_id");

CREATE INDEX "idx_user_organization_org" ON "user_organization" ("organization_id");

CREATE UNIQUE INDEX ON "task_user" ("task_id", "user_id");

CREATE UNIQUE INDEX ON "task_category" ("task_id", "category_id");

CREATE UNIQUE INDEX ON "response" ("task_id", "user_id");

ALTER TABLE "organization" ADD CONSTRAINT "fk_status" FOREIGN KEY ("status_id") REFERENCES "organization_statuses" ("id") ON DELETE RESTRICT;

ALTER TABLE "user_organization" ADD CONSTRAINT "fk_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "user_organization" ADD CONSTRAINT "fk_organization" FOREIGN KEY ("organization_id") REFERENCES "organization" ("id") ON DELETE CASCADE;

ALTER TABLE "task" ADD FOREIGN KEY ("organization_id") REFERENCES "organization" ("id") ON DELETE CASCADE;

ALTER TABLE "task_user" ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE;

ALTER TABLE "task_user" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "task_category" ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE;

ALTER TABLE "task_category" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id") ON DELETE CASCADE;

ALTER TABLE "response" ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE;

ALTER TABLE "response" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "comment" ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE;

ALTER TABLE "comment" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "statistic" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "notification" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "chat" ADD FOREIGN KEY ("user1_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "chat" ADD FOREIGN KEY ("user2_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "message" ADD FOREIGN KEY ("chat_id") REFERENCES "chat" ("id") ON DELETE CASCADE;

ALTER TABLE "message" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "task_status" ADD FOREIGN KEY ("id") REFERENCES "task" ("status_id") ON DELETE CASCADE;

ALTER TABLE "response_status" ADD FOREIGN KEY ("id") REFERENCES "response" ("status_id") ON DELETE CASCADE;

ALTER TABLE "approve_task" ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE;

ALTER TABLE "approve_task" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "approve_task" ADD FOREIGN KEY ("status_id") REFERENCES "approve_task_status" ("id") ON DELETE CASCADE;

ALTER TABLE "task" ADD FOREIGN KEY ("type_id") REFERENCES "task_type" ("id");

ALTER TABLE "approve_task" ADD FOREIGN KEY ("approved") REFERENCES "user" ("id");

ALTER TABLE "approve_file" ADD FOREIGN KEY ("approve_task_id") REFERENCES "approve_task" ("id");

ALTER TABLE "approve_file" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "approve_file" ADD FOREIGN KEY ("file_id") REFERENCES "file" ("id");
