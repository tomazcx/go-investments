CREATE TABLE IF NOT EXISTS account (
	id UUID PRIMARY KEY,
	forename VARCHAR(255),
	surname VARCHAR(255),
	email VARCHAR(255) UNIQUE,
	balance DOUBLE,
	created_at TIMESTAMP DEFAULT(NOW())
);

CREATE TABLE IF NOT EXISTS document (
	type INTEGER,
	document VARCHAR(20),
	accountId UUID,
	CONSTRAINT pk_document PRIMARY KEY(type, document),
	CONSTRAINT fk_account_document FOREIGN KEY(accountId) REFERENCES account(id) ON DELETE CASCADE
);
