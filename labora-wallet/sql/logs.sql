CREATE TABLE public.logs
(
    id SERIAL NOT NULL PRIMARY KEY,
	dni VARCHAR(255) NOT NULL,
	country VARCHAR(255) NOT NULL,
	status_request VARCHAR(255) NOT NULL,
	date_request date NOT NULL,
	type_request VARCHAR(255) NOT NULL
);