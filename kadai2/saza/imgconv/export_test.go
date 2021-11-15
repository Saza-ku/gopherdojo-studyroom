package imgconv

type ExportFileType fileType

func (t ExportFileType) String() string {
	ft := fileType(t)
	return ft.String()
}

const (
	ExportJpegType = jpegType
	ExportPngType = pngType
	ExportGifType = gifType
	ExportOthers =  others
)

func ExportChangeExt(path string, destExt ExportFileType) string {
	return changeExt(path, fileType(destExt))
}
