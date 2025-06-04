CREATE TABLE students (
    id uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
    full_name varchar NOT NULL UNIQUE,
    age integer NOT NULL,
    group_name varchar NOT NULL
);