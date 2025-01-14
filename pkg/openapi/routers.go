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

package openapi

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// A Route defines the parameters for an api endpoint
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes are a collection of defined api endpoints
type Routes []Route

// Router defines the required methods for retrieving api routes
type Router interface {
	Routes() Routes
}

const errMsgRequiredMissing = "required parameter is missing"

// NewRouter creates a new router for any number of api routers
func NewRouter(routers ...Router) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, api := range routers {
		for _, route := range api.Routes() {
			var handler http.Handler
			handler = route.HandlerFunc
			handler = Logger(handler, route.Name)

			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	}

	return router
}

// EncodeJSONResponse uses the json encoder to write an interface to the http response with an optional status code
func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return json.NewEncoder(w).Encode(i)
}

// ReadFormFileToTempFile reads file data from a request form and writes it to a temporary file
func ReadFormFileToTempFile(r *http.Request, key string) (*os.File, error) {
	_, fileHeader, err := r.FormFile(key)
	if err != nil {
		return nil, err
	}

	return readFileHeaderToTempFile(fileHeader)
}

// ReadFormFilesToTempFiles reads files array data from a request form and writes it to a temporary files
func ReadFormFilesToTempFiles(r *http.Request, key string) ([]*os.File, error) {
	var maxMemory int64 = 32 << 20 // 32MB
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		return nil, err
	}

	files := make([]*os.File, 0, len(r.MultipartForm.File[key]))

	for _, fileHeader := range r.MultipartForm.File[key] {
		file, err := readFileHeaderToTempFile(fileHeader)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

// readFileHeaderToTempFile reads multipart.FileHeader and writes it to a temporary file
func readFileHeaderToTempFile(fileHeader *multipart.FileHeader) (*os.File, error) {
	formFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	defer formFile.Close()

	fileBytes, err := ioutil.ReadAll(formFile)
	if err != nil {
		return nil, err
	}

	// Note by Myungjin: fileHeader.Filename is not used because it can have
	//                   path separator (e.g., /),  which cannot be used as prefix
	//                   in tempfile generation, thus causing
	//                   "pattern contains path separator" error
	// file, err := ioutil.TempFile("", fileHeader.Filename)
	file, err := ioutil.TempFile("", "flame")
	if err != nil {
		return nil, err
	}

	// Note by Myungjin: the tmp file shouldn't be closed; otherwise, access to
	//                   the temp file will fail close it only after the tempfile
	//                   is processed
	// defer file.Close()

	file.Write(fileBytes)

	// Note by Myungjin: This line is added to move the file descriptor to the start
	//                   of the file so that the file can be accessed later on
	// Rewind to the start of file
	file.Seek(0, io.SeekStart)

	return file, nil
}

/*
// parseInt64Parameter parses a string parameter to an int64.
func parseInt64Parameter(param string, required bool) (int64, error) {
	if param == "" {
		if required {
			return 0, errors.New(errMsgRequiredMissing)
		}

		return 0, nil
	}

	return strconv.ParseInt(param, 10, 64)
}
*/

// parseInt32Parameter parses a string parameter to an int32.
func parseInt32Parameter(param string, required bool) (int32, error) {
	if param == "" {
		if required {
			return 0, errors.New(errMsgRequiredMissing)
		}

		return 0, nil
	}

	base := 10
	bitSize := 32
	val, err := strconv.ParseInt(param, base, bitSize)
	if err != nil {
		return -1, err
	}

	return int32(val), nil
}

/*
// parseBoolParameter parses a string parameter to a bool
func parseBoolParameter(param string) (bool, error) {
	val, err := strconv.ParseBool(param)
	if err != nil {
		return false, err
	}

	return bool(val), nil
}

// parseInt64ArrayParameter parses a string parameter containing array of values to []int64.
func parseInt64ArrayParameter(param, delim string, required bool) ([]int64, error) {
	if param == "" {
		if required {
			return nil, errors.New(errMsgRequiredMissing)
		}

		return nil, nil
	}

	str := strings.Split(param, delim)
	ints := make([]int64, len(str))

	for i, s := range str {
		if v, err := strconv.ParseInt(s, 10, 64); err != nil {
			return nil, err
		} else {
			ints[i] = v
		}
	}

	return ints, nil
}

// parseInt32ArrayParameter parses a string parameter containing array of values to []int32.
func parseInt32ArrayParameter(param, delim string, required bool) ([]int32, error) {
	if param == "" {
		if required {
			return nil, errors.New(errMsgRequiredMissing)
		}

		return nil, nil
	}

	str := strings.Split(param, delim)
	ints := make([]int32, len(str))

	for i, s := range str {
		if v, err := strconv.ParseInt(s, 10, 32); err != nil {
			return nil, err
		} else {
			ints[i] = int32(v)
		}
	}

	return ints, nil
}
*/
