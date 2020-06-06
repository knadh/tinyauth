package pgstore

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/goyesql/v2"
	goyesqlx "github.com/knadh/goyesql/v2/sqlx"
)

type Queries struct {
	CreateUser          *sqlx.Stmt `db:"create-user"`
	GetUserByID         *sqlx.Stmt `db:"get-user-by-id"`
	GetUserByGUID       *sqlx.Stmt `db:"get-user-by-guid"`
	GetUserByIdentifier *sqlx.Stmt `db:"get-user-by-identifier"`
	UpdateUser          *sqlx.Stmt `db:"update-user"`
	DeleteUser          *sqlx.Stmt `db:"delete-user"`
	LockUser            *sqlx.Stmt `db:"lock-user"`
	UnlockUser          *sqlx.Stmt `db:"unlock-user"`
}

var sqlQueries = `
-- name: create-user
INSERT INTO {tbl}
	guid, identifier, identifier_type, password, require_password, email, display_name, permissions, status
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: get-user-by-id
SELECT * FROM {tbl} WHERE id = $1;

-- name: get-user-by-guid
SELECT * FROM {tbl} WHERE guid = $1;

-- name: get-user-by-identifier
SELECT * FROM {tbl} WHERE identifier = $1;

-- name: delete-user
DELETE FROM {tbl} WHERE identifier = $1 CASCADE;
`

func prepareQueries(db *sqlx.DB, tbl string) (*Queries, error) {
	qMap, err := goyesql.ParseBytes([]byte(strings.ReplaceAll(sqlQueries, `{tbl}`, tbl)))
	if err != nil {
		return nil, err
	}

	// Prepare queries.
	var q Queries
	if err := goyesqlx.ScanToStruct(&q, qMap, db.Unsafe()); err != nil {
		return nil, err
	}
	return &q, nil
}
