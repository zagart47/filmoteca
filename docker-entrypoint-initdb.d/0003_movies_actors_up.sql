CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TABLE IF NOT EXISTS public.movies_actors
(
    id         SERIAL      NOT NULL PRIMARY KEY,
    movie_id   INT,
    FOREIGN KEY (movie_id) REFERENCES public.movies (id),
    actor_id   INT,
    FOREIGN KEY (actor_id) REFERENCES public.actors (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    is_del     BOOL
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON public.movies_actors
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();