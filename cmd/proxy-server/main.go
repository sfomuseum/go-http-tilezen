package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-http-ping/v2"
	"github.com/aaronland/go-http-server"		
	"github.com/sfomuseum/go-http-tilezen/http"
	"github.com/whosonfirst/go-cache"
	"log"
	gohttp "net/http"
	"time"
)

func main() {

	server_uri := flag.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")
	timeout_seconds := flag.Int("timeout", 30, "The maximum number of seconds to allow for fetching a given tile")

	flag.Parse()

	ctx := context.Background()

	mux := gohttp.NewServeMux()

	ping_handler, err := ping.PingPongHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/ping", ping_handler)

	go_cache, err := cache.NewCache(ctx, "gocache://")

	timeout := time.Duration(*timeout_seconds) * time.Second

	proxy_opts := &http.TilezenProxyHandlerOptions{
		Cache:   go_cache,
		Timeout: timeout,
	}

	proxy_handler, err := http.TilezenProxyHandler(proxy_opts)

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/", proxy_handler)

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create server for '%s', %v", *server_uri, err)
	}
	
	log.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatal(err)
	}

}
