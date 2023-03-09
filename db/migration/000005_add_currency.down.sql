ALTER TABLE "accounts" DROP FOREIGN KEY ("currency") REFERENCES "currency" ("desc");

DROP TABLE IF EXISTS "currency";
