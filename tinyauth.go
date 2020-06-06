package tinyauth

import (
	"time"
)

// User statuses.
const (
	UserStatusPending  = "pending"
	UserStatusActive   = "active"
	UserStatusInactive = "inactive"
	UserStatusBlocked  = "blocked"
)

// TinyAuth is the controller instance for the tinyauth lib.
type TinyAuth struct {
	opt   Options
	store Store
}

// Options represents the configuration for TinyAuth.
type Options struct {
	UsersTable string `json:"users_table"`
}

type Store interface {
	CreateUser(User) (int64, error)
	GetUserByID(id int64) (User, error)
	GetUserByGUID(guid string) (User, error)
	GetUserByIdentifier(identifier string) (User, error)
	UpdateUser(User) error
	DeleteUser(id int64) error
}

// User represents a user in the store.
type User struct {
	ID int64 `db:"id" json:"id"`

	// A globally random unique identifier to refer to the user's account in
	// public contexts (eg: URLs). This can be a random hash or a GUID or any
	// unique string.
	GUID string `db:"guid" json:"guid"`

	// Identifier is the primary login identifier. It can be a username,
	// email, OAuth ID or any identifier string.
	Identifier string `db:"identifier" json:"identifier"`

	// IdentifierType represents the type of the identifier. It is also a
	// string, for eg: userid, email, google (in case of OAuth) etc.
	IdentifierType string `db:"identifier_type" json:"identifier_type"`

	// Password is the option bcrypt password for this account.
	Password string `db:"password" json:"-"`

	// RequiredPassword indicates whether the login method requires a password.
	// For instance, case of OAuth, this will be false.
	RequirePassword bool `db:"require_password" json:"-"`

	// Email is the e-mail ID of the user. This is used for password resets
	// and other communication. This is used irrespective of whether the
	// identifier is also an e-mail string or not.
	Email string `db:"email" json:"email"`

	// DisplayName is the name of the user that's displayed publicly.
	DisplayName string `db:"name" json:"display_name"`

	// Permissions is the list of arbitrary permissions assigned to the user.
	// Permissions are just strings, eg: campaign.create, campaign.delete etc.
	Permissions []string `db:"permissions" json:"permission"`

	// Status is the status enum (pending, active, inactive, disabled).
	Status string `db:"status" json:"status"`

	// ConfirmationCode holds temporary codes for different confirmation
	// scenarios such as e-mail confirmation or as pssword reset confirmation.
	ConfirmationCode   string    `db:"confirmation_code" json:"confirmation_code`
	ConfirmationExpiry time.Time `db:"confirmation_expiry" json:"confirmation_expiry"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"created_at" json:"created_at"`
}

// New returns a new instance of TinyAuth.
func New(o Options, s Store) (*TinyAuth, error) {
	return &TinyAuth{
		opt:   o,
		store: s,
	}, nil
}

// CreateUser creates a new user and returns the numeric ID and the GUID
// both of which are automatically generated.
func (t *TinyAuth) CreateUser(u User) (int64, error) {
	id, err := t.store.CreateUser(u)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t *TinyAuth) GetUserByID(id int64) (User, error) {
	return t.store.GetUserByID(id)
}

func (t *TinyAuth) GetUserGUID(guid string) (User, error) {
	return t.store.GetUserByGUID(guid)
}

func (t *TinyAuth) GetUserByIdentifier(identifier string) (User, error) {
	return t.store.GetUserByIdentifier(identifier)
}

func (t *TinyAuth) UpdateUser(id int64, u User) error {
	return nil
}

func (t *TinyAuth) DeleteUser(id int64) error {
	return t.store.DeleteUser(id)
}

func (t *TinyAuth) LoginWithPassword(u User, password string) error {
	return nil
}

func (t *TinyAuth) LockUser(uuid string) error {
	return nil
}

func (t *TinyAuth) UnlockUser(uuid string) error {
	return nil
}
