package binding

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func keys(m map[string][]string) []string {
	k := make([]string, len(m))
	for oneKey := range m {
		k = append(k, oneKey)
	}
	return k
}

func TestJsonUtilities(t *testing.T) {
	goodJson := "{\"L1-A\": {\"L2\": \"V2\"}, \"L1-B\": \"V1-B\"}"

	Convey("Given well-formed JSON input", t, func() {
		ioReader := bytes.NewBufferString(goodJson)
		result, err := FlatDecode(ioReader)

		Convey("nested structures are unmarshalled and flattened", func() {
			So(err, ShouldBeNil)
			So(len(result), ShouldEqual, 2)
			resultKeys := keys(result)
			So(resultKeys, ShouldContain, "L1-B")
			So(resultKeys, ShouldContain, "L1-A.L2")
		})

		Convey("JSON values get unquoted", func() {
			So(result["L1-A.L2"][0], ShouldEqual, "V2")
			So(result["L1-A.L2"][0], ShouldNotEqual, "\"V2\"")
			So(result["L1-B"][0], ShouldEqual, "V1-B")
		})

	})
}
