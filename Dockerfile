FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod *.go ./

RUN CGO_ENABLED=0 go build
RUN strip minredir

# Deploy the application binary into a lean image
FROM scratch AS release

WORKDIR /

COPY --from=build /app/minredir /minredir

EXPOSE 8080

USER 2048:2048

ENTRYPOINT ["/minredir"]