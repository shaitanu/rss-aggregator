package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rss-aggregator/internal/database"
)

type apiConfig struct{
	DB *database.Queries
}

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}

	dbURL:=os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Database url is not found in the environment")
	}

	conn,err:=sql.Open("postgres",dbURL)
	if err!=nil{
		log.Fatal("Can't connect to databse:",err)
	}

	apiCfg:=apiConfig{
	DB: database.New(conn),
	}
	fmt.Println("Port:", portString)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users",apiCfg.handlerCreateUser)
	v1Router.Get("/users",apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Post("/feeds",apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds",apiCfg.hadnlerGetFeeds)

	v1Router.Post("/feed_follows",apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows",apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
