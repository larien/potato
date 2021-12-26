FROM golang:1.17-alpine AS build

WORKDIR /src

COPY . .

RUN go build -o /out/app .

RUN ls

FROM alpine

LABEL org.opencontainers.image.source="https://github.com/larien/potato"

WORKDIR /app

COPY --from=build /out/app /app/

EXPOSE 8080

ENTRYPOINT ./app