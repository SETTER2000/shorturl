test:
	go test -v -count 1 ./...

test100:
	go test -v -count 100 ./...

race:
	go test -v -race -count 1 ./...

# Путь где создать бинарник
BIN_PATH=cmd/shortener

# Наименование бинарника
BIN_NAME=shortener

# Покрытие тестами
COVER_OUT=profiles/coverage.out

#.PHONY: gen
#gen:
#	mockgen -source=internal/usecase/interfaces.go -destination=internal/usecase/mocks/mock_interfaces.go

# Запустить сервис shorturl (shortener) in Memory
short_m:
	go build -o $(BIN_PATH)/$(BIN_NAME) $(BIN_PATH)/*.go
	./$(BIN_PATH)/$(BIN_NAME) -f= -d=

# Запустить сервис shorturl (shortener) in File
short_f:
	go build -o $(BIN_PATH)/$(BIN_NAME) $(BIN_PATH)/*.go
	./$(BIN_PATH)/$(BIN_NAME) -f=storage.txt -d=

# Запустить сервис shorturl (shortener) in DB
short_d:
	go build -o $(BIN_PATH)/$(BIN_NAME) $(BIN_PATH)/*.go
	./$(BIN_PATH)/$(BIN_NAME) -d postgres://shorturl:DBshorten-2023@127.0.0.1:5432/shorturl?sslmode=disable

short:
	go build -o $(BIN_PATH)/$(BIN_NAME) $(BIN_PATH)/*.go
	./$(BIN_PATH)/$(BIN_NAME)

short2:
	go run -ldflags "-X main.Version=v1.0.1 -X 'main.BuildTime=$(date +'%Y/%m/%d %H:%M:%S')'" $(BIN_PATH)/main.go


cover:
	go test -v -count 1 -race -coverpkg=./... -coverprofile=$(COVER_OUT) ./...
	go tool cover -func $(COVER_OUT)
	go tool cover -html=$(COVER_OUT)
	rm $(COVER_OUT)

cover1:
	go test -v -count 1  -coverpkg=./... -coverprofile=cover.out.tmp ./...
	cat cover.out.tmp | grep -v mocks/*  > cover.out2.tmp
	cat cover.out2.tmp | grep -v log/*  > $(COVER_OUT)
	go tool cover -func $(COVER_OUT)
	go tool cover -html=$(COVER_OUT)
	rm cover.out.tmp cover.out2.tmp
	rm $(COVER_OUT)


####################################
# Для Win7
CC7=shortenertest-windows-amd64.exe
BIN_NAME_WIN=shortener.exe

# Запустить сервис shorturl (shortener) in Memory
short7_m:
	go build -o $(BIN_PATH)/$(BIN_NAME_WIN) cmd/shortener/main.go
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/$(BIN_PATH)/shortener -f=

# Запустить сервис shorturl (shortener) in File
short7_f:
	go build -o $(BIN_PATH)/$(BIN_NAME_WIN) cmd/shortener/main.go
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/$(BIN_PATH)/shortener -f=storage.txt

# Запустить сервис shorturl (shortener) in DB
short7_d:
	go build -o $(BIN_PATH)/$(BIN_NAME_WIN) cmd/shortener/main.go
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/$(BIN_PATH)/shortener -d postgres://postgres:123456@localhost:5432/postgres
