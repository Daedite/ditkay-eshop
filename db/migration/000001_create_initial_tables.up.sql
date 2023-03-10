
CREATE TABLE IF NOT EXISTS "product"(
    "id" varchar PRIMARY KEY,
    "name" varchar NOT NULL,
    "quantity" int,
    "buy_price" float,
    "sell_price" float,
    "description" varchar,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "media_type" (
   "id" varchar PRIMARY KEY,
   "name" varchar NOT NULL,
   "description" varchar,
   "created_at" timestamptz NOT NULL,
   "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "product_media" (
  "id" varchar PRIMARY KEY,
  "media_id" varchar NOT NULL,
  "product_id" varchar NOT NULL,
  "media_type" varchar NOT NULL,
   "created_at" timestamptz NOT NULL,
   "updated_at" timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS "media" (
    "id" varchar PRIMARY KEY,
    "image" varchar,
    "description" varchar,
    "created_at" timestamptz NOT NULL,
    "updated_at" timestamptz NOT NULL
    );

ALTER TABLE "product_media" add FOREIGN KEY ("product_id") REFERENCES "product" ("id");
ALTER TABLE "product_media" add FOREIGN KEY ("media_id") REFERENCES "media" ("id");