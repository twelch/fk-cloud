// Code generated by goa v3.1.2, DO NOT EDIT.
//
// user HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	user "github.com/fieldkit/cloud/server/api/gen/user"
	userviews "github.com/fieldkit/cloud/server/api/gen/user/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeRolesResponse returns an encoder for responses returned by the user
// roles endpoint.
func EncodeRolesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*userviews.AvailableRoles)
		enc := encoder(ctx, w)
		body := NewRolesResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRolesRequest returns a decoder for requests sent to the user roles
// endpoint.
func DecodeRolesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
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
		payload := NewRolesPayload(auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeRolesError returns an encoder for errors returned by the roles user
// endpoint.
func EncodeRolesError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(user.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRolesBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			res := v.(user.Forbidden)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRolesForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(user.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRolesNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(user.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRolesUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteResponse returns an encoder for responses returned by the user
// delete endpoint.
func EncodeDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteRequest returns a decoder for requests sent to the user delete
// endpoint.
func DecodeDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			userID int32
			auth   string
			err    error

			params = mux.Vars(r)
		)
		{
			userIDRaw := params["userId"]
			v, err2 := strconv.ParseInt(userIDRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("userID", userIDRaw, "integer"))
			}
			userID = int32(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeletePayload(userID, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeDeleteError returns an encoder for errors returned by the delete user
// endpoint.
func EncodeDeleteError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(user.BadRequest)
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
		case "forbidden":
			res := v.(user.Forbidden)
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
			res := v.(user.NotFound)
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
		case "unauthorized":
			res := v.(user.Unauthorized)
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
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUploadPhotoResponse returns an encoder for responses returned by the
// user upload photo endpoint.
func EncodeUploadPhotoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeUploadPhotoRequest returns a decoder for requests sent to the user
// upload photo endpoint.
func DecodeUploadPhotoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			contentType   string
			contentLength int64
			auth          string
			err           error
		)
		contentType = r.Header.Get("Content-Type")
		if contentType == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Content-Type", "header"))
		}
		{
			contentLengthRaw := r.Header.Get("Content-Length")
			if contentLengthRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("Content-Length", "header"))
			}
			v, err2 := strconv.ParseInt(contentLengthRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("contentLength", contentLengthRaw, "integer"))
			}
			contentLength = v
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewUploadPhotoPayload(contentType, contentLength, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeUploadPhotoError returns an encoder for errors returned by the upload
// photo user endpoint.
func EncodeUploadPhotoError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(user.BadRequest)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadPhotoBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "forbidden":
			res := v.(user.Forbidden)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadPhotoForbiddenResponseBody(res)
			}
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "not-found":
			res := v.(user.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadPhotoNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not-found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(user.Unauthorized)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadPhotoUnauthorizedResponseBody(res)
			}
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDownloadPhotoResponse returns an encoder for responses returned by the
// user download photo endpoint.
func EncodeDownloadPhotoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*user.DownloadPhotoResult)
		val := res.Length
		lengths := strconv.FormatInt(val, 10)
		w.Header().Set("Content-Length", lengths)
		w.Header().Set("Content-Type", res.ContentType)
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDownloadPhotoRequest returns a decoder for requests sent to the user
// download photo endpoint.
func DecodeDownloadPhotoRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			userID int32
			err    error

			params = mux.Vars(r)
		)
		{
			userIDRaw := params["userId"]
			v, err2 := strconv.ParseInt(userIDRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("userID", userIDRaw, "integer"))
			}
			userID = int32(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewDownloadPhotoPayload(userID)

		return payload, nil
	}
}

// EncodeDownloadPhotoError returns an encoder for errors returned by the
// download photo user endpoint.
func EncodeDownloadPhotoError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad-request":
			res := v.(user.BadRequest)
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
		case "forbidden":
			res := v.(user.Forbidden)
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
			res := v.(user.NotFound)
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
		case "unauthorized":
			res := v.(user.Unauthorized)
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
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalUserviewsAvailableRoleViewToAvailableRoleResponseBody builds a value
// of type *AvailableRoleResponseBody from a value of type
// *userviews.AvailableRoleView.
func marshalUserviewsAvailableRoleViewToAvailableRoleResponseBody(v *userviews.AvailableRoleView) *AvailableRoleResponseBody {
	res := &AvailableRoleResponseBody{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}
