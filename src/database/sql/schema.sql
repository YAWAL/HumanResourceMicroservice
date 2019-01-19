CREATE TABLE IF NOT EXISTS "employees" (
  "employee_id"        BIGSERIAL   NOT NULL,
  "name"               VARCHAR(20) NOT NULL,
  "last_name"          VARCHAR(30) NOT NULL,
  "midl_name"          VARCHAR(30),
  "pass_series_num"    VARCHAR(10),
  "identification_num" VARCHAR(20),
  "address_id"         INTEGER     NOT NULL,
  "birth_date"         DATE        NOT NULL,
  "employee_type"      VARCHAR(20),
  "position"           VARCHAR(100),
  "join_date"          DATE        NOT NULL,
  "quit_date"          DATE,
  "is_quit"            BOOLEAN DEFAULT FALSE,
  CONSTRAINT "employee_pk" PRIMARY KEY ("employee_id")
);

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

CREATE TABLE IF NOT EXISTS "departments" (
  "department_id"   INTEGER NOT NULL,
  "department_name" VARCHAR(50),
  "chairman_id"     INTEGER NOT NULL,
  CONSTRAINT "department_pk" PRIMARY KEY ("department_id")
);

CREATE TABLE IF NOT EXISTS "workshops" (
  "workshop_id"   INTEGER NOT NULL,
  "workshop_name" VARCHAR(50),
  "chairman_id"   INTEGER NOT NULL,
  CONSTRAINT "workshop_pk" PRIMARY KEY ("workshop_id")
);

CREATE TABLE IF NOT EXISTS "department_employees" (
  "employee_id"   BIGSERIAL NOT NULL,
  "department_id" INTEGER   NOT NULL
);

CREATE TABLE IF NOT EXISTS "workshop_employees" (
  "employee_id" BIGSERIAL NOT NULL,
  "workshop_id" INTEGER   NOT NULL
);

ALTER TABLE "employees"
  ADD CONSTRAINT "address_employee_fk"
FOREIGN KEY ("address_id")
REFERENCES "address" ("address_id")
ON DELETE NO ACTION
ON UPDATE NO ACTION;
