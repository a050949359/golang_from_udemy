CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" bigserial PRIMARY KEY,
  "form_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "account" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfer" ("form_account_id");

CREATE INDEX ON "transfer" ("to_account_id");

CREATE INDEX ON "transfer" ("form_account_id", "to_account_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("form_account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");
