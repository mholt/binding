package binding

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type AllTypes struct {
	Uint8            uint8
	PointerToUint8   *uint8
	Uint8Slice       []uint8
	Uint16           uint16
	PointerToUint16  *uint16
	Uint16Slice      []uint16
	Uint32           uint32
	PointerToUint32  *uint32
	Uint32Slice      []uint32
	Uint64           uint64
	PointerToUint64  *uint64
	Uint64Slice      []uint64
	Int8             int8
	PointerToInt8    *int8
	Int8Slice        []int8
	Int16            int16
	PointerToInt16   *int16
	Int16Slice       []int16
	Int32            int32
	PointerToInt32   *int32
	Int32Slice       []int32
	Int64            int64
	PointerToInt64   *int64
	Int64Slice       []int64
	Float32          float32
	PointerToFloat32 *float32
	Float32Slice     []float32
	Float64          float64
	PointerToFloat64 *float64
	Float64Slice     []float64
	Uint             uint
	PointerToUint    *uint
	UintSlice        []uint
	Int              int
	PointerToInt     *int
	IntSlice         []int
	Bool             bool
	PointerToBool    *bool
	BoolSlice        []bool
	String           string
	PointerToString  *string
	StringSlice      []string
	Time             time.Time
	PointerToTime    *time.Time
	TimeSlice        []time.Time
}

func (at *AllTypes) FieldMap() FieldMap {
	return FieldMap{
		&at.Uint8:            Field{Form: "uint8", Required: true},
		&at.PointerToUint8:   Field{Form: "pointerToUint8", Required: true},
		&at.Uint8Slice:       Field{Form: "uint8Slice", Required: true},
		&at.Uint16:           Field{Form: "uint16", Required: true},
		&at.PointerToUint16:  Field{Form: "pointerToUint16", Required: true},
		&at.Uint16Slice:      Field{Form: "uint16Slice", Required: true},
		&at.Uint32:           Field{Form: "uint32", Required: true},
		&at.PointerToUint32:  Field{Form: "pointerToUint32", Required: true},
		&at.Uint32Slice:      Field{Form: "uint32Slice", Required: true},
		&at.Uint64:           Field{Form: "uint64", Required: true},
		&at.PointerToUint64:  Field{Form: "pointerToUint64", Required: true},
		&at.Uint64Slice:      Field{Form: "uint64Slice", Required: true},
		&at.Int8:             Field{Form: "int8", Required: true},
		&at.PointerToInt8:    Field{Form: "pointerToInt8", Required: true},
		&at.Int8Slice:        Field{Form: "int8Slice", Required: true},
		&at.Int16:            Field{Form: "int16", Required: true},
		&at.PointerToInt16:   Field{Form: "pointerToInt16", Required: true},
		&at.Int16Slice:       Field{Form: "int16Slice", Required: true},
		&at.Int32:            Field{Form: "int32", Required: true},
		&at.PointerToInt32:   Field{Form: "pointerToInt32", Required: true},
		&at.Int32Slice:       Field{Form: "int32Slice", Required: true},
		&at.Int64:            Field{Form: "int64", Required: true},
		&at.PointerToInt64:   Field{Form: "pointerToInt64", Required: true},
		&at.Int64Slice:       Field{Form: "int64Slice", Required: true},
		&at.Float32:          Field{Form: "float32", Required: true},
		&at.PointerToFloat32: Field{Form: "pointerToFloat32", Required: true},
		&at.Float32Slice:     Field{Form: "float32Slice", Required: true},
		&at.Float64:          Field{Form: "float64", Required: true},
		&at.PointerToFloat64: Field{Form: "pointerToFloat64", Required: true},
		&at.Float64Slice:     Field{Form: "float64Slice", Required: true},
		&at.Uint:             Field{Form: "uint", Required: true},
		&at.PointerToUint:    Field{Form: "pointerToUint", Required: true},
		&at.UintSlice:        Field{Form: "uintSlice", Required: true},
		&at.Int:              Field{Form: "int", Required: true},
		&at.PointerToInt:     Field{Form: "pointerToInt", Required: true},
		&at.IntSlice:         Field{Form: "intSlice", Required: true},
		&at.Bool:             Field{Form: "bool", Required: true},
		&at.PointerToBool:    Field{Form: "pointerToBool", Required: true},
		&at.BoolSlice:        Field{Form: "boolSlice", Required: true},
		&at.String:           Field{Form: "string", Required: true},
		&at.PointerToString:  Field{Form: "pointerToString", Required: true},
		&at.StringSlice:      Field{Form: "stringSlice", Required: true},
		&at.Time:             Field{Form: "time", Required: true},
		&at.PointerToTime:    Field{Form: "pointerToTime", Required: true},
		&at.TimeSlice:        Field{Form: "timeSlice", Required: true},
	}
}

