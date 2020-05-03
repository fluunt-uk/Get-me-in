package templates

import (
	"encoding/json"
	"fmt"
	s "github.com/ProjectReferral/Get-me-in/customer-api/models"
	"github.com/matcornic/hermes"
	"log"
)

var staticTemplates = make(map[string]*s.BaseEmail)

// Configure templates by setting a theme and your product info
var h = hermes.Hermes{
    // Optional Theme
    // Theme: new(Default) 
    Product: hermes.Product{
        // Appears in header & footer of e-mails
        Name: "GMI Team",
        // Optional product logo
        Logo: "https://www.clipartmax.com/png/middle/425-4252869_blank-raffle-tickets-template-free-ticket-booking-icon-png.png",
        Copyright: "Copyright © 2020 GMI. All rights reserved.",
        TroubleText: "",
    },
}

func ToStruct(d []byte, p interface{}){
	err := json.Unmarshal(d, &p)
	if err != nil {
		fmt.Println(err)
	}
}

func StringParsedHTML(e hermes.Email) string {
	emailBody, err := h.GenerateHTML(e)
	if err != nil {
		failOnError(err, "Failed to generate HTML email")
	}
	return emailBody
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}