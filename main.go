package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRESQL_ADDON_HOST"),
		os.Getenv("POSTGRESQL_ADDON_PORT"),
		os.Getenv("POSTGRESQL_ADDON_USER"),
		os.Getenv("POSTGRESQL_ADDON_PASSWORD"),
		os.Getenv("POSTGRESQL_ADDON_DB"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

/*
func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "", "config file")
	flag.Parse()

	cfg, err := app.NewConfig(configFile)
	if err != nil {
		panic(err)
	}

	app, err := app.New(cfg)
	if err != nil {
		panic(err)
	}

	err = server.Run(context.Background(), app)
	app.Logger.Info(err.Error())
}
*/
