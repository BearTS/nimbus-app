package constellation

import (
	"github.com/bearts/nimbus/src/utils"
	"net/http"
)

func ConstellationAPIDevices(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		DeviceList(w, req)
	} else if req.Method == "POST" {
		DeviceCreate(w, req)
	} else {
		utils.Error("UserRoute: Method not allowed"+req.Method, nil)
		utils.HTTPError(w, "Method not allowed", http.StatusMethodNotAllowed, "HTTP001")
		return
	}
}
