CREATE TABLE IF NOT EXISTS patron (
    id integer PRIMARY KEY,
    name character varying  NULL,
    number character varying  NULL,
    email character varying  NULL
);

CREATE TABLE IF NOT EXISTS restaurant (
    id integer PRIMARY KEY,
    name character varying  NULL,
    sunday character varying  NULL,
    monday character varying  NULL,
    tuesday character varying  NULL,
    wednesday character varying  NULL,
    thursday character varying  NULL,
    friday character varying  NULL,
    saturday character varying  NULL,
    hours character varying  NULL
);

CREATE TABLE IF NOT EXISTS reservation (
    restaurant_id integer  NOT NULL,
    patron_id integer  NOT NULL,
    start_time timestamp without time zone  NOT NULL,
    CONSTRAINT reservation_pkey
        PRIMARY KEY (restaurant_id, patron_id, start_time),
    CONSTRAINT reservation_paton_id_fkey
        FOREIGN KEY (patron_id)
            REFERENCES patron(id),
    CONSTRAINT reservation_restaurant_id_fkey
        FOREIGN KEY (restaurant_id)
            REFERENCES restaurant(id)
);
