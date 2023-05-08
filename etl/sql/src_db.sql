CREATE TABLE "old"."public"."user_attributes"
(
    "client_id"    BIGINT,
    "id"           SERIAL,
    "utm_campaign" VARCHAR(30),
    "firstname"    TEXT,
    "lastname"     TEXT,
    CONSTRAINT "user_attributes_pk" PRIMARY KEY ("client_id")
);


CREATE TABLE "orders_attributes"
(
    "id"             SERIAL,
    "client_id"      BIGINT,
    "itemcode"       BIGINT,
    "category"       TEXT,
    "stock"          BOOLEAN,
    "vendorcode"     BIGINT,
    "payment_amount" BIGINT,
    "description"    TEXT,
    "num"            NUMERIC(14, 2),
    "datetime"       TIMESTAMP,
    "priceproduct"   NUMERIC(14, 2),
    CONSTRAINT "orders_attributes_fk" FOREIGN KEY ("client_id") REFERENCES "user_attributes" ("client_id"),
    CONSTRAINT "orders_attributes_pk" PRIMARY KEY ("id", "client_id")
);

CREATE TABLE "public"."user_activity_log"
(
    "id"          SERIAL,
    "client_id"   BIGINT,
    "hitdatetime" TIMESTAMP,
    "action"      VARCHAR(20),
    CONSTRAINT "client_id_fk" FOREIGN KEY ("client_id") REFERENCES "user_attributes" ("client_id"),
    CONSTRAINT "user_activity_log_id_pk" PRIMARY KEY ("id", "client_id")
);

CREATE TABLE "user_payment_log"
(
    "id"             SERIAL UNIQUE,
    "client_id"      BIGINT,
    "order_id"       BIGINT,
    "hitdatetime"    TIMESTAMP,
    "action"         VARCHAR(20),
    "payment_amount" NUMERIC(14, 2),
    CONSTRAINT "user_payment_log_client_id_fk" FOREIGN KEY ("client_id") REFERENCES "user_attributes" ("client_id"),
    CONSTRAINT "user_payment_log_order_id_fk" FOREIGN KEY ("order_id") REFERENCES "orders_attributes" ("id"),
    CONSTRAINT "user_payment_log_pk" PRIMARY KEY ("id", "client_id")
);