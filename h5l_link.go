package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// Create a soft link under linkLocId with name of LinkName pointing to targetPath
func lCreateSoft(targetPath string, linkLocId C.hid_t, LinkName string) error {

	c_target_path := C.CString(targetPath)
	defer C.free(unsafe.Pointer(c_target_path))

	c_linkname := C.CString(LinkName)
	defer C.free(unsafe.Pointer(c_linkname))

	lcpl_id := C.hid_t(C.H5P_DEFAULT)
	lapl_id := C.hid_t(C.H5P_DEFAULT)

	rc := C.H5Lcreate_soft(c_target_path, linkLocId, c_linkname, lcpl_id, lapl_id)
	if rc < 0 {
		return fmt.Errorf("Error creating soft link %s -> %s", LinkName, targetPath)
	}

	return nil
}
