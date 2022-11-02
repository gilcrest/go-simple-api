package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"golang.org/x/text/language"

	"github.com/gilcrest/diy-go-api"
	"github.com/gilcrest/diy-go-api/errs"
	"github.com/gilcrest/diy-go-api/secure"
	"github.com/gilcrest/diy-go-api/sqldb/datastore"
)

// createPersonTx creates a Person in the database
// Any Users attached to the Person will also be created.
// The created User will be associated to any Orgs attached.
func createPersonTx(ctx context.Context, tx pgx.Tx, p diy.Person, adt diy.Audit) error {
	var err error

	createPersonParams := datastore.CreatePersonParams{
		PersonID:        p.ID,
		PersonExtlID:    p.ExternalID.String(),
		CreateAppID:     adt.App.ID,
		CreateUserID:    adt.User.NullUUID(),
		CreateTimestamp: adt.Moment,
		UpdateAppID:     adt.App.ID,
		UpdateUserID:    adt.User.NullUUID(),
		UpdateTimestamp: adt.Moment,
	}

	// create Person db record
	var rowsAffected int64
	rowsAffected, err = datastore.New(tx).CreatePerson(ctx, createPersonParams)
	if err != nil {
		return errs.E(errs.Database, err)
	}

	if rowsAffected != 1 {
		return errs.E(errs.Database, fmt.Sprintf("person rows affected should be 1, actual: %d", rowsAffected))
	}

	// loop through all users associated to the Person and create them
	for _, u := range p.Users {
		cuTxParams := createUserTxParams{
			PersonID: p.ID,
			User:     u,
			Audit:    adt,
		}
		err = createUserTx(ctx, tx, cuTxParams)
		if err != nil {
			return err
		}
	}

	return nil

}

type createUserTxParams struct {
	// The ID of the Person the User is associated to
	PersonID uuid.UUID
	// The User to be created
	User *diy.User
	// The details for which app and user created/updated the User
	Audit diy.Audit
}

// createUserTx creates a User in the database
func createUserTx(ctx context.Context, tx pgx.Tx, params createUserTxParams) error {
	var err error

	var birthYear, birthMonth, birthDay sql.NullInt64
	if !params.User.BirthDate.IsZero() {
		birthYear = diy.NewNullInt64(int64(params.User.BirthDate.Year()))
		birthMonth = diy.NewNullInt64(int64(params.User.BirthDate.Month()))
		birthDay = diy.NewNullInt64(int64(params.User.BirthDate.Day()))
	}

	cuParams := datastore.CreateUserParams{
		UserID:          params.User.ID,
		UserExtlID:      params.User.ExternalID.String(),
		PersonID:        params.PersonID,
		NamePrefix:      diy.NewNullString(params.User.NamePrefix),
		FirstName:       params.User.FirstName,
		MiddleName:      diy.NewNullString(params.User.MiddleName),
		LastName:        params.User.LastName,
		NameSuffix:      diy.NewNullString(params.User.NameSuffix),
		Nickname:        diy.NewNullString(params.User.Nickname),
		Email:           diy.NewNullString(params.User.Email),
		CompanyName:     diy.NewNullString(params.User.CompanyName),
		CompanyDept:     diy.NewNullString(params.User.CompanyDepartment),
		JobTitle:        diy.NewNullString(params.User.JobTitle),
		BirthDate:       diy.NewNullTime(params.User.BirthDate),
		BirthYear:       birthYear,
		BirthMonth:      birthMonth,
		BirthDay:        birthDay,
		CreateAppID:     params.Audit.App.ID,
		CreateUserID:    params.Audit.User.NullUUID(),
		CreateTimestamp: params.Audit.Moment,
		UpdateAppID:     params.Audit.App.ID,
		UpdateUserID:    params.Audit.User.NullUUID(),
		UpdateTimestamp: params.Audit.Moment,
	}

	// create User db record
	var rowsAffected int64
	rowsAffected, err = datastore.New(tx).CreateUser(ctx, cuParams)
	if err != nil {
		return errs.E(errs.Database, err)
	}

	if rowsAffected != 1 {
		return errs.E(errs.Database, fmt.Sprintf("user rows affected should be 1, actual: %d", rowsAffected))
	}

	return nil

}

// FindUserByID finds a User in the datastore given their User ID
func FindUserByID(ctx context.Context, dbtx datastore.DBTX, id uuid.UUID) (*diy.User, error) {
	dbUser, err := datastore.New(dbtx).FindUserByID(ctx, id)
	if err != nil {
		return nil, errs.E(errs.Database, err)
	}

	var ulp []datastore.UsersLangPref
	ulp, err = datastore.New(dbtx).FindUserLanguagePreferencesByUserID(ctx, id)
	if err != nil {
		return nil, errs.E(errs.Database, err)
	}

	var langPrefs []language.Tag
	for _, p := range ulp {
		tag := language.Make(p.LanguageTag)
		langPrefs = append(langPrefs, tag)
	}

	u := &diy.User{
		ID:                  dbUser.UserID,
		ExternalID:          secure.MustParseIdentifier(dbUser.UserExtlID),
		NamePrefix:          dbUser.NamePrefix.String,
		FirstName:           dbUser.FirstName,
		MiddleName:          dbUser.MiddleName.String,
		LastName:            dbUser.LastName,
		FullName:            "", // TODO - add FullName to users table (and structs)
		NameSuffix:          dbUser.NameSuffix.String,
		Nickname:            dbUser.Nickname.String,
		Gender:              "", // TODO - add Gender to db (full list)
		Email:               dbUser.Email.String,
		CompanyName:         dbUser.CompanyName.String,
		CompanyDepartment:   dbUser.CompanyDept.String,
		JobTitle:            dbUser.JobTitle.String,
		BirthDate:           dbUser.BirthDate.Time,
		LanguagePreferences: langPrefs,
		HostedDomain:        "", // TODO - add a bunch of fields to db
		PictureURL:          "",
		ProfileLink:         "",
		Source:              "",
	}

	return u, nil
}

type attachOrgAssociationParams struct {
	Org   *diy.Org
	User  *diy.User
	Audit diy.Audit
}

// attachOrgAssociation associates an Org with a User in the database.
func attachOrgAssociation(ctx context.Context, tx pgx.Tx, params attachOrgAssociationParams) error {

	createUsersOrgParams := datastore.CreateUsersOrgParams{
		UsersOrgID:      uuid.New(),
		OrgID:           params.Org.ID,
		UserID:          params.User.ID,
		CreateAppID:     params.Audit.App.ID,
		CreateUserID:    diy.NewNullUUID(params.Audit.User.ID),
		CreateTimestamp: params.Audit.Moment,
		UpdateAppID:     params.Audit.App.ID,
		UpdateUserID:    diy.NewNullUUID(params.Audit.User.ID),
		UpdateTimestamp: params.Audit.Moment,
	}

	// create database record using datastore
	rowsAffected, err := datastore.New(tx).CreateUsersOrg(ctx, createUsersOrgParams)
	if err != nil {
		return errs.E(errs.Database, err)
	}

	// should only impact exactly one record
	if rowsAffected != 1 {
		return errs.E(errs.Database, fmt.Sprintf("CreateUsersOrg() should insert 1 row, actual: %d", rowsAffected))
	}

	return nil
}
