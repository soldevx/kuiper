FROM golang:1.16.6-buster

LABEL maintainer="EPI"
LABEL description="image supporting kuiper-andro api framework"
LABEL version="1.0"

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


# -------------------------------------
# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY ./andro/go.mod .
COPY ./andro/go.sum .
RUN go mod download

# Copy the code into the container
COPY ./andro .


# -------------------------------------
# Move to /cmd/api directory and build app
WORKDIR /build/cmd/api

# Build the application
RUN go build 


# -------------------------------------
# Move to /cmd/migration directory and build migration
WORKDIR /build/cmd/migration

# Build the application
RUN go build 

# -------------------------------------
# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/cmd/api/api .
RUN cp /build/cmd/api/conf.local.yaml .
RUN cp /build/cmd/migration/migration .
RUN cp -R /build/assets .

# -------------------------------------
# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/dist/api"]













