package tmdb

import (
	"log"
	"testing"
)

func TestTmdb(t *testing.T) {
	client := NewClient("e610ded10c3f47d05fe797961d90fea6")
	log.Println("yeah client is created")

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
