# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.15.2 as matrix-rotage

# Install dependencies
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors

# copy the local package files to the container workspace
ADD . /matrix-rotage
#COPY .env /matrix-rotate

# Setting up working directory
WORKDIR /matrix-rotage

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o matrix-rotage .

# alpine image
FROM alpine:3.9.2 as prod
# Setting up working directory
WORKDIR /root/
# copy movies binary
COPY --from=matrix-rotage /matrix-rotage .
#COPY .env .

# Service listens on port 8080
EXPOSE 8080
# Run the movies microservice when the container starts.
ENTRYPOINT ["./matrix-rotage"]
