package ghs

import (
	"fmt"
	"net/http"
)
type GHS struct {
	id int
	title string
	chorus string
	stanza_1 string
	stanza_2 string
	stanza_3 string
	stanza_4 string
	stanza_5 string
	stanza_6 string
	stanza_7 string
	stanza_8 string
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	fetchAllSongs, err := db.Query("SELECT * FROM songs")
	if err != nil {
		panic(err.Error())
	}
	ghs := GHS{}

	for fetchAllSongs.Next() {
		var id int
		var title, chorus, stanza_1, stanza_2, stanza_3, stanza_4, stanza_5, stanza_6, stanza_7, stanza_8 string
		err = fetchAllSongs.Scan(&id, &title, &chorus, &stanza_1, &stanza_2, &stanza_3, &stanza_4, &stanza_5, &stanza_6, &stanza_7, &stanza_8)

		if err != nil {
			panic(err.Error())
		}
		ghs.id = id
		ghs.title = title
		ghs.chorus = chorus
		ghs.stanza_1= stanza_1
		ghs.stanza_2 = stanza_2
		ghs.stanza_3 = stanza_3
		ghs.stanza_4= stanza_4
		ghs.stanza_5 = stanza_5
		ghs.stanza_6 = stanza_6
		ghs.stanza_7= stanza_7
		ghs.stanza_8 = stanza_8

	}
	fmt.Println(ghs)
	render(w, "index.html", nil)
	
}

