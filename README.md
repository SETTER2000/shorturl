# shorturl

демо: https://rooder.ru
--

# API
* baseurl = https://rooder.ru/
* domain = {{baseurl}}/api
### POST
Отправить текстом длинный url например:  http://example.com/sdg/dfg/dfgh
```azure
POST {{baseurl}}/
```
или JSON например:

{
    "url":"http://example.com/sdg/dfg/dfgh"
}
```azure
POST {{domain}}/shorten
```
### GET
Получить все свои адреса
```azure
GET {{domain}}/user/urls
```


# Начало работы


## Compile
```azure
go build -o cmd/shortener/shortener cmd/shortener/main.go
```

## Run
Запуск сервера без поддержки https. 

В этом случаи сервер будет доступен 
только в 
локальной сети на порте :8080 
```azure
./cmd/shortener/shortener
```
Запрос со второго терминала
```azure
curl http://localhost:8080
```
## Запуск сервера с поддержкой https
Здесь для примера используется домен rooder.ru.

Для того чтоб сервис правильно отвечал на запросы по https протоколу, его нужно разместить на хосте с белым IP.

Для начала нужно сгенерировать сертификат.

### Генерация сертификата SSL/TLS
1. Генерим RSA ключи
```azure
mkdir certs && openssl genrsa -out certs/dev_rsa.key 4096
```
2. Генерим CSR (certificate signing request) 
(что указывать - не имеет значения):
```azure
openssl req -new -key certs/dev_rsa.key -out certs/dev.csr
```
3. Генерим self-signed сертификат
```azure
openssl x509 -req -days 365 -in certs/dev.csr -signkey certs/dev_rsa.key -out certs/dev.crt
```
4. Получаем инфо. сертификата (опционально)
```azure
openssl x509 -in certs/dev.crt -text -noout
```
5. Компилируем и запускаем
```azure
go build -o cmd/shortener/shortener cmd/shortener/main.go && sudo ./cmd/shortener/shortener -s
```
6. Запрос по сети
```azure
curl https://rooder.ru
```

### Запустить как сервис
```azure
sudo ./cmd/shortener/shortener -s >/dev/null &
```


### Make
Для удобства на сервер установить make
```azure
sudo apt-get update && sudo apt-get install make
```

#### Compile
```azure
make short
```

#### Run
Запустить сервис shorturl
```azure
make run 
```

#### HTTPS
Запустить сервис shorturl с протоколом HTTPS
```azure
make hs
```
### Документация
Запустить сервис с документацией, доступен здесь: http://rooder.ru:6060/pkg/github.com/SETTER2000/shorturl/?m=all	
```azure
make godoc
```

## Флаги
Включить поддержку HTTPS протокола
```azure
sudo ./cmd/shortener/shortener -s
```



### Check ports
************************
Проверить порты открытые на удалённом сервере
в данном примере это сервер mo.ru
```azure
nc -zv rooder.ru 1-9999 2>&1 | grep succeeded!
```

### AST проверка проекта
```
go vet -vettool=$(which cmd/staticlint/staticlint) ./...
```

### Генерация AST в графическом представлении  
```
./cmd/staticlint/ast4 cmd/shortener/main.go | dot -Tsvg -o shorturl.svg
```

# Обновление шаблона

Чтобы иметь возможность получать обновления автотестов и других частей шаблона выполните следующую команды:

```
git remote add -m main template https://github.com/yandex-praktikum/go-musthave-shortener-tpl.git
```

Для обновления кода автотестов выполните команду:

```
git fetch template && git checkout template/main .github
```

затем добавьте полученые изменения в свой репозиторий.

# Запуск автотестов

Для успешного запуска автотестов вам необходимо давать вашим веткам названия вида `iter<number>`, где `<number>` -
порядковый номер итерации.

Например в ветке с названием `iter4` запустятся автотесты для итераций с первой по четвертую.

При мерже ветки с итерацией в основную ветку (`main`) будут запускаться все автотесты.