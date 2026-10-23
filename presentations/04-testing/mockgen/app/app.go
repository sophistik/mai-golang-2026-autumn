package app

import (
	"fmt"

	"go-course/04-testing/mockgen/party"
)

// PartyService зависит только от интерфейсов — тестируется без БД и сети.
type PartyService struct {
	visitorLister party.VisitorLister
	greeter       party.Greeter
}

func NewPartyService(l party.VisitorLister, g party.Greeter) *PartyService {
	return &PartyService{visitorLister: l, greeter: g}
}

func (s *PartyService) GreetVisitors(justNice bool) error {
	visitors, err := s.visitorLister.ListVisitors(party.Nice)
	if err != nil {
		return fmt.Errorf("nice visitors: %w", err)
	}
	if !justNice {
		notNice, err := s.visitorLister.ListVisitors(party.NotNice)
		if err != nil {
			return fmt.Errorf("not-nice visitors: %w", err)
		}
		visitors = append(visitors, notNice...)
	}
	for _, v := range visitors {
		s.greeter.Hello(v.String())
	}
	return nil
}
