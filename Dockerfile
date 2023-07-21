FROM golang:alpine AS builder

COPY . /app

# Change working directory in the container
WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache make gcc g++

RUN make build

# Remove unused files
RUN apk del make gcc g++ git
RUN rm -rf /var/cache/apk/*
RUN find / -name "*_test.go" -delete


FROM alpine

RUN apk add --no-cache tzdata

COPY --from=builder /app/bin/application /bin/application

EXPOSE 8000
ENTRYPOINT ["/bin/application"]
