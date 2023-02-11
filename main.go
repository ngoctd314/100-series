package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://172.18.5.44:8000/analytics/SspDealidReport?FromDate=2023-02-04&ToDate=2023-02-05&dspids=1007,1008,1010,1100,2032,2033&group=date&timezone=7")
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(data))

}
