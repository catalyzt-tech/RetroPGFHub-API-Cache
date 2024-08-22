package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
)

type Server struct {
	app          *fiber.App
	url          string
	cache        *cache.Cache
	cron         *cron.Cron
	authorizeKey string
	corsOrigin   string
}

func main() {

	var authorKey string = os.Getenv("AUTHORIZE_KEY")
	if authorKey == "" {
		panic("Can't find AUTHORIZE_KEY in environment")
	}
	var port string = os.Getenv("PORT")
	if port == "" {
		panic("Can't find PORT in environment")
	}

	// can be * or http://1, http://2
	var corsOrigin string = os.Getenv("CORS_ORIGIN")
	if corsOrigin == "" {
		panic("Can't find corsOrigin in environment")
	}

	var s *Server = &Server{
		app: fiber.New(fiber.Config{
			BodyLimit:      4 * 1024 * 1024,
			RequestMethods: []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE"},
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusNotFound).SendString("hello")
				// return utils.Render(c, layout.NotFoundComponents())
			},
		}),
		url:   port,
		cache: cache.New(cache.NoExpiration, cache.NoExpiration),
		cron: cron.New(
			cron.WithParser(
				cron.NewParser(
					cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow))),
		authorizeKey: authorKey,
		corsOrigin:   corsOrigin,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Saving data to cache . . .")
	_, err := handleCache(s.cache, s.authorizeKey)
	if err != nil {
		panic(err)
	}

	_, err = s.cron.AddFunc("0 */15 * * * *", func() {
		fmt.Println("Cronjob saving data into cache: ", time.Now())
		_, err := handleCache(s.cache, s.authorizeKey)
		if err != nil {
			return
		}
	})

	if err != nil {
		panic(err)
	}

	s.cron.Start()

	cron.New(cron.WithSeconds())

	go s.gracefulShutdown(quit)

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: corsOrigin,
	}))

	cacheApi := newCacheApi(s.cache, s.authorizeKey)

	s.app.Get("/api", cacheApi.HandleCacheApi)

	s.httpListening()
}

func (s *Server) httpListening() {
	if err := s.app.Listen(s.url); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v", err)
	}
}

func (s *Server) gracefulShutdown(quit <-chan os.Signal) {

	log.Printf("Starting service v6")
	<-quit

	log.Printf("Shutting down service")

	if err := s.app.Shutdown(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
