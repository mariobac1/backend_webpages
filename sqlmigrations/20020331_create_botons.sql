CREATE TABLE buttons (
    id UUID NOT NULL,
	name VARCHAR(25) NOT NULL,
    color VARCHAR(25) NOT NULL,
    shape VARCHAR(25) NOT NULL,
    details JSONB NOT NULL,
    created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
    updated_at INTEGER,
    CONSTRAINT buttons_id_pk PRIMARY KEY (id),
    CONSTRAINT buttons_name_uk UNIQUE (name)
);

COMMENT ON TABLE buttons IS 'Storage the admins and buttons for the Webpage';
