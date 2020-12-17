package helpers

import (
	"fmt"
	"github.com/carlescere/scheduler"
	_ "github.com/joho/godotenv/autoload"
	"io/ioutil"
	"net/http"
	"os"
)

func AutomaticEmail() {
	job := func() {
		response, err := http.Get(os.Getenv("CRON_ENDPOINT"))
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			fmt.Println(string(data))
		}
	}

	scheduler.Every().Day().At("06:30").Run(job)

}
