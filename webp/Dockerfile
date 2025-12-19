FROM debian:bookworm-slim AS builder

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        curl \
        ca-certificates \
        webp \
        parallel \
        libavif-bin && \
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["cwebp"]
