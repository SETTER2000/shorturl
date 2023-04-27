# Для Win7
CC7=shortenertest-windows-amd64.exe
# Бинарник для windows
BUILD_BIN7=build/shortener.exe
# Покрытие тестами
COVER_OUT=profiles/coverage.out

# Запустить сервис shorturl (shortener) in Memory
short_m:
	go build -o build/shortener.exe cmd/shortener/main.go
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/build/shortener -f=

# Запустить сервис shorturl (shortener) in File
short_f:
	go build -o build/shortener.exe cmd/shortener/main.go
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/build/shortener -f=storage.txt

# Запустить сервис shorturl (shortener) in DB
short_db:
	go build -o build/shortener.exe cmd/shortener/main.go
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/build/shortener -d postgres://postgres:123456@localhost:5432/postgres

cover7:
	go test -v -count 1 -race -coverpkg=./... -coverprofile=$(COVER_OUT) ./...
	go tool cover -func $(COVER_OUT)
	go tool cover -html=$(COVER_OUT)
	rm $(COVER_OUT)

race:
	go test -v -race -count 1 ./...

iter1:
	#$(CC7) -test.v -test.run=^TestIteration1$ -binary-path=$(BUILD_BIN7)
	#shortenertestwindowsamd64.exe -test.v -test.run=^TestIteration1$ -binary-path=build/shortener.exe
