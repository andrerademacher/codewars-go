FROM golang:1.20-bookworm
LABEL maintainer="andre.rademacher.business@googlemail.com"

ARG GID=1000
ARG UID=1000
ARG UNAME=codewars

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
RUN mkdir -p "/home/${UNAME}/go/pkg"
VOLUME /home/${UNAME}/go/pkg
ENV GOPATH="/home/${UNAME}/go"
ENV PATH="${PATH}:/home/${UNAME}/go/bin"

# provide qa tooling with golangci-lint
COPY --from=golangci/golangci-lint /usr/bin/golangci-lint /usr/bin/golangci-lint

# install gingko to run tests
RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest

# download and verify dependencies
COPY go.mod ../go.sum  /codewars/go120/
RUN go mod download -x \
    && go mod verify
