FROM golang:1.23.4
#ARG TARGET_URL
#ARG INTERVAL
#ARG BODY_FILE
WORKDIR /app
COPY go.mod ./
#go.sum ./
RUN go mod download
COPY main.go ./
COPY requests/*.json ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-scheduler
CMD [ "/golang-scheduler" ]