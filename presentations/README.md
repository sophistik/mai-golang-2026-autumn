# Презентации

Слайды курса Go в формате [golang.org/x/tools/present](https://pkg.go.dev/golang.org/x/tools/present).

## Структура

- `01-basics/lecture.slide` — лекция 1: введение и основы языка
- `02-interfaces/lecture.slide` — лекция 2: интерфейсы и обработка ошибок
- `03-concurrency/lecture.slide` — лекция 3: горутины, каналы, sync и context
- `04-testing/lecture.slide` — лекция 4: тестирование, бенчмарки и тулинг
- `05-generics-json/lecture.slide` — лекция 5: дженерики, JSON и стандартная библиотека
- `06-database/lecture.slide` — лекция 6: работа с базами данных (Postgres)

Примеры лекции 6 требуют PostgreSQL только при запуске; схема для
практики — `06-database/repo/schema.sql` (`docker run postgres:16`).

Рядом со слайдами лежат каталоги с примерами: `.play`-файлы — полные программы
(со кнопкой Run), `.code` — фрагменты для показа. Импорты локальных пакетов
ссылаются на модуль `go-course` (см. `go.mod`).

## Как показать слайды

Установить present и запустить из этой папки:

    go get -tool golang.org/x/tools/cmd/present
    go tool present

Откроется http://localhost:3999 — перейти в каталог лекции и открыть `lecture.slide`.

Ключи:

- `-play` — запуск `.play`-примеров локально (по умолчанию включён; `-play=0` отключает)
- `-use_playground` — отправлять `.play`-код на go.dev/play вместо локального запуска
- `-notes` — заметки докладчика (строки с префиксом `: `); нажатие `N` открывает второе окно с заметками

## Проверка примеров

Все `.go`-файлы компилируются вместе как один модуль:

    go build ./...
    go vet ./...
