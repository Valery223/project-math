package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
)

type Server struct {
	Multiplier float64
}

func NewServer(rtp float64) *Server {
	multiplier := math.Sqrt(2*9999*rtp + 1)
	return &Server{Multiplier: multiplier}
}

type Response struct {
	Result float64 `json:"result"`
}

func (s *Server) GenerateMultiplier() float64 {
	return s.Multiplier

}

func (s *Server) HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Result: s.GenerateMultiplier()})
}

func main() {
	rtp := flag.Float64("rtp", 0.0, "RTP value")
	flag.Parse()

	server := NewServer(*rtp)

	http.HandleFunc("/get", server.HandleGet)

	fmt.Printf("Server starting on :64333 with RTP = %.2f\n", *rtp)
	log.Fatal(http.ListenAndServe(":64333", nil))
}
