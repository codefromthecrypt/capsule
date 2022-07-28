// Package hf host functions
package hostfunctions

import (
	"errors"
	"strconv"
	_ "unsafe"
)

//export hostReadFile
//go:linkname hostReadFile
func hostReadFile(filePathPtrPos uint32, size uint32, retBuffPtrPos **byte, retBuffSize *int)

// ReadFile :
// Call host function: hostReadFile
// Pass a string as parameter
// Get a string from the host

func ReadFile(filePath string) (string, error) {

	filePathPtrPos, size := GetStringPtrPositionAndSize(filePath)

	var buffPtr *byte
	var buffSize int

	hostReadFile(filePathPtrPos, size, &buffPtr, &buffSize)

	var resultStr = ""
	var err error
	valueStr := GetStringResult(buffPtr, buffSize)

	// check the return value
	if IsErrorString(valueStr) {
		errorMessage, errorCode := GetErrorStringInfo(valueStr)
		if errorCode == 0 {
			err = errors.New(errorMessage)
		} else {
			err = errors.New(errorMessage + " (" + strconv.Itoa(errorCode) + ")")
		}

	} else {
		resultStr = valueStr
	}
	return resultStr, err

}