FROM golang AS builder

WORKDIR /coredns
RUN git clone --depth=1 https://github.com/coredns/coredns /coredns && \
    echo "unifi:github.com/xruins/unifi-coredns" >> plugin.cfg && \
    echo "tailscale:github.com/xruins/tailscale-coredns" >> plugin.cfg && \
    echo "alias:github.com/serverwentdown/alias" >> plugin.cfg && \
    make

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder --chmod=755 /coredns/coredns /usr/local/bin/coredns
ENTRYPOINT ["coredns"]
