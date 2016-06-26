package dateutils

import (
	"errors"
	"regexp"
	"time"
)

// Regular expressions By Bauke Scholtz:
// http://balusc.omnifaces.org/2007/09/dateutil.html
var DateFormats = map[string]string{
	`^\d{8}$`:                                        `20060102`,
	`^\d{1,2}-\d{1,2}-\d{4}$`:                        `02-01-2006`,
	`^\d{4}-\d{1,2}-\d{1,2}$`:                        `2006-01-02`,
	`^\d{1,2}/\d{1,2}/\d{4}$`:                        `01/02/2006`,
	`^\d{4}/\d{1,2}/\d{1,2}$`:                        `2006/01/02`,
	`(?i)^\d{1,2}\s[a-z]{3}\s\d{4}$`:                 `02 Jan 2006`,
	`(?i)^\d{1,2}\s[a-z]{4,}\s\d{4}$`:                `02 January 2006`,
	`^\d{12}$`:                                       `200601021504`,
	`^\d{8}\s\d{4}$`:                                 `20060102 1504`,
	`^\d{1,2}-\d{1,2}-\d{4}\s\d{1,2}:\d{2}$`:         `02-01-2006 15:04`,
	`^\d{4}-\d{1,2}-\d{1,2}\s\d{1,2}:\d{2}$`:         `2006-01-02 15:04`,
	`^\d{1,2}/\d{1,2}/\d{4}\s\d{1,2}:\d{2}$`:         `01/02/2006 15:04`,
	`^\d{4}/\d{1,2}/\d{1,2}\s\d{1,2}:\d{2}$`:         `2006/01/02 15:04`,
	`(?i)^\d{1,2}\s[a-z]{3}\s\d{4}\s\d{1,2}:\d{2}$`:  `02 Jan 2006 15:04`,
	`(?i)^\d{1,2}\s[a-z]{4,}\s\d{4}\s\d{1,2}:\d{2}$`: `02 January 2006 15:04`,
	`^\d{14}$`:                                             `20060102150405`,
	`^\d{8}\s\d{6}$`:                                       `20060102 150405`,
	`^\d{1,2}-\d{1,2}-\d{4}\s\d{1,2}:\d{2}:\d{2}$`:         `02-01-2006 15:04:05`,
	`^\d{4}-\d{1,2}-\d{1,2}\s\d{1,2}:\d{2}:\d{2}$`:         `2006-01-02 15:04:05`,
	`^\d{1,2}/\d{1,2}/\d{4}\s\d{1,2}:\d{2}:\d{2}$`:         `01/02/2006 15:04:05`,
	`^\d{4}/\d{1,2}/\d{1,2}\s\d{1,2}:\d{2}:\d{2}$`:         `2006/01/02 15:04:05`,
	`(?i)^\d{1,2}\s[a-z]{3}\s\d{4}\s\d{1,2}:\d{2}:\d{2}$`:  `02 Jan 2006 15:04:05`,
	`(?i)^\d{1,2}\s[a-z]{4,}\s\d{4}\s\d{1,2}:\d{2}:\d{2}$`: `02 January 2006 15:04:05`,
}

type DateTime struct {
	date *time.Time
}

// ParseDateTime takes input string and tries to parse it into a date object.
func ParseDateTime(inputDate string) (*DateTime, error) {
	var parsedDate *time.Time
	var err error

	for key, layout := range DateFormats {
		matchResult := regexp.MustCompile(key).MatchString(inputDate)
		if matchResult {
			var date time.Time
			date, err = time.Parse(layout, inputDate)
			parsedDate = &date
			break
		}
	}

	var dateStruct *DateTime

	if (err != nil) || (parsedDate == nil) {
		dateStruct = nil
		err = errors.New("Unable to parse timestamp")
	} else {
		dateStruct = &DateTime{
			date: parsedDate,
		}
		err = nil
	}

	return dateStruct, err
}
