package service

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// Writes to table with specific id, starting from "from". Only 3 columns are used/
func FillSpreadsheet(from byte, rows [][]interface{}) {
	// We use google service account to manipulate app. It is a robot with GRANT EDIT
	// After we created this account, we can use creds to do stuff
	b, err := ioutil.ReadFile("./src/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := config.Client(oauth2.NoContext)

	// erdauletbatalov's spreadsheet id (named web3space_task1)
	spreadsheetId := "1mqBbMNEonPIPozZ9EA-kzzon81uXP9MKsmEzim0URHA"

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// From [0][0] to [2][2]
	rangeData := fmt.Sprintf("%c1:%c%d", from, from+2, len(rows))
	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}
	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:  rangeData,
		Values: rows,
	})
	_, err = srv.Spreadsheets.Values.BatchUpdate(spreadsheetId, rb).Do()
	if err != nil {
		log.Fatal(err)
	}
}
