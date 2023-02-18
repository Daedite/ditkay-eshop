
CREATE TABLE IF NOT EXISTS "product"(
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "quantity" int,
    "buy_price" float,
    "sell_price" float,
    "description" varchar,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "media_type" (
   "id" SERIAL PRIMARY KEY,
   "name" varchar NOT NULL,
   "description" varchar,
   "created_at" timestamptz NOT NULL,
   "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "product_media" (
  "id" SERIAL PRIMARY KEY,
  "media_id" int NOT NULL,
  "product_id" int NOT NULL,
  "media_type" int NOT NULL,
   "created_at" timestamptz NOT NULL,
   "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "media" (
    "id" SERIAL PRIMARY KEY,
    "image" varchar,
    "description" varchar,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL
    );

ALTER TABLE "product_media" add FOREIGN KEY ("product_id") REFERENCES "product" ("id");
ALTER TABLE "product_media" add FOREIGN KEY ("media_id") REFERENCES "media" ("id");