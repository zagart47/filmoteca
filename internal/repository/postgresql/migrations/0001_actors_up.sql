CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TABLE IF NOT EXISTS public.actors
(
    id         SERIAL      NOT NULL PRIMARY KEY,
    name       VARCHAR     NOT NULL,
    gender     VARCHAR,
    birthdate  VARCHAR,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_del     BOOL
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON public.actors
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp()


