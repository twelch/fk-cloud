// Code generated by goa v3.1.2, DO NOT EDIT.
//
// station HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"

	stationviews "github.com/fieldkit/cloud/server/api/gen/station/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeAddResponse returns an encoder for responses returned by the station
// add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*stationviews.StationFull)
		enc := encoder(ctx, w)
		body := NewAddResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the station add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AddRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			auth string
		)
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewAddPayload(&body, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeAddError returns an encoder for errors returned by the add station
// endpoint.
func EncodeAddError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "station-owner-conflict":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewAddStationOwnerConflictResponseBody(res)
			}
			w.Header().Set("goa-error", "station-owner-conflict")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad-request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewAddBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewAddUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewAddForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewAddNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetResponse returns an encoder for responses returned by the station
// get endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*stationviews.StationFull)
		enc := encoder(ctx, w)
		body := NewGetResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRequest returns a decoder for requests sent to the station get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id   int32
			auth string
			err  error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int32(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetPayload(id, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeGetError returns an encoder for errors returned by the get station
// endpoint.
func EncodeGetError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad-request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateResponse returns an encoder for responses returned by the
// station update endpoint.
func EncodeUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*stationviews.StationFull)
		enc := encoder(ctx, w)
		body := NewUpdateResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateRequest returns a decoder for requests sent to the station
// update endpoint.
func DecodeUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id   int32
			auth string

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int32(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdatePayload(&body, id, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeUpdateError returns an encoder for errors returned by the update
// station endpoint.
func EncodeUpdateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad-request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListMineResponse returns an encoder for responses returned by the
// station list mine endpoint.
func EncodeListMineResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*stationviews.StationsFull)
		enc := encoder(ctx, w)
		body := NewListMineResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListMineRequest returns a decoder for requests sent to the station
// list mine endpoint.
func DecodeListMineRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			auth string
			err  error
		)
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewListMinePayload(auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeListMineError returns an encoder for errors returned by the list mine
// station endpoint.
func EncodeListMineError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListMineUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListMineForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListMineNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad-request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListMineBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListProjectResponse returns an encoder for responses returned by the
// station list project endpoint.
func EncodeListProjectResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*stationviews.StationsFull)
		enc := encoder(ctx, w)
		body := NewListProjectResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListProjectRequest returns a decoder for requests sent to the station
// list project endpoint.
func DecodeListProjectRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id   int32
			auth *string
			err  error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int32(v)
		}
		authRaw := r.Header.Get("Authorization")
		if authRaw != "" {
			auth = &authRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewListProjectPayload(id, auth)
		if payload.Auth != nil {
			if strings.Contains(*payload.Auth, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Auth, " ", 2)[1]
				payload.Auth = &cred
			}
		}

		return payload, nil
	}
}

// EncodeListProjectError returns an encoder for errors returned by the list
// project station endpoint.
func EncodeListProjectError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListProjectUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListProjectForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListProjectNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad-request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListProjectBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDownloadPhotoResponse returns an encoder for responses returned by the
// station download photo endpoint.
func EncodeDownloadPhotoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*stationviews.DownloadedPhoto)
		enc := encoder(ctx, w)
		body := NewDownloadPhotoResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDownloadPhotoRequest returns a decoder for requests sent to the
