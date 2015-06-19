package eventbrite

import(
	"log"
	"strings"
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
		"?token",
		config.Config.Event.Oathtoken,
	}

	url := strings.Join(urlParts, "")

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err);
	}

	defer res.Body.Close();

	body, err := ioutil.ReadAll(res.Body)

	log.Print(body);
	return 0
}
