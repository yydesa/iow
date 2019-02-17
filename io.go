package iow

// #include "iow/iowkit.c"
import "C"

import "unsafe"

type Handle C.IOWKIT_HANDLE

func Open() Handle {
	var r C.IOWKIT_HANDLE = C.IowKitOpenDevice()
	return Handle(r)
}

func Set(io Handle, x int) {
	var d C.DWORD = (C.DWORD)(x)

	C.IowKitWrite(
		C.IOWKIT_HANDLE(io),
		C.IOW_PIPE_IO_PINS,
		(*C.char)(unsafe.Pointer(&d)), 4)
}
