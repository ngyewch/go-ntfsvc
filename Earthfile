FROM goreleaser/goreleaser:latest
WORKDIR /workspace

build:
    # Download deps before copying code.
    COPY go.mod go.sum .
    RUN go mod download
    # Copy and build code.
    COPY . .
    RUN goreleaser --snapshot --rm-dist
    SAVE ARTIFACT dist/ntfsvc-client_linux_amd64/ntfsvc-client
