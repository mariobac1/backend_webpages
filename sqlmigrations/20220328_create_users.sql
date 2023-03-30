CREATE TABLE users (
    id UUID NOT NULL,
	name VARCHAR(25) NOT NULL,
    email VARCHAR(254) NOT NULL,
    password VARCHAR(72) NOT NULL,
    avatar VARCHAR(100) NOT NULL,
    details JSONB NOT NULL,
    created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
    updated_at INTEGER,
    CONSTRAINT users_id_pk PRIMARY KEY (id),
    CONSTRAINT users_email_uk UNIQUE (email)
);

COMMENT ON TABLE users IS 'Storage the admins and users for the Webpage';
