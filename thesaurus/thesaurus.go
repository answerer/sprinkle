package thesaurus

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type BigHug struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHug) Synonyms (term string) ([]string, error) {
	var syns []string

	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")

	if err != nil {
		return syns, fmt.Errorf("Bighug: failed when looking for synonyms for \"" + term + "\" " + err.Error())
	}

	defer response.Body.Close()
	var data synonyms
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}

	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)
	return syns, nil
}
