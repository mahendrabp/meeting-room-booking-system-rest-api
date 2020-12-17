FROM golang:latest
LABEL maintainer="mahendrabp <>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o /bin/meeting-room-booking-system-rest-api -v .
EXPOSE 8202
CMD ["./bin/meeting-room-booking-system-rest-api"]