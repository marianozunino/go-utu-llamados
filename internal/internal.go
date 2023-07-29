package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/marianozunino/go-utu-llamados/ent"
	"github.com/marianozunino/go-utu-llamados/ent/joboffer"
)

type service struct {
	client *ent.Client
	bot    *tgbotapi.BotAPI
	chatID int64
}

// NewService creates a new instance of the service.
func NewService(client *ent.Client, token string, chatID int64) *service {
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return &service{
		client: client,
		bot:    bot,
		chatID: chatID,
	}
}

// Run starts the service.
func (s *service) Run() {
	// Fetch all inputs from the table
	inputs, err := s.client.SearchInput.Query().All(context.Background())
	if err != nil {
		panic(err)
	}

	// Process each input
	for _, input := range inputs {
		jobOffers := s.processInput(input)
		s.saveJobOffers(jobOffers, input)
	}
}

// processInput processes an individual input record.
func (s *service) processInput(input *ent.SearchInput) []*ent.JobOffer {
	// Construct the URL based on the input data
	url := s.buildURL(input.Search)

	// Request the HTML page.
	doc, err := s.fetchHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	return s.parseHTML(doc)
}

// buildURL constructs the URL using the input data.
func (s *service) buildURL(someField string) string {
	u, err := url.Parse("https://www.utu.edu.uy/concursos-y-llamados")
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("combine", someField)
	u.RawQuery = q.Encode()
	return u.String()
}

// fetchHTML retrieves the HTML document from the provided URL.
func (s *service) fetchHTML(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// parseHTML extracts job offers from the HTML document.
func (s *service) parseHTML(doc *goquery.Document) []*ent.JobOffer {
	var jobOffers []*ent.JobOffer

	doc.Find("#collapseContent .row").Each(func(i int, s *goquery.Selection) {
		publicationDateStr := strings.TrimSpace(s.Find("p.mb-0").First().Text())
		publicationDateStr = strings.Replace(publicationDateStr, "Publicado: ", "", 1)
		dateLayout := "02/01/2006"

		publicationDate, err := time.Parse(dateLayout, publicationDateStr)
		if err != nil {
			panic(err)
		}

		title := s.Find("h5.my-1 a").First().Text()
		title = strings.TrimSpace(title)

		description := s.Find("p.mb-0").Eq(1).Text()
		description = strings.TrimSpace(description)

		fileURL, _ := s.Find("p.mb-0 a").First().Attr("href")
		fileURL = strings.TrimSpace(fileURL)

		postURL, _ := s.Find("h5.my-1 a").First().Attr("href")
		postURL = strings.TrimSpace(postURL)

		jobOffer := &ent.JobOffer{
			Title:       title,
			Description: description,
			PublishedAt: publicationDate,
			URL:         postURL,
			FileURL:     fileURL,
		}

		jobOffers = append(jobOffers, jobOffer)
	})

	return jobOffers
}

// saveJobOffers saves the job offers to the database.
func (s *service) saveJobOffers(jobOffers []*ent.JobOffer, input *ent.SearchInput) {
	for _, jobOffer := range jobOffers {
		if !s.client.JobOffer.Query().
			Where(joboffer.URLEQ(jobOffer.URL), joboffer.PublishedAtEQ(jobOffer.PublishedAt)).
			ExistX(context.Background()) {

			_, err := s.client.JobOffer.Create().
				SetTitle(jobOffer.Title).
				SetDescription(jobOffer.Description).
				SetPublishedAt(jobOffer.PublishedAt).
				SetURL(jobOffer.URL).
				SetFileURL(jobOffer.FileURL).
				SetUpdatedAt(time.Now()).
				AddSearchInputs(input).
				Save(context.Background())
			if err != nil {
				panic(err)
			}
			s.bot.Send(tgbotapi.NewMessage(s.chatID, jobOffer.Title))
		} else {
			fmt.Println("Job offer already exists")
		}
	}
}
