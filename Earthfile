VERSION 0.6
FROM golang:1.20

WORKDIR /src

src:
    COPY . .

# run the linter
lint:
    FROM +src
    RUN make lint

# run all tests
test:
    FROM +src
    RUN make test

#build the api executable and create a debug image
build:
    FROM +src
    RUN make build
    RUN mv bin/todo /usr/local/bin/todo

    CMD ["/usr/local/bin/todo"]

    SAVE ARTIFACT /usr/local/bin/todo
    SAVE IMAGE todo:latest


# copy the api executable from +build and package it into a distroless image
distroless:
    FROM gcr.io/distroless/static:nonroot

    COPY +build/todo /usr/local/bin/todo

    USER nonroot:nonroot
    CMD ["/usr/local/bin/todo"]

    SAVE IMAGE todo:latest
