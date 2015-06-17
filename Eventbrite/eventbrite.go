package eventbrite

import(
	"log"
	"net/http"
	"io/ioutil"
	"github.com/WSU-ACM/hackathon-go-api-server"
)

var eventbrite_url = "https://www.eventbriteapi.com/v3/"

func GetTeams() Team[] {

}
func GetRemainingSpots() int {
	url := eventbrite_url + main.Config.Eventbrite.Event_id +
			+ "?token=" + main.Config.Eventbrite.oathToken;

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err);
	}

	defer res.Body.Close();

	body, err := ioutil.ReadAll(res.Body)

	log.Print(body);
}
