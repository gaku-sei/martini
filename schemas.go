package main

import (
	"encoding/xml"
)

// The schema for the GetGamesList endpoint
type Game struct {
	XMLName xml.Name `xml:"Data" json:"-"`
	Game    struct {
		XMLName     xml.Name `xml:"Game" json:"-"`
		Id          int      `xml:"id" json:"id"`
		GameTitle   string   `json:"title"`
		Platform    string   `json:"platform"`
		PlatformId  int      `json:"platformId"`
		Coop        GameBool `xml:"Co-op" json:"coop"`
		ReleaseDate GameTime `json:"releaseDate"`
		Overview    string   `json:"overview,omitempty"`
		Genres      []string `xml:"Genres>genre" json:"genres"`
		Players     string   `json:"players,omitempty"`
		Publisher   string   `json:"publisher,omitempty"`
		Developer   string   `json:"developer,omitempty"`
		ESRB        string   `json:"esrb,omitempty"`
		Similar     struct {
			XMLName      xml.Name `xml:"Similar" json:"-"`
			SimilarCount int      `json:"count"`
			Game         []struct {
				XMLName    xml.Name `xml:"Game" json:"-"`
				Id         int      `xml:"id" json:"id"`
				PlatformId int      `json:"platformId"`
			} `json:"games"`
		} `json:"similar"`
	} `json:"game"`
}

// The schema for the GetGame endpoint
type Games struct {
	XMLName xml.Name `xml:"Data" json:"-"`
	Games   []struct {
		XMLName     xml.Name `xml:"Game" json:"-"`
		Id          int      `xml:"id" json:"id"`
		GameTitle   string   `json:"title"`
		ReleaseDate GameTime `json:"releaseDate"`
		Platform    string   `json:"platform,omitempty"`
	} `xml:"Game" json:"games"`
}

// The schema for the GetGamesList endpoint
type Platform struct {
	XMLName  xml.Name `xml:"Data" json:"-"`
	Platform struct {
		XMLName        xml.Name `xml:"Platform" json:"-"`
		Id             int      `xml:"id" json:"id"`
		Platform       string   `json:"platform"`
		Overview       string   `xml:"overview" json:"overview"`
		Developer      string   `xml:"developer" json:"developer",omitempty`
		Manufacturer   string   `xml:"manufacturer" json:"manufacturer,omitempty"`
		CPU            string   `xml:"cpu" json:"cpu,omitempty"`
		Memory         string   `xml:"memory" json:"memory,omitempty"`
		Graphics       string   `xml:"graphics" json:"graphics,omitempty"`
		Sound          string   `xml:"sound" json:"sound,omitempty"`
		Display        string   `xml:"display" json:"display,omitempty"`
		Media          string   `xml:"media" json:"media,omitempty"`
		MaxControllers string   `xml:"maxcontrollers" json:"maxControllers,omitempty"`
	} `json:"platform"`
}

// The schema for the GetPlatformsList endpoint
type Platforms struct {
	XMLName   xml.Name `xml:"Data" json:"-"`
	Platforms struct {
		XMLName   xml.Name `xml:"Platforms" json:"-"`
		Platforms []struct {
			XMLName xml.Name `xml:"Platform" json:"-"`
			Id      int      `xml:"id" json:"id"`
			Name    string   `xml:"name" json:"name"`
			Alias   string   `xml:"alias" json:"alias"`
		} `xml:"Platform" json:"platforms"`
	} `xml:"Platforms" json:"data"`
}
