package d3d9

//#include <windows.h>
import "C"
import "strconv"

func toErr(result C.HRESULT) error {
	if result >= 0 {
		return nil
	}
	return hResultError(result)
}

type hResultError C.HRESULT

func (r hResultError) Error() string {
	switch r {
	case ERR_WRONGTEXTUREFORMAT:
		return "wrong texture format"
	case ERR_UNSUPPORTEDCOLOROPERATION:
		return "unsupported color operation"
	case ERR_UNSUPPORTEDCOLORARG:
		return "unsupported color argument"
	case ERR_UNSUPPORTEDALPHAOPERATION:
		return "unsupported alpha operation"
	case ERR_UNSUPPORTEDALPHAARG:
		return "unsupported alpha argument"
	case ERR_TOOMANYOPERATIONS:
		return "too many operations"
	case ERR_CONFLICTINGTEXTUREFILTER:
		return "conflicting texture filter"
	case ERR_UNSUPPORTEDFACTORVALUE:
		return "unsupported factor value"
	case ERR_CONFLICTINGRENDERSTATE:
		return "conflicting render state"
	case ERR_UNSUPPORTEDTEXTUREFILTER:
		return "unsupported texture filter"
	case ERR_CONFLICTINGTEXTUREPALETTE:
		return "conflicting texture palette"
	case ERR_DRIVERINTERNALERROR:
		return "driver internal error"
	case ERR_NOTFOUND:
		return "not found"
	case ERR_MOREDATA:
		return "more data"
	case ERR_DEVICELOST:
		return "device lost"
	case ERR_DEVICENOTRESET:
		return "device not reset"
	case ERR_NOTAVAILABLE:
		return "not available"
	case ERR_OUTOFVIDEOMEMORY:
		return "out of video memory"
	case ERR_INVALIDDEVICE:
		return "invalid device"
	case ERR_INVALIDCALL:
		return "invalid call"
	case ERR_DRIVERINVALIDCALL:
		return "driver invalid call"
	case ERR_WASSTILLDRAWING:
		return "was still drawing"
	default:
		return "unknown error code " + strconv.Itoa(int(r))
	}
}
