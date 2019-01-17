# GeoIP2 API
JSON API service for retrieving info for IP or domain from maxmind's GeoLite2 and GeoIP2 database

This project was made just for practicing while learning Golang

# Dependencies
* https://github.com/oschwald/geoip2-golang
* https://maxmind.com

# Usage
## 1. Configure
Configuration is performed through following environment variables:
* `GEOIP_DB` - path to maxmind's city mmdb database file
* `HOST` - host or ip of web server listener
* `PORT` - port of web server listener

## 2. Run Example
```
env PORT=7777 HOST=127.0.0.1 GEOIP_DB=GeoLite2-City.mmdb bin/geoip
``` 

## 3. Call
### Request Example
```
curl http://localhost:7777/?ip=github.com
curl http://localhost:7777/?ip=140.82.118.3
```
### Response Example
```json
{
  "result": {
    "ip": "140.82.118.3",
    "city": {
      "id": 0,
      "name": ""
    },
    "continent": {
      "id": 6255149,
      "code": "NA",
      "name": "North America"
    },
    "country": {
      "id": 6252001,
      "code": "US",
      "name": "United States"
    },
    "location": {
      "accuracy_radius": 1000,
      "latitude": 37.751,
      "longitude": -97.822,
      "metro_code": 0,
      "time_zone": ""
    }
  }
}
```

### Error Example

```json
{
  "result": null,
  "message": "Internal Server Error"
}
```
