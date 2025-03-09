CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name STRING NOT NULL,
                       surname STRING NOT NULL,
                       role SMALLINT NOT NULL
);

CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       organization_id INTEGER NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
                       title VARCHAR(255) NOT NULL,
                       type SMALLINT NOT NULL,
                       description TEXT NOT NULL,
                       location VARCHAR(255) NOT NULL,
                       task_date TIMESTAMP NOT NULL,
                       participants_count INTEGER CHECK (participants_count > 0),
                       max_score INTEGER CHECK (max_score >= 0),
                       status SMALLINT DEFAULT,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks_users (
                             id SERIAL PRIMARY KEY,
                             task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
                             user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                             UNIQUE (task_id, user_id)
);

CREATE TABLE tasks_organizations (
                                     id SERIAL PRIMARY KEY,
                                     task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
                                     organization_id INTEGER NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
                                     UNIQUE (task_id, organization_id)
);

CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE task_category (
                               id SERIAL PRIMARY KEY,
                               task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
                               category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
                               UNIQUE (task_id, category_id)
);

CREATE TABLE responses (
                           id SERIAL PRIMARY KEY,
                           task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
                           user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                           status VARCHAR(50) DEFAULT 'Откликнулся',
                           UNIQUE (task_id, user_id)
);

CREATE TABLE comments (
                          id SERIAL PRIMARY KEY,
                          task_id INTEGER NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
                          user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          comment TEXT NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_tasks_updated_at()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_tasks_updated_at_trigger
    BEFORE UPDATE ON tasks
    FOR EACH ROW
    EXECUTE FUNCTION update_tasks_updated_at();
