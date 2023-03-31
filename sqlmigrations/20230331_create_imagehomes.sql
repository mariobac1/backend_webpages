CREATE TABLE imagehomes (
    id UUID NOT NULL,
	name VARCHAR(25) NOT NULL,
    color VARCHAR(25) NOT NULL,
    description TEXT,
    details JSONB NOT NULL,
    created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
    updated_at INTEGER,
    CONSTRAINT imagehomes_id_pk PRIMARY KEY (id),
    CONSTRAINT imagehomes_name_uk UNIQUE (name)
);

COMMENT ON TABLE imagehomes IS 'Storage the admins and imagehomes for the Webpage';
