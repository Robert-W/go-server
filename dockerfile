##############################
# Layer for building our image
##############################
FROM golang:1.25 AS build

WORKDIR /

ENV CGO_ENABLED=0 \
    GOCACHE=/go-cache \
    GOMODCACHE=/gomod-cache

# Copy everything that is not explicitly called out in the .dockerignore
# if you are adding sensitive files, you must add them to the .dockerignore
COPY . .

# Create a user here that we intend to run the scratch image with
RUN useradd -u 42000 scooby

# Use a cache to ensure repeated builds run faster
RUN --mount=type=cache,target=/go-cache \
    --mount=type=cache,target=/gomod-cache \
    go build -ldflags="-s -w" -o api cmd/api/main.go

############################
# Layer for production image
############################
FROM scratch

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /api /bin/api

# This is the port our API serves on
EXPOSE 3000

USER scooby

ENTRYPOINT ["/bin/api"]
