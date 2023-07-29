package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/go-co-op/gocron"
	_ "github.com/lib/pq"
	"github.com/marianozunino/go-utu-llamados/ent"
	"github.com/marianozunino/go-utu-llamados/internal"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	chatIDStr := os.Getenv("CHAT_ID")
	connStr := os.Getenv("DATABASE_URL")

	chatId, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ent.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	service := internal.NewService(client, token, chatId)
	scheduler := gocron.NewScheduler(time.UTC)

	healthCheck(client)

	scheduler.Every(1).Day().At("16:20").Do(service.Run)
	scheduler.StartBlocking()
}

func healthCheck(client *ent.Client) {
	srv := http.Server{
		Addr: "0.0.0.0:8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// serve index template
			w.WriteHeader(http.StatusOK)
			jobs := client.JobOffer.Query().AllX(r.Context())
			temp := template.Must(template.ParseFiles("static/index.tmpl"))
			temp.Execute(w, jobs)
		}),
	}

	go srv.ListenAndServe()
	log.Println("Server started")
}
