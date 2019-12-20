// Copyright 2019 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package servicebroker provides access to the Service Broker API.
//
// For product documentation, see: https://cloud.google.com/kubernetes-engine/docs/concepts/add-on/service-broker
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/servicebroker/v1"
//   ...
//   ctx := context.Background()
//   servicebrokerService, err := servicebroker.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   servicebrokerService, err := servicebroker.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   servicebrokerService, err := servicebroker.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package servicebroker // import "google.golang.org/api/servicebroker/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled

const apiId = "servicebroker:v1"
const apiName = "servicebroker"
const apiVersion = "v1"
const basePath = "https://servicebroker.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/cloud-platform",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.V1 = NewV1Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	V1 *V1Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewV1Service(s *Service) *V1Service {
	rs := &V1Service{s: s}
	return rs
}

type V1Service struct {
	s *Service
}

// GoogleIamV1__Binding: Associates `members` with a `role`.
type GoogleIamV1__Binding struct {
	// Condition: The condition that is associated with this binding.
	// NOTE: An unsatisfied condition will not allow user access via
	// current
	// binding. Different bindings, including their conditions, are
	// examined
	// independently.
	Condition *GoogleType__Expr `json:"condition,omitempty"`

	// Members: Specifies the identities requesting access for a Cloud
	// Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents
	// anyone
	//    who is authenticated with a Google account or a service
	// account.
	//
	// * `user:{emailid}`: An email address that represents a specific
	// Google
	//    account. For example, `alice@example.com` .
	//
	//
	// * `serviceAccount:{emailid}`: An email address that represents a
	// service
	//    account. For example,
	// `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google
	// group.
	//    For example, `admins@example.com`.
	//
	//
	// * `domain:{domain}`: The G Suite domain (primary) that represents all
	// the
	//    users of that domain. For example, `google.com` or
	// `example.com`.
	//
	//
	Members []string `json:"members,omitempty"`

	// Role: Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role string `json:"role,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Condition") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Condition") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__Binding) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__Binding
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__Policy: Defines an Identity and Access Management (IAM)
// policy. It is used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` is a collection of `bindings`. A `binding` binds one or
// more
// `members` to a single `role`. Members can be user accounts, service
// accounts,
// Google groups, and domains (such as G Suite). A `role` is a named
// list of
// permissions (defined by IAM or configured by users). A `binding`
// can
// optionally specify a `condition`, which is a logic expression that
// further
// constrains the role binding based on attributes about the request
// and/or
// target resource.
//
// **JSON Example**
//
//     {
//       "bindings": [
//         {
//           "role": "roles/resourcemanager.organizationAdmin",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//
// "serviceAccount:my-project-id@appspot.gserviceaccount.com"
//           ]
//         },
//         {
//           "role": "roles/resourcemanager.organizationViewer",
//           "members": ["user:eve@example.com"],
//           "condition": {
//             "title": "expirable access",
//             "description": "Does not grant access after Sep 2020",
//             "expression": "request.time <
//             timestamp('2020-10-01T00:00:00.000Z')",
//           }
//         }
//       ]
//     }
//
// **YAML Example**
//
//     bindings:
//     - members:
//       - user:mike@example.com
//       - group:admins@example.com
//       - domain:google.com
//       - serviceAccount:my-project-id@appspot.gserviceaccount.com
//       role: roles/resourcemanager.organizationAdmin
//     - members:
//       - user:eve@example.com
//       role: roles/resourcemanager.organizationViewer
//       condition:
//         title: expirable access
//         description: Does not grant access after Sep 2020
//         expression: request.time <
// timestamp('2020-10-01T00:00:00.000Z')
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam/docs).
type GoogleIamV1__Policy struct {
	// Bindings: Associates a list of `members` to a `role`. Optionally may
	// specify a
	// `condition` that determines when binding is in effect.
	// `bindings` with no members will result in an error.
	Bindings []*GoogleIamV1__Binding `json:"bindings,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a policy from overwriting each
	// other.
	// It is strongly suggested that systems make use of the `etag` in
	// the
	// read-modify-write cycle to perform policy updates in order to avoid
	// race
	// conditions: An `etag` is returned in the response to `getIamPolicy`,
	// and
	// systems are expected to put that etag in the request to
	// `setIamPolicy` to
	// ensure that their change will be applied to the same version of the
	// policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the
	// existing
	// policy is overwritten. Due to blind-set semantics of an etag-less
	// policy,
	// 'setIamPolicy' will not fail even if the incoming policy version does
	// not
	// meet the requirements for modifying the stored policy.
	Etag string `json:"etag,omitempty"`

	// Version: Specifies the format of the policy.
	//
	// Valid values are 0, 1, and 3. Requests specifying an invalid value
	// will be
	// rejected.
	//
	// Operations affecting conditional bindings must specify version 3.
	// This can
	// be either setting a conditional policy, modifying a conditional
	// binding,
	// or removing a binding (conditional or unconditional) from the
	// stored
	// conditional policy.
	// Operations on non-conditional policies may specify any valid value
	// or
	// leave the field unset.
	//
	// If no etag is provided in the call to `setIamPolicy`, version
	// compliance
	// checks against the stored policy is skipped.
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Bindings") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bindings") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__Policy) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__Policy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__SetIamPolicyRequest: Request message for `SetIamPolicy`
// method.
type GoogleIamV1__SetIamPolicyRequest struct {
	// Policy: REQUIRED: The complete policy to be applied to the
	// `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as
	// Projects)
	// might reject them.
	Policy *GoogleIamV1__Policy `json:"policy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Policy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Policy") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__SetIamPolicyRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__SetIamPolicyRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__TestIamPermissionsRequest: Request message for
// `TestIamPermissions` method.
type GoogleIamV1__TestIamPermissionsRequest struct {
	// Permissions: The set of permissions to check for the `resource`.
	// Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For
	// more
	// information see
	// [IAM
	// Overview](https://cloud.google.com/iam/docs/overview#permissions).
	Permissions []string `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__TestIamPermissionsRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__TestIamPermissionsRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleIamV1__TestIamPermissionsResponse: Response message for
// `TestIamPermissions` method.
type GoogleIamV1__TestIamPermissionsResponse struct {
	// Permissions: A subset of `TestPermissionsRequest.permissions` that
	// the caller is
	// allowed.
	Permissions []string `json:"permissions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Permissions") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleIamV1__TestIamPermissionsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleIamV1__TestIamPermissionsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GoogleType__Expr: Represents an expression text. Example:
//
//     title: "User account presence"
//     description: "Determines whether the request has a user account"
//     expression: "size(request.user) > 0"
type GoogleType__Expr struct {
	// Description: An optional description of the expression. This is a
	// longer text which
	// describes the expression, e.g. when hovered over it in a UI.
	Description string `json:"description,omitempty"`

	// Expression: Textual representation of an expression in
	// Common Expression Language syntax.
	//
	// The application context of the containing message determines
	// which
	// well-known feature set of CEL is supported.
	Expression string `json:"expression,omitempty"`

	// Location: An optional string indicating the location of the
	// expression for error
	// reporting, e.g. a file name and a position in the file.
	Location string `json:"location,omitempty"`

	// Title: An optional title for the expression, i.e. a short string
	// describing
	// its purpose. This can be used e.g. in UIs which allow to enter
	// the
	// expression.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GoogleType__Expr) MarshalJSON() ([]byte, error) {
	type NoMethod GoogleType__Expr
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "servicebroker.getIamPolicy":

type V1GetIamPolicyCall struct {
	s            *Service
	resource     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetIamPolicy: Gets the access control policy for a resource.
// Returns an empty policy if the resource exists and does not have a
// policy
// set.
func (r *V1Service) GetIamPolicy(resource string) *V1GetIamPolicyCall {
	c := &V1GetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	return c
}

// OptionsRequestedPolicyVersion sets the optional parameter
// "options.requestedPolicyVersion": The policy format version to be
// returned.
//
// Valid values are 0, 1, and 3. Requests specifying an invalid value
// will be
// rejected.
//
// Requests for policies with any conditional bindings must specify
// version 3.
// Policies without any conditional bindings may specify any valid value
// or
// leave the field unset.
func (c *V1GetIamPolicyCall) OptionsRequestedPolicyVersion(optionsRequestedPolicyVersion int64) *V1GetIamPolicyCall {
	c.urlParams_.Set("options.requestedPolicyVersion", fmt.Sprint(optionsRequestedPolicyVersion))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1GetIamPolicyCall) Fields(s ...googleapi.Field) *V1GetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *V1GetIamPolicyCall) IfNoneMatch(entityTag string) *V1GetIamPolicyCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1GetIamPolicyCall) Context(ctx context.Context) *V1GetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V1GetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1GetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/1.11.0 gdcl/20191203")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "servicebroker.getIamPolicy" call.
// Exactly one of *GoogleIamV1__Policy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleIamV1__Policy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *V1GetIamPolicyCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleIamV1__Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a resource.\nReturns an empty policy if the resource exists and does not have a policy\nset.",
	//   "flatPath": "v1/{v1Id}:getIamPolicy",
	//   "httpMethod": "GET",
	//   "id": "servicebroker.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "options.requestedPolicyVersion": {
	//       "description": "Optional. The policy format version to be returned.\n\nValid values are 0, 1, and 3. Requests specifying an invalid value will be\nrejected.\n\nRequests for policies with any conditional bindings must specify version 3.\nPolicies without any conditional bindings may specify any valid value or\nleave the field unset.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "response": {
	//     "$ref": "GoogleIamV1__Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.setIamPolicy":

type V1SetIamPolicyCall struct {
	s                                *Service
	resource                         string
	googleiamv1__setiampolicyrequest *GoogleIamV1__SetIamPolicyRequest
	urlParams_                       gensupport.URLParams
	ctx_                             context.Context
	header_                          http.Header
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any
// existing policy.
//
// Can return Public Errors: NOT_FOUND, INVALID_ARGUMENT and
// PERMISSION_DENIED
func (r *V1Service) SetIamPolicy(resource string, googleiamv1__setiampolicyrequest *GoogleIamV1__SetIamPolicyRequest) *V1SetIamPolicyCall {
	c := &V1SetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.googleiamv1__setiampolicyrequest = googleiamv1__setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1SetIamPolicyCall) Fields(s ...googleapi.Field) *V1SetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1SetIamPolicyCall) Context(ctx context.Context) *V1SetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V1SetIamPolicyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1SetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/1.11.0 gdcl/20191203")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleiamv1__setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "servicebroker.setIamPolicy" call.
// Exactly one of *GoogleIamV1__Policy or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleIamV1__Policy.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *V1SetIamPolicyCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__Policy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleIamV1__Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified resource. Replaces any\nexisting policy.\n\nCan return Public Errors: NOT_FOUND, INVALID_ARGUMENT and PERMISSION_DENIED",
	//   "flatPath": "v1/{v1Id}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "servicebroker.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "GoogleIamV1__SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIamV1__Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "servicebroker.testIamPermissions":

type V1TestIamPermissionsCall struct {
	s                                      *Service
	resource                               string
	googleiamv1__testiampermissionsrequest *GoogleIamV1__TestIamPermissionsRequest
	urlParams_                             gensupport.URLParams
	ctx_                                   context.Context
	header_                                http.Header
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource.
// If the resource does not exist, this will return an empty set
// of
// permissions, not a NOT_FOUND error.
//
// Note: This operation is designed to be used for building
// permission-aware
// UIs and command-line tools, not for authorization checking. This
// operation
// may "fail open" without warning.
func (r *V1Service) TestIamPermissions(resource string, googleiamv1__testiampermissionsrequest *GoogleIamV1__TestIamPermissionsRequest) *V1TestIamPermissionsCall {
	c := &V1TestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.googleiamv1__testiampermissionsrequest = googleiamv1__testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1TestIamPermissionsCall) Fields(s ...googleapi.Field) *V1TestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1TestIamPermissionsCall) Context(ctx context.Context) *V1TestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V1TestIamPermissionsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V1TestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/1.11.0 gdcl/20191203")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googleiamv1__testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "servicebroker.testIamPermissions" call.
// Exactly one of *GoogleIamV1__TestIamPermissionsResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleIamV1__TestIamPermissionsResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *V1TestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*GoogleIamV1__TestIamPermissionsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GoogleIamV1__TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified resource.\nIf the resource does not exist, this will return an empty set of\npermissions, not a NOT_FOUND error.\n\nNote: This operation is designed to be used for building permission-aware\nUIs and command-line tools, not for authorization checking. This operation\nmay \"fail open\" without warning.",
	//   "flatPath": "v1/{v1Id}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "servicebroker.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "GoogleIamV1__TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleIamV1__TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
