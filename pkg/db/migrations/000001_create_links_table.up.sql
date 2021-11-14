CREATE TABLE IF NOT EXISTS links (
	id serial NOT NULL,
	hash varchar(16) NULL,
	original_url varchar(512) NULL,
	creation_date timestamp NULL,
	expiration_date timestamp NULL,
	CONSTRAINT links_pk PRIMARY KEY (id)
);
