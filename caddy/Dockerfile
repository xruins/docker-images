FROM quay.io/curl/curl AS fetch-stage

ARG CADDY_HASH
ARG CADDY_URL="https://caddyserver.com/api/download?os=linux&arch=amd64&p=github.com%2Fcaddy-dns%2Fcloudflare&p=github.com%2Fmholt%2Fcaddy-dynamicdns&p=github.com%2FWeidiDeng%2Fcaddy-cloudflare-ip&p=github.com%2Fmentimeter%2Fcaddy-storage-cf-kv"
RUN curl -o /home/curl_user/caddy -fsSL --compressed "${CADDY_URL}"

FROM gcr.io/distroless/static-debian12:nonroot
COPY --chmod=755 --from=fetch-stage /home/curl_user/caddy /bin/caddy
LABEL org.opencontainers.image.source="https://github.com/xruins/docker-images"
