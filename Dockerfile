FROM golang:1.23.4
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
COPY requests/*.json ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-scheduler
CMD [ "/golang-scheduler" ]