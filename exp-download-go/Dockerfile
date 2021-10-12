# syntax=docker/dockerfile:1

#
# See https://docs.docker.com/language/golang/build-images/ for more information
# on how this Dockerfile was constructed.
#

#
# Build container
#
FROM golang:1.16-buster AS build-image

WORKDIR /exp_download_go

#
# Add build files to the build container and pull
# down dependencies.
#
COPY Makefile ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download

#
# Copy the source into the build container.
#
COPY *.go ./

#
# Buld the application.
#
RUN make

#
# Construct the runtime container.
#
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build-image /exp_download_go/go_app /go_app

#
# The API sever listens on 8081.
#
EXPOSE 8081

ENTRYPOINT ["/go_app"]