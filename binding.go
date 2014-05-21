package binding

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Bind takes data out of the request and deserializes into a struct according
// to the Content-Type of the request. If no Content-Type is specified, there
// better be data in the query string, otherwise an error will be produced.
func Bind(req *http.Request, userStruct FieldMapper) (errors Errors) {
	contentType := req.Header.Get("Content-Type")

	if strings.Contains(contentType, "form-urlencoded") {
		return Form(req, userStruct)
	} else if strings.Contains(contentType, "multipart/form-data") {
		return MultipartForm(req, userStruct)
	} else if strings.Contains(contentType, "json") {
		return Json(req, userStruct)
	} else {
		if contentType == "" {
			if len(req.URL.Query()) > 0 {
				return Form(req, userStruct)
			} else {
				errors.Add([]string{}, ContentTypeError, "Empty Content-Type")
			}
		} else {
			errors.Add([]string{}, ContentTypeError, "Unsupported Content-Type")
		}
	}

	return
}

// Form deserializes form data out of the request into a struct you provide.
// This function invokes data validation after deserialization.
func Form(req *http.Request, userStruct FieldMapper) (errors Errors) {
	parseErr := req.ParseForm()
	if parseErr != nil {
		errors.Add([]string{}, DeserializationError, parseErr.Error())
		return
	}

	fm := userStruct.FieldMap()
	for fieldName, fieldPointer := range fm {
		str := req.Form.Get(fieldName)

		errorHandler := func(err error) {
			if err != nil {
				errors.Add([]string{fieldName}, TypeError, err.Error())
			}
		}

		switch t := fieldPointer.(type) {
		case *uint8:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseUint(str, 10, 8) // TODO: Should bases be 0 so they are auto-detected, or should we keep assuming base 10?
			errorHandler(err)
			*t = uint8(val)
		case *uint16:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseUint(str, 10, 16)
			errorHandler(err)
			*t = uint16(val)
		case *uint32:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseUint(str, 10, 32)
			errorHandler(err)
			*t = uint32(val)
		case *uint64:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseUint(str, 10, 64)
			errorHandler(err)
			*t = val
		case *int8:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseInt(str, 10, 8)
			errorHandler(err)
			*t = int8(val)
		case *int16:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseInt(str, 10, 16)
			errorHandler(err)
			*t = int16(val)
		case *int32:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseInt(str, 10, 32)
			errorHandler(err)
			*t = int32(val)
		case *int64:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseInt(str, 10, 64)
			errorHandler(err)
			*t = val
		case *float32:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseFloat(str, 32)
			errorHandler(err)
			*t = float32(val)
		case *float64:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseFloat(str, 64)
			errorHandler(err)
			*t = val
		case *uint:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseUint(str, 10, 0)
			errorHandler(err)
			*t = uint(val)
		case *int:
			if str == "" {
				str = "0"
			}
			val, err := strconv.ParseInt(str, 10, 0)
			errorHandler(err)
			*t = int(val)
		case *bool:
			val, err := strconv.ParseBool(str)
			errorHandler(err)
			*t = val
		case *string:
			*t = str
		case *time.Time:
			//*t, _ = time.Parse()
		}
	}

	errors = append(errors, Validate(req, userStruct)...)

	return
}

// MultipartForm reads a multipart form request and deserializes its data into
// a struct you provide. It then calls Form to get the rest of the form data
// out of the request.
func MultipartForm(req *http.Request, userStruct FieldMapper) (errors Errors) {
	multipartReader, err := req.MultipartReader()
	if err != nil {
		errors.Add([]string{}, DeserializationError, err.Error())
		return
	} else {
		form, parseErr := multipartReader.ReadForm(MaxMemory)
		if parseErr != nil {
			errors.Add([]string{}, DeserializationError, parseErr.Error())
			return
		}
		req.MultipartForm = form
	}
	return Form(req, userStruct)
}

// Json deserializes a JSON request body into a struct you specify
// using the standard encoding/json package (which uses reflection).
// This function invokes data validation after deserialization.
func Json(req *http.Request, userStruct FieldMapper) (errors Errors) {
	if req.Body != nil {
		defer req.Body.Close()
		err := json.NewDecoder(req.Body).Decode(userStruct)
		if err != nil && err != io.EOF {
			errors.Add([]string{}, DeserializationError, err.Error())
			return
		}
	} else {
		errors.Add([]string{}, DeserializationError, "Empty request body")
		return
	}

	errors = append(errors, Validate(req, userStruct)...)

	return
}

// Validate ensures that all conditions have been met on every field in the
// populated struct. Validation should occur after the request has been
// deserialized into the struct.
func Validate(req *http.Request, userStruct FieldMapper) (errors Errors) {
	fm := userStruct.FieldMap()

	for fieldName, fieldSpec := range fm {
		addRequiredError := func() {
			errors.Add([]string{fieldName}, RequiredError, "Required")
		}

		if field, ok := fieldSpec.(Field); ok {
			if field.Required {
				switch t := field.Target.(type) {
				case *uint8:
					if *t == 0 {
						addRequiredError()
					}
				case *uint16:
					if *t == 0 {
						addRequiredError()
					}
				case *uint32:
					if *t == 0 {
						addRequiredError()
					}
				case *uint64:
					if *t == 0 {
						addRequiredError()
					}
				case *int8:
					if *t == 0 {
						addRequiredError()
					}
				case *int16:
					if *t == 0 {
						addRequiredError()
					}
				case *int32:
					if *t == 0 {
						addRequiredError()
					}
				case *int64:
					if *t == 0 {
						addRequiredError()
					}
				case *float32:
					if *t == 0 {
						addRequiredError()
					}
				case *float64:
					if *t == 0 {
						addRequiredError()
					}
				case *uint:
					if *t == 0 {
						addRequiredError()
					}
				case *int:
					if *t == 0 {
						addRequiredError()
					}
				case *bool:
					if *t == false {
						addRequiredError()
					}
				case *string:
					if *t == "" {
						addRequiredError()
					}
				case *time.Time:
					if t.IsZero() {
						addRequiredError()
					}
				}
			}
		}
	}

	if validator, ok := userStruct.(Validator); ok {
		errors = validator.Validate(errors, req)
	}

	return
}

type (
	// Only types that are FieldMappers can be used to bind requests.
	FieldMapper interface {

		// FieldMap returns a map that keys field names from the request
		// to pointers into which the values will be deserialized.
		FieldMap() FieldMap
	}

	// FieldMap is a map of field names in the request to pointers into
	// which the values will be deserialized.
	FieldMap map[string]interface{}

	// Field describes the properties of a field.
	Field struct {
		Target   interface{}
		Required bool
	}

	// Validator can be implemented by your type to handle some
	// rudimentary request validation separately from your
	// application logic.
	Validator interface {
		// Validate validates that the request is OK. It is recommended
		// that validation be limited to checking values for syntax and
		// semantics, enough to know that you can make sense of the request
		// in your application. For example, you might verify that a credit
		// card number matches a valid pattern, but you probably wouldn't
		// perform an actual credit card authorization here.
		Validate(Errors, *http.Request) Errors
	}
)

var (
	// Maximum amount of memory to use when parsing a multipart form.
	// Set this to whatever value you prefer; default is 10 MB.
	MaxMemory = int64(1024 * 1024 * 10)
)

const (
	jsonContentType           = "application/json; charset=utf-8"
	StatusUnprocessableEntity = 422
)
