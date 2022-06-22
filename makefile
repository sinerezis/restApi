 # Собираем бинарный файл
 .PHONY: build
 build: 
	go build -v ./cmd/apiServer

# Запуск тестов
.PHONY: test
test:
		go test -v -race -timeout 10s ./...

#по умолчанию - собираем бинарник
.DEFAULT_GOAL := build