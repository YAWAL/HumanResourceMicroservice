CREATE TABLE IF NOT EXISTS "employees" (
  "employee_id"        BIGSERIAL   NOT NULL,
  "name"               VARCHAR(20) NOT NULL,
  "last_name"          VARCHAR(30) NOT NULL,
  "middle_name"          VARCHAR(30),
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

INSERT INTO employees (
  "employee_id", "name", "last_name", "pass_series_num", "identification_num", "address_id",
  "birth_date", "employee_type", "position", "join_date", "quit_date", "is_quit")
VALUES (
  20110101001, 'JOHN', 'SMITH', 'KN491002', '9870465738', 1, '01/01/1988', 'ENGINEER', 'DESIGN ENGINEER',
               '01/01/2011', '01/01/2012', TRUE
);
INSERT INTO employees (
  "employee_id", "name", "last_name", "middle_name", "pass_series_num", "identification_num", "address_id",
  "birth_date", "employee_type", "position", "join_date", "quit_date", "is_quit")
VALUES (
  20070101002, 'JOHN', 'DIR', '', 'BN581092', '9580762396', 2, '02 / 02 / 1977', 'worker', 'mechanic',
               '03 / 03 / 2007', '04 / 04 / 2015', TRUE
);
INSERT INTO employees (
  "employee_id", "name", "last_name", "middle_name", "pass_series_num", "identification_num", "address_id",
  "birth_date", "employee_type", "position", "join_date", "is_quit")
VALUES (
  20170102001, 'jek', 'black', 'shean', 'rz391992', '3870965738', 3, '01 / 01 / 1988', 'ENGINEER', 'PRODUCT ENGINEER',
               '01 / 01 / 2017', FALSE
);
INSERT INTO employees (
  "employee_id", "name", "last_name", "middle_name", "pass_series_num", "identification_num", "address_id",
  "birth_date", "employee_type", "position", "join_date", "is_quit")
VALUES (
  20160301001, 'ivan', 'petro', 'vasyl', 'gl781092', '7580722396', 4, '11 / 12 / 1987', 'worker', 'CARPENTER',
               '08 / 08 / 2016', FALSE
);