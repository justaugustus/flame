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
	"context"
	"errors"
	"net/http"

	"wwwin-github.cisco.com/eti/fledge/cmd/controller/database"
	"wwwin-github.cisco.com/eti/fledge/pkg/objects"
)

// DesignApiService is a service that implents the logic for the DesignApiServicer
// This service should implement the business logic for every endpoint for the DesignApi API.
// Include any external packages or services that will be required by this service.
type DesignApiService struct {
}

// NewDesignApiService creates a default api service
func NewDesignApiService() DesignApiServicer {
	return &DesignApiService{}
}

// CreateDesign - Create a new design template.
func (s *DesignApiService) CreateDesign(ctx context.Context, user string, designInfo objects.DesignInfo) (ImplResponse, error) {
	var d = objects.Design{
		UserId:      user,
		Name:        designInfo.Name,
		Description: designInfo.Description,
		Schemas:     []objects.DesignSchema{},
	}
	err := database.CreateDesign(user, d)
	if err != nil {
		return Response(http.StatusInternalServerError, nil), errors.New("create new design request failed")
	}
	return Response(http.StatusCreated, nil), nil
}

// GetDesign - Get design template information
func (s *DesignApiService) GetDesign(ctx context.Context, user string, designId string) (ImplResponse, error) {
	info, err := database.GetDesign(user, designId)
	if err != nil {
		return Response(http.StatusInternalServerError, nil), errors.New("get design template information request failed")
	}
	return Response(http.StatusOK, info), nil
}