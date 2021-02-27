package nlopt

import "C"
import "unsafe"

//export nloptFunc
func nloptFunc(n uint, x *C.double, gradient *C.double, fData unsafe.Pointer) C.double {
	f := getFunc(uintptr(fData))
	cX := (*[(1 << 29) - 1]C.double)(unsafe.Pointer(x))[:n:n]
	goX := toGoArray(cX)

	var goGrad []float64
	var cGrad []C.double
	if gradient != nil {
		cGrad = (*[(1 << 29) - 1]C.double)(unsafe.Pointer(gradient))[:n:n]
		goGrad = toGoArray(cGrad)
	}
	v := (C.double)(f(goX, goGrad))
	if gradient != nil {
		for i := 0; i < len(cGrad); i++ {
			cGrad[i] = (C.double)((goGrad)[i])
		}
	}
	return v
}

//export nloptMfunc
func nloptMfunc(m uint, result *C.double, n uint, x *C.double, gradient *C.double, fData unsafe.Pointer) {
	f := getMfunc(uintptr(fData))
	cX := (*[(1 << 29) - 1]C.double)(unsafe.Pointer(x))[:n:n]
	goX := toGoArray(cX)

	var goGrad []float64
	var cGrad []C.double
	if gradient != nil {
		cGrad = (*[(1 << 29) - 1]C.double)(unsafe.Pointer(gradient))[:n:n]
		goGrad = toGoArray(cGrad)
	}

	var goResult []float64
	var cResult []C.double
	if result != nil {
		cResult = (*[(1 << 29) - 1]C.double)(unsafe.Pointer(result))[: m*n : m*n]
		goResult = toGoArray(cResult)
	}

	f(goResult, goX, goGrad)

	if gradient != nil {
		for i := 0; i < len(cGrad); i++ {
			cGrad[i] = (C.double)((goGrad)[i])
		}
	}

	if result != nil {
		for i := 0; i < len(cResult); i++ {
			cResult[i] = (C.double)((goResult)[i])
		}
	}

}
