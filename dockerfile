FROM golang:1.25-alpine AS build

WORKDIR /

ENV CGO_ENABLED=0 \
    GOCACHE=/go-cache \
    GOMODCACHE=/gomod-cache

# Copy everything that is not explicitly called out in the .dockerignore
# if you are adding sensitive files, you must add them to the .dockerignore
COPY . .

RUN --mount=type=cache,target=/go-cache \
    --mount=type=cache,target=/gomod-cache \
    go build -o api cmd/api/main.go


# Create our final production image
FROM scratch

COPY --from=build /api /bin/api

# This is the port our API serves on
EXPOSE 3000

ENTRYPOINT ["/bin/api"]
