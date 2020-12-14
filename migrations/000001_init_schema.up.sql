CREATE TABLE "users"(
                        "id" serial PRIMARY KEY,
                        "login" varchar(128) UNIQUE NOT NULL,
                        "password" varchar(64) NOT NULL
);


CREATE TABLE "languages"(
                            "id" serial PRIMARY KEY,
                            "name" varchar UNIQUE NOT NULL
);

CREATE TABLE "tasks"(
                        "id" serial PRIMARY KEY,
                        "name" varchar NOT NULL,
                        "description" varchar NOT NULL,
                        "time_limit" int NOT NULL,
                        "memory" int NOT NULL
);

CREATE TABLE "test_cases"(
                             "id" serial PRIMARY KEY,
                             "task_id" serial,
                             "test_data" text NOT NULL,
                             "answer" text NOT NULL,
                             FOREIGN KEY ("task_id") references "tasks"("id") ON DELETE CASCADE
);

CREATE TABLE "history_user"(
                               "id" serial PRIMARY KEY,
                               "user_id" serial,
                               "task_id" serial,
                               "language_id" serial,
                               "solution" text,
                               "success" bool NOT NULL,
                               FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
                               FOREIGN KEY ("task_id") REFERENCES "tasks"("id") ON DELETE CASCADE,
                               FOREIGN KEY ("language_id") REFERENCES "languages"("id") ON DELETE CASCADE
);
