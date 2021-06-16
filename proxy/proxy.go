package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var servers = make(map[string]*httputil.ReverseProxy)

func findServer(url *url.URL) *httputil.ReverseProxy {
	nice := url.String()

	if server, ok := servers[nice]; ok {
		return server
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	servers[nice] = proxy

	return proxy
}

func proxyRequest(url *url.URL, res http.ResponseWriter, req *http.Request) {
	proxy := findServer(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	proxy.ServeHTTP(res, req)
}

func findEndpoint(req *http.Request) *url.URL {
	endpoints := []string{
		"http://server_1",
		"http://server_2",
		"http://server_3",
	}

	endpoint, _ := url.Parse(endpoints[rand.Intn(3)])

	return endpoint
}

func onRequest(res http.ResponseWriter, req *http.Request) {
	endpoint := findEndpoint(req)
	log.Println("This is a request")
	log.Printf("I'll redirect it to: %v", endpoint.String())

	proxyRequest(endpoint, res, req)
}

func main() {
	log.Println("main")

	http.HandleFunc("/", onRequest)
	http.HandleFunc("/ws", onSocketRequest)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		panic(err)
	}
}

func onSocketRequest(res http.ResponseWriter, req *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(req, res)
	if err != nil {
		log.Println("conn")
		log.Println(err)
	}

	go func() {
		defer conn.Close()

		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				log.Println("reading error")
				log.Panic(err)
			}
			log.Println(string(msg))

			// err = wsutil.WriteServerMessage(conn, op, msg)
			// if err != nil {
			// 	log.Println("write error")
			// 	log.Println(err)
			// }
		}
	}()
}
