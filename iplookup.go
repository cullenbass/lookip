package main

import (
	"errors"
	"github.com/oschwald/maxminddb-golang"
	"log"
)

type PostedJson struct {
	Ip string `json:"ip"`
}

// Tells maxminddb what to return instead of querying entire row.
type Record struct {
	Location struct {
		Latitude  float32 `maxminddb:"latitude"`
		Longitude float32 `maxminddb:"longitude"`
	} `maxminddb:"location"`
}

// JSON encoding- the vars have to be exported, so capitals are needed.
type LookupData struct {
	Longitude  float32 `json:longitude`
	Latitude   float32 `json:latitude`
	StatusCode uint32  `json:statusCode`
}

// Yes, I used global variables here, but I didn't want to have to keep
// opening and closing the DB. Since all access happens inside this file,
// I'm not too worried.
var started bool = false
var db *maxminddb.Reader

func initDb() error {
	var err error
	db, err = maxminddb.Open("./GeoLite2-City.mmdb")
	return err
}

// I could have used a goroutine and some channeles, but I didn't. Not sure why.
func lookup(logData *LogData) (LookupData, error) {
	if !started {
		err := initDb()
		if err != nil {
			log.Fatal(err)
		}
	}
	var record Record
	err := db.Lookup(logData.Ip, &record)
	var dat LookupData
	if err == nil {
		dat = LookupData{
			Longitude:  record.Location.Longitude,
			Latitude:   record.Location.Latitude,
			StatusCode: logData.StatusCode,
		}
		return dat, nil
	} else {
		return LookupData{}, errors.New("ERROR PARSING")
	}
}
