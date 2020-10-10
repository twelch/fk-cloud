// Code generated by goa v3.2.4, DO NOT EDIT.
//
// discussion HTTP server encoders and decoders
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

	discussion "github.com/fieldkit/cloud/server/api/gen/discussion"
	discussionviews "github.com/fieldkit/cloud/server/api/gen/discussion/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeProjectResponse returns an encoder for responses returned by the
// discussion project endpoint.
func EncodeProjectResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*discussionviews.Discussion)
		enc := encoder(ctx, w)
		body := NewProjectResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeProjectRequest returns a decoder for requests sent to the discussion
// project endpoint.
func DecodeProjectRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID int32
			auth      string
			err       error

			params = mux.Vars(r)
		)
		{
			projectIDRaw := params["projectID"]
			v, err2 := strconv.ParseInt(projectIDRaw, 10, 32)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("projectID", projectIDRaw, "integer"))
			}
			projectID = int32(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewProjectPayload(projectID, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeProjectError returns an encoder for errors returned by the project
// discussion endpoint.
func EncodeProjectError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewProjectUnauthorizedResponseBody(res)
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
				body = NewProjectForbiddenResponseBody(res)
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
				body = NewProjectNotFoundResponseBody(res)
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
				body = NewProjectBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDataResponse returns an encoder for responses returned by the
// discussion data endpoint.
func EncodeDataResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*discussionviews.Discussion)
		enc := encoder(ctx, w)
		body := NewDataResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDataRequest returns a decoder for requests sent to the discussion data
// endpoint.
func DecodeDataRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			bookmark string
			auth     string
			err      error
		)
		bookmark = r.URL.Query().Get("bookmark")
		if bookmark == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("bookmark", "query string"))
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDataPayload(bookmark, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodeDataError returns an encoder for errors returned by the data
// discussion endpoint.
func EncodeDataError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewDataUnauthorizedResponseBody(res)
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
				body = NewDataForbiddenResponseBody(res)
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
				body = NewDataNotFoundResponseBody(res)
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
				body = NewDataBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodePostMessageResponse returns an encoder for responses returned by the
// discussion post message endpoint.
func EncodePostMessageResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*discussion.PostMessageResult)
		enc := encoder(ctx, w)
		body := NewPostMessageResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodePostMessageRequest returns a decoder for requests sent to the
// discussion post message endpoint.
func DecodePostMessageRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body PostMessageRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidatePostMessageRequestBody(&body)
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
		payload := NewPostMessagePayload(&body, auth)
		if strings.Contains(payload.Auth, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Auth, " ", 2)[1]
			payload.Auth = cred
		}

		return payload, nil
	}
}

// EncodePostMessageError returns an encoder for errors returned by the post
// message discussion endpoint.
func EncodePostMessageError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
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
				body = NewPostMessageUnauthorizedResponseBody(res)
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
				body = NewPostMessageForbiddenResponseBody(res)
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
				body = NewPostMessageNotFoundResponseBody(res)
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
				body = NewPostMessageBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad-request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalDiscussionviewsThreadedPostViewToThreadedPostResponseBody builds a
// value of type *ThreadedPostResponseBody from a value of type
// *discussionviews.ThreadedPostView.
func marshalDiscussionviewsThreadedPostViewToThreadedPostResponseBody(v *discussionviews.ThreadedPostView) *ThreadedPostResponseBody {
	res := &ThreadedPostResponseBody{
		ID:        *v.ID,
		CreatedAt: *v.CreatedAt,
		UpdatedAt: *v.UpdatedAt,
		Body:      *v.Body,
		Bookmark:  v.Bookmark,
	}
	if v.Author != nil {
		res.Author = marshalDiscussionviewsPostAuthorViewToPostAuthorResponseBody(v.Author)
	}
	if v.Replies != nil {
		res.Replies = make([]*ThreadedPostResponseBody, len(v.Replies))
		for i, val := range v.Replies {
			res.Replies[i] = marshalDiscussionviewsThreadedPostViewToThreadedPostResponseBody(val)
		}
	}

	return res
}

// marshalDiscussionviewsPostAuthorViewToPostAuthorResponseBody builds a value
// of type *PostAuthorResponseBody from a value of type
// *discussionviews.PostAuthorView.
func marshalDiscussionviewsPostAuthorViewToPostAuthorResponseBody(v *discussionviews.PostAuthorView) *PostAuthorResponseBody {
	res := &PostAuthorResponseBody{
		ID:       *v.ID,
		Name:     *v.Name,
		MediaURL: v.MediaURL,
	}

	return res
}

// unmarshalNewPostRequestBodyToDiscussionNewPost builds a value of type
// *discussion.NewPost from a value of type *NewPostRequestBody.
func unmarshalNewPostRequestBodyToDiscussionNewPost(v *NewPostRequestBody) *discussion.NewPost {
	res := &discussion.NewPost{
		ThreadID:  v.ThreadID,
		Body:      *v.Body,
		ProjectID: v.ProjectID,
		Bookmark:  v.Bookmark,
	}

	return res
}

// marshalDiscussionThreadedPostToThreadedPostResponseBody builds a value of
// type *ThreadedPostResponseBody from a value of type *discussion.ThreadedPost.
func marshalDiscussionThreadedPostToThreadedPostResponseBody(v *discussion.ThreadedPost) *ThreadedPostResponseBody {
	res := &ThreadedPostResponseBody{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Body:      v.Body,
		Bookmark:  v.Bookmark,
	}
	if v.Author != nil {
		res.Author = marshalDiscussionPostAuthorToPostAuthorResponseBody(v.Author)
	}
	if v.Replies != nil {
		res.Replies = make([]*ThreadedPostResponseBody, len(v.Replies))
		for i, val := range v.Replies {
			res.Replies[i] = marshalDiscussionThreadedPostToThreadedPostResponseBody(val)
		}
	}

	return res
}

// marshalDiscussionPostAuthorToPostAuthorResponseBody builds a value of type
// *PostAuthorResponseBody from a value of type *discussion.PostAuthor.
func marshalDiscussionPostAuthorToPostAuthorResponseBody(v *discussion.PostAuthor) *PostAuthorResponseBody {
	res := &PostAuthorResponseBody{
		ID:       v.ID,
		Name:     v.Name,
		MediaURL: v.MediaURL,
	}

	return res
}