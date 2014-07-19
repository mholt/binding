package binding

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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

		Convey("Missing one required field", func() {

			Convey("A Required error should be produced", nil)

		})

		Convey("Missing multiple required fields", func() {

			Convey("As many Required errors should be produced", nil)

		})
	})

}
