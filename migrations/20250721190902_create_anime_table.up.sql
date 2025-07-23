CREATE TABLE IF NOT EXISTS anime (
                                     id SERIAL PRIMARY KEY,
                                     created_at TIMESTAMP,
                                     updated_at TIMESTAMP,
                                     deleted_at TIMESTAMP,
                                     title TEXT NOT NULL,
                                     studio TEXT NOT NULL,
                                     type TEXT NOT NULL,
                                     episodes INTEGER NOT NULL,
                                     author TEXT,
                                     genres TEXT[],
                                     rating DECIMAL(3,1),
    description TEXT,
    cover_url TEXT,
    release_year INTEGER
    );

select * from anime;

SELECT current_database();