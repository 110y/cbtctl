FROM golang:1.17.3-buster AS builder

WORKDIR /go/src/github.com/110y/cbtctl

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 go build -tags "netgo,osusergo" -installsuffix netgo -o /usr/local/src/cbtctl .

FROM gcr.io/distroless/base@sha256:4f25af540d54d0f43cd6bc1114b7709f35338ae97d29db2f9a06012e3e82aba8

COPY --from=builder /usr/local/src/cbtctl /cbtctl
ENTRYPOINT ["/cbtctl"]
