FROM golang AS builder

WORKDIR /app
RUN git clone --depth=1 https://github.com/natesales/gpsd-exporter /app && \
    CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/gpsd-exporter

FROM gcr.io/distroless/base-debian12:nonroot

COPY --from=builder --chmod=755 /app/gpsd-exporter /usr/local/bin/gpsd-exporter
ENTRYPOINT ["/usr/local/bin/gpsd-exporter"]
