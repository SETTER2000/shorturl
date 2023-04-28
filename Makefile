test:
	go test -v -count 1 ./...

# Для Win7
CC7=shortenertest-windows-amd64.exe
# Наименование бинарника
BIN_NAME=shortener.exe
# Бинарник для windows
BUILD_BIN=build

# Путь где создать бинарник
BIN_PATH=cmd/shortener

# Покрытие тестами
COVER_OUT=profiles/coverage.out

# Запустить сервис shorturl (shortener) in Memory
short_m:
	go build -o $(BIN_PATH)/$(BIN_NAME) cmd/shortener/main.go
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/$(BIN_PATH)/shortener -f=

# Запустить сервис shorturl (shortener) in File
short_f:
	go build -o $(BIN_PATH)/$(BIN_NAME) cmd/shortener/main.go
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/$(BIN_PATH)/shortener -f=storage.txt

# Запустить сервис shorturl (shortener) in DB
short_db:
	go build -o $(BIN_PATH)/$(BIN_NAME) cmd/shortener/main.go
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/$(BIN_PATH)/shortener -d postgres://postgres:123456@localhost:5432/postgres

cover7:
	go test -v -count 1 -race -coverpkg=./... -coverprofile=$(COVER_OUT) ./...
	go tool cover -func $(COVER_OUT)
	go tool cover -html=$(COVER_OUT)
	rm $(COVER_OUT)

cover1:
	go test -v -count 1  -coverpkg=./... -coverprofile=cover.out.tmp ./...
	cat cover.out.tmp | grep -v "mock_*.go" > $(COVER_OUT)
	rm cover.out.tmp
	go tool cover -func $(COVER_OUT)
	go tool cover -html=$(COVER_OUT)
	rm $(COVER_OUT)

race:
	go test -v -race -count 1 ./...

.PHONY: gen
gen:
	mockgen -source=internal/usecase/interfaces.go -destination=internal/usecase/mocks/mock_interfaces.go