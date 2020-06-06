package pgstore

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/tinyauth"
	"github.com/lib/pq"
)

// PGStore represents an SQL database store for tinyauth.
type PGStore struct {
	db      *sqlx.DB
	queries *Queries
	tbl     string
}

// New returns a new instance of SQLStore.
func New(db *sql.DB, tbl string) (*PGStore, error) {
	s := PGStore{
		db:  sqlx.NewDb(db, "postgres"),
		tbl: tbl,
	}

	// Parse and prepare the SQL queries.
	q, err := prepareQueries(s.db, tbl)
	if err != nil {
		return nil, err
	}
	s.queries = q

	return &s, nil
}

// CreateUser creates a new user store and returns the autoincrement ID.
func (s *PGStore) CreateUser(u tinyauth.User) (int64, error) {
	var id int64
	if err := s.queries.CreateUser.Get(&id,
		u.GUID,
		u.Identifier,
		u.IdentifierType,
		u.Password,
		u.RequirePassword,
		u.Email,
		u.DisplayName,
		u.Permissions,
		u.Status,
	); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && strings.HasSuffix(pqErr.Constraint, "_key") {
			return 0, errors.New("user already exists")
		}
	}
	return id, nil
}

// GetUserByID gets a user record by its ID.
func (s *PGStore) GetUserByID(id string) (tinyauth.User, error) {
	var u tinyauth.User
	err := s.queries.GetUserByID.Select(&u, id)
	return u, err
}

// GetUserByGUID gets a user record by its GUID.
func (s *PGStore) GetUserByGUID(id string) (tinyauth.User, error) {
	var u tinyauth.User
	err := s.queries.GetUserByGUID.Select(&u, id)
	return u, err
}

// GetUserByIdentifier gets a user record its identifier.
func (s *PGStore) GetUserByIdentifier(id string) (tinyauth.User, error) {
	var u tinyauth.User
	err := s.queries.GetUserByIdentifier.Select(&u, id)
	return u, err
}
