-- name: CreateUser :one
INSERT INTO users (
	id,
	name,
	email,
	password,
	profile_image,
	created_at,
	created_by,
	is_deleted
) VALUES (
	 $1, $2, $3, $4, $5, $6, $7, $8
) 
RETURNING * ;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET
	name = $2,
	email = $3,
	password = $4,
	profile_image = $5,
	updated_at = $6,
	updated_by = $7
WHERE
	id = $1
RETURNING *;

-- name: GetListUserWithPagination :many
SELECT * FROM users LIMIT $1 OFFSET $2;