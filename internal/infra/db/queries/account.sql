-- name: GetAccount :one
SELECT * FROM account WHERE id = $1;

-- name: GetAccountByEmail :one
SELECT * FROM account WHERE email = $1;

-- name: GetAccountByDocument :one
SELECT a.id, a.forename, a.surname, a.email, a.balance, d.document, d.type FROM document d INNER JOIN account a ON a.id = d.account_id WHERE d.document = $1;

-- name: DocumentExists :one
SELECT * FROM document WHERE document = $1;

-- name: CreateAccount :one
INSERT INTO account (id, forename, surname, email, balance) VALUES ($1, $2, $3, $4, $5) RETURNING created_at;

-- name: CreateDocument :exec
INSERT INTO document(type, document, account_id) VALUES ($1, $2, $3);

-- name: UpdateAccount :exec
UPDATE account SET forename = $2, surname  = $3 WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM account WHERE id = $1;
