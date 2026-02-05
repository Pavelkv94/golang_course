FROM golang:1.25-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN chmod +x entrypoint.sh
# разбиваем на 2 этапа и компилируем заранее для последующего корректного завершения программы
RUN go build -o /app/exe main.go

ENTRYPOINT ["/app/entrypoint.sh"]
# разбиваем на 2 этапа и компилируем заранее для последующего корректного завершения программы. 
# так как при поступлении сигнала sigterm, контейнер завершается и мы не можем завершить горутины без последствий
# CMD ["make" "service-run"] - неправильный запуск программы, так как мы не можем завершить горутины без последствий
CMD ["/app/exe"] 
