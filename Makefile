# Copyright 2011 Utkan Güngördü. All rights reserved.
# Use of this source code is governed by a BSD-style
# See the LICENSE file of the official Go distrubtion.

include $(GOROOT)/src/Make.inc

TARG=extractor

CGOFILES=extractor.go const.go meta.go util.go
CGO_LDFLAGS=-lextractor

include $(GOROOT)/src/Make.pkg
