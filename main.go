package main

import (
	"context"
	"fmt"
	"github.com/IuryAlves/cleaning-robot/app/server"
	"github.com/IuryAlves/cleaning-robot/app/storage/migrations"
	"github.com/IuryAlves/cleaning-robot/app/svc"
	"net/http"
	"os"
)

func runServer(port string) {
	http.HandleFunc("/tibber-developer-test/enter-path", server.EnterPathHandler)
	fmt.Printf("listening on port: %s\n", port)
	http.ListenAndServe(":8080", nil)
}

func migrate(ctx context.Context) error {
	s := svc.New(svc.WithDefaultStorageClient())
	db := s.StorageClient.GetDB()
	if db == nil {
		return fmt.Errorf("storage client not initialized")
	}
	return migrations.Migrate(ctx, db)
}

func usage() {
	fmt.Printf(`Usage: %s [--server|--migrate]
	--server Run the server
	--migrate Run the database migrations
	--help print this helper
`, os.Args[0])
}

func main() {
	ctx := context.Background()
	option := "--server"
	if len(os.Args) >= 2 {
		option = os.Args[1]
	} else {
		usage()
		os.Exit(1)
	}

	switch option {
	case "--help":
		usage()
		os.Exit(0)
	case "--server":
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		if err := migrate(ctx); err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		runServer(port)
	case "--migrate":
		if err := migrate(ctx); err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	default:
		fmt.Println(fmt.Errorf("invalid program argument: %s", option))
		usage()
		os.Exit(1)
	}
}
