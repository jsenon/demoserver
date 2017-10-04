// Package api demoserver.
//
// the purpose of this package is to provide Api Interface
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost:9010
//     BasePath: /api
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package api

import (
	"net/http"

	"encoding/json"
	"fmt"
)

// Input failed validation.
// swagger:response validationError
type ValidationError struct {
	// The error message
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// Success Answer
// swagger:response successAnswer
type successAnswer struct {
	// The error message
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// Status Answer
// swagger:response statusjson
type statusjson struct {
	// The Status message
	// in: body
	Code    int32  `json:"statuscode"`
	Message string `json:"statusmessage"`
}

// swagger:route GET /healthy/am-i-up health amiup
//
// Check health of platform
//
// This will sent information if service is up and running.
//
//     Responses:
//       default: validationError
//       200: statusjson
func Statusamiup(w http.ResponseWriter, req *http.Request) {
	answerjson := statusjson{
		Code:    200,
		Message: "Im awake !",
	}
	b, err := json.Marshal(answerjson)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write(b)
}

// swagger:route GET /healthy/about health about
//
// Check service availability
//
// This will details status availability
//
//     Responses:
//       default: validationError
//       200: statusjson
func Statusabout(w http.ResponseWriter, req *http.Request) {
	answerjson := statusjson{
		Code:    200,
		Message: "Made by Somebody",
	}
	b, err := json.Marshal(answerjson)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write(b)
}
