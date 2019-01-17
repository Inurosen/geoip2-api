package geolite

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

type GeoLite struct {
	db string
}

func New(dbPath string) GeoLite {
	return GeoLite{db: dbPath}
}

func (g GeoLite) GetInfo(ip net.IP) (*IpInfo, error) {
	db, err := geoip2.Open(g.db)
	if err != nil {
		log.Println("geoip2.Open:", err)
		return nil, err
	}
	defer db.Close()

	log.Println("IP:", ip)

	info, err := db.City(ip)
	if err != nil {
		log.Println("db.City: ", err)
		return nil, err
	}

	ipInfo := &IpInfo{
		ip.String(),
		City{Id: info.City.GeoNameID, Name: info.City.Names["en"]},
		Continent{Id: info.Continent.GeoNameID, Code: info.Continent.Code, Name: info.Continent.Names["en"]},
		Country{Id: info.Country.GeoNameID, Code: info.Country.IsoCode, Name: info.Country.Names["en"]},
		Location{
			AccuracyRadius: info.Location.AccuracyRadius,
			Latitude:       info.Location.Latitude,
			Longitude:      info.Location.Longitude,
			MetroCode:      info.Location.MetroCode,
			TimeZone:       info.Location.TimeZone,
		},
	}

	return ipInfo, nil
}

type IpInfo struct {
	Ip        string `json:"ip"`
	City      `json:"city"`
	Continent `json:"continent"`
	Country   `json:"country"`
	Location  `json:"location"`
}
type Continent struct {
	Id   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
type Country struct {
	Id   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Location struct {
	AccuracyRadius uint16  `json:"accuracy_radius"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	MetroCode      uint    `json:"metro_code"`
	TimeZone       string  `json:"time_zone"`
}

type City struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
