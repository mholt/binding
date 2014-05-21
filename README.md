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

However, this handler ignores any errors. Suppose that `message` was not filled out and is an empty string. There is a built-in error handler that writes the errors to the response for you. You can use it like this:

```go
if errors.Handle(resp) {
	return
}
```

As you can see, `errors.Handle()` returns `true` if there were errors that it wrote to the response.


Supported types
----------------

- uint, uint8, uint16, uint32, uint64
- int, int8, int16, int32, int64
- float32, float64
- bool
- string
- time.Time
