package cyberobservableobjects

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
)

// NewWrapperFileCyberObservableObjectSTIX формирует новый объект 'file'
func NewWrapperFileCyberObservableObjectSTIX() *WrapperFile {
	return &WrapperFile{
		methodstixobjects.NewFileCyberObservableObjectSTIX(),
	}
}

func (e *WrapperFile) Get() *WrapperFile {
	return e
}

func (e *WrapperFile) ToStringBeautiful(num int) string {
	return e.FileCyberObservableObjectSTIX.ToStringBeautiful(num)
}

func (e *WrapperFile) MarshalBSON() ([]byte, error) {
	ffo := FinalyFileObjects{
		CommonPropertiesObjectSTIX:                        e.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: e.OptionalCommonPropertiesCyberObservableObjectSTIX,
		Size:               e.Size,
		Name:               e.Name,
		NameEnc:            e.NameEnc,
		MagicNumberHex:     e.MagicNumberHex,
		MimeType:           e.MimeType,
		ParentDirectoryRef: e.ParentDirectoryRef,
		Hashes:             e.Hashes,
		ContentRef:         e.ContentRef,
		ContainsRefs:       e.ContainsRefs,
		Extensions:         e.Extensions,
	}

	if ctime, err := time.Parse(time.RFC3339, e.Ctime); err == nil {
		ffo.Ctime = ctime
	}

	if mtime, err := time.Parse(time.RFC3339, e.Mtime); err == nil {
		ffo.Mtime = mtime
	}

	if atime, err := time.Parse(time.RFC3339, e.Atime); err == nil {
		ffo.Atime = atime
	}

	return bson.Marshal(ffo)
}
