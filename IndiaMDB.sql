CREATE TABLE "userdata" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL,
  "passwordhash" varchar NOT NULL
);

CREATE TABLE "movie" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "summary" varchar NOT NULL,
  "language" varchar NOT NULL,
  "genre" varchar NOT NULL,
  "release_date" timestamptz NOT NULL
);

CREATE TABLE "rating" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "movie_id" int NOT NULL,
  "user_id" int NOT NULL,
  "rating" int NOT NULL
);

ALTER TABLE "rating" ADD FOREIGN KEY ("user_id") REFERENCES "userdata" ("id");

ALTER TABLE "rating" ADD FOREIGN KEY ("movie_id") REFERENCES "movie" ("id");
