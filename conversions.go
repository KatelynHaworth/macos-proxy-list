package proxylist

/*
#cgo LDFLAGS: -framework CoreFoundation -framework SystemConfiguration
#include <SystemConfiguration/SystemConfiguration.h>
 */
import "C"
import (
	"unsafe"
)

// goToCFString will convert a Golang
// string into a CFStringRef for use in
// macOS calls
func goToCFString(s string) C.CFStringRef {
	return C.CFStringCreateWithBytes(0, *(**C.UInt8)(unsafe.Pointer(&s)), C.CFIndex(len(s)), C.kCFStringEncodingUTF8, 0)
}

// cfTypeRefToGo will attempt to convert a
// Core Foundation reference to it's Golang
// counterpart.
//
// This method is not exhaustive, it currently
// on supports strings, numbers (limited) and
// arrays, all other types will result in a value
// value
func cfTypeRefToGo(ref C.CFTypeRef) interface{} {
	switch C.CFGetTypeID(ref) {
	case C.CFStringGetTypeID():
		return cfStringToGoString((C.CFStringRef)(ref))

	case C.CFNumberGetTypeID():
		return cfNumberToGoNumber((C.CFNumberRef)(ref))

	case C.CFArrayGetTypeID():
		return cfArrayToGoArray((C.CFArrayRef)(ref))

	default:
		return nil
	}
}

// cfStringToGoString converts a OSX CFStringRef
// into a Golang string object
func cfStringToGoString(ref C.CFStringRef) string {
	/*
	Check if the string is already an plan
	C string, if so we can make the conversion
	using CGO
	*/
	pointer := C.CFStringGetCStringPtr(ref, C.kCFStringEncodingUTF8)
	if pointer != nil {
		return C.GoString(pointer)
	}

	/*
	CFString wasn't backed by a C string, need to
	manually get the bytes and size to convert it
	*/
	stringLength := C.CFStringGetLength(ref)
	if stringLength == 0 {
		return "" // String was already empty
	}

	bufLen := C.CFStringGetMaximumSizeForEncoding(stringLength, C.kCFStringEncodingUTF8)
	if bufLen == 0 {
		return "" // String isn't encoded using UTF-8, unable to convert it
	}

	var bufferFilled C.CFIndex
	buffer := make([]byte, bufLen)

	C.CFStringGetBytes(ref, C.CFRange{0, stringLength}, C.kCFStringEncodingUTF8, C.UInt8(0), C.false, (*C.UInt8)(&buffer[0]), bufLen, &bufferFilled)

	return string(buffer[:bufferFilled])
}

// cfNumberToGoNumber will take a CFNumberRef
// and match it's type to the appropriate
// conversion method to result in a Golang representation.
//
// This method is not exhaustive, there are a few CFNumber
// types that haven't been implemented due to the lack of
// need for the purpose of this package.
func cfNumberToGoNumber(ref C.CFNumberRef) interface{} {
	numberType := C.CFNumberGetType(ref)

	switch numberType {
	case C.kCFNumberSInt8Type:
		var number int8
		C.CFNumberGetValue(ref, C.kCFNumberSInt8Type, unsafe.Pointer(&number))
		return number

	case C.kCFNumberSInt16Type:
		var number int16
		C.CFNumberGetValue(ref, C.kCFNumberSInt16Type, unsafe.Pointer(&number))
		return number

	case C.kCFNumberSInt32Type:
		var number int32
		C.CFNumberGetValue(ref, C.kCFNumberSInt32Type, unsafe.Pointer(&number))
		return number

	case C.kCFNumberSInt64Type:
		var number int64
		C.CFNumberGetValue(ref, C.kCFNumberSInt64Type, unsafe.Pointer(&number))
		return number

	case C.kCFNumberFloat32Type:
		var number float32
		C.CFNumberGetValue(ref, C.kCFNumberFloat32Type, unsafe.Pointer(&number))
		return number

	case C.kCFNumberFloat64Type:
		var number float64
		C.CFNumberGetValue(ref, C.kCFNumberFloat64Type, unsafe.Pointer(&number))
		return number

	case C.kCFNumberIntType:
		var number int
		C.CFNumberGetValue(ref, C.kCFNumberIntType, unsafe.Pointer(&number))
		return number

	default:
		return nil
	}
}

// cfArrayToGoArray will convert a CFArray
// to a Golang slice of interfaces with the values
// of the CFArray being converted to their Golang
// version
func cfArrayToGoArray(ref C.CFArrayRef) []interface{} {
	size := C.CFArrayGetCount(ref)
	elements := make([]C.CFTypeRef, size)

	C.CFArrayGetValues(ref, C.CFRange{0, size}, (*unsafe.Pointer)(unsafe.Pointer(&elements[0])))

	array := make([]interface{}, size)
	for i, elm := range elements {
		array[i] = cfTypeRefToGo(elm)
	}

	return array
}