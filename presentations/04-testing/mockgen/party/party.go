// Package party — домен «вечеринки» по мотивам статьи Авито про gomock:
// https://habr.com/ru/companies/avito/articles/658907/
package party

import "fmt"

//go:generate mockgen -package mocks -destination=../app/mocks/party.go go-course/04-testing/mockgen/party Greeter,VisitorLister

type Visitor struct {
	Name    string
	Surname string
}

func (v Visitor) String() string {
	return fmt.Sprintf("%s %s", v.Name, v.Surname)
}

type Group string

const (
	Nice    Group = "nice"
	NotNice Group = "not-nice"
)

// Интерфейсы объявляет потребитель — сервис (лекция 2).
type Greeter interface {
	Hello(name string) string
}

type VisitorLister interface {
	ListVisitors(group Group) ([]Visitor, error)
}
