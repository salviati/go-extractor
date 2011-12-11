// Copyright 2011 Utkan Güngördü. All rights reserved.
// Use of this source code is governed by a BSD-style
// See the LICENSE file of the official Go distrubtion.

package extractor

import "C"
import "unsafe"
import "reflect"

func copy_cdata(cdata *C.char, csize uint) []byte {
	var data []byte

	if cdata != nil && csize > 0 {
		clen := int(csize)
		cdataGo := (*[1 << 30]byte)(unsafe.Pointer(cdata))[:clen]
		hdr := (*reflect.SliceHeader)(unsafe.Pointer(&cdataGo))
		hdr.Cap = clen
		hdr.Len = clen

		data = make([]byte, len(cdataGo))
		copy(data, cdataGo)
	}

	return data
}

