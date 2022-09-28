// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createPaste = `-- name: CreatePaste :one
insert into paste (
    id, created_at, expires_in, deletes_at, title, owner_id, private
) values (
    $1, $2, $3, $4, $5, $6, $7
)
returning id, created_at, expires_in, deletes_at, title, owner_id, private
`

type CreatePasteParams struct {
	ID        string         `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	ExpiresIn ExpiresIn      `db:"expires_in"`
	DeletesAt sql.NullTime   `db:"deletes_at"`
	Title     string         `db:"title"`
	OwnerID   sql.NullString `db:"owner_id"`
	Private   bool           `db:"private"`
}

func (q *Queries) CreatePaste(ctx context.Context, arg CreatePasteParams) (Paste, error) {
	row := q.queryRow(ctx, q.createPasteStmt, createPaste,
		arg.ID,
		arg.CreatedAt,
		arg.ExpiresIn,
		arg.DeletesAt,
		arg.Title,
		arg.OwnerID,
		arg.Private,
	)
	var i Paste
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.ExpiresIn,
		&i.DeletesAt,
		&i.Title,
		&i.OwnerID,
		&i.Private,
	)
	return i, err
}

const createPasty = `-- name: CreatePasty :one
insert into pasty (
    id, paste_id, title, content, language
) values (
    $1, $2, $3, $4, $5
)
returning id, paste_id, title, content, language
`

type CreatePastyParams struct {
	ID       string `db:"id"`
	PasteID  string `db:"paste_id"`
	Title    string `db:"title"`
	Content  string `db:"content"`
	Language string `db:"language"`
}

func (q *Queries) CreatePasty(ctx context.Context, arg CreatePastyParams) (Pasty, error) {
	row := q.queryRow(ctx, q.createPastyStmt, createPasty,
		arg.ID,
		arg.PasteID,
		arg.Title,
		arg.Content,
		arg.Language,
	)
	var i Pasty
	err := row.Scan(
		&i.ID,
		&i.PasteID,
		&i.Title,
		&i.Content,
		&i.Language,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
insert into "user" (
    id, created_at, username, avatar_url, provider_name, provider_id
) values (
    $1, $2, $3, $4, $5, $6
)
returning id, created_at, username, avatar_url, contributor, supporter, provider_name, provider_id
`

type CreateUserParams struct {
	ID           string    `db:"id"`
	CreatedAt    time.Time `db:"created_at"`
	Username     string    `db:"username"`
	AvatarUrl    string    `db:"avatar_url"`
	ProviderName string    `db:"provider_name"`
	ProviderID   string    `db:"provider_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.Username,
		arg.AvatarUrl,
		arg.ProviderName,
		arg.ProviderID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Username,
		&i.AvatarUrl,
		&i.Contributor,
		&i.Supporter,
		&i.ProviderName,
		&i.ProviderID,
	)
	return i, err
}

const deleteExpiredPastes = `-- name: DeleteExpiredPastes :one
with deleted as
    (delete from paste where expires_in != 'never' and deletes_at < now() returning id, created_at, expires_in, deletes_at, title, owner_id, private)
select count(*) from deleted
`

func (q *Queries) DeleteExpiredPastes(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.deleteExpiredPastesStmt, deleteExpiredPastes)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deletePaste = `-- name: DeletePaste :exec
delete from paste where id = $1
`

func (q *Queries) DeletePaste(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deletePasteStmt, deletePaste, id)
	return err
}

const existsPaste = `-- name: ExistsPaste :one
select exists(select 1 from paste where id = $1)
`

