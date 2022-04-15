VERSION --explicit-global 0.6
FROM debian:stretch
RUN apt-get -y update && apt-get install -y curl git
RUN git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.9.0
RUN echo ". $HOME/.asdf/asdf.sh" >> /root/.bashrc
ENV PATH="${PATH}:/root/.asdf/shims:/root/.asdf/bin"

RUN asdf plugin add golang
RUN asdf plugin add goreleaser
RUN asdf plugin add task

WORKDIR /workspace

setup:
    COPY .tool-versions .
    RUN asdf install
    COPY go.mod go.sum .
    RUN go mod download
    COPY . .

build:
    FROM +setup

    ARG EARTHLY_GIT_HASH
    ARG EARTHLY_GIT_COMMIT_TIMESTAMP

    RUN GIT_COMMIT=$EARTHLY_GIT_HASH GIT_COMMIT_TIMESTAMP=$EARTHLY_GIT_COMMIT_TIMESTAMP goreleaser --snapshot --rm-dist
    SAVE ARTIFACT dist
    SAVE ARTIFACT dist/ntfsvc-client_linux_amd64/ntfsvc-client
