package db

import (
	"context"
	"database/sql"
	"github.com/Brandon689/appStructure/types"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type Database struct {
	db *bun.DB
}

func NewDatabase() (*Database, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, "./mydatabase.db")
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	return &Database{db: db}, nil
}

func (d *Database) Migrate(ctx context.Context) error {
	_, err := d.db.NewCreateTable().Model(&types.User{}).IfNotExists().Exec(ctx)
	return err
}

func (d *Database) InsertUser(ctx context.Context, user *types.User) error {
	_, err := d.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (d *Database) GetUserByName(ctx context.Context, name string) (*types.User, error) {
	user := new(types.User)
	err := d.db.NewSelect().Model(user).Where("name = ?", name).Scan(ctx)
	return user, err
}
