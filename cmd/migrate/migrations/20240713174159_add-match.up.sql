DO $$
BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_catalog.pg_class c
        JOIN   pg_catalog.pg_namespace n ON n.oid = c.relnamespace
        WHERE  n.nspname = 'public'
        AND    c.relname = 'matches'
        AND    c.relkind = 'r'
    ) THEN
        CREATE TABLE public.matches (
            id SERIAL PRIMARY KEY,
            week INT NOT NULL,
            team1_id INT REFERENCES teams(id),
            team2_id INT REFERENCES teams(id),
            team1_score INT DEFAULT 0,
            team2_score INT DEFAULT 0,
            played BOOLEAN DEFAULT FALSE
        );
    END IF;
END $$;
