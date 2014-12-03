package mem

/*
#define WINVER 0x0500
#include <stdlib.h>
#include <windows.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func physicalMem() (mem uint64, err error) {
	var statex C.struct__MEMORYSTATUSEX
	statex.dwLength = C.DWORD(unsafe.Sizeof(statex))
	succeeded := C.GlobalMemoryStatusEx(&statex)
	if succeeded == C.FALSE {
		lastError := C.GetLastError()
		return 0, fmt.Errorf("Unable to get memory. Error: %d", int(lastError))
	}
	return uint64(statex.ullTotalPhys), nil
}
