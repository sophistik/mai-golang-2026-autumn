package app

import (
	"errors"
	"testing"

	"go-course/04-testing/mockgen/app/mocks"
	"go-course/04-testing/mockgen/party"
	"go.uber.org/mock/gomock"
)

// Паттерн «моки в табличных тестах» по мотивам статьи Авито:
// https://habr.com/ru/companies/avito/articles/658907/
func TestPartyService_GreetVisitors(t *testing.T) {
	type fields struct {
		visitorLister *mocks.MockVisitorLister
		greeter       *mocks.MockGreeter
	}
	type args struct {
		justNice bool
	}

	tests := []struct {
		name    string
		prepare func(f *fields) // ожидания конкретного кейса
		args    args
		wantErr bool
	}{
		{
			name: "список nice упал — ошибка",
			prepare: func(f *fields) {
				f.visitorLister.EXPECT().
					ListVisitors(party.Nice).
					Return(nil, errors.New("db down"))
			},
			args:    args{justNice: true},
			wantErr: true,
		},
		{
			name: "ошибка на not-nice — тоже ошибка",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.visitorLister.EXPECT().ListVisitors(party.Nice).
						Return([]party.Visitor{{Name: "Пётр"}}, nil),
					f.visitorLister.EXPECT().ListVisitors(party.NotNice).
						Return(nil, errors.New("db down")),
				)
			},
			args:    args{justNice: false},
			wantErr: true,
		},
		{
			name: "списки ок — приветствия по порядку",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.visitorLister.EXPECT().ListVisitors(party.Nice).
						Return([]party.Visitor{{Name: "Пётр", Surname: "Иванов"}}, nil),
					f.visitorLister.EXPECT().ListVisitors(party.NotNice).
						Return([]party.Visitor{{Name: "Зина", Surname: "Петрова"}}, nil),
					f.greeter.EXPECT().Hello("Пётр Иванов"),
					f.greeter.EXPECT().Hello("Зина Петрова"),
				)
			},
			args:    args{justNice: false},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t) // свой контроллер на каждый кейс
			f := fields{
				visitorLister: mocks.NewMockVisitorLister(ctrl),
				greeter:       mocks.NewMockGreeter(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f) // EXPECT() навешиваются здесь, не в таблице
			}

			s := NewPartyService(f.visitorLister, f.greeter)
			if err := s.GreetVisitors(tt.args.justNice); (err != nil) != tt.wantErr {
				t.Errorf("GreetVisitors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
