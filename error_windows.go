package d3d9

import "strconv"

// Error is returned by all Direct3D9 functions. It encapsulates the error code
// returned by Direct3D. If a function succeeds it will return nil as the Error
// and if it fails you can retrieve the error code using the Code() function.
// You can check the result against the predefined error codes (like
// ERR_DEVICELOST, E_OUTOFMEMORY etc).
type Error interface {
	error
	// Code returns the Direct3D error code for a function. Call this function
	// only if the Error is not nil, if the error code is D3D_OK or any other
	// code that signifies success, a function will return nil as the Error
	// instead of a non-nil error with that code in it. This way, functions
	// behave in a standard Go way, returning nil as the error in case of
	// success and only returning non-nil errors if something went wrong.
	Code() int32
}

func toErr(result uintptr) Error {
	res := hResultError(result) // cast to signed int
	if res >= 0 {
		return nil
	}
	return res
}

type hResultError int32

func (r hResultError) Code() int32 { return int32(r) }

func (r hResultError) Error() string {
	switch r {
	case ERR_CONFLICTINGRENDERSTATE:
		return "conflicting render state"
	case ERR_CONFLICTINGTEXTUREFILTER:
		return "conflicting texture filter"
	case ERR_CONFLICTINGTEXTUREPALETTE:
		return "conflicting texture palette"
	case ERR_DEVICEHUNG:
		return "device hung"
	case ERR_DEVICELOST:
		return "device lost"
	case ERR_DEVICENOTRESET:
		return "device not reset"
	case ERR_DEVICEREMOVED:
		return "device removed"
	case ERR_DRIVERINTERNALERROR:
		return "driver internal error"
	case ERR_DRIVERINVALIDCALL:
		return "driver invalid call"
	case ERR_INVALIDCALL:
		return "invalid call"
	case ERR_INVALIDDEVICE:
		return "invalid device"
	case ERR_MOREDATA:
		return "more data"
	case ERR_NOTAVAILABLE:
		return "not available"
	case ERR_NOTFOUND:
		return "not found"
	case ERR_OUTOFVIDEOMEMORY:
		return "out of video memory"
	case ERR_TOOMANYOPERATIONS:
		return "too many operations"
	case ERR_UNSUPPORTEDALPHAARG:
		return "unsupported alpha argument"
	case ERR_UNSUPPORTEDALPHAOPERATION:
		return "unsupported alpha operation"
	case ERR_UNSUPPORTEDCOLORARG:
		return "unsupported color argument"
	case ERR_UNSUPPORTEDCOLOROPERATION:
		return "unsupported color operation"
	case ERR_UNSUPPORTEDFACTORVALUE:
		return "unsupported factor value"
	case ERR_UNSUPPORTEDTEXTUREFILTER:
		return "unsupported texture filter"
	case ERR_WASSTILLDRAWING:
		return "was still drawing"
	case ERR_WRONGTEXTUREFORMAT:
		return "wrong texture format"
	case ERR_UNSUPPORTEDOVERLAY:
		return "unsupported overlay"
	case ERR_UNSUPPORTEDOVERLAYFORMAT:
		return "unsupported overlay format"
	case ERR_CANNOTPROTECTCONTENT:
		return "cannot protect content"
	case ERR_UNSUPPORTEDCRYPTO:
		return "unsupported crypto"

	case E_FAIL:
		return "fail"
	case E_INVALIDARG:
		return "invalid argument"
	case E_NOINTERFACE:
		return "no interface"
	case E_NOTIMPL:
		return "not implemented"
	case E_OUTOFMEMORY:
		return "out of memory"

	case S_NOT_RESIDENT:
		return "not resident"
	case S_RESIDENT_IN_SHARED_MEMORY:
		return "resident in shared memory"

	default:
		return "unknown error code " + strconv.Itoa(int(r))
	}
}
