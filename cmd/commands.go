package cmd

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func NewMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Run db migration",
		Long:  ``,
		Run: func(commandServe *cobra.Command, args []string) {
			// Open the SQLite database.
			db, err := sql.Open("sqlite3", "links.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			// Create a "links" table if it doesn't exist.
			_, err = db.Exec(`CREATE TABLE IF NOT EXISTS links (id INTEGER PRIMARY KEY, url TEXT)`)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
}

var (
	argUrl         string
	argDescription string
)

func NewAddUrlCmd() *cobra.Command {
	newUrl := &cobra.Command{
		Use:   "new",
		Short: "Run db migration",
		Long:  ``,
		Run: func(commandServe *cobra.Command, args []string) {
			// Open the SQLite database.
			db, err := sql.Open("sqlite3", "links.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			log.Printf("adding url: %s\n", argUrl)

			// Add a link to the database
			_, err = db.Exec("INSERT INTO links (url) VALUES (?)", argUrl)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	newUrl.Flags().StringVar(&argUrl, "url", "", "new url string")
	// newUrl.Flags().StringVar(&argDescription, "description", "", "url description")

	return newUrl
}

func NewListUrlCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all urls",
		Long:  ``,
		Run: func(commandServe *cobra.Command, args []string) {
			// Open the SQLite database.
			db, err := sql.Open("sqlite3", "links.db")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			// Retrieve links from the database
			rows, err := db.Query("SELECT id, url FROM links")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			for rows.Next() {
				var id int
				var url string
				if err := rows.Scan(&id, &url); err != nil {
					log.Fatal(err)
				}
				log.Printf("ID: %d, URL: %s\n", id, url)
			}
		},
	}
}
