CREATE TABLE variablevalues (
    id UUID NOT NULL,
	name VARCHAR(25) NOT NULL,
	title VARCHAR(25),
	paragraph VARCHAR(250),
    color VARCHAR(25),
    bgColor VARCHAR(25),
    font VARCHAR(25),
    icon VARCHAR(25),
    description TEXT,
    details JSONB NOT NULL,
    created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
    updated_at INTEGER,
    CONSTRAINT variablevalues_id_pk PRIMARY KEY (id),
    CONSTRAINT variablevalues_name_uk UNIQUE (name)
);

COMMENT ON TABLE variablevalues IS 'Storage the admins and variablevalues for the Webpage';
