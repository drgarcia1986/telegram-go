package server

import (
	"github.com/drgarcia1986/telegram-go/client"
	"github.com/drgarcia1986/telegram-go/router"
)

type Server struct {
	router router.Router
	client *client.Client
}

func (s *Server) HandleFunc(pattern string, handler router.Handler) {
	s.router.Register(pattern, handler)
}

func (s Server) processMessage(message client.Message) {
	resp, err := s.router.Resolve(message.Text)
	if err != nil {
		return
	}
	s.client.SendMessage(message.Chat.ID, message.ID, resp)
}

func (s Server) Run() error {
	offset := 0
	for {
		updates, err := s.client.GetUpdates(offset + 1)
		if err != nil {
			continue
		}
		for _, update := range updates {
			if update.ID > offset {
				offset = update.ID
			}
			go s.processMessage(update.Message)
		}
	}
	return nil
}

func New(token string) *Server {
	router := router.New()
	client := client.New(token)
	return &Server{router: router, client: client}
}
