![binding is reflectionless](http://mholt.github.io/binding/resources/images/binding-sm.png)

binding
=======

Reflectionless data binding for Go's net/http



Features
---------

- Deserializes form, multipart form, and JSON data from requests
- Not middleware: just a function call
- Built-in error handling
- Performs data validation
- Usable in any setting where `net/http` is present (Negroni, gocraft/web, std lib, etc.)


Usage example
--------------

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/mholt/binding"
)

// Define types to hold request data; you can also decorate
// them with struct tags for JSON deserialization.
// (For a convenient way to convert JSON to Go structs,
// see: http://mholt.github.io/json-to-go)
type (
	ContactForm struct {
		User struct {
			ID   int
			Name string
		}
		Email   string
		Message string
	}
)

// Pointer receiver is vital here
func (cf *ContactForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		"user_id": &cf.User.ID,
		"name":    &cf.User.Name,
		"email":   &cf.Email,
		"message": binding.Field{
			Target:   &cf.Message,
			Required: true,
		},
	}
}

// You may optionally implement the binding.Validator interface
// for custom data validation
func (cf ContactForm) Validate(errors binding.Errors, req *http.Request) binding.Errors {
	if cf.Message == "Go needs generics" {
		errors = append(errors, binding.Error{
			FieldNames:     []string{"message"},
			Classification: "ComplaintError",
			Message:        "Go has generics. They're called interfaces.",
		})
	}
	return errors
}

// Now data binding, validation, and error handling is taken care of while
// keeping your application handler clean and simple.
func handler(resp http.ResponseWriter, req *http.Request) {
	contactForm := new(ContactForm)
	errs := binding.Bind(req, contactForm)
	if errs.Handle(resp) {
		return
	}
	fmt.Fprintf(resp, "From:    %s\n", contactForm.User.Name)
	fmt.Fprintf(resp, "Message: %s\n", contactForm.Message)
}

func main() {
	http.HandleFunc("/contact", handler)
	http.ListenAndServe(":3000", nil)
}
```



Error Handling
---------------

`binding.Bind()` and the other deserializers return errors. You don't have to use them, but the `binding.Errors` type comes with a kind of built-in "handler" to write the errors to the response as JSON for you. For example, you might do this in your HTTP handler:

```go
if binding.Bind(req, contactForm).Handle(resp) {
	return
}
```

As you can see, if `.Handle()` wrote errors to the response, your handler may gracefully exit.




Supported types (forms)
------------------------

The following types are supported in form deserialization. (JSON requests are delegated to `encoding/json`.)

- uint, []uint, uint8, uint16, uint32, uint64
- int, []int, int8, int16, int32, int64
- float32, []float32, float64, []float64
- bool, []bool
- string, []string
- time.Time, []time.Time
