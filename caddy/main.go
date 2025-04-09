package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// plug in Caddy modules here
	_ "github.com/caddyserver/caddy/v2/modules/standard"
        _ "github.com/caddy-dns/cloudflare"
        _ "github.com/mholt/caddy-dynamicdns"
        _ "github.com/gerolf-vent/caddy-vault-storage"
        _ "github.com/mentimeter/caddy-storage-cf-kv"
)

func main() {
	caddycmd.Main()
}
