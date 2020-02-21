FROM golang:1.12 as builder
WORKDIR /go/src/github.com/promoboxx/<packageName>/
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/<packageName> main.go

FROM alpine:3.7
# Install some common tools often needed during deploys
RUN apk -v --update add ca-certificates bash jq git openssh python py-pip && \
    pip install --upgrade awscli==1.14.5 && \
    apk -v --purge del py-pip && \
    rm /var/cache/apk/*
WORKDIR /
COPY --from=builder /go/src/github.com/promoboxx/<packageName>/bin/<packageName> .
