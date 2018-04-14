package database

import (
	"context"
	"time"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

// User is a user of pixel.horse. We do not store passwords.
type User struct {
	ID        int       `db:"id"`
	UUID      uuid.UUID `db:"uuid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"` // postgres updates this for you, don't worry about it
	Username  string    `db:"username"`
	DiscordID string    `db:"discord_id"`
	IsAdmin   bool      `db:"is_admin"`
}

// UserOperations is the set of operations that can be done on the users table in
// postgres.
type UserOperations interface {
	CreateUser(context.Context, User) (User, error)
	ReadUserByDiscordID(context.Context, string) (User, error)
	ReadUserByUUID(context.Context, uuid.UUID) (User, error)
	UpdateUser(context.Context, User) (User, error)
	DeleteUser(context.Context, int) error
}

// CreateUser does the needful to create users.
func (p PostgresOperations) CreateUser(ctx context.Context, u User) (User, error) {
	qry := `INSERT INTO users
                ( username
                , discord_id
                , is_admin
                ) VALUES
                ( $1
                , $2
                , $3
                )
                RETURNING id, uuid, created_at, upodated_at;`
	row := p.db.QueryRowContext(ctx, qry, u.Username, u.DiscordID, u.IsAdmin)

	if err := row.Scan(&u.ID, &u.UUID, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return User{}, err
	}

	return u, nil
}

func (p PostgresOperations) ReadUserByDiscordID(ctx context.Context, did string) (User, error) {
	qry := `SELECT
                  u.id
                , u.uuid
                , u.created_at
                , u.updated_at
                , u.username
                , u.discord_id
                , u.is_admin
                FROM users AS u
                WHERE u.discord_id = $1
                LIMIT 1;`

	resp := p.db.QueryRowxContext(ctx, qry, did)
	var u User
	err := resp.StructScan(&u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (p PostgresOperations) ReadUserByUUID(ctx context.Context, uid uuid.UUID) (User, error) {
	qry := `SELECT
                  u.id
                , u.uuid
                , u.created_at
                , u.updated_at
                , u.username
                , u.discord_id
                , u.is_admin
                FROM users AS u
                WHERE u.uuid = $1
                LIMIT 1;`

	resp := p.db.QueryRowxContext(ctx, qry, uid)
	var u User
	err := resp.StructScan(&u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (p PostgresOperations) UpdateUser(ctx context.Context, u User) (User, error) {
	qry := `UPDATE users
                SET username = $1
                    is_admin = $2
                WHERE id = $3
                RETURNING updated_at;`
	row := p.db.QueryRowContext(ctx, qry, u.Username, u.IsAdmin, u.ID)
	err := row.Scan(&u.UpdatedAt)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (p PostgresOperations) DeleteUser(ctx context.Context, uid int) error {
	qry := `DELETE FROM users
                WHERE id = $1
                LIMIT 1;`
	res, err := p.db.ExecContext(ctx, qry, uid)
	if err != nil {
		return err
	}
	if rs, err := res.RowsAffected(); err != nil && rs != 1 {
		return errors.New("no user actually deleted?")
	}

	return nil
}
