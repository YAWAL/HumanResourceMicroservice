CREATE TABLE IF NOT EXISTS "address" (
  "address_id"      INTEGER NOT NULL,
  "employee_id"     INTEGER NOT NULL,
  "phone_num"       VARCHAR,
  "residence_place" VARCHAR,
  "street"          VARCHAR,
  "building_num"    VARCHAR,
  "flat_num"        VARCHAR,
  "zip"             VARCHAR,
  CONSTRAINT "address_pk" PRIMARY KEY ("address_id")
);