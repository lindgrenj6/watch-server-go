##########
# Dockerfile only here as an exercise. 
#
# This build wouldn't be able to do anything unless I mounted scripts inside the container, 
# but hey its a 8.2MB image at the end. 
##########

FROM golang:1.15 as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 go build -o watch-server

FROM scratch
COPY --from=builder /build/watch-server /watch-server
ENTRYPOINT ["/watch-server"]
