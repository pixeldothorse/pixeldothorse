package database

import (
	"database/sql"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	bindata "github.com/mattes/migrate/source/go-bindata"
	"github.com/pixeldothorse/pixeldothorse/internal/database/dmigrations"
)

// Migrate attempts to migrate the given database by URL to the most recent
// version of the schema.
func Migrate(durl string) error {
	s := bindata.Resource(dmigrations.AssetNames(),
		func(name string) ([]byte, error) {
			return dmigrations.Asset(name)
		})

	d, err := bindata.WithInstance(s)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, durl)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}

// Destroy unmigrates the given database by URL from the given schema. This will
// DESTROY DATA. ENSURE DATA YOU WANT TO KEEP IS BACKED UP.
func Destroy(durl string) error {
	s := bindata.Resource(dmigrations.AssetNames(),
		func(name string) ([]byte, error) {
			return dmigrations.Asset(name)
		})

	d, err := bindata.WithInstance(s)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, durl)
	if err != nil {
		return err
	}

	err = m.Down()
	if err != nil {
		return err
	}

	db, err := sql.Open("postgres", durl)
	if err != nil {
		return err
	}
	defer db.Close()

	db.Exec("DROP TABLE schema_migrations")

	return nil
}
