package binding

import (
	"encoding/json"
	"net/http"
	"strings"
)

// This file shamelessly adapted from martini-contrib/binding

type (
	// Errors may be generated during deserialization, binding,
	// or validation. It implements the built-in error interface.
	Errors []FieldError

	// Error is a powerful implementation of the built-in error
	// interface that allows for error classification, custom error
	// messages associated with specific fields, or with no
	// associations at all.
	fieldError struct {
		// An error supports a field names because an
		// error can morph two ways: (1) it can indicate something
		// wrong with the request as a whole, or (2) it can point to a
		// specific problem with a particular input field.
		fieldName string `json:"fieldNames,omitempty"`

		// The classification is like an error code, convenient to
		// use when processing or categorizing an error programmatically.
		// It may also be called the "kind" of error.
		classification string `json:"classification,omitempty"`

		// Message should be human-readable and detailed enough to
		// pinpoint and resolve the problem, but it should be brief. For
		// example, a payload of 100 objects in a JSON array might have
		// an error in the 41st object. The message should help the
		// end user find and fix the error with their request.
		message string `json:"message,omitempty"`
	}

	FieldError interface {
		error
		Field() string
		Kind() string
	}
)

// Add adds an fieldError associated with the field indicated
// by fieldName, with the given classification and message.
func (e *Errors) Add(fieldName string, classification, message string) {
	*e = append(*e, fieldError{
		fieldName:      fieldName,
		classification: classification,
		message:        message,
	})
}

// Len returns the number of errors.
func (e *Errors) Len() int {
	return len(*e)
}

// Has determines whether an e has an Error with
// a given classification in it; it does not search on messages
// or field names.
func (e *Errors) Has(class string) bool {
	for _, err := range *e {
		if err.Kind() == class {
			return true
		}
	}
	return false
}

// Handle writes the errors to response in JSON form if any errors
// are contained, and it will return true. Otherwise, nothing happens
// and false is returned.
// (The value receiver is due to issue 8: https://github.com/mholt/binding/issues/8)
func (e *Errors) Handle(response http.ResponseWriter) bool {
	if e.Len() > 0 {
		response.Header().Set("Content-Type", jsonContentType)
		if e.Has(DeserializationError) {
			response.WriteHeader(http.StatusBadRequest)
		} else if e.Has(ContentTypeError) {
			response.WriteHeader(http.StatusUnsupportedMediaType)
		} else {
			response.WriteHeader(StatusUnprocessableEntity)
		}
		errOutput, _ := json.Marshal(e)
		response.Write(errOutput)
		return true
	}
	return false
}

// Error returns a concatenation of all its error messages.
func (e Errors) Error() string {
	messages := []string{}
	for _, err := range e {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, ", ")
}

// Fields returns the list of field names this error is
// associated with.
func (e fieldError) Field() string {
	return e.fieldName
}

// Kind returns this error's classification.
func (e fieldError) Kind() string {
	return e.classification
}

// Error returns this error's message.
func (e fieldError) Error() string {
	return e.message
}

const (
	RequiredError        = "RequiredError"
	ContentTypeError     = "ContentTypeError"
	DeserializationError = "DeserializationError"
	TypeError            = "TypeError"
)
