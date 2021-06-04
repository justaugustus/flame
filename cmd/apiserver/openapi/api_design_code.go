/*
 * Fledge REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

// A DesignCodeApiController binds http requests to an api service and writes the service results to the http response
type DesignCodeApiController struct {
	service DesignCodeApiServicer
}

// NewDesignCodeApiController creates a default api controller
func NewDesignCodeApiController(s DesignCodeApiServicer) Router {
	return &DesignCodeApiController{service: s}
}

// Routes returns all of the api route for the DesignCodeApiController
func (c *DesignCodeApiController) Routes() Routes {
	return Routes{
		{
			"GetDesignCode",
			strings.ToUpper("Get"),
			"/designs/{user}/{design}/code",
			c.GetDesignCode,
		},
		{
			"UpdateDesignCode",
			strings.ToUpper("Post"),
			"/designs/{user}/{design}/code",
			c.UpdateDesignCode,
		},
	}
}

// GetDesignCode - Get a zipped design code file owned by user
func (c *DesignCodeApiController) GetDesignCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	design := params["design"]

	result, err := c.service.GetDesignCode(r.Context(), user, design)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateDesignCode - Update a design doce
func (c *DesignCodeApiController) UpdateDesignCode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["user"]

	design := params["design"]

	body := os.File{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.UpdateDesignCode(r.Context(), user, design, &body)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}