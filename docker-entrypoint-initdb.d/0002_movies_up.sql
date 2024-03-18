CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TABLE IF NOT EXISTS public.movies
(
    id           SERIAL       NOT NULL PRIMARY KEY,
    title        VARCHAR(150) NOT NULL,
    description  VARCHAR(1000),
    release_date VARCHAR,
    rating       DECIMAL(3, 1),
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    is_del       BOOL
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON public.movies
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();