// Code generated by protoc-gen-go-cloudfunction. DO NOT EDIT.

package pixivv1

type GetArtworkHandlerFunc = func(context.Context, *GetArtworkRequest) (*GetArtworkResponse, error)

func GetArtworkHandler(w http.ResponseWriter, r *http.Request, do GetArtworkHandlerFunc) {
	if err := cfutil.ApplyCors(w, r); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	if err := cfutil.ApplyContentType(w, r); err != nil {
		cfutil.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}
	var request GetArtworkRequest
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
