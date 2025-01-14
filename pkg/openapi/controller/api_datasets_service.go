// Copyright (c) 2021 Cisco Systems, Inc. and its affiliates
// All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Flame REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cisco-open/flame/cmd/controller/app/database"
	"github.com/cisco-open/flame/pkg/openapi"
)

// DatasetsApiService is a service that implents the logic for the DatasetsApiServicer
// This service should implement the business logic for every endpoint for the DatasetsApi API.
// Include any external packages or services that will be required by this service.
type DatasetsApiService struct {
	dbService database.DBService
}

// NewDatasetsApiService creates a default api service
func NewDatasetsApiService(dbService database.DBService) openapi.DatasetsApiServicer {
	return &DatasetsApiService{
		dbService: dbService,
	}
}

// CreateDataset - Create meta info for a new dataset.
func (s *DatasetsApiService) CreateDataset(ctx context.Context, user string,
	datasetInfo openapi.DatasetInfo) (openapi.ImplResponse, error) {
	datasetId, err := s.dbService.CreateDataset(user, datasetInfo)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("failed to create new dataset: %v", err)
	}

	return openapi.Response(http.StatusCreated, datasetId), nil
}

// GetAllDatasets - Get the meta info on all the datasets
func (s *DatasetsApiService) GetAllDatasets(ctx context.Context, limit int32) (openapi.ImplResponse, error) {
	datasetList, err := s.dbService.GetDatasets("", limit)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("failed to get list of datasets: %v", err)
	}

	return openapi.Response(http.StatusOK, datasetList), nil
}

// GetDataset - Get dataset meta information
func (s *DatasetsApiService) GetDataset(ctx context.Context, user string, datasetId string) (openapi.ImplResponse, error) {
	// TODO - update GetDataset with the required logic for this service method.
	// Add api_datasets_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, DatasetInfo{}) or use other options such as http.Ok ...
	//return Response(200, DatasetInfo{}), nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("GetDataset method not implemented")
}

// GetDatasets - Get the meta info on all the datasets owned by user
func (s *DatasetsApiService) GetDatasets(ctx context.Context, user string, limit int32) (openapi.ImplResponse, error) {
	datasetList, err := s.dbService.GetDatasets(user, limit)
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, nil), fmt.Errorf("failed to get list of datasets: %v", err)
	}

	return openapi.Response(http.StatusOK, datasetList), nil
}

// UpdateDataset - Update meta info for a given dataset
func (s *DatasetsApiService) UpdateDataset(ctx context.Context, user string, datasetId string,
	datasetInfo openapi.DatasetInfo) (openapi.ImplResponse, error) {
	// TODO - update UpdateDataset with the required logic for this service method.
	// Add api_datasets_service.go to the .openapi-generator-ignore to avoid overwriting this service
	// implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(0, Error{}) or use other options such as http.Ok ...
	//return Response(0, Error{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("UpdateDataset method not implemented")
}
