# Build.
FROM golang:1.23.2 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Deploy.
FROM alpine:latest AS release-stage
WORKDIR /
COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/project_pages /project_pages
COPY --from=build-stage /app/data /data
COPY --from=build-stage /app/static /static
EXPOSE 8081
ENTRYPOINT ["/entrypoint"]
CMD ["-p", "8081"]

# build cmd: docker build -t counter-basic:latest .
# run cmd:   docker run -p 8081:8081 counter-basic:latest