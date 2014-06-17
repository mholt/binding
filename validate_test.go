package binding

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestValidate(t *testing.T) {

	Convey("Given a struct populated properly and as expected", t, func() {

		Convey("No errors should be produced", nil)

	})

	Convey("Given a struct that is a Validator", t, func() {

		Convey("The user's Validate method should be invoked and its errors appended", nil)

	})

	Convey("Each case in the type switch should be tested", t, func() {

		Convey("uint8", nil)

		Convey("uint16", nil)

		Convey("uint32", nil)

		Convey("uint64", nil)

		Convey("int8", nil)

		Convey("int16", nil)

		Convey("int32", nil)

		Convey("int64", nil)

		Convey("float32", nil)

		Convey("[]float32", nil)

		Convey("float64", nil)

		Convey("[]float64", nil)

		Convey("uint", nil)

		Convey("[]uint", nil)

		Convey("int", nil)

		Convey("[]int", nil)

		Convey("bool", nil)

		Convey("[]bool", nil)

		Convey("string", nil)

		Convey("[]string", nil)

		Convey("time.Time", nil)

		Convey("[]time.Time", nil)

	})

}
