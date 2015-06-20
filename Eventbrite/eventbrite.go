package eventbrite

import(
	"log"
	"fmt"
	"strings"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"github.com/WSU-ACM/hackathon-go-api-server/Config"
)

var eventbrite_url = "https://www.eventbriteapi.com/v3/"

/*func GetTeams() []Team {

}*/

func GetRemainingSpots() int {
	urlParts := []string{
		eventbrite_url,
		config.Config.Event.Event_id,
		"?format=json",
		"&token=",
		config.Config.Event.Oathtoken,
	}

	url := strings.Join(urlParts, "")
	var body map[string]interface{}

	fmt.Printf("url: %s\n", url)

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("get err: %s\n", err);
	}
	defer res.Body.Close()


	data, err := ioutil.ReadAll(res.Body);
	if err != nil {
		fmt.Printf("reader err: %s\n", err);
	}

	//fmt.Printf("data: %s\n", data);

	err = json.Unmarshal(data, &body)
	if err != nil {
		fmt.Printf("Unmarshal err: %s\n", err);
	}

	log.Print(body);
	return 0
}
