package configapi

import (
	"encoding/json"
	"github.com/bearts/nimbus/src/utils"
	"net/http"
)

func ConfigApiRestart(w http.ResponseWriter, req *http.Request) {
	if utils.AdminOnly(w, req) != nil {
		return
	}

	if req.Method == "GET" {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "OK",
		})
		utils.RestartServer()
	} else {
		utils.Error("Restart: Method not allowed"+req.Method, nil)
		utils.HTTPError(w, "Method not allowed", http.StatusMethodNotAllowed, "HTTP001")
		return
	}
}
