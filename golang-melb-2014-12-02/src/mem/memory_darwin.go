package mem

/*
#include <stdlib.h>
#include <sys/sysctl.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

func physicalMem() (mem uint64, err error) {
	cstr := C.CString("hw.memsize")
	defer C.free(unsafe.Pointer(cstr))
	var oldlen C.size_t
	if _, err = C.sysctlbyname(cstr, nil, &oldlen, nil, 0); err != nil {
		return 0, err
	}
	if oldlen != C.size_t(8) {
		return 0, errors.New("Not a 64-bit int!")
	}
	buf := make([]byte, 8)
	_, err = C.sysctlbyname(cstr, unsafe.Pointer(&buf[0]), &oldlen, nil, 0)
	if err != nil {
		return 0, err
	}
	return uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 |
		uint64(buf[3])<<24 | uint64(buf[4])<<32 | uint64(buf[5])<<40 |
		uint64(buf[6])<<48 | uint64(buf[7])<<56, nil
}
