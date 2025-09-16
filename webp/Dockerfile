FROM debian:bookworm-slim AS builder

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        wget \
        build-essential \
        cmake \
        libjpeg-dev \
        libpng-dev \
        libtiff-dev \
        libgif-dev && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /build

# Install libwebp with cwebp
RUN wget https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.3.2.tar.gz && \
    tar -xzf libwebp-1.3.2.tar.gz && \
    cd libwebp-1.3.2 && \
    ./configure --prefix=/usr/local && \
    make -j$(nproc) && \
    make install

# Install GNU parallel
RUN wget https://ftp.gnu.org/gnu/parallel/parallel-20230722.tar.bz2 && \
    tar -xjf parallel-20230722.tar.bz2 && \
    cd parallel-20230722 && \
    ./configure --prefix=/usr/local && \
    make -j$(nproc) && \
    make install

FROM gcr.io/distroless/base-debian12:nonroot

COPY --from=builder --chmod=755 /usr/local/bin/cwebp /usr/local/bin/cwebp
COPY --from=builder --chmod=755 /usr/local/bin/dwebp /usr/local/bin/dwebp
COPY --from=builder --chmod=755 /usr/local/bin/parallel /usr/local/bin/parallel
COPY --from=builder /usr/local/lib/libwebp.so* /usr/local/lib/
COPY --from=builder /usr/local/lib/libwebpdemux.so* /usr/local/lib/
COPY --from=builder /usr/local/lib/libwebpmux.so* /usr/local/lib/

ENV PATH="/usr/local/bin:${PATH}"
ENV LD_LIBRARY_PATH="/usr/local/lib"

ENTRYPOINT ["cwebp"]