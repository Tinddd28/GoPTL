-- name: CreateUser :one
INSERT INTO users (
    name,
    lastname,
    email,
    country,
    hashpass
) VALUES (
             $1,
             $2,
             $3,
             $4,
             $5
         ) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
set name = $2,
    lastname = $3,
    email = $4,
    country = $5
WHERE id = $1
returning *;

-- name: UpdatePassword :one
update users
set hashpass = $2
where email = $1
returning *;

-- name: SetUserActive :one
update users
set isactive = $2
where email = $1
returning *;

-- name: SetUserVerified :one
update users
set isverified = $2
where email = $1
returning *;

-- name: SetUserSuperuser :one
update users
set issuperuser = $2
where email = $1
returning *;

-- name: GetUserForAuth :one
select email, hashpass from users
where email = $1 limit 1;
