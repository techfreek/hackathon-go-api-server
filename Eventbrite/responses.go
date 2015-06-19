package eventbrite

import (
	"time"
)

type TeamsResponse struct {
	Pagination Pages `json:"pagination"`
	Teams []Team `json:"teams"`
}

type Pages struct {
	Count int `json:"object_count"`
	PageNum int `json:"page_number"`
	PageSize int `json:"page_size"`
	PageCount int `json:"page_count"`
}

type Email struct {
	Verified string `json:"verified"`
	Primary string `json:"primary"`
	Email string `json:"email"`
}

type Team struct {
	ResourceURI string `json:"resource_uri"`
	AttendeesURI string `json:"attendees_uri"`
	ID string `json:"id"`
	TeamCreator Creator `json:"creator"`
	Changed time.Time `json:"changed"`
	Created time.Time `json:"created"`
	Count int `json:"attendee_count"`
}

//I am not creative at naming
type TextHtml struct {
	Text string `json:"text"`
	Html string `json:"html"`
}

type TimeInfo struct {
	Timezone string `json:"timezone"`
	Local time.Time `json:"local"`
	End time.Time `json:"end"`
}

type Event struct {
	ResourceURI string `json:"resource_uri"`
	Name TextHtml `json:"name"`
	Description TextHtml `json:"Description"`
	ID string `json:"id"`
	URL string `json:"url"`
	Changed time.Time `json:"changed"`
	Created time.Time `json:"created"`
	Capacity int `json:"capacity"`
	Status string `json:"status"`
	Currency string `json:"currency"`
	Listed bool `json:"listed"`
	Shareable bool `json:"shareable"` 
	InviteOnly bool `json:"invite_only"`
	OnlineEvent bool `json:"online_event"`
	ShowRemaining bool `json:"show_remaining"`
	TxTimeLimit int `json:"tx_time_limit"`
	LogoID string `json:"logo_id"`
	OrganizerID string `json:"organizer_id"`
	VenueID string `json:"venue_id"`
	CategoryID string `json:"category_id"`
	SubCategoryID string `json:"subcategory_id"`
	FormatID string `json:"format_id"`
}

type Creator struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	ID string `json:"id"`
	Name string `json:"name"`
	Emails []Email `json:"Email"`
}