package main

import (
	"fmt"
	"geoip/pkg/configuration"
	"geoip/pkg/geoip"
	"geoip/pkg/response"
	"log"
	"net"
	"net/http"
	"os"
)

const GEOIP_DB string = "GEOIP_DB"
const HOST string = "HOST"
const PORT string = "PORT"

var config configuration.Configuration

func main() {
	server := config.Get(HOST) + ":" + config.Get(PORT)
	http.HandleFunc("/", getIpInfo)
	fmt.Println("Starting server at " + server)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func getIpInfo(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("ip")
	if query == "" {
		response.Json(w, response.Response{Code: 200, Result: nil, Message: "Missing parameter: ip"})
		return
	}

	ip := net.ParseIP(query)
	if ip == nil {
		ipv4, err := net.LookupIP(query)
		if err != nil {
			log.Println("net.LookupIP:", err)
			response.Json(w, response.Response{Code: 200, Result: nil, Message: err.Error()})
			return
		}
		if len(ipv4) > 0 {
			ip = net.ParseIP(ipv4[0].String())
		}
		if ip == nil {
			response.Json(w, response.Response{Code: 200, Result: nil, Message: "can't resolve IP"})
			return
		}
	}

	service := geoip.New(config.Get("GEOIP_DB"))

	ipInfo, err := service.GetInfo(ip)
	if err != nil {
		response.Json(w, response.Response{Code: 200, Result: nil, Message: "Internal Server Error"})
		return
	}

	response.Json(w, response.Response{Code: 200, Result: ipInfo})
	return
}

func init() {
	config = configuration.New()
	config.Set(GEOIP_DB, getenv(GEOIP_DB, "GeoLite2-City.mmdb"))
	config.Set(HOST, getenv(HOST, "localhost"))
	config.Set(PORT, getenv(PORT, "7777"))
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
