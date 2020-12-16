package helpers

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"io/ioutil"
	"net/http"
)

func AutomaticEmail() {
	job := func() {
		response, err := http.Get("http://localhost:8202/api/v1/")
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}
	//scheduler.Every(2).Seconds().Run(job)
	//scheduler.Every().Day().Run(job)
	scheduler.Every().Day().At("06:30").Run(job)

}
