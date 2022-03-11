FROM golang:1.17

WORKDIR /restaurant_reservation_api

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go install github.com/gthomas0/restaurant_reservation_api

CMD ["restaurant_reservation_api"]
