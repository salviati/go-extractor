// Copyright 2011 Utkan Güngördü. All rights reserved.
// Use of this source code is governed by a BSD-style
// See the LICENSE file of the official Go distrubtion.

package extractor

//#include <extractor.h>
import "C"

type Meta struct {
	PluginName string
	Type uint
	Format uint
	DataMimeType string
	Data []byte
}

// Returns "" if the type is not known, otherwise
// an English (locale: C) string describing the type;
// translate using 'dgettext ("libextractor", rval)'
func MetaTypeToString(m uint) string {
	cstring := C.EXTRACTOR_metatype_to_string(uint32(m));
	return C.GoString(cstring)
}

func MetaTypeMax() uint {
	return uint(C.EXTRACTOR_metatype_get_max())
}

