// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: auth.sql

package sqlc

import (
	"context"
)

const createOrReplaceUser = `-- name: CreateOrReplaceUser :exec
REPLACE INTO user (id, email, firstname, lastname, role_id)
VALUES (?, ?, ?, ?, ?)
`

type CreateOrReplaceUserParams struct {
	ID        string `db:"id"`
	Email     string `db:"email"`
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	RoleID    int64  `db:"role_id"`
}

func (q *Queries) CreateOrReplaceUser(ctx context.Context, arg CreateOrReplaceUserParams) error {
	_, err := q.db.ExecContext(ctx, createOrReplaceUser,
		arg.ID,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.RoleID,
	)
	return err
}

const findAllUser = `-- name: FindAllUser :many
SELECT id, email, firstname, lastname, active, role_id
FROM user
`

func (q *Queries) FindAllUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, findAllUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Firstname,
			&i.Lastname,
			&i.Active,
			&i.RoleID,
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

const findUserById = `-- name: FindUserById :one
SELECT id, email, firstname, lastname, active, role_id
FROM user
WHERE id = ?
`

func (q *Queries) FindUserById(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, findUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
		&i.Active,
		&i.RoleID,
	)
	return i, err
}

const updateUserActive = `-- name: UpdateUserActive :exec
UPDATE user
SET active = ?
WHERE id = ?
`

type UpdateUserActiveParams struct {
	Active bool   `db:"active"`
	ID     string `db:"id"`
}

func (q *Queries) UpdateUserActive(ctx context.Context, arg UpdateUserActiveParams) error {
	_, err := q.db.ExecContext(ctx, updateUserActive, arg.Active, arg.ID)
	return err
}

const updateUserRole = `-- name: UpdateUserRole :exec
UPDATE user
SET role_id = ?
WHERE id = ?
`

type UpdateUserRoleParams struct {
	RoleID int64  `db:"role_id"`
	ID     string `db:"id"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) error {
	_, err := q.db.ExecContext(ctx, updateUserRole, arg.RoleID, arg.ID)
	return err
}
