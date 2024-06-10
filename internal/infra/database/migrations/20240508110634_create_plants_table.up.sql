CREATE TABLE IF NOT EXISTS public.plants
(
    id              SERIAL PRIMARY KEY,
    user_id         INTEGER REFERENCES public.users(id),
    name            TEXT NOT NULL,
    address         TEXT NOT NULL,
    lat             DOUBLE PRECISION NOT NULL,
    lon             DOUBLE PRECISION NOT NULL,
    type            TEXT NOT NULL,
    created_date    TIMESTAMP NOT NULL,
    updated_date    TIMESTAMP NOT NULL,
    deleted_date    TIMESTAMP
);