// Code generated by protoc-gen-go-cloudfunction. DO NOT EDIT.

package wibuwallpaperv1

import (
	"context"
	"net/http"
	"github.com/aiocean/cfutil"
)

type ListTopicsHandlerFunc = func(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error)

func ListTopicsHandler(w http.ResponseWriter, r *http.Request, do ListTopicsHandlerFunc) {
	if err := cfutil.ApplyCors(w, r); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := cfutil.ApplyContentType(w, r); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	var request ListTopicsRequest
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
