 # Собираем бинарный файл
 .PHONY: buil
 build: 
	go build -v ./cmd/apiServer

# Запуск тестов
.PHONY: test
test:
		go test -v -race -timeout 30s ./...

#по умолчанию - собираем бинарник
.DEFAULT_GOAL := build