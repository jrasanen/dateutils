package dateutils

import (
	"fmt"
	"testing"
)

var expectedTimestamps = map[string]string{
	`20181231`:                  `2018-12-31 00:00:00 +0000 UTC`,
	`31-12-2018`:                `2018-12-31 00:00:00 +0000 UTC`,
	`2018-12-31`:                `2018-12-31 00:00:00 +0000 UTC`,
	`12/31/2018`:                `2018-12-31 00:00:00 +0000 UTC`,
	`2018/12/31`:                `2018-12-31 00:00:00 +0000 UTC`,
	`24 Dec 2019`:               `2019-12-24 00:00:00 +0000 UTC`,
	`02 December 2018`:          `2018-12-02 00:00:00 +0000 UTC`,
	`201812311939`:              `2018-12-31 19:39:00 +0000 UTC`,
	`20181231 1959`:             `2018-12-31 19:59:00 +0000 UTC`,
	`31-12-2018 15:04`:          `2018-12-31 15:04:00 +0000 UTC`,
	`2006-01-02 15:04`:          `2006-01-02 15:04:00 +0000 UTC`,
	`12/31/2018 15:04`:          `2018-12-31 15:04:00 +0000 UTC`,
	`2018/12/31 15:04`:          `2018-12-31 15:04:00 +0000 UTC`,
	`31 Dec 2018 15:04`:         `2018-12-31 15:04:00 +0000 UTC`,
	`31 December 2018 15:04`:    `2018-12-31 15:04:00 +0000 UTC`,
	`20181231150405`:            `2018-12-31 15:04:05 +0000 UTC`,
	`20181231 150405`:           `2018-12-31 15:04:05 +0000 UTC`,
	`31-12-2018 15:04:05`:       `2018-12-31 15:04:05 +0000 UTC`,
	`2018-12-31 15:04:05`:       `2018-12-31 15:04:05 +0000 UTC`,
	`12/31/2018 15:04:05`:       `2018-12-31 15:04:05 +0000 UTC`,
	`2018/12/31 15:04:05`:       `2018-12-31 15:04:05 +0000 UTC`,
	`31 Dec 2018 15:04:05`:      `2018-12-31 15:04:05 +0000 UTC`,
	`02 December 2018 15:04:05`: `2018-12-02 15:04:05 +0000 UTC`,
}

func TestParseDateTime(t *testing.T) {

	// Try executing every test case for valid formats
	for tz, expected := range expectedTimestamps {
		parsedTime, err := ParseDateTime(tz)
		if err != nil {
			panic(err)
		}

		// Returned timestamp matches what we expect it to be?
		if parsedTime.date.String() != expected {
			panic(fmt.Sprintf("expected %v, got %v", expected, parsedTime.date))
		}
	}

}

func TestParseDateTimeError(t *testing.T) {
	_, err := ParseDateTime("mehe?")

	if err == nil {
		panic(fmt.Sprintf("expected an error, got nil"))
	}
}

func TestFuzzyTime(t *testing.T) {
	fmt.Println(FuzzyTime("2019-06-26 12:00"))
}
