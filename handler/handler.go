package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/uzimihsr/dice-api/dice"
)

type diceResult struct {
	Number int `json:"number"`
	Faces  int `json:"faces"`
}

func DiceHandler(d dice.DiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = diceGet(w, r, d)
		default:
			diceNotFound(w)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CheatDiceHandler(d dice.DiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = cheatDiceGet(w, r, d)
		default:
			diceNotFound(w)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func diceGet(w http.ResponseWriter, r *http.Request, d dice.DiceInterface) (err error) {
	faces := 6
	q := r.URL.Query()
	queryFaces := q.Get("faces")
	if queryFaces != "" {
		faces, err = strconv.Atoi(queryFaces)
		if err != nil {
			return err
		}
		if faces <= 0 {
			faces = 6
		}
	}

	d.SetFaces(faces)
	number := d.Roll()
	result := diceResult{number, d.GetFaces()}
	j, err := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	return
}

func cheatDiceGet(w http.ResponseWriter, r *http.Request, d dice.DiceInterface) (err error) {
	faces := 6
	q := r.URL.Query()
	queryFaces := q.Get("faces")
	queryNumber := q.Get("number")
	if queryFaces != "" {
		faces, err = strconv.Atoi(queryFaces)
		if err != nil {
			return err
		}
		if faces <= 0 {
			faces = 6
		}
	}
	num := faces
	if queryNumber != "" {
		num, err = strconv.Atoi(queryNumber)
		if err != nil {
			return err
		}
	}

	d.SetFaces(faces)
	number := d.Cheat(num)
	result := diceResult{number, d.GetFaces()}
	j, err := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
	return
}

func diceNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
