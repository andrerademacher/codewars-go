FROM golang:1.20-bullseye
LABEL maintainer="andre.rademacher.business@googlemail.com"

RUN useradd \
    --create-home \
    --groups root \
    --home-dir /home/codewars \
    --shell /bin/bash \
    --system \
    --uid 1000 \
    codewars
USER codewars

# install gingko to run tests
RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest
COPY ../go.mod ../go.sum  /codewars/go120/
WORKDIR /codewars/go120

RUN go mod download -x \
    && go mod verify
