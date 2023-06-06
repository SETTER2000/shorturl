test:
	go test -v -count 1 ./...

test100:
	go test -v -count 100 ./...

race:
	go test -v -race -count 1 ./...

# Название файла с точкой входа
MAIN=main.go

# Путь где создать бинарник
APP_DIR=cmd/shortener

# Наименование бинарника
BIN_NAME=shortener

# Покрытие тестами
COVER_OUT=profiles/coverage.out

# Запечь версию и время создания сборки
BAKING="-X 'github.com/SETTER2000/shorturl/internal/app.dateString=`date`' -X 'github.com/SETTER2000/shorturl/internal/app.versionString=`git describe --tags`' -X 'github.com/SETTER2000/shorturl/internal/app.commitString=`git rev-parse HEAD`'"
# Подключение к базе данных
DB=postgres://shorturl:DBshorten-2023@127.0.0.1:5432/shorturl?sslmode=disable

# Запустить сервис shorturl сконфигурировав его от файла json (config/config.json)
conf:
	./$(APP_DIR)/$(BIN_NAME) -c config/config_tt.json

# Запустить сервис shorturl с протоколом HTTPS
hs:
	sudo ./$(APP_DIR)/$(BIN_NAME) -s

# Запустить сервис shorturl и с протоколом HTTPS в фоновом режиме
hsf:
	sudo ./$(APP_DIR)/$(BIN_NAME) -s >/dev/null &

# Запустить сервис shorturl с подключением к DB
# FILE_STORAGE_PATH=;DATABASE_DSN=postgres://shorturl:DBshorten-2023@127.0.0.1:5432/shorturl?sslmode=disable
run:
	./$(APP_DIR)/$(BIN_NAME) -d $(DB)

# Скомпилировать и запустить бинарник сервиса shorturl (shortener) с подключением к DB и запечёнными аргументами сборки
short:
	go build -ldflags $(BAKING) -o cmd/shortener/shortener cmd/shortener/$(MAIN)
	./$(APP_DIR)/$(BIN_NAME)

# Запустить сервис shorturl (shortener) in Memory
short_m:
	go build -ldflags $(BAKING) -o $(APP_DIR)/$(BIN_NAME) $(APP_DIR)/$(MAIN)
	./$(APP_DIR)/$(BIN_NAME) -f= -d=

# Запустить сервис shorturl (shortener) in File
short_f:
	go build -ldflags $(BAKING) -o $(APP_DIR)/$(BIN_NAME) $(APP_DIR)/$(MAIN)
	./$(APP_DIR)/$(BIN_NAME) -f=storage.txt -d=

# Скомпилировать и запустить бинарник сервиса shorturl (shortener) с подключением к DB
short_d:
	go build -tags pro -ldflags $(BAKING) -o $(APP_DIR)/$(BIN_NAME) $(APP_DIR)/$(MAIN)
	./$(APP_DIR)/$(BIN_NAME) -d $(DB)

# Скомпилировать и запустить бинарник сервиса shorturl (shortener) с подключением к DB
# и разрешением чтения заголовка X-Real-IP или X-Forwarded-For
short_x:
	go build -tags pro -ldflags $(BAKING) -o $(APP_DIR)/$(BIN_NAME) $(APP_DIR)/$(MAIN)
	./$(APP_DIR)/$(BIN_NAME) -d $(DB) -resolve_ip_using_header true

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

# Запустить сервис с документацией
# Доступен здесь: http://rooder.ru:6060/pkg/github.com/SETTER2000/shorturl/?m=all
godoc:
	godoc -http rooder.ru:6060

# Запустить сервис с документацией в фоновом режиме
doc:
	godoc -http=rooder.ru:6060  -play >/dev/null &

####################################
# Для Win7
CC7=shortenertest-windows-amd64.exe
BIN_NAME_WIN=shortener.exe

# Запустить сервис shorturl (shortener) in Memory
short7_m:
	go build -o $(APP_DIR)/$(BIN_NAME_WIN) cmd/shortener/$(MAIN)
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/$(APP_DIR)/shortener -f=

# Запустить сервис shorturl (shortener) in File
short7_f:
	go build -o $(APP_DIR)/$(BIN_NAME_WIN) cmd/shortener/$(MAIN)
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/$(APP_DIR)/shortener -f=storage.txt

# Запустить сервис shorturl (shortener) in DB
short7_d:
	go build -o $(APP_DIR)/$(BIN_NAME_WIN) cmd/shortener/$(MAIN)
	D:\__PROJECTS\GoProjects\Y.Praktikum\Projects\shorturl/$(APP_DIR)/shortener -d postgres://postgres:123456@localhost:5432/postgres