func (q *Queries) ExistsPaste(ctx context.Context, id string) (bool, error) {
	row := q.queryRow(ctx, q.existsPasteStmt, existsPaste, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsUserById = `-- name: ExistsUserById :one
select exists(select 1 from "user" where id = $1)
`

func (q *Queries) ExistsUserById(ctx context.Context, id string) (bool, error) {
	row := q.queryRow(ctx, q.existsUserByIdStmt, existsUserById, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsUserByProvider = `-- name: ExistsUserByProvider :one
select exists(select 1 from "user" where provider_name = $1 and provider_id = $2)
`

type ExistsUserByProviderParams struct {
	ProviderName string `db:"provider_name"`
	ProviderID   string `db:"provider_id"`
}

func (q *Queries) ExistsUserByProvider(ctx context.Context, arg ExistsUserByProviderParams) (bool, error) {
	row := q.queryRow(ctx, q.existsUserByProviderStmt, existsUserByProvider, arg.ProviderName, arg.ProviderID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsUserByUsername = `-- name: ExistsUserByUsername :one
select exists(select 1 from "user" where lower(username) = lower($1))
`

func (q *Queries) ExistsUserByUsername(ctx context.Context, lower string) (bool, error) {
	row := q.queryRow(ctx, q.existsUserByUsernameStmt, existsUserByUsername, lower)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getPaste = `-- name: GetPaste :one
select id, created_at, expires_in, deletes_at, title, owner_id, private from paste
where id = $1 limit 1
`

func (q *Queries) GetPaste(ctx context.Context, id string) (Paste, error) {
	row := q.queryRow(ctx, q.getPasteStmt, getPaste, id)
	var i Paste
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.ExpiresIn,
		&i.DeletesAt,
		&i.Title,
		&i.OwnerID,
		&i.Private,
	)
	return i, err
}

const getPasteCount = `-- name: GetPasteCount :one
select count(*) from paste
`

func (q *Queries) GetPasteCount(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.getPasteCountStmt, getPasteCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getPastePasties = `-- name: GetPastePasties :many
select id, paste_id, title, content, language from pasty
where paste_id = $1
`

func (q *Queries) GetPastePasties(ctx context.Context, pasteID string) ([]Pasty, error) {
	rows, err := q.query(ctx, q.getPastePastiesStmt, getPastePasties, pasteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pasty
	for rows.Next() {
		var i Pasty
		if err := rows.Scan(
			&i.ID,
			&i.PasteID,
			&i.Title,
			&i.Content,
			&i.Language,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPasteStarCount = `-- name: GetPasteStarCount :one
select count(*) from star where paste_id = $1
`

func (q *Queries) GetPasteStarCount(ctx context.Context, pasteID string) (int64, error) {
	row := q.queryRow(ctx, q.getPasteStarCountStmt, getPasteStarCount, pasteID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUserAllPastes = `-- name: GetUserAllPastes :many
select id, created_at, expires_in, deletes_at, title, owner_id, private from paste
where owner_id = $1
order by paste.created_at desc
limit $2
offset $3
`

type GetUserAllPastesParams struct {
	OwnerID sql.NullString `db:"owner_id"`
	Limit   int32          `db:"limit"`
	Offset  int32          `db:"offset"`
}

func (q *Queries) GetUserAllPastes(ctx context.Context, arg GetUserAllPastesParams) ([]Paste, error) {
	rows, err := q.query(ctx, q.getUserAllPastesStmt, getUserAllPastes, arg.OwnerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Paste
	for rows.Next() {
		var i Paste
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.ExpiresIn,
			&i.DeletesAt,
			&i.Title,
			&i.OwnerID,
			&i.Private,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserAllPastesCount = `-- name: GetUserAllPastesCount :one
select count(*) from paste
where owner_id = $1
`

func (q *Queries) GetUserAllPastesCount(ctx context.Context, ownerID sql.NullString) (int64, error) {
	row := q.queryRow(ctx, q.getUserAllPastesCountStmt, getUserAllPastesCount, ownerID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUserById = `-- name: GetUserById :one
select id, created_at, username, avatar_url, contributor, supporter, provider_name, provider_id from "user"
where id = $1 limit 1
`

func (q *Queries) GetUserById(ctx context.Context, id string) (User, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Username,
		&i.AvatarUrl,
		&i.Contributor,
		&i.Supporter,
		&i.ProviderName,
		&i.ProviderID,
	)
	return i, err
}

const getUserByProvider = `-- name: GetUserByProvider :one
select id, created_at, username, avatar_url, contributor, supporter, provider_name, provider_id from "user"
where provider_name = $1 and provider_id = $2 limit 1
`

type GetUserByProviderParams struct {
	ProviderName string `db:"provider_name"`
	ProviderID   string `db:"provider_id"`
}

func (q *Queries) GetUserByProvider(ctx context.Context, arg GetUserByProviderParams) (User, error) {
	row := q.queryRow(ctx, q.getUserByProviderStmt, getUserByProvider, arg.ProviderName, arg.ProviderID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Username,
		&i.AvatarUrl,
		&i.Contributor,
		&i.Supporter,
		&i.ProviderName,
		&i.ProviderID,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
select id, created_at, username, avatar_url, contributor, supporter, provider_name, provider_id from "user"
where lower(username) = lower($1) limit 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, lower string) (User, error) {
	row := q.queryRow(ctx, q.getUserByUsernameStmt, getUserByUsername, lower)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Username,
		&i.AvatarUrl,
		&i.Contributor,
		&i.Supporter,
		&i.ProviderName,
		&i.ProviderID,
	)
	return i, err
}

const getUserPublicPastes = `-- name: GetUserPublicPastes :many
select id, created_at, expires_in, deletes_at, title, owner_id, private from paste
where owner_id = $1 and private = false
order by paste.created_at desc
limit $2
offset $3
`

type GetUserPublicPastesParams struct {
	OwnerID sql.NullString `db:"owner_id"`
	Limit   int32          `db:"limit"`
	Offset  int32          `db:"offset"`
}

func (q *Queries) GetUserPublicPastes(ctx context.Context, arg GetUserPublicPastesParams) ([]Paste, error) {
	rows, err := q.query(ctx, q.getUserPublicPastesStmt, getUserPublicPastes, arg.OwnerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Paste
	for rows.Next() {
		var i Paste
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.ExpiresIn,
			&i.DeletesAt,
			&i.Title,
			&i.OwnerID,
			&i.Private,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserPublicPastesCount = `-- name: GetUserPublicPastesCount :one
select count(*) from paste
where owner_id = $1 and private = false
`

func (q *Queries) GetUserPublicPastesCount(ctx context.Context, ownerID sql.NullString) (int64, error) {
	row := q.queryRow(ctx, q.getUserPublicPastesCountStmt, getUserPublicPastesCount, ownerID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const isPasteStarred = `-- name: IsPasteStarred :one
select exists(select 1 from star where user_id = $1 and paste_id = $2)
`

type IsPasteStarredParams struct {
	UserID  string `db:"user_id"`
	PasteID string `db:"paste_id"`
}

func (q *Queries) IsPasteStarred(ctx context.Context, arg IsPasteStarredParams) (bool, error) {
	row := q.queryRow(ctx, q.isPasteStarredStmt, isPasteStarred, arg.UserID, arg.PasteID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const starPaste = `-- name: StarPaste :exec
insert into star (user_id, paste_id) values ($1, $2) returning user_id, paste_id
`

type StarPasteParams struct {
	UserID  string `db:"user_id"`
	PasteID string `db:"paste_id"`
}

func (q *Queries) StarPaste(ctx context.Context, arg StarPasteParams) error {
	_, err := q.exec(ctx, q.starPasteStmt, starPaste, arg.UserID, arg.PasteID)
	return err
}

const unstarPaste = `-- name: UnstarPaste :exec
delete from star where user_id = $1 and paste_id = $2
`

type UnstarPasteParams struct {
	UserID  string `db:"user_id"`
	PasteID string `db:"paste_id"`
}

func (q *Queries) UnstarPaste(ctx context.Context, arg UnstarPasteParams) error {
	_, err := q.exec(ctx, q.unstarPasteStmt, unstarPaste, arg.UserID, arg.PasteID)
	return err
}
