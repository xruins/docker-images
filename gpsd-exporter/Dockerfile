FROM golang AS builder

WORKDIR /src
RUN git clone --depth=1 https://github.com/natesales/gpsd-exporter /src && \
    go build -ldflags="-w -s" -o /gpsd-exporter

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder --chmod=755 /gpsd-exporter /usr/local/bin/gpsd-exporter
ENTRYPOINT ["gps-exporter"]