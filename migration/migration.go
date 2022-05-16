package migration

import (
	"context"
	"log"

	"github.com/deusdat/arangomigo"
)

func Migrate(endpoint, user, pass, db string) {
	config := arangomigo.Config{
		Endpoints: []string{endpoint},
		Username:  user,
		Password:  pass,
		Db:        db,
	}

	migrations := []arangomigo.Migration{
		&arangomigo.Database{
			Operation: arangomigo.Operation{Type: "database", Action: arangomigo.CREATE, Name: db},
		},
	}
	// for _, table := range tables {
	// 	migrations = append(migrations, &arangomigo.Collection{
	// 		Operation:   arangomigo.Operation{Type: "collection", Action: arangomigo.CREATE, Name: table},
	// 		JournalSize: &journalSize,
	// 		WaitForSync: &waitForSync,
	// 	})
	// }
	err := arangomigo.PerformMigrations(context.Background(), config, migrations)
	if e(err) {
		log.Fatal(err)
	}
}

func e(err error) bool {
	return err != nil
}
