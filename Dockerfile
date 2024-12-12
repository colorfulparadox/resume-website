# Build.
FROM golang:1.23.2 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/data /data
COPY --from=build-stage /app/static /static
COPY --from=build-stage /app/project.html /project.html
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]