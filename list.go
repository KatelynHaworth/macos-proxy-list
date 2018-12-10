package proxylist

/*
#cgo LDFLAGS: -framework CoreFoundation -framework SystemConfiguration
#include <SystemConfiguration/SystemConfiguration.h>
 */
import "C"
import (
	"fmt"
	"unsafe"
)

// GetProxyList will query the macOS system for
// the current proxy configuration and return a
// Golang map of proxy settings and their values
func GetProxyList() (map[Key]interface{}, error) {
	store := C.SCDynamicStoreCreateWithOptions(0, goToCFString("app"), 0, nil, nil)
	if store == 0 {
		return nil, fmt.Errorf("unable to open dynamic store: app")
	}

	proxySettings := C.SCDynamicStoreCopyProxies(store)

	return dictionaryToMap(proxySettings), nil
}

// dictionaryToMap will convert a Core Foundation
// Dictionary returned by the SCDynamicStoreCopyProxies
// method into a Golang friendly map that can be accessed
// without the need for type conversions
func dictionaryToMap(ref C.CFDictionaryRef) map[Key]interface{} {
	size := C.CFDictionaryGetCount(ref)
	cfMap := make(map[Key]interface{}, 0)

	keys := make([]C.CFTypeRef, size)
	values := make([]C.CFTypeRef, size)

	C.CFDictionaryGetKeysAndValues(ref, (*unsafe.Pointer)(unsafe.Pointer(&keys[0])), (*unsafe.Pointer)(unsafe.Pointer(&values[0])))

	for i := 0; i < int(size); i++ {
		key := Key(cfStringToGoString((C.CFStringRef)(keys[i])))
		cfMap[key] = cfTypeRefToGo(values[i])
	}

	return cfMap
}