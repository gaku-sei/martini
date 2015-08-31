package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// The DB struct
// Wraps the gorm.DB and adds methods
type DB struct {
	gorm.DB
}

func (db *DB) Fetch(i interface{}) Fetcher {
	return Fetcher{db, i}
}

func OpenDatabase(dbname string) *DB {
	if db, err := gorm.Open("sqlite3", fmt.Sprintf("./%s.db", dbname)); err != nil {
		panic(err)
	} else {
		db.AutoMigrate(&Favorite{})
		db.AutoMigrate(&OwnedGame{})
		return &DB{db}
	}
}

// The EndPoint enum
// Typesafe endpoints for the API
type EndPoint int

const (
	GetGamesList EndPoint = iota
	GetGame
	GetPlatformsList
	GetPlatform
	GetPlatformGames
)

var endpoints = [...]string{
	"GetGamesList",
	"GetGame",
	"GetPlatformsList",
	"GetPlatform",
	"GetPlatformGames",
}

func (endpoint EndPoint) String() string {
	return endpoints[endpoint]
}

// The Fetcher type
// Convenient wrapper for sql simple select queries
type Fetcher struct {
	db *DB
	i  interface{}
}

func (f Fetcher) All(out io.Writer) (e error) {
	f.db.Find(f.i)
	return WriteMarshal(&f.i, out)
}

func (f Fetcher) First(out io.Writer) (e error) {
	f.db.First(f.i)
	return WriteMarshal(&f.i, out)
}

func (f Fetcher) Last(out io.Writer) (e error) {
	f.db.Last(f.i)
	return WriteMarshal(&f.i, out)
}

// Game bool used during XML unmarshalling
// Converts "yes" "no" strings to boolean
type GameBool bool

// Implementation of the Unmarshaler interface for GameBool
func (gb *GameBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var v string
	d.DecodeElement(&v, &start)
	*gb = strings.ToLower(v) == "yes"
	return
}

// Game time used during XML unmarshalling
type GameTime struct {
	time.Time
}

// Implementation of the xml.Unmarshaler interface for GameTime
func (gt *GameTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (e error) {
	const shortForm = "01/02/2006"
	var v string
	d.DecodeElement(&v, &start)
	if v != "" {
		if parse, err := time.Parse(shortForm, v); err != nil {
			return err
		} else {
			*gt = GameTime{parse}
		}
	}
	return
}

// Implementation of the json.Marshaler interface for GameTime
func (gt GameTime) MarshalJSON() ([]byte, error) {
	if gt.IsZero() {
		return json.Marshal(nil)
	} else {
		return json.Marshal(gt.String())
	}
}

// The XMLRequest structure allowing xml streaming in json to an io.Writer output
type XMLRequest struct {
	Url string
}

// Streams from fetched xml data to an io.Writer (ex: http.ResponseWriter)
func (p XMLRequest) Stream(t interface{}, out io.Writer) (e error) {
	if resp, err := http.Get(p.Url); err != nil {
		return err
	} else {
		defer resp.Body.Close()
		xml.NewDecoder(resp.Body).Decode(&t)
		json.NewEncoder(out).Encode(&t)
		return
	}
}

// Query map, stringifies as a proper URL query search
type Query map[string]string

// Implementation of the Stringer interface for Query
func (q Query) String() (ret string) {
	if len(q) > 0 {
		v := url.Values{}
		for key, val := range q {
			v.Add(key, val)
		}
		return fmt.Sprintf("?%s", v.Encode())
	} else {
		return
	}

}
