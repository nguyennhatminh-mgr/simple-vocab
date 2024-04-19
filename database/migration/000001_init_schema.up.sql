CREATE TABLE "Users" (
  "id" SERIAL PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL default (now())
);

CREATE TABLE "Categories" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar NOT NULL,
   "created_at" timestamptz NOT NULL default (now())
);

CREATE TABLE "Vocabularies" (
  "id" SERIAL PRIMARY KEY,
  "word" varchar NOT NULL,
  "meaning" varchar NOT NULL,
  "category_id" integer,
  "created_at" timestamptz NOT NULL default (now())
);

ALTER TABLE "Vocabularies" ADD FOREIGN KEY ("category_id") REFERENCES "Categories" ("id");
