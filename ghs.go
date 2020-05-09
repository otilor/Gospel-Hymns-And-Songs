package ghs

import (
	"fmt"
	"net/http"
)


func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	fetchAllSongs, err := db.Query("SELECT * FROM songs")
	if err != nil {
		panic(err.Error())
	}
	ghs := GHS{}
	ghsSlice := []GHS{}
	for fetchAllSongs.Next() {
		var id int
		var title, chorus, stanza_1, stanza_2, stanza_3, stanza_4, stanza_5, stanza_6, stanza_7, stanza_8 string
		err = fetchAllSongs.Scan(&id, &title, &chorus, &stanza_1, &stanza_2, &stanza_3, &stanza_4, &stanza_5, &stanza_6, &stanza_7, &stanza_8)

		if err != nil {
			panic(err.Error())
		}
		ghs.Id = id
		ghs.Title = title
		ghs.Chorus = chorus
		ghs.Stanza_1= stanza_1
		ghs.Stanza_2 = stanza_2
		ghs.Stanza_3 = stanza_3
		ghs.Stanza_4= stanza_4
		ghs.Stanza_5 = stanza_5
		ghs.Stanza_6 = stanza_6
		ghs.Stanza_7= stanza_7
		ghs.Stanza_8 = stanza_8
		
		ghsSlice = append(ghsSlice, ghs)

	}

	render(w, "index.html", ghsSlice)
	
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Id := r.URL.Query().Get("id")
	findGhs, err := db.Query("SELECT * FROM songs WHERE id = ?", Id)
	if err != nil {
		panic(err.Error())
	}

	ghs := GHS{}
	ghsSlice := []GHS{}

	for findGhs.Next() {
		var id int
		var title, chorus, stanza_1, stanza_2, stanza_3, stanza_4, stanza_5, stanza_6, stanza_7, stanza_8 string
		err = findGhs.Scan(&id, &title, &chorus, &stanza_1, &stanza_2, &stanza_3, &stanza_4, &stanza_5, &stanza_6, &stanza_7, &stanza_8)

		if err != nil {
			panic(err.Error())
		}
		ghs.Id = id
		ghs.Title = title
		ghs.Chorus = chorus
		ghs.Stanza_1= stanza_1
		ghs.Stanza_2 = stanza_2
		ghs.Stanza_3 = stanza_3
		ghs.Stanza_4= stanza_4
		ghs.Stanza_5 = stanza_5
		ghs.Stanza_6 = stanza_6
		ghs.Stanza_7= stanza_7
		ghs.Stanza_8 = stanza_8
		
		ghsSlice = append(ghsSlice, ghs)
		
	}
	render(w, "show.html", ghsSlice)
}