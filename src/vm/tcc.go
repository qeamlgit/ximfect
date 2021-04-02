package vm

// #cgo LDFLAGS: -L. -llibtcc
// #include "libtcc.h"
import "C"
import "unsafe"

const (
	OutputMemory     = 1
	OutputExe        = 2
	OutputDll        = 3
	OutputObj        = 4
	OutputPreprocess = 5
)

type TCCState struct {
	tcc *C.struct_TCCState
}

func NewTCC() TCCState {
	tcc := C.tcc_new()
	return TCCState{tcc}
}

func (s TCCState) SetOutputType(outputType int) bool {
	res := C.tcc_set_output_type(s.tcc, C.int(outputType))
	return res != -1
}

func (s TCCState) CompileString(src string) bool {
	res := C.tcc_compile_string(s.tcc, C.CString(src))
	return res != -1
}

func (s TCCState) Relocate() bool {
	res := C.tcc_relocate(s.tcc, unsafe.Pointer(uintptr(1)))
	return res != -1
}

func (s TCCState) GetSymbol(name string) unsafe.Pointer {
	return C.tcc_get_symbol(s.tcc, C.CString(name))
}

func (s TCCState) Delete() {
	C.tcc_delete(s.tcc)
}
