package handler

import (
	"encoding/json"
	"net/http"
)

// Index handler
func (b *Base) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal([]string{b.Config().GetString("environment")})
	w.Write(j)
}
