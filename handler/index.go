package handler

import (
	"encoding/json"
	"github.com/harymindiar/mini/core"
	"github.com/harymindiar/mini/provider"
	"net/http"
)

// Index handler
func (b *Base) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	config, _ := b.Container.Get(provider.ConfigServiceName)
	j, _ := json.Marshal([]string{config.(core.Config).GetString("environment")})
	w.Write(j)
}
