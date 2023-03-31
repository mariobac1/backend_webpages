CREATE TABLE products (
    id UUID NOT NULL,
	name VARCHAR(25) NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    description TEXT NOT NULL,
    details JSONB NOT NULL,
    created_at INTEGER NOT NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
    updated_at INTEGER,
    CONSTRAINT products_id_pk PRIMARY KEY (id),
    CONSTRAINT products_name_uk UNIQUE (name)
);

COMMENT ON TABLE products IS 'Storage the admins and products for the Webpage';
