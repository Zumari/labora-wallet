CREATE TABLE public.wallets
(
    id SERIAL NOT NULL PRIMARY KEY,
	dni VARCHAR(255) NOT NULL,
	country VARCHAR(255) NOT NULL,
	order_request DATE NOT NULL,
	balance DECIMAL(15,2) NOT NULL
);