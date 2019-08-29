package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/sfomuseum/go-http-tilezen/http"
	"github.com/sfomuseum/go-http-tilezen/server"	
	"github.com/whosonfirst/go-whosonfirst-cache"
	gohttp "net/http"
	gourl "net/url"
)

func main() {

	var proto = flag.String("protocol", "http", "The protocol for placeholder-client server to listen on. Valid protocols are: http, lambda.")
	host := flag.String("host", "localhost", "The host to listen for requests on.")
	port := flag.Int("port", 8080, "The port to listen for requests on.")

	flag.Parse()

	mux := gohttp.NewServeMux()

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/ping", ping_handler)

	cache_opts, err := cache.DefaultGoCacheOptions()

	if err != nil {
		log.Fatal(err)
	}
	
	go_cache, err := cache.NewGoCache(cache_opts)
		
	proxy_opts := &http.TilezenProxyHandlerOptions{
		Cache: go_cache,
	}
	
	proxy_handler, err := http.TilezenProxyHandler(proxy_opts)

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/", proxy_handler)
	
	address := fmt.Sprintf("http://%s:%d", *host, *port)

	u, err := gourl.Parse(address)

	if err != nil {
		log.Fatal(err)
	}
	
	s, err := server.NewStaticServer(*proto, u)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s\n", s.Address())

	err = s.ListenAndServe(mux)

	if err != nil {
		log.Fatal(err)
	}
	
}
