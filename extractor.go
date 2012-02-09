// Copyright 2011 Utkan Güngördü. All rights reserved.
// Use of this source code is governed by a BSD-style
// See the LICENSE file of the official Go distrubtion.

package extractor

// #cgo pkg-config: libextractor
//
//#include <extractor.h>
//#include <stdlib.h>
//extern void addMeta(void*, char*, unsigned long, unsigned long, char*, char*, unsigned int);
/* int
_extractor_callback(void *cls,
			 const char *plugin_name,
			 enum EXTRACTOR_MetaType type,
			 enum EXTRACTOR_MetaFormat format,
			 const char *data_mime_type,
			 const char *data,
			 size_t data_len)
{
	addMeta(cls, (char*)plugin_name, (unsigned long)type, (unsigned long)format, (char*)data_mime_type, (char*)data, data_len);
	return 0;
}
void _extract_wrap(struct EXTRACTOR_PluginList *plugins,
		  const char *filename,
		  const void *data,
		  size_t size,
		  void *proc_cls) {
	return EXTRACTOR_extract(plugins, filename, data, size, _extractor_callback, proc_cls);
}
*/
import "C"
import "unsafe"

// TODO(utkan): Add a String funciton for Meta struct.
// TODO(utkan): Documentation.


type Extractor struct {
	clist *[0]byte
}

func (ex *Extractor) AddLib(library, options string, flags uint) {
	clibrary := C.CString(library)
	defer C.free(unsafe.Pointer(clibrary))
	coptions := C.CString(options)
	defer C.free(unsafe.Pointer(coptions))
	cflags := uint32(flags)

	ex.clist = C.EXTRACTOR_plugin_add(ex.clist, clibrary, coptions, cflags)
}

func (ex *Extractor) AddConfig(config string, flags uint) {
	cconfig := C.CString(config)
	defer C.free(unsafe.Pointer(cconfig))
	cflags := uint32(flags)

	ex.clist = C.EXTRACTOR_plugin_add_config(ex.clist, cconfig, cflags)
}

func (ex *Extractor) RemoveLib(library string) {
	clibrary := C.CString(library)
	defer C.free(unsafe.Pointer(clibrary))

	ex.clist = C.EXTRACTOR_plugin_remove(ex.clist, clibrary)
}

func (ex *Extractor) RemoveAllLibs() {
	C.EXTRACTOR_plugin_remove_all(ex.clist)
}

//export addMeta
func addMeta(cls unsafe.Pointer, cpluginname *C.char, ctype uint32, cformat uint32, cdatamimetype, cdata *C.char, csize uint) {
	metas := (*[]Meta)(cls)

	data := copy_cdata(cdata, csize)

	pluginname := C.GoString(cpluginname)
	datamimetype := C.GoString(cdatamimetype)

	meta := &Meta{
		PluginName: pluginname,
		Type: uint(ctype),
		Format: uint(cformat),
		DataMimeType: datamimetype,
		Data: data,
	}

	*metas = append(*metas, *meta)
}

func (ex *Extractor) File(filename string) []Meta {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	metas := make([]Meta, 0)
	C._extract_wrap(ex.clist, cfilename, nil, 0, unsafe.Pointer(&metas))
	return metas
}

func (ex *Extractor) Memory(data []byte) []Meta {
	metas := make([]Meta, 0)
	C._extract_wrap(ex.clist, nil, unsafe.Pointer(&data[0]), C.size_t(len(data)), unsafe.Pointer(&metas))
	return metas
}

func New(options uint) *Extractor {
	ex := &Extractor{}
	ex.clist = C.EXTRACTOR_plugin_add_defaults(uint32(options))
	return ex
}