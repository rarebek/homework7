CREATE TABLE IF NOT EXISTS users(
                        "id" UUID PRIMARY KEY NOT NULL,
                        "name" TEXT NOT NULL,
                        "age" INTEGER NOT NULL,
                        "username" TEXT UNIQUE NOT NULL,
                        "email" TEXT UNIQUE NOT NULL,
                        "password" TEXT NOT NULL,
                        "refresh_token" TEXT NOT NULL,
                        "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        "updated_at" TIMESTAMP,
                        "deleted_at" TIMESTAMP
);