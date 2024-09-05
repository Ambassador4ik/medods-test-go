package dbclient

import (
	"context"
	"log"
	"sync"

	"github.com/Ambassador4ik/medods-test-go/ent"
	"github.com/Ambassador4ik/medods-test-go/internal/config"
	_ "github.com/lib/pq"
)

// Global database client object, can only be initialised once
var (
	Client *ent.Client
	once   sync.Once
)

func InitEntClient(cfg *config.Config) {
	once.Do(func() {
		var err error
		Client, err = ent.Open("postgres", cfg.DBSource)
		if err != nil {
			log.Fatalf("Failed opening connection to postgres: %v", err)
		}

		// Auto apply migrations
		if err := Client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("Failed creating schema resources: %v", err)
		}
	})
}
