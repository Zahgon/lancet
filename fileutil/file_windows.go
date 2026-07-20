//go:build windows

package fileutil

type tagVS_FIXEDFILEINFO struct {
	Signature        uint32
	StructVersion    uint32
	FileVersionMS    uint32
	FileVersionLS    uint32
	ProductVersionMS uint32
	ProductVersionLS uint32
	FileFlagsMask    uint32
	FileFlags        uint32
	FileOS           uint32
	FileType         uint32
	FileSubtype      uint32
	FileDateMS       uint32
	FileDateLS       uint32
}

func GetExeOrDllVersion(filePath string) (string, error) { _ = "STUB: not implemented"; return "", nil }
