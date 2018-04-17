package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const ERROR_FORBIDDEN string = "Forbidden"
const ERROR_NOT_IMPLEMENTED string = "Not implemented"
const ERROR_FORM_DATA_INCOMPLETE string = "Incomplete form data."
const ERROR_EXTRACT_ID_FROM_GET_COULD_NOT_PARSE_ID_STRING string = "Could not parse ID. ID must be unsigned integer value"
const ERROR_EXTRACT_ID_FROM_GET_NO_ID_GIVEN string = "Invalid or no ID given"
const DEFAULT_HEADER_CONTENT_TYPE_KEY string = "Content-Type"
const DEFAULT_HEADER_CONTENT_TYPE_VALUE_JSON string = "application/json; charset=UTF-8"

const FORM_KEY_ID string = "ID"
const FORM_KEY_NAME string = "name"
const FORM_KEY_USERNAME string = "username"
const FORM_KEY_PASSWORD string = "password"
const FORM_KEY_PASSWORD_NEW string = "passwordnew"
const FORM_KEY_PASSWORD_REPEAT string = "passwordrepeat"
const FORM_KEY_TOKEN_VALIDITY_SECONDS string = "validityseconds"

//TODO make these "extract id" methods usable for all to avoid code duplication ... wonky implementation here
func extractIDFromModelGet(urlPath string) (uint32, error) {
	path := strings.Split(urlPath, "/")
	var rawID int
	var id uint32
	var err error
	if len(path) > 0 {
		rawID, err = strconv.Atoi(path[3])
		if err != nil {
			err = errors.New(ERROR_EXTRACT_ID_FROM_GET_COULD_NOT_PARSE_ID_STRING)
		} else {
			id = uint32(rawID) //TODO those "blind" uint32 casts should probably be handled better...
		}
	} else {
		err = errors.New(ERROR_EXTRACT_ID_FROM_GET_NO_ID_GIVEN)
	}
	return id, err
}

func formGetInt32(r *http.Request, key string) (int32, error) {
	var int32Val int32
	var err error
	rawString, err := formGet(r, key)
	if err == nil {
		intVal, err := strconv.Atoi(rawString)
		if err == nil {
			int32Val = int32(intVal) //TODO those "blind" int32 casts should probably be handled better...
		}
	}
	return int32Val, err
}

func formGet(r *http.Request, key string) (string, error) {
	var stringVal string
	var err error

	rawVal := r.Form[key]
	if len(rawVal) > 0 {
		stringVal = strings.Join(rawVal, "")
	} else {
		err = errors.New(ERROR_FORM_DATA_INCOMPLETE)
	}
	return stringVal, err
}

func getResponse(successState int, objectToSerialize interface{}) (int, []byte) {
	status := successState
	response, err := json.Marshal(objectToSerialize)
	if err != nil {
		status = http.StatusInternalServerError
		response, err = json.Marshal(Error{Message: err.Error()})
		if err != nil {
			response = []byte(err.Error())
		}
	}
	return status, response
}

func getErrorResponse(successState int, message string) (int, []byte) {
	status := successState
	response, err := json.Marshal(Error{Message: message})
	if err != nil {
		response = []byte(message)
	}
	return status, response
}

func getErrorForbiddenResponse() (int, []byte) {
	return getErrorResponse(http.StatusForbidden, ERROR_FORBIDDEN)
}
