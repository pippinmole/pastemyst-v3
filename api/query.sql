-- name: GetPaste :one
select * from pastes
where id = $1 limit 1;

-- name: GetPastePasties :many
select * from pasties
where paste_id = $1;

-- name: ExistsPaste :one
select exists(select 1 from pastes where id = $1);

-- name: CreatePaste :one
insert into pastes (
    id, created_at, expires_in, deletes_at, title, owner_id, private
) values (
    $1, $2, $3, $4, $5, $6, $7
)
returning *;

-- name: CreatePasty :one
insert into pasties (
    id, paste_id, title, content, language
) values (
    $1, $2, $3, $4, $5
)
returning *;

-- name: DeletePaste :exec
delete from pastes where id = $1;

-- name: GetPasteCount :one
select count(*) from pastes;

-- name: ExistsUserByProvider :one
select exists(select 1 from users where provider_name = $1 and provider_id = $2);

-- name: ExistsUserByUsername :one
select exists(select 1 from users where lower(username) = lower($1));

-- name: ExistsUserById :one
select exists(select 1 from users where id = $1);

-- name: GetUserById :one
select * from users
where id = $1 limit 1;

-- name: GetUserByUsername :one
select * from users
where lower(username) = lower($1) limit 1;

-- name: GetUserByProvider :one
select * from users
where provider_name = $1 and provider_id = $2 limit 1;

-- name: CreateUser :one
insert into users (
    id, created_at, username, avatar_url, provider_name, provider_id
) values (
    $1, $2, $3, $4, $5, $6
)
returning *;

-- name: GetUserPublicPastes :many
select * from pastes
where owner_id = $1 and private = false
order by pastes.created_at desc
limit $2
offset $3;

-- name: GetUserPublicPastesCount :one
select count(*) from pastes
where owner_id = $1 and private = false;

-- name: GetUserAllPastes :many
select * from pastes
where owner_id = $1
order by pastes.created_at desc
limit $2
offset $3;

-- name: GetUserAllPastesCount :one
select count(*) from pastes
where owner_id = $1;

-- name: DeleteExpiredPastes :one
with deleted as
    (delete from pastes where expires_in != 'never' and deletes_at < now() returning *)
select count(*) from deleted;

-- name: IsPasteStarred :one
select exists(select 1 from stars where user_id = $1 and paste_id = $2);

-- name: StarPaste :exec
insert into stars (user_id, paste_id) values ($1, $2) returning *;

-- name: UnstarPaste :exec
delete from stars where user_id = $1 and paste_id = $2;

-- name: GetPasteStarCount :one
select count(*) from stars where paste_id = $1;
