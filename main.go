package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/xaaaaaanny/gatorrss/internal/database"
)

import (
	"github.com/xaaaaaanny/gatorrss/internal/config"
	"log"
	"os"
)

type state struct {
	db     *database.Queries
	Config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("can`t read config file: %v", err)
	}

	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("cant connect to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	appState := &state{
		db:     dbQueries,
		Config: &cfg,
	}

	existCommands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	if len(os.Args) < 2 {
		log.Fatal("less than 2 arguments")
	}

	existCommands.register("login", handlerLogin)
	existCommands.register("register", handlerRegister)
	existCommands.register("reset", handlerReset)
	existCommands.register("users", handlerListUsers)
	existCommands.register("agg", handlerAggregate)
	existCommands.register("addfeed", handlerCreateFeed)
	existCommands.register("feeds", handlerListFeeds)
	existCommands.register("follow", handlerFeedFollow)
	existCommands.register("following", handlerListFollowFeeds)

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = existCommands.run(appState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
