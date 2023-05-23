package main

import (
	"bytes"
	"encoding/json"
	"github.com/MihajloJankovic/Alati/Dao"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
func decodeBody(r io.Reader) (*poststore.Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt poststore.Config
	if err := json.Unmarshal(StreamToByte(r), &rt); err != nil {
		return nil, err
	}
	return &rt, nil
}

func decodeGroupBody(r io.Reader) (*poststore.ConfigGroup, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt poststore.ConfigGroup
	if err := json.Unmarshal(StreamToByte(r), &rt); err != nil {
		return nil, err
	}
	return &rt, nil
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createId() string {
	return uuid.New().String()
}