func TestValidate(t *testing.T) {

	Convey("Given a struct populated properly and as expected", t, func() {

		Convey("No errors should be produced", FailureContinues, func() {
			req, err := http.NewRequest("POST", "http://www.example.com", nil)
			if err != nil {
				t.Fatal(err)
			}
			model := AllTypes{}
			model.Uint8 = 255
			model.PointerToUint8 = &model.Uint8
			model.Uint8Slice = []uint8{model.Uint8}
			model.Uint16 = 65535
			model.PointerToUint16 = &model.Uint16
			model.Uint16Slice = []uint16{model.Uint16}
			model.Uint32 = 4294967295
			model.PointerToUint32 = &model.Uint32
			model.Uint32Slice = []uint32{model.Uint32}
			model.Uint64 = 18446744073709551615
			model.PointerToUint64 = &model.Uint64
			model.Uint64Slice = []uint64{model.Uint64}
			model.Int8 = 127
			model.PointerToInt8 = &model.Int8
			model.Int8Slice = []int8{model.Int8}
			model.Int16 = 32767
			model.PointerToInt16 = &model.Int16
			model.Int16Slice = []int16{model.Int16}
			model.Int32 = 2147483647
			model.PointerToInt32 = &model.Int32
			model.Int32Slice = []int32{model.Int32}
			model.Int64 = 9223372036854775807
			model.PointerToInt64 = &model.Int64
			model.Int64Slice = []int64{model.Int64}
			model.Float32 = 3.14
			model.PointerToFloat32 = &model.Float32
			model.Float32Slice = []float32{model.Float32}
			model.Float64 = 2.718
			model.PointerToFloat64 = &model.Float64
			model.Float64Slice = []float64{model.Float64}
			model.Uint = 4294967295
			model.PointerToUint = &model.Uint
			model.UintSlice = []uint{model.Uint}
			model.Int = 4294967295
			model.PointerToInt = &model.Int
			model.IntSlice = []int{model.Int}
			model.Bool = true
			model.PointerToBool = &model.Bool
			model.BoolSlice = []bool{model.Bool}
			model.String = "I'm a little teapot"
			model.PointerToString = &model.String
			model.StringSlice = []string{model.String}
			model.Time = time.Now()
			model.PointerToTime = &model.Time
			model.TimeSlice = []time.Time{model.Time}

			errs := Validate(req, &model)

			expectedErrs := make(map[string]bool)
			for _, v := range model.FieldMap() {
				f, ok := v.(Field)
				if !ok {
					t.Fatal("unexpected value in FieldMap")
				}

				expectedErrs[f.Form] = false
			}

			for _, bindErr := range errs {
				for _, name := range bindErr.FieldNames {
					if bindErr.Classification == RequiredError {
						expectedErrs[name] = true
					}
				}
			}

			for k, v := range expectedErrs {
				Convey(fmt.Sprintf("A Required error for %s should not be produced", k), func() {
					if v {
						Println(k, "has an unexpected Required error")
					}
					So(v, ShouldBeFalse)
				})
			}
		})

	})

	Convey("Given a populated struct missing one required field", t, func() {

		Convey("A Required error should be produced", nil)

	})

	Convey("Given a populated struct missing multiple required fields", t, func() {

		Convey("As many Required errors should be produced", FailureContinues, func() {
			req, err := http.NewRequest("POST", "http://www.example.com", nil)
			if err != nil {
				t.Fatal(err)
			}
			model := new(AllTypes)
			errs := Validate(req, model)

			expectedErrs := make(map[string]bool)
			for _, v := range model.FieldMap() {
				f, ok := v.(Field)
				if !ok {
					t.Fatal("unexpected value in FieldMap")
				}

				expectedErrs[f.Form] = false
			}

			for _, bindErr := range errs {
				for _, name := range bindErr.FieldNames {
					if bindErr.Classification == RequiredError {
						expectedErrs[name] = true
					}
				}
			}

			for k, v := range expectedErrs {
				Convey(fmt.Sprintf("A Required error for %s should be produced", k), func() {
					if !v {
						Println(k, "is missing the expected Required error")
					}
					So(v, ShouldBeTrue)
				})
			}
		})
	})

	Convey("Given a struct that is a Validator", t, func() {

		Convey("The user's Validate method should be invoked and its errors appended", nil)

	})

	Convey("Each case in the type switch should be tested", t, func() {

		Convey("uint8", nil)
		Convey("*uint8", nil)
		Convey("[]uint8", nil)

		Convey("uint16", nil)
		Convey("*uint16", nil)
		Convey("[]uint16", nil)

		Convey("uint32", nil)
		Convey("*uint32", nil)
		Convey("[]uint32", nil)

		Convey("uint64", nil)
		Convey("*uint64", nil)
		Convey("[]uint64", nil)

		Convey("int8", nil)
		Convey("*int8", nil)
		Convey("[]int8", nil)

		Convey("int16", nil)
		Convey("*int16", nil)
		Convey("[]int16", nil)

		Convey("int32", nil)
		Convey("*int32", nil)
		Convey("[]int32", nil)

		Convey("int64", nil)
		Convey("*int64", nil)
		Convey("[]int64", nil)

		Convey("float32", nil)
		Convey("*float32", nil)
		Convey("[]float32", nil)

		Convey("float64", nil)
		Convey("*float64", nil)
		Convey("[]float64", nil)

		Convey("uint", nil)
		Convey("*uint", nil)
		Convey("[]uint", nil)

		Convey("int", nil)
		Convey("*int", nil)
		Convey("[]int", nil)

		Convey("bool", nil)
		Convey("*bool", nil)
		Convey("[]bool", nil)

		Convey("string", nil)
		Convey("*string", nil)
		Convey("[]string", nil)

		Convey("time.Time", nil)
		Convey("*time.Time", nil)
		Convey("[]time.Time", nil)

	})

}
