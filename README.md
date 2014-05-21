![binding is reflectionless](http://mholt.github.io/binding/resources/images/binding-sm.png)

binding
=======

Reflectionless data binding for Go's net/http



Features
---------

- Deserializes form, multipart form, or JSON data from requests
- Not middleware: just a function call or two
- Usable in any setting where `net/http` is present


Usage
------

Suppose you have a contact form on your site that takes a user ID, email, and a message, and at least the message is required. Make a struct to hold the data:

```go
type ContactForm struct {
    UserID  int
    Email   string
    Message string
}
```

Then we have it implement `binding.FieldMapper` so we can bind to it. This is nearly as easy as struct tags, doesn't require reflection, and is more flexible:

```go
func (cf *ContactForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		"user_id": &cf.UserID,
		"email":   &cf.Email,
		"message": binding.Field{
			Target:   &cf.Message,
			Required: true,
		}
	}
}
```

Notice that "message" was mapped differently. To add properties to a field, we use `binding.Field`. Otherwise, it is not necessary.

Then in your HTTP handler, you can use `binding.Bind`:

```go
func handler(resp http.ResponseWriter, req *http.Request) {
	contactForm := new(ContactForm)
	binding.Bind(req, contactForm)

	fmt.Fprintf(resp, "Message from: %s", contactForm.Email)
}
```


Validation
-----------

Validating the data is supported out-of-the-box. Just implement the `binding.Validator` interface on your type:

```go
func (cf ContactForm) Validate(errors binding.Errors, req *http.Request) binding.Errors {
	if len(cf.Message) < 5 {
		errors.Add([]string{"message"}, "LengthError", "Message should be at least 5 characters")
	}
	return errors
}
```

Your errors will be combined with the ones produced by `Form` or `Json` deserializers.



Error Handling
---------------

Errors are returned from calls to `binding.Bind()` and the other deserializers. You can ignore them if you want, or you can use them. The `binding.Errors` type comes with a kind of built-in "handler" to write the errors to the response as JSON for you. For example, you might do this in your HTTP handler:

```go
errors := binding.Bind(req, contactForm)
if errors.Handle(resp) {
	return
}
```

As you can see, if `errors.Handle()` wrote errors to the response, your handler may gracefully exit.




Supported types
----------------

The following list is for form deserialization. (JSON requests are delegated to `encoding/json` so any type that can be marshalled/unmarshalled is supported.)

- uint, uint8, uint16, uint32, uint64
- int, int8, int16, int32, int64
- float32, float64
- bool
- string
- time.Time
