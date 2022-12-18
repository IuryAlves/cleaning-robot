FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . ./

RUN go build -o cleaning-robot

FROM  alpine:3.16

WORKDIR /app
COPY --from=build /app/cleaning-robot /app/cleaning-robot
EXPOSE 8080
ENTRYPOINT [ "./cleaning-robot" ]