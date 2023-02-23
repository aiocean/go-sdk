// Code generated by protoc-gen-go-cloudfunction. DO NOT EDIT.

package pinterestv1

import (
	"context"
	"net/http"
	"github.com/aiocean/cfutil"
)

type GetVideoPinHandlerFunc = func(context.Context, *GetVideoPinRequest) (*GetVideoPinResponse, error)

func GetVideoPinHandler(w http.ResponseWriter, r *http.Request, do GetVideoPinHandlerFunc) {
	if err := cfutil.ApplyCors(w, r); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := cfutil.ApplyContentType(w, r); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	var request GetVideoPinRequest
	if err := cfutil.ReadRequest(r, &request); err != nil {
		cfutil.WriteError(w, r, http.StatusBadRequest, err)
		return
	}
	response, err := do(r.Context(), &request)
	if err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := cfutil.WriteResponse(w, r, response); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
	}
}