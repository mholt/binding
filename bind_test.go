package binding

import (
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
