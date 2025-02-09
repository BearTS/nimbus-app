package authorizationserver

import (
	"github.com/bearts/nimbus/src/utils"
	"net/http"
)

func introspectionEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	mySessionData := newSession("", req)
	ir, err := oauth2.NewIntrospectionRequest(ctx, req, mySessionData)
	if err != nil {
		utils.Error("Error occurred in NewIntrospectionRequest", err)
		oauth2.WriteIntrospectionError(ctx, rw, err)
		return
	}

	oauth2.WriteIntrospectionResponse(ctx, rw, ir)
}
