package main

import (
	"log"

	"github.com/pizza-nz/go-social/internal/db"
	"github.com/pizza-nz/go-social/internal/env"
	"github.com/pizza-nz/go-social/internal/store"
)

var version string = "0.0.1"

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR",
				"postgres://mock:mock@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENVIRO", "development"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxIdleConns,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Print("DB connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
