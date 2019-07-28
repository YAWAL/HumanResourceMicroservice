CREATE TABLE IF NOT EXISTS "employees"
(
    "employee_id" UUID        NOT NULL,
    "name"        VARCHAR(20) NOT NULL,
    "last_name"   VARCHAR(30) NOT NULL,
    "middle_name" VARCHAR(30),
    "position"    VARCHAR(100),
    "is_quit"     BOOLEAN DEFAULT FALSE,
    CONSTRAINT "employee_pk" PRIMARY KEY ("employee_id")
);

CREATE TABLE IF NOT EXISTS "employee_infos"
(
    "employee_info_id"   UUID NOT NULL,
    "pass_series_num"    VARCHAR(10),
    "identification_num" VARCHAR(20),
    "birth_date"         DATE,
    "join_date"          DATE,
    "quit_date"          DATE,
    CONSTRAINT "employee_infos_pk" PRIMARY KEY ("employee_info_id")
);

CREATE TABLE IF NOT EXISTS "employee_addresses"
(
    "address_id"      UUID NOT NULL,
    "phone_num"       VARCHAR,
    "residence_place" VARCHAR,
    "street"          VARCHAR,
    "building_num"    VARCHAR,
    "flat_num"        VARCHAR,
    "zip"             VARCHAR,
    CONSTRAINT "address_pk" PRIMARY KEY ("address_id")
);
