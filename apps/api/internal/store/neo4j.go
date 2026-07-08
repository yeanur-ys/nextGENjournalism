package store

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func NewNeo4j(ctx context.Context, uri, user, password string) (neo4j.DriverWithContext, error) {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(user, password, ""))
	if err != nil {
		return nil, err
	}
	if err := driver.VerifyConnectivity(ctx); err != nil {
		return nil, err
	}

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err = session.Run(ctx, `
		CREATE CONSTRAINT article_id IF NOT EXISTS
		FOR (a:Article) REQUIRE a.id IS UNIQUE
	`, nil)
	if err != nil {
		return nil, err
	}
	_, err = session.Run(ctx, `
		CREATE CONSTRAINT user_id IF NOT EXISTS
		FOR (u:User) REQUIRE u.id IS UNIQUE
	`, nil)
	if err != nil {
		return nil, err
	}

	return driver, nil
}
