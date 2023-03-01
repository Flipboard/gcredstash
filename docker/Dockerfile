FROM golang:buster AS builder

ARG TARGETARCH

ENV GOOS=linux
ENV GOARCH=${TARGETARCH}
ENV GOPATH=/go

# build gcredstash
COPY . /go/src/github.com/Flipboard/gcredstash
WORKDIR /go/src/github.com/Flipboard/gcredstash
RUN make go-get
RUN go mod tidy
RUN CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo -o gcredstash-linux-${TARGETARCH}
RUN ls -la

# Copy the build artifact so we can dump it using the --output argument
FROM scratch AS export
ARG TARGETARCH
COPY --from=builder /go/src/github.com/Flipboard/gcredstash/gcredstash-linux-${TARGETARCH} .
