package tmdb

import (
	"encoding/json"
	"log"
	"testing"
)

type Tornado struct {
	Images struct {
		Base_Url        string `json: base_url`
		Secure_Base_Url string `json: secure_base_url`
	} `json: images`
	ChangeKeys []string `json: change_keys`
}

func TestJson(t *testing.T) {
	encoded := `{
		"images": {
			"base_url": "http://d3gtl9l2a4fn1j.cloudfront.net/t/p/",
			"secure_base_url": "https://d3gtl9l2a4fn1j.cloudfront.net/t/p/"
		},
		"change_keys": []
	}`

	tornado := &Tornado{}
	err := json.Unmarshal([]byte(encoded), &tornado)
	if err != nil {
		log.Fatalf("mother %v", err)
	}

	log.Printf("tornado.base is %v", tornado.Images.Base_Url)
	log.Printf("tornado.secure is %v", tornado.Images.Secure_Base_Url)
}

func DontTestConfig(t *testing.T) {
	client, err := NewClient("e610ded10c3f47d05fe797961d90fea6", true)
	// client, err := NewClient("e610ded10cee797961d90fea6", true)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("BaseUrl = %s", client.BaseUrl)
	log.Printf("SecureBaseUrl = %s", client.SecureBaseUrl)
}

func DontTestTmdb(t *testing.T) {
	client, err := NewClient("e610ded10c3f47d05fe797961d90fea6", true)
	// client, err := NewClient("e610ded10cee797961d90fea6", true)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("yeah client is created")

	// res, err := client.SearchMovie("2010")
	res, err := client.SearchMovie("10 things i hate about you")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)

	if res.Total_Results != 1 {
		log.Println("more than one")
	}

	id := res.Results[0].Id

	gmr, err := client.GetMovie(id)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(gmr)
}
