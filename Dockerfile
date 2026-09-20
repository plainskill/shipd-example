# shipd example app: tiny Go HTTP server (no framework, stdlib only)
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/app .

FROM alpine:3.20
RUN adduser -D -u 10007 app
USER app
COPY --from=build /out/app /app
EXPOSE 8000
CMD ["/app"]
