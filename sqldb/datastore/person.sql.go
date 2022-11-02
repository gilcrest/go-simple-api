// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: person.sql

package datastore

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPerson = `-- name: CreatePerson :execrows
INSERT INTO person (person_id, person_extl_id, create_app_id, create_user_id,
                    create_timestamp, update_app_id, update_user_id, update_timestamp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type CreatePersonParams struct {
	PersonID        uuid.UUID
	PersonExtlID    string
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) (int64, error) {
	result, err := q.db.Exec(ctx, createPerson,
		arg.PersonID,
		arg.PersonExtlID,
		arg.CreateAppID,
		arg.CreateUserID,
		arg.CreateTimestamp,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const createUser = `-- name: CreateUser :execrows
INSERT INTO users (user_id, user_extl_id, person_id, name_prefix, first_name, middle_name, last_name, name_suffix,
                   nickname, email, company_name, company_dept, job_title, birth_date, birth_year, birth_month, birth_day,
                   create_app_id, create_user_id, create_timestamp,
                   update_app_id, update_user_id, update_timestamp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
`

type CreateUserParams struct {
	UserID          uuid.UUID
	UserExtlID      string
	PersonID        uuid.UUID
	NamePrefix      sql.NullString
	FirstName       string
	MiddleName      sql.NullString
	LastName        string
	NameSuffix      sql.NullString
	Nickname        sql.NullString
	Email           sql.NullString
	CompanyName     sql.NullString
	CompanyDept     sql.NullString
	JobTitle        sql.NullString
	BirthDate       sql.NullTime
	BirthYear       sql.NullInt64
	BirthMonth      sql.NullInt64
	BirthDay        sql.NullInt64
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	result, err := q.db.Exec(ctx, createUser,
		arg.UserID,
		arg.UserExtlID,
		arg.PersonID,
		arg.NamePrefix,
		arg.FirstName,
		arg.MiddleName,
		arg.LastName,
		arg.NameSuffix,
		arg.Nickname,
		arg.Email,
		arg.CompanyName,
		arg.CompanyDept,
		arg.JobTitle,
		arg.BirthDate,
		arg.BirthYear,
		arg.BirthMonth,
		arg.BirthDay,
		arg.CreateAppID,
		arg.CreateUserID,
		arg.CreateTimestamp,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const createUserLanguagePreference = `-- name: CreateUserLanguagePreference :execrows
INSERT INTO users_lang_prefs (user_id, language_tag, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type CreateUserLanguagePreferenceParams struct {
	UserID          uuid.UUID
	LanguageTag     string
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

func (q *Queries) CreateUserLanguagePreference(ctx context.Context, arg CreateUserLanguagePreferenceParams) (int64, error) {
	result, err := q.db.Exec(ctx, createUserLanguagePreference,
		arg.UserID,
		arg.LanguageTag,
		arg.CreateAppID,
		arg.CreateUserID,
		arg.CreateTimestamp,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const createUsersOrg = `-- name: CreateUsersOrg :execrows
INSERT INTO users_org (users_org_id, org_id, user_id,
                       create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

type CreateUsersOrgParams struct {
	UsersOrgID      uuid.UUID
	OrgID           uuid.UUID
	UserID          uuid.UUID
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

func (q *Queries) CreateUsersOrg(ctx context.Context, arg CreateUsersOrgParams) (int64, error) {
	result, err := q.db.Exec(ctx, createUsersOrg,
		arg.UsersOrgID,
		arg.OrgID,
		arg.UserID,
		arg.CreateAppID,
		arg.CreateUserID,
		arg.CreateTimestamp,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const deletePerson = `-- name: DeletePerson :execrows
DELETE
FROM person
WHERE person_id = $1
`

func (q *Queries) DeletePerson(ctx context.Context, personID uuid.UUID) (int64, error) {
	result, err := q.db.Exec(ctx, deletePerson, personID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const deleteUserByID = `-- name: DeleteUserByID :execrows
DELETE FROM users
WHERE user_id = $1
`

func (q *Queries) DeleteUserByID(ctx context.Context, userID uuid.UUID) (int64, error) {
	result, err := q.db.Exec(ctx, deleteUserByID, userID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const deleteUserLanguagePreferences = `-- name: DeleteUserLanguagePreferences :execrows
DELETE FROM users_lang_prefs
WHERE user_id = $1
`

func (q *Queries) DeleteUserLanguagePreferences(ctx context.Context, userID uuid.UUID) (int64, error) {
	result, err := q.db.Exec(ctx, deleteUserLanguagePreferences, userID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findPersonByUserExternalID = `-- name: FindPersonByUserExternalID :one
SELECT p.person_id,
       p.person_extl_id,
       pa.user_id,
       pa.user_extl_id,
       pa.name_prefix,
       pa.first_name,
       pa.middle_name,
       pa.last_name,
       pa.name_suffix,
       pa.nickname,
       pa.email,
       pa.company_name,
       pa.company_dept,
       pa.job_title,
       pa.birth_date,
       pa.birth_year,
       pa.birth_month,
       pa.birth_day
FROM person p
         inner join users pa on pa.person_id = p.person_id
WHERE pa.user_extl_id = $1
`

type FindPersonByUserExternalIDRow struct {
	PersonID     uuid.UUID
	PersonExtlID string
	UserID       uuid.UUID
	UserExtlID   string
	NamePrefix   sql.NullString
	FirstName    string
	MiddleName   sql.NullString
	LastName     string
	NameSuffix   sql.NullString
	Nickname     sql.NullString
	Email        sql.NullString
	CompanyName  sql.NullString
	CompanyDept  sql.NullString
	JobTitle     sql.NullString
	BirthDate    sql.NullTime
	BirthYear    sql.NullInt64
	BirthMonth   sql.NullInt64
	BirthDay     sql.NullInt64
}

func (q *Queries) FindPersonByUserExternalID(ctx context.Context, userExtlID string) (FindPersonByUserExternalIDRow, error) {
	row := q.db.QueryRow(ctx, findPersonByUserExternalID, userExtlID)
	var i FindPersonByUserExternalIDRow
	err := row.Scan(
		&i.PersonID,
		&i.PersonExtlID,
		&i.UserID,
		&i.UserExtlID,
		&i.NamePrefix,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.NameSuffix,
		&i.Nickname,
		&i.Email,
		&i.CompanyName,
		&i.CompanyDept,
		&i.JobTitle,
		&i.BirthDate,
		&i.BirthYear,
		&i.BirthMonth,
		&i.BirthDay,
	)
	return i, err
}

const findPersonByUserID = `-- name: FindPersonByUserID :one
SELECT p.person_id,
       p.person_extl_id,
       pa.user_id,
       pa.user_extl_id,
       pa.name_prefix,
       pa.first_name,
       pa.middle_name,
       pa.last_name,
       pa.name_suffix,
       pa.nickname,
       pa.email,
       pa.company_name,
       pa.company_dept,
       pa.job_title,
       pa.birth_date,
       pa.birth_year,
       pa.birth_month,
       pa.birth_day
FROM person p
         inner join users pa on pa.person_id = p.person_id
WHERE pa.user_id = $1
`

type FindPersonByUserIDRow struct {
	PersonID     uuid.UUID
	PersonExtlID string
	UserID       uuid.UUID
	UserExtlID   string
	NamePrefix   sql.NullString
	FirstName    string
	MiddleName   sql.NullString
	LastName     string
	NameSuffix   sql.NullString
	Nickname     sql.NullString
	Email        sql.NullString
	CompanyName  sql.NullString
	CompanyDept  sql.NullString
	JobTitle     sql.NullString
	BirthDate    sql.NullTime
	BirthYear    sql.NullInt64
	BirthMonth   sql.NullInt64
	BirthDay     sql.NullInt64
}

func (q *Queries) FindPersonByUserID(ctx context.Context, userID uuid.UUID) (FindPersonByUserIDRow, error) {
	row := q.db.QueryRow(ctx, findPersonByUserID, userID)
	var i FindPersonByUserIDRow
	err := row.Scan(
		&i.PersonID,
		&i.PersonExtlID,
		&i.UserID,
		&i.UserExtlID,
		&i.NamePrefix,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.NameSuffix,
		&i.Nickname,
		&i.Email,
		&i.CompanyName,
		&i.CompanyDept,
		&i.JobTitle,
		&i.BirthDate,
		&i.BirthYear,
		&i.BirthMonth,
		&i.BirthDay,
	)
	return i, err
}

const findUserByExternalID = `-- name: FindUserByExternalID :one
SELECT user_id, user_extl_id, person_id, name_prefix, first_name, middle_name, last_name, name_suffix, nickname, email, company_name, company_dept, job_title, birth_date, birth_year, birth_month, birth_day, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp FROM users
WHERE user_extl_id = $1
`

func (q *Queries) FindUserByExternalID(ctx context.Context, userExtlID string) (User, error) {
	row := q.db.QueryRow(ctx, findUserByExternalID, userExtlID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserExtlID,
		&i.PersonID,
		&i.NamePrefix,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.NameSuffix,
		&i.Nickname,
		&i.Email,
		&i.CompanyName,
		&i.CompanyDept,
		&i.JobTitle,
		&i.BirthDate,
		&i.BirthYear,
		&i.BirthMonth,
		&i.BirthDay,
		&i.CreateAppID,
		&i.CreateUserID,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateUserID,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findUserByID = `-- name: FindUserByID :one
SELECT user_id, user_extl_id, person_id, name_prefix, first_name, middle_name, last_name, name_suffix, nickname, email, company_name, company_dept, job_title, birth_date, birth_year, birth_month, birth_day, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp FROM users
WHERE user_id = $1
`

func (q *Queries) FindUserByID(ctx context.Context, userID uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, findUserByID, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserExtlID,
		&i.PersonID,
		&i.NamePrefix,
		&i.FirstName,
		&i.MiddleName,
		&i.LastName,
		&i.NameSuffix,
		&i.Nickname,
		&i.Email,
		&i.CompanyName,
		&i.CompanyDept,
		&i.JobTitle,
		&i.BirthDate,
		&i.BirthYear,
		&i.BirthMonth,
		&i.BirthDay,
		&i.CreateAppID,
		&i.CreateUserID,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateUserID,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findUserLanguagePreferencesByUserID = `-- name: FindUserLanguagePreferencesByUserID :many
SELECT user_id, language_tag, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp
FROM users_lang_prefs
WHERE user_id = $1
`

func (q *Queries) FindUserLanguagePreferencesByUserID(ctx context.Context, userID uuid.UUID) ([]UsersLangPref, error) {
	rows, err := q.db.Query(ctx, findUserLanguagePreferencesByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UsersLangPref
	for rows.Next() {
		var i UsersLangPref
		if err := rows.Scan(
			&i.UserID,
			&i.LanguageTag,
			&i.CreateAppID,
			&i.CreateUserID,
			&i.CreateTimestamp,
			&i.UpdateAppID,
			&i.UpdateUserID,
			&i.UpdateTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
