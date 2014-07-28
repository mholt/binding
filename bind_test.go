package binding

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBind(t *testing.T) {
	Convey("A request", t, func() {

		Convey("Without a Content-Type", func() {

			Convey("But with a query string", func() {

				Convey("Should invoke the Form deserializer", nil)

			})

			Convey("And without a query string", func() {

				Convey("Should yield an error", nil)

			})

		})

		Convey("With a form-urlencoded Content-Type", func() {

			Convey("Should invoke the Form deserializer", nil)

		})

		Convey("With a multipart/form-data Content-Type", func() {

			Convey("Should invoke the MultipartForm deserializer", nil)

		})

		Convey("With a json Content-Type", func() {

			Convey("Should invoke the Json deserializer", nil)

		})

		Convey("With an unsupported Content-Type", func() {

			Convey("Should yield an error", nil)

		})

	})
}

func TestBindForm(t *testing.T) {
	Convey("Given a struct reference and complete form data", t, func() {
		cm := NewCompleteModel()
		formData := cm.FormValues()

		Convey("Given that all of the struct's fields are required", func() {
			model := AllTypes{}
			Convey("When bindForm is called", func() {
				req, err := http.NewRequest("POST", "http://www.example.com", nil)
				So(err, ShouldBeNil)
				var errs Errors
				errs = bindForm(req, &model, formData, nil, errs)
				Convey("Then all of the struct's fields should be populated", func() {

				})

				Convey("Then no errors should be produced", FailureContinues, func() {
					So(errs.Len(), ShouldEqual, 0)
					if errs.Len() > 0 {
						for _, e := range errs {
							Println(fmt.Sprintf("%v. %s", e.FieldNames, e.Message))
						}
					}
				})
			})
		})
	})
}
