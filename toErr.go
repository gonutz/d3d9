package d3d9

//#include <d3d9.h>
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
	case C.D3DERR_WRONGTEXTUREFORMAT:
		return "wrong texture format"
	case C.D3DERR_UNSUPPORTEDCOLOROPERATION:
		return "unsupported color operation"
	case C.D3DERR_UNSUPPORTEDCOLORARG:
		return "unsupported color argument"
	case C.D3DERR_UNSUPPORTEDALPHAOPERATION:
		return "unsupported alpha operation"
	case C.D3DERR_UNSUPPORTEDALPHAARG:
		return "unsupported alpha argument"
	case C.D3DERR_TOOMANYOPERATIONS:
		return "too many operations"
	case C.D3DERR_CONFLICTINGTEXTUREFILTER:
		return "conflicting texture filter"
	case C.D3DERR_UNSUPPORTEDFACTORVALUE:
		return "unsupported factor value"
	case C.D3DERR_CONFLICTINGRENDERSTATE:
		return "conflicting render state"
	case C.D3DERR_UNSUPPORTEDTEXTUREFILTER:
		return "unsupported texture filter"
	case C.D3DERR_CONFLICTINGTEXTUREPALETTE:
		return "conflicting texture palette"
	case C.D3DERR_DRIVERINTERNALERROR:
		return "driver internal error"
	case C.D3DERR_NOTFOUND:
		return "not found"
	case C.D3DERR_MOREDATA:
		return "more data"
	case C.D3DERR_DEVICELOST:
		return "device lost"
	case C.D3DERR_DEVICENOTRESET:
		return "device not reset"
	case C.D3DERR_NOTAVAILABLE:
		return "not available"
	case C.D3DERR_OUTOFVIDEOMEMORY:
		return "out of video memory"
	case C.D3DERR_INVALIDDEVICE:
		return "invalid device"
	case C.D3DERR_INVALIDCALL:
		return "invalid call"
	case C.D3DERR_DRIVERINVALIDCALL:
		return "driver invalid call"
	case C.D3DERR_WASSTILLDRAWING:
		return "was still drawing"
	case C.D3DERR_DEVICEREMOVED:
		return "device removed"
	case C.D3DERR_DEVICEHUNG:
		return "device hung"
	case C.D3DERR_UNSUPPORTEDOVERLAY:
		return "unsupported overlay"
	case C.D3DERR_UNSUPPORTEDOVERLAYFORMAT:
		return "unsupported ovleray format"
	case C.D3DERR_CANNOTPROTECTCONTENT:
		return "cannot protect content"
	case C.D3DERR_UNSUPPORTEDCRYPTO:
		return "unsupported crypto"
	case C.D3DERR_PRESENT_STATISTICS_DISJOINT:
		return "present statistics disjoint"

	default:
		return "unknown error code " + strconv.Itoa(int(r))
	}
}
