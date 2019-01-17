package response

import (
	"encoding/json"
	"fmt"
	"geoip/pkg/geolite"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Println("json.Marshal: ", err)
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"result\": null, \"message\": \"Internal Server Error\"}")
	}
	w.WriteHeader(response.Code)
	fmt.Fprintf(w, string(responseJson))
}

type Response struct {
	Code    int             `json:"-"`
	Result  *geolite.IpInfo `json:"result"`
	Message string          `json:"message,omitempty"`
}
