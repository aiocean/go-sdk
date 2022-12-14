// Code generated by protoc-gen-go-cloudfunction. DO NOT EDIT.

package wibuzonev1

import (
	"context"
	"net/http"
	"github.com/aiocean/cfutil"
)

type ListTagsHandlerFunc = func(context.Context, *ListTagsRequest) (*ListTagsResponse, error)

func ListTagsHandler(w http.ResponseWriter, r *http.Request, do ListTagsHandlerFunc) {
	if err := cfutil.ApplyCors(w, r); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := cfutil.ApplyContentType(w, r); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	var request ListTagsRequest
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

type GetTagHandlerFunc = func(context.Context, *GetTagRequest) (*GetTagResponse, error)

func GetTagHandler(w http.ResponseWriter, r *http.Request, do GetTagHandlerFunc) {
	if err := cfutil.ApplyCors(w, r); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := cfutil.ApplyContentType(w, r); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	var request GetTagRequest
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
