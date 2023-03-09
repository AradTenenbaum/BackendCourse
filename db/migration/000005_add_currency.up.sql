CREATE TABLE "currency" (
  "desc" varchar PRIMARY KEY
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("currency") REFERENCES "currency" ("desc");