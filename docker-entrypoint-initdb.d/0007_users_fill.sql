BEGIN;

INSERT INTO public.users (name, role, is_del) VALUES
                                                         ('vasya', 'admin', false),
                                                         ('admin', 'admin', false),
                                                         ('petya', 'admin', false),
                                                         ('sasha', 'admin', false);
COMMIT;