package s2s

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sberryman/playground-golang/utils"
)

// Op provides an S2S operation request and response
type Op struct {
	// LogCtx      *logrus.Entry
	BearerToken string // optional bearer token to replace default S2S
	Hostname    string
	Route       string
	ReqBody     interface{}       `json:"req_body,omitempty"`  // request body data. Required for POST
	ReqQuery    map[string]string `json:"req_query,omitempty"` // query inputs if applicable
	Resp        interface{}       // pointer to structure to unmarshal response into
	RespCode    int
	// internal usage
	url      string
	name     string
	httpResp *http.Response
	err      error
}

// NewOp creates a new S2S operation
func NewOp(op *Op) *Op {
	return op
}

// Get sends an S2S GET request. After the operation the LogCtx will be updated to contain the request and response
func (op *Op) Get() error {
	op.name = "s2s.Get"
	op.url = fmt.Sprintf("%s/%s", strings.TrimRight(op.Hostname, "/"), strings.TrimLeft(op.Route, "/"))
	// op.LogCtx = op.LogCtx.WithFields(logrus.Fields{"url": op.url, "req_query": op.ReqQuery})

	op.httpResp, op.err = singleShared.client.Get(
		op.url,
		func(req *http.Request) {
			if op.BearerToken != "" {
				// Replace S2S authorization with user's provided
				req.Header.Del("Authorization")
				req.Header.Set("Authorization", "Bearer "+op.BearerToken)
			} else if op.Hostname == utils.Config.Powerline.ClusterHostname {
				// Replace S2S authorization with powerline's
				req.Header.Del("Authorization")
				req.Header.Set("Authorization", "Bearer "+singleShared.pltoken)
			}

			if op.ReqQuery == nil || len(op.ReqQuery) == 0 {
				return
			}

			q := req.URL.Query()
			for key, val := range op.ReqQuery {
				q.Add(key, val)
			}

			req.URL.RawQuery = q.Encode()
		},
	)

	// process response
	return op.parseResponse()
}

// Post sends an S2S POST request. After the operation the LogCtx will be updated to contain the request and response
func (op *Op) Post() error {
	op.name = "s2s.Post"
	op.url = fmt.Sprintf("%s/%s", strings.TrimRight(op.Hostname, "/"), strings.TrimLeft(op.Route, "/"))
	// op.LogCtx = op.LogCtx.WithFields(logrus.Fields{"url": op.url, "req_body": op.ReqBody, "req_query": op.ReqQuery})

	var reqBody []byte
	var err error

	if op.ReqBody != nil {
		reqBody, err = json.Marshal(op.ReqBody)
		if err != nil {
			// op.LogCtx.WithError(err).Error(op.name + "-Marshal")
			return err
		}
	}

	// send request
	op.httpResp, op.err = singleShared.client.Post(
		op.url,
		"application/json",
		bytes.NewReader(reqBody),
		func(req *http.Request) {
			if op.BearerToken != "" {
				// Replace S2S authorization with user's provided
				req.Header.Del("Authorization")
				req.Header.Set("Authorization", "Bearer "+op.BearerToken)
			} else if op.Hostname == utils.Config.Powerline.ClusterHostname {
				// Replace S2S authorization with powerline's
				req.Header.Del("Authorization")
				req.Header.Set("Authorization", "Bearer "+singleShared.pltoken)
			}
		},
	)

	// process response
	return op.parseResponse()
}

// parses the response of the s2s request - common for get/post
func (op *Op) parseResponse() error {
	if op.err != nil {
		// op.LogCtx.WithError(op.err).Error(op.name + "-Req")
		return op.err
	}

	op.RespCode = op.httpResp.StatusCode
	// op.LogCtx = op.LogCtx.WithField("resp_status", op.httpResp.Status)

	// Get body
	defer op.httpResp.Body.Close()
	respBody, err := ioutil.ReadAll(op.httpResp.Body)
	if err != nil {
		// op.LogCtx.WithError(err).Error(op.name + "-ReadAll")
		return err
	}

	// op.LogCtx = op.LogCtx.WithField("resp_body", string(respBody))

	// Unmarshall into response
	if strings.Contains(op.httpResp.Header.Get("Content-Type"), "application/json") {
		err = json.Unmarshal(respBody, op.Resp)
		if err != nil {
			// op.LogCtx.WithError(err).Error(op.name + "-Unmarshal")
			return err
		}
	}

	// Autofail for != 2xx codes
	if err == nil && op.httpResp.StatusCode >= 300 {
		err = fmt.Errorf(op.httpResp.Status)
		// op.LogCtx.Error(op.name + "-Code")
		return err
	}

	return nil
}
