package group

import (
	"context"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
)

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GroupModel struct {
	DB *bun.DB
}

func (m GroupModel) All(ctx context.Context) ([]Group, error) {
	groups := make([]Group, 0)
	if err := m.DB.NewSelect().Model(&groups).OrderExpr("id ASC").Scan(ctx); err != nil {
		return nil, err
	}
	return groups, nil
}

func stubGroups() []Group {
	var groups []Group
	var g1 Group
	g1.ID = 1
	g1.Name = "all"
	groups = append(groups, g1)
	return groups
}

func (m GroupModel) LoadFixtures() {
	ctx := context.Background()
	// Register models for the fixture.
	m.DB.RegisterModel((*Group)(nil))

	// WithRecreateTables tells Bun to drop existing tables and create new ones.
	fixture := dbfixture.New(m.DB, dbfixture.WithRecreateTables())

	// Load fixture.yml which contains data for User and Story models.
	if err := fixture.Load(ctx, os.DirFS("./group"), "fixture.yaml"); err != nil {
		panic(err)
	}
}
