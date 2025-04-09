FROM golang AS builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -ldflags="-w -s" -o /app/caddy

FROM gcr.io/distroless/static-debian12
COPY --chmod=755 --from=builder /app/caddy /bin/caddy
LABEL org.opencontainers.image.source="https://github.com/xruins/docker-images"
ENTRYPOINT ["/bin/caddy"]
CMD ["run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]