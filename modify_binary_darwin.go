package gomonkey

import (
	"fmt"
	"syscall"
)

func modifyBinary(target uintptr, bytes []byte) {
	kkk := 88888
	fmt.Println(kkk)
	fmt.Println("VVVVVVVVV")
	function := entryAddress(target, len(bytes))

	page := entryAddress(pageStart(target), syscall.Getpagesize())
	err := syscall.Mprotect(page, syscall.PROT_WRITE)
	if err != nil {
		panic(err)
	}
	copy(function, bytes)

	err = syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_EXEC)
	if err != nil {
		panic(err)
	}
}
