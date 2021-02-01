package main

import (
	"encoding/json"
	"net/http"
)

type club struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Stadium  string `json:"stadium"`
	Location string `json:"location"`
	Klasemen int32  `json:"klasemen"`
}

var clubs = []*club{}

func init() {
	clubs = append(clubs, &club{ID: "epl001", Name: "Manchester City", Stadium: "Etihad Stadium", Location: "Manchester", Klasemen: 1})
	clubs = append(clubs, &club{ID: "epl002", Name: "Manchester United", Stadium: "Old Trafford", Location: "Manchester", Klasemen: 2})
	clubs = append(clubs, &club{ID: "epl003", Name: "Leichester City", Stadium: "King Power Stadium", Location: "Leichester", Klasemen: 3})
}

func getClubData(id string) <-chan club {
	result := make(chan club)
	go func() {
		for _, data := range clubs {
			if (*data).ID == id {
				result <- *data
				break
			}
		}
		result <- club{}
	}()
	return result
}

func getClubsData() <-chan []*club {
	result := make(chan []*club)
	go func() {
		result <- clubs
	}()
	return result
}

func outputClubData(w http.ResponseWriter, data interface{}) {
	responseJSON, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
