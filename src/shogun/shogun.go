package main

import (
	"encoding/json"
	"net/http"
)

func writeJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		println("err", err)
	}
}
type slackResponse struct {
	Text string `json:"text",omitempty`
	Username string `json:"username",omitempty`
}
func main() {
	http.HandleFunc("/trin", func(w http.ResponseWriter, r *http.Request) {
		writeJson(w, slackResponse{
			"Hi, I'm Shogun Trin, I will help you count your pull-up, push-up, hook... I mean hiccup. But in the mean time, I have to learn math :alien:",
			"Shogun Trin",
		})
	})
	http.ListenAndServe(":3000", nil)
}