// station download photo endpoint.
func DecodeDownloadPhotoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			stationID   int32
			size        *int32
			ifNoneMatch *string
			auth        string
			err         error

			params = mux.Vars(r)
		)
		{
			stationIDRaw := params["stationId"]
			v, err2 := strconv.ParseInt(stationIDRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("stationID", stationIDRaw, "integer"))
			}
			stationID = int32(v)
		}
		{
			sizeRaw := r.URL.Query().Get("size")
			if sizeRaw != "" {
				v, err2 := strconv.ParseInt(sizeRaw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("size", sizeRaw, "integer"))
				}
				pv := int32(v)
				size = &pv
			}
		}
		ifNoneMatchRaw := r.Header.Get("If-None-Match")
		if ifNoneMatchRaw != "" {
			ifNoneMatch = &ifNoneMatchRaw
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDownloadPhotoPayload(stationID, size, ifNoneMatch, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeDownloadPhotoError returns an encoder for errors returned by the
// download photo station endpoint.
func EncodeDownloadPhotoError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDownloadPhotoUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDownloadPhotoForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDownloadPhotoNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad-request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDownloadPhotoBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListAllResponse returns an encoder for responses returned by the
// station list all endpoint.
func EncodeListAllResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*stationviews.PageOfStations)
		enc := encoder(ctx, w)
		body := NewListAllResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListAllRequest returns a decoder for requests sent to the station list
// all endpoint.
func DecodeListAllRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			page     *int32
			pageSize *int32
			ownerID  *int32
			query    *string
			sortBy   *string
			auth     string
			err      error
		)
		{
			pageRaw := r.URL.Query().Get("page")
			if pageRaw != "" {
				v, err2 := strconv.ParseInt(pageRaw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("page", pageRaw, "integer"))
				}
				pv := int32(v)
				page = &pv
			}
		}
		{
			pageSizeRaw := r.URL.Query().Get("pageSize")
			if pageSizeRaw != "" {
				v, err2 := strconv.ParseInt(pageSizeRaw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("pageSize", pageSizeRaw, "integer"))
				}
				pv := int32(v)
				pageSize = &pv
			}
		}
		{
			ownerIDRaw := r.URL.Query().Get("ownerId")
			if ownerIDRaw != "" {
				v, err2 := strconv.ParseInt(ownerIDRaw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("ownerID", ownerIDRaw, "integer"))
				}
				pv := int32(v)
				ownerID = &pv
			}
		}
		queryRaw := r.URL.Query().Get("query")
		if queryRaw != "" {
			query = &queryRaw
		}
		sortByRaw := r.URL.Query().Get("sortBy")
		if sortByRaw != "" {
			sortBy = &sortByRaw
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewListAllPayload(page, pageSize, ownerID, query, sortBy, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeListAllError returns an encoder for errors returned by the list all
// station endpoint.
func EncodeListAllError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListAllUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListAllForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListAllNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad-request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewListAllBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteResponse returns an encoder for responses returned by the
// station delete endpoint.
func EncodeDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteRequest returns a decoder for requests sent to the station
// delete endpoint.
func DecodeDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			stationID int32
			auth      string
			err       error

			params = mux.Vars(r)
		)
		{
			stationIDRaw := params["stationId"]
			v, err2 := strconv.ParseInt(stationIDRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("stationID", stationIDRaw, "integer"))
			}
			stationID = int32(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeletePayload(stationID, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeDeleteError returns an encoder for errors returned by the delete
// station endpoint.
func EncodeDeleteError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "bad-request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalStationviewsStationOwnerViewToStationOwnerResponseBody builds a value
// of type *StationOwnerResponseBody from a value of type
// *stationviews.StationOwnerView.
func marshalStationviewsStationOwnerViewToStationOwnerResponseBody(v *stationviews.StationOwnerView) *StationOwnerResponseBody {
	res := &StationOwnerResponseBody{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// marshalStationviewsStationUploadViewToStationUploadResponseBody builds a
// value of type *StationUploadResponseBody from a value of type
// *stationviews.StationUploadView.
func marshalStationviewsStationUploadViewToStationUploadResponseBody(v *stationviews.StationUploadView) *StationUploadResponseBody {
	res := &StationUploadResponseBody{
		ID:       *v.ID,
		Time:     *v.Time,
		UploadID: *v.UploadID,
		Size:     *v.Size,
		URL:      *v.URL,
		Type:     *v.Type,
	}
	if v.Blocks != nil {
		res.Blocks = make([]int64, len(v.Blocks))
		for i, val := range v.Blocks {
			res.Blocks[i] = val
		}
	}

	return res
}

// marshalStationviewsStationPhotosViewToStationPhotosResponseBody builds a
// value of type *StationPhotosResponseBody from a value of type
// *stationviews.StationPhotosView.
func marshalStationviewsStationPhotosViewToStationPhotosResponseBody(v *stationviews.StationPhotosView) *StationPhotosResponseBody {
	res := &StationPhotosResponseBody{
		Small: *v.Small,
	}

	return res
}

// marshalStationviewsStationConfigurationsViewToStationConfigurationsResponseBody
// builds a value of type *StationConfigurationsResponseBody from a value of
// type *stationviews.StationConfigurationsView.
func marshalStationviewsStationConfigurationsViewToStationConfigurationsResponseBody(v *stationviews.StationConfigurationsView) *StationConfigurationsResponseBody {
	res := &StationConfigurationsResponseBody{}
	if v.All != nil {
		res.All = make([]*StationConfigurationResponseBody, len(v.All))
		for i, val := range v.All {
			res.All[i] = marshalStationviewsStationConfigurationViewToStationConfigurationResponseBody(val)
		}
	}

	return res
}

// marshalStationviewsStationConfigurationViewToStationConfigurationResponseBody
// builds a value of type *StationConfigurationResponseBody from a value of
// type *stationviews.StationConfigurationView.
func marshalStationviewsStationConfigurationViewToStationConfigurationResponseBody(v *stationviews.StationConfigurationView) *StationConfigurationResponseBody {
	res := &StationConfigurationResponseBody{
		ID:           *v.ID,
		Time:         *v.Time,
		ProvisionID:  *v.ProvisionID,
		MetaRecordID: v.MetaRecordID,
		SourceID:     v.SourceID,
	}
	if v.Modules != nil {
		res.Modules = make([]*StationModuleResponseBody, len(v.Modules))
		for i, val := range v.Modules {
			res.Modules[i] = marshalStationviewsStationModuleViewToStationModuleResponseBody(val)
		}
	}

	return res
}

// marshalStationviewsStationModuleViewToStationModuleResponseBody builds a
// value of type *StationModuleResponseBody from a value of type
// *stationviews.StationModuleView.
func marshalStationviewsStationModuleViewToStationModuleResponseBody(v *stationviews.StationModuleView) *StationModuleResponseBody {
	res := &StationModuleResponseBody{
		ID:           *v.ID,
		HardwareID:   v.HardwareID,
		MetaRecordID: v.MetaRecordID,
		Name:         *v.Name,
		Position:     *v.Position,
		Flags:        *v.Flags,
		Internal:     *v.Internal,
		FullKey:      *v.FullKey,
	}
	if v.Sensors != nil {
		res.Sensors = make([]*StationSensorResponseBody, len(v.Sensors))
		for i, val := range v.Sensors {
			res.Sensors[i] = marshalStationviewsStationSensorViewToStationSensorResponseBody(val)
		}
	}

	return res
}

// marshalStationviewsStationSensorViewToStationSensorResponseBody builds a
// value of type *StationSensorResponseBody from a value of type
// *stationviews.StationSensorView.
func marshalStationviewsStationSensorViewToStationSensorResponseBody(v *stationviews.StationSensorView) *StationSensorResponseBody {
	res := &StationSensorResponseBody{
		Name:          *v.Name,
		UnitOfMeasure: *v.UnitOfMeasure,
		Key:           *v.Key,
		FullKey:       *v.FullKey,
	}
	if v.Reading != nil {
		res.Reading = marshalStationviewsSensorReadingViewToSensorReadingResponseBody(v.Reading)
	}
	if v.Ranges != nil {
		res.Ranges = make([]*SensorRangeResponseBody, len(v.Ranges))
		for i, val := range v.Ranges {
			res.Ranges[i] = marshalStationviewsSensorRangeViewToSensorRangeResponseBody(val)
		}
	}

	return res
}

// marshalStationviewsSensorReadingViewToSensorReadingResponseBody builds a
// value of type *SensorReadingResponseBody from a value of type
// *stationviews.SensorReadingView.
func marshalStationviewsSensorReadingViewToSensorReadingResponseBody(v *stationviews.SensorReadingView) *SensorReadingResponseBody {
	if v == nil {
		return nil
	}
	res := &SensorReadingResponseBody{
		Last: *v.Last,
		Time: *v.Time,
	}

	return res
}

// marshalStationviewsSensorRangeViewToSensorRangeResponseBody builds a value
// of type *SensorRangeResponseBody from a value of type
// *stationviews.SensorRangeView.
func marshalStationviewsSensorRangeViewToSensorRangeResponseBody(v *stationviews.SensorRangeView) *SensorRangeResponseBody {
	res := &SensorRangeResponseBody{
		Minimum: *v.Minimum,
		Maximum: *v.Maximum,
	}

	return res
}

// marshalStationviewsStationLocationViewToStationLocationResponseBody builds a
// value of type *StationLocationResponseBody from a value of type
// *stationviews.StationLocationView.
func marshalStationviewsStationLocationViewToStationLocationResponseBody(v *stationviews.StationLocationView) *StationLocationResponseBody {
	if v == nil {
		return nil
	}
	res := &StationLocationResponseBody{}
	if v.Precise != nil {
		res.Precise = make([]float64, len(v.Precise))
		for i, val := range v.Precise {
			res.Precise[i] = val
		}
	}
	if v.Region != nil {
		res.Region = make([][][]float64, len(v.Region))
		for i, val := range v.Region {
			res.Region[i] = make([][]float64, len(val))
			for j, val := range val {
				res.Region[i][j] = make([]float64, len(val))
				for k, val := range val {
					res.Region[i][j][k] = val
				}
			}
		}
	}

	return res
}

// marshalStationviewsStationDataSummaryViewToStationDataSummaryResponseBody
// builds a value of type *StationDataSummaryResponseBody from a value of type
// *stationviews.StationDataSummaryView.
func marshalStationviewsStationDataSummaryViewToStationDataSummaryResponseBody(v *stationviews.StationDataSummaryView) *StationDataSummaryResponseBody {
	if v == nil {
		return nil
	}
	res := &StationDataSummaryResponseBody{
		Start:           *v.Start,
		End:             *v.End,
		NumberOfSamples: *v.NumberOfSamples,
	}

	return res
}

// marshalStationviewsStationFullViewToStationFullResponseBody builds a value
// of type *StationFullResponseBody from a value of type
// *stationviews.StationFullView.
func marshalStationviewsStationFullViewToStationFullResponseBody(v *stationviews.StationFullView) *StationFullResponseBody {
	res := &StationFullResponseBody{
		ID:                 *v.ID,
		Name:               *v.Name,
		DeviceID:           *v.DeviceID,
		ReadOnly:           *v.ReadOnly,
		Battery:            v.Battery,
		RecordingStartedAt: v.RecordingStartedAt,
		MemoryUsed:         v.MemoryUsed,
		MemoryAvailable:    v.MemoryAvailable,
		FirmwareNumber:     v.FirmwareNumber,
		FirmwareTime:       v.FirmwareTime,
		UpdatedAt:          *v.UpdatedAt,
		LocationName:       v.LocationName,
		PlaceNameOther:     v.PlaceNameOther,
		PlaceNameNative:    v.PlaceNameNative,
	}
	if v.Owner != nil {
		res.Owner = marshalStationviewsStationOwnerViewToStationOwnerResponseBody(v.Owner)
	}
	if v.Uploads != nil {
		res.Uploads = make([]*StationUploadResponseBody, len(v.Uploads))
		for i, val := range v.Uploads {
			res.Uploads[i] = marshalStationviewsStationUploadViewToStationUploadResponseBody(val)
		}
	}
	if v.Photos != nil {
		res.Photos = marshalStationviewsStationPhotosViewToStationPhotosResponseBody(v.Photos)
	}
	if v.Configurations != nil {
		res.Configurations = marshalStationviewsStationConfigurationsViewToStationConfigurationsResponseBody(v.Configurations)
	}
	if v.Location != nil {
		res.Location = marshalStationviewsStationLocationViewToStationLocationResponseBody(v.Location)
	}
	if v.Data != nil {
		res.Data = marshalStationviewsStationDataSummaryViewToStationDataSummaryResponseBody(v.Data)
	}

	return res
}

// marshalStationviewsEssentialStationViewToEssentialStationResponseBody builds
// a value of type *EssentialStationResponseBody from a value of type
// *stationviews.EssentialStationView.
func marshalStationviewsEssentialStationViewToEssentialStationResponseBody(v *stationviews.EssentialStationView) *EssentialStationResponseBody {
	res := &EssentialStationResponseBody{
		ID:                 *v.ID,
		DeviceID:           *v.DeviceID,
		Name:               *v.Name,
		CreatedAt:          *v.CreatedAt,
		UpdatedAt:          *v.UpdatedAt,
		RecordingStartedAt: v.RecordingStartedAt,
		MemoryUsed:         v.MemoryUsed,
		MemoryAvailable:    v.MemoryAvailable,
		FirmwareNumber:     v.FirmwareNumber,
		FirmwareTime:       v.FirmwareTime,
		LastIngestionAt:    v.LastIngestionAt,
	}
	if v.Owner != nil {
		res.Owner = marshalStationviewsStationOwnerViewToStationOwnerResponseBody(v.Owner)
	}
	if v.Location != nil {
		res.Location = marshalStationviewsStationLocationViewToStationLocationResponseBody(v.Location)
	}

	return res
}
