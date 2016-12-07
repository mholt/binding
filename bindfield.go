package binding

import (
	"errors"
	"mime/multipart"
	"strconv"
	"time"
)

type errorHandler func(err error)

func bindFormField(
	fieldPointer interface{},
	fieldName string,
	fieldSpec Field,
	strs []string,
	formData map[string][]string,
	formFile map[string][]*multipart.FileHeader,
	errorHandler errorHandler,
) {
	switch t := fieldPointer.(type) {
	case **multipart.FileHeader:
		if files, ok := formFile[fieldName]; ok {
			*t = files[0]
		}
	case *[]**multipart.FileHeader:
		if files, ok := formFile[fieldName]; ok {
			for _, file := range files {
				*t = append(*t, &file)
			}
		}
	default:
		bindField(fieldPointer, fieldName, fieldSpec, strs, errorHandler)
	}
}

func bindField(
	fieldPointer interface{},
	fieldName string,
	fieldSpec Field,
	strs []string,
	errorHandler errorHandler,
) {
	switch t := fieldPointer.(type) {
	case *uint8:
		val, err := strconv.ParseUint(strs[0], 10, 8)
		errorHandler(err)
		*t = uint8(val)
	case **uint8:
		parsed, err := strconv.ParseUint(strs[0], 10, 8)
		if err != nil {
			errorHandler(err)
			return
		}
		val := uint8(parsed)
		*t = &val
	case *[]uint8:
		for _, str := range strs {
			val, err := strconv.ParseUint(str, 10, 8)
			errorHandler(err)
			*t = append(*t, uint8(val))
		}
	case *uint16:
		val, err := strconv.ParseUint(strs[0], 10, 16)
		errorHandler(err)
		*t = uint16(val)
	case **uint16:
		parsed, err := strconv.ParseUint(strs[0], 10, 16)
		if err != nil {
			errorHandler(err)
			return
		}
		val := uint16(parsed)
		*t = &val
	case *[]uint16:
		for _, str := range strs {
			val, err := strconv.ParseUint(str, 10, 16)
			errorHandler(err)
			*t = append(*t, uint16(val))
		}
	case *uint32:
		val, err := strconv.ParseUint(strs[0], 10, 32)
		errorHandler(err)
		*t = uint32(val)
	case **uint32:
		parsed, err := strconv.ParseUint(strs[0], 10, 32)
		if err != nil {
			errorHandler(err)
			return
		}
		val := uint32(parsed)
		*t = &val
	case *[]uint32:
		for _, str := range strs {
			val, err := strconv.ParseUint(str, 10, 32)
			errorHandler(err)
			*t = append(*t, uint32(val))
		}
	case *uint64:
		val, err := strconv.ParseUint(strs[0], 10, 64)
		errorHandler(err)
		*t = val
	case **uint64:
		parsed, err := strconv.ParseUint(strs[0], 10, 64)
		if err != nil {
			errorHandler(err)
			return
		}
		val := uint64(parsed)
		*t = &val
	case *[]uint64:
		for _, str := range strs {
			val, err := strconv.ParseUint(str, 10, 64)
			errorHandler(err)
			*t = append(*t, uint64(val))
		}
	case *int8:
		val, err := strconv.ParseInt(strs[0], 10, 8)
		errorHandler(err)
		*t = int8(val)
	case **int8:
		parsed, err := strconv.ParseInt(strs[0], 10, 8)
		if err != nil {
			errorHandler(err)
			return
		}
		val := int8(parsed)
		*t = &val
	case *[]int8:
		for _, str := range strs {
			val, err := strconv.ParseInt(str, 10, 8)
			errorHandler(err)
			*t = append(*t, int8(val))
		}
	case *int16:
		val, err := strconv.ParseInt(strs[0], 10, 16)
		errorHandler(err)
		*t = int16(val)
	case **int16:
		parsed, err := strconv.ParseInt(strs[0], 10, 16)
		if err != nil {
			errorHandler(err)
			return
		}
		val := int16(parsed)
		*t = &val
	case *[]int16:
		for _, str := range strs {
			val, err := strconv.ParseInt(str, 10, 16)
			errorHandler(err)
			*t = append(*t, int16(val))
		}
	case *int32:
		val, err := strconv.ParseInt(strs[0], 10, 32)
		errorHandler(err)
		*t = int32(val)
	case **int32:
		parsed, err := strconv.ParseInt(strs[0], 10, 32)
		if err != nil {
			errorHandler(err)
			return
		}
		val := int32(parsed)
		*t = &val
	case *[]int32:
		for _, str := range strs {
			val, err := strconv.ParseInt(str, 10, 32)
			errorHandler(err)
			*t = append(*t, int32(val))
		}
	case *int64:
		val, err := strconv.ParseInt(strs[0], 10, 64)
		errorHandler(err)
		*t = val
	case **int64:
		parsed, err := strconv.ParseInt(strs[0], 10, 64)
		if err != nil {
			errorHandler(err)
			return
		}
		val := int64(parsed)
		*t = &val
	case *[]int64:
		for _, str := range strs {
			val, err := strconv.ParseInt(str, 10, 64)
			errorHandler(err)
			*t = append(*t, int64(val))
		}
	case *float32:
		val, err := strconv.ParseFloat(strs[0], 32)
		errorHandler(err)
		*t = float32(val)
	case **float32:
		parsed, err := strconv.ParseFloat(strs[0], 32)
		if err != nil {
			errorHandler(err)
			return
		}
		val := float32(parsed)
		*t = &val
	case *[]float32:
		for _, str := range strs {
			val, err := strconv.ParseFloat(str, 32)
			errorHandler(err)
			*t = append(*t, float32(val))
		}
	case *float64:
		val, err := strconv.ParseFloat(strs[0], 64)
		errorHandler(err)
		*t = val
	case **float64:
		parsed, err := strconv.ParseFloat(strs[0], 64)
		if err != nil {
			errorHandler(err)
			return
		}
		val := float64(parsed)
		*t = &val
	case *[]float64:
		for _, str := range strs {
			val, err := strconv.ParseFloat(str, 64)
			errorHandler(err)
			*t = append(*t, val)
		}
	case *uint:
		val, err := strconv.ParseUint(strs[0], 10, 0)
		errorHandler(err)
		*t = uint(val)
	case **uint:
		parsed, err := strconv.ParseUint(strs[0], 10, 0)
		if err != nil {
			errorHandler(err)
			return
		}
		val := uint(parsed)
		*t = &val
	case *[]uint:
		for _, str := range strs {
			val, err := strconv.ParseUint(str, 10, 0)
			errorHandler(err)
			*t = append(*t, uint(val))
		}
	case *int:
		val, err := strconv.ParseInt(strs[0], 10, 0)
		errorHandler(err)
		*t = int(val)
	case **int:
		parsed, err := strconv.ParseInt(strs[0], 10, 0)
		if err != nil {
			errorHandler(err)
			return
		}
		val := int(parsed)
		*t = &val
	case *[]int:
		for _, str := range strs {
			val, err := strconv.ParseInt(str, 10, 0)
			errorHandler(err)
			*t = append(*t, int(val))
		}
	case *bool:
		val, err := strconv.ParseBool(strs[0])
		errorHandler(err)
		*t = val
	case **bool:
		val, err := strconv.ParseBool(strs[0])
		if err != nil {
			errorHandler(err)
			return
		}
		*t = &val
	case *[]bool:
		for _, str := range strs {
			val, err := strconv.ParseBool(str)
			errorHandler(err)
			*t = append(*t, val)
		}
	case *string:
		*t = strs[0]
	case **string:
		s := strs[0]
		*t = &s
	case *[]string:
		*t = strs
	case *time.Time:
		timeFormat := TimeFormat
		if fieldSpec.TimeFormat != "" {
			timeFormat = fieldSpec.TimeFormat
		}
		val, err := time.Parse(timeFormat, strs[0])
		errorHandler(err)
		*t = val
	case **time.Time:
		timeFormat := TimeFormat
		if fieldSpec.TimeFormat != "" {
			timeFormat = fieldSpec.TimeFormat
		}
		val, err := time.Parse(timeFormat, strs[0])
		if err != nil {
			errorHandler(err)
			return
		}
		*t = &val
	case *[]time.Time:
		timeFormat := TimeFormat
		if fieldSpec.TimeFormat != "" {
			timeFormat = fieldSpec.TimeFormat
		}
		for _, str := range strs {
			val, err := time.Parse(timeFormat, str)
			errorHandler(err)
			*t = append(*t, val)
		}
	default:
		errorHandler(errors.New("Field type is unsupported by the application"))
	}
}
