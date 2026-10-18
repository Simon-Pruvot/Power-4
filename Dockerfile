# syntax=docker/dockerfile:1
# Puissance 4 Web — serveur Go (templates HTML), un seul conteneur.

FROM golang:1.25-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY src ./src
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/power4 ./src

FROM alpine:3.22
RUN apk add --no-cache tzdata && adduser -D -H app
WORKDIR /app
COPY --from=build /out/power4 ./power4
# Le serveur lit templates, CSS et images relativement à /app (src/...).
COPY src/template ./src/template
COPY src/CSS ./src/CSS
COPY src/images ./src/images
# Photos prises avec la webcam (privées à chaque visiteur, éphémères).
RUN mkdir -p /app/uploads && chown app:app /app/uploads
USER app
ENV PORT=8080 \
    UPLOAD_DIR=/app/uploads \
    TZ=Europe/Paris
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget -qO /dev/null "http://127.0.0.1:${PORT}/regle" || exit 1
CMD ["./power4"]
