-- name: CreateUser :one
INSERT INTO Users (username, email, email_verified, image)
VALUES ($1,$2,$3,$4)
RETURNING username;

-- name: UpdateUser :one
UPDATE Users
SET
  username = COALESCE(sqlc.narg(username), username),
  email = COALESCE(sqlc.narg(email), email),
  email_verified = COALESCE(sqlc.narg(email_verified), email_verified)
WHERE
  id = $1
RETURNING id;

-- name: GetUser :one
SELECT *
FROM Users
WHERE username = $1;

-- name: GetAllUsers :many
SELECT *
FROM Users;

-- name: DeleteUser :one
DELETE
FROM Users
WHERE username = $1
RETURNING username;

-- name: CreateAccount :one
INSERT INTO Accounts (user_id, type, provider, provider_account_id, refresh_token, access_token, expires_at, token_type, scope, id_token, session_state)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
RETURNING user_id;

-- name: UpdateAccount :one
UPDATE Accounts
SET
  type = COALESCE(sqlc.narg(type), type),
  provider = COALESCE(sqlc.narg(provider), provider),
  provider_account_id = COALESCE(sqlc.narg(provider_account_id), provider_account_id),
  refresh_token = COALESCE(sqlc.narg(refresh_token), refresh_token),
  access_token = COALESCE(sqlc.narg(access_token), access_token),
  expires_at = COALESCE(sqlc.narg(expires_at), expires_at),
  token_type = COALESCE(sqlc.narg(token_type), token_type),
  scope = COALESCE(sqlc.narg(scope), scope),
  id_token = COALESCE(sqlc.narg(id_token), id_token),
  session_state = COALESCE(sqlc.narg(session_state), session_state)
WHERE
  user_id = $1
RETURNING id;

-- name: GetAllAccountsFromUser :many
SELECT *
FROM Accounts
JOIN Users
ON Accounts.user_id = Users.id
WHERE Users.username = $1;

-- name: CreateRefreshToken :one
INSERT INTO UserRefreshTokens (user_id, refresh_token)
VALUES ($1,$2)
RETURNING user_id;

-- name: UpdateRefreshToken :one
UPDATE UserRefreshTokens
SET refresh_token = $1
RETURNING user_id;

-- name: DeleteRefreshToken :one
DELETE
FROM UserRefreshTokens
WHERE user_id = $1
RETURNING user_id;

-- name: GetRefreshToken :one
SELECT *
FROM UserRefreshTokens
WHERE user_id = $1;