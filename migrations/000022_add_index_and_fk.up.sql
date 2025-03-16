CREATE UNIQUE INDEX ON "task_user" ("task_id", "user_id");
CREATE UNIQUE INDEX ON "task_category" ("task_id", "category_id");
CREATE UNIQUE INDEX ON "response" ("task_id", "user_id");

ALTER TABLE "task"
    ADD FOREIGN KEY ("type_id") REFERENCES "task_type" ("id") ON DELETE CASCADE,
    ADD FOREIGN KEY ("status_id") REFERENCES "task_status" ("id") ON DELETE CASCADE;

ALTER TABLE "task_user"
    ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE;

ALTER TABLE "task_category"
    ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE,
    ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id") ON DELETE CASCADE;

ALTER TABLE "response"
    ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE,
    ADD FOREIGN KEY ("status_id") REFERENCES "response_status" ("id") ON DELETE CASCADE;

ALTER TABLE "comment"
    ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE;

ALTER TABLE "message"
    ADD FOREIGN KEY ("chat_id") REFERENCES "chat" ("id") ON DELETE CASCADE;

ALTER TABLE "approve_task"
    ADD FOREIGN KEY ("task_id") REFERENCES "task" ("id") ON DELETE CASCADE,
    ADD FOREIGN KEY ("status_id") REFERENCES "approve_task_status" ("id") ON DELETE CASCADE,
    ADD FOREIGN KEY ("approved") REFERENCES "approve_task_status" ("id") ON DELETE SET NULL;

ALTER TABLE "approve_file"
    ADD FOREIGN KEY ("approve_task_id") REFERENCES "approve_task" ("id") ON DELETE CASCADE,
    ADD FOREIGN KEY ("file_id") REFERENCES "file" ("id") ON DELETE CASCADE;
