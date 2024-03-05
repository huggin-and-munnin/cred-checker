FROM golang:alpine AS build

WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache make
RUN make build


FROM alpine:3.16.0

WORKDIR /go/src/app
COPY --from=build /go/src/app/app ./

CMD ["./app"]


