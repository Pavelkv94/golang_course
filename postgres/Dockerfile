FROM golang:1.25-bookworm

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN chmod +x entrypoint.sh

RUN go build -o /app/exe main.go

ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["/app/exe"]
