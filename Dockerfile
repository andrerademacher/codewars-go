FROM golang:1.20-bookworm
LABEL maintainer="andre.rademacher.business@googlemail.com"

ARG GID=1000
ARG UID=1000
ARG UNAME=codewars

ARG GOCACHE="/home/${UNAME}/.cache/go-build"
ARG GOENV="/home/${UNAME}/.config/go/env"
ARG GOPATH="/home/${UNAME}/go"

ENV PATH="${PATH}:${GOPATH}/bin"

# setup user and home directory
RUN groupadd \
    --gid ${GID} \
    --non-unique \
    ${UNAME}

RUN useradd \
    --create-home \
    --gid ${GID} \
    --home-dir /home/codewars \
    --shell /bin/bash \
    --uid ${UID} \
    ${UNAME}

USER ${UNAME}
WORKDIR /codewars/go120

# setup go
RUN mkdir -p "${GOCACHE}"
VOLUME "${GOCACHE}"
ENV GOCACHE="${GOCACHE}"

RUN mkdir -p "${GOENV}"
VOLUME "${GOENV}"
ENV GOENV="${GOENV}"

RUN mkdir -p "${GOPATH}/pkg"
VOLUME "${GOPATH}/pkg"
ENV GOPATH="${GOPATH}"

# provide qa tooling with golangci-lint
COPY --from=golangci/golangci-lint /usr/bin/golangci-lint /usr/bin/golangci-lint

# install gingko to run tests
RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest

# download and verify dependencies
COPY go.mod ../go.sum  /codewars/go120/
RUN go mod download -x \
    && go mod verify
