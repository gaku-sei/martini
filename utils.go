package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// The constants used when building the url
const apiUrl = "http://thegamesdb.net/api/%s.php%s"

// Marshals anything then writes the content into the writer
func WriteMarshal(i interface{}, out io.Writer) (e error) {
	if bytes, err := json.Marshal(i); err != nil {
		return err
	} else {
		out.Write(bytes)
	}
	return
}

// Prepares an XML request
func PrepareXML(endpoint EndPoint, params Query) XMLRequest {
	return XMLRequest{
		Url: fmt.Sprintf(apiUrl, endpoint, params),
	}
}
