package cyberobservableobjects

import (
	"time"

	"github.com/av-belyakov/methodstixobjects/datamodels/commonproperties"
	"github.com/av-belyakov/methodstixobjects/datamodels/commonpropertiesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/cyberobservableobjectsstix"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

// WrapperDomainName тип содержит дополнительный метод MarshalBSON
type WrapperDomainName struct {
	*cyberobservableobjectsstix.DomainNameCyberObservableObjectSTIX
}

type FinalyDomainNameObjects struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                           `bson:"value" required:"true"`
	ResolvesToRefs []stixhelpers.IdentifierTypeSTIX `bson:"resolves_to_refs"`
}

// WrapperURL тип содержит дополнительный метод MarshalBSON
type WrapperURL struct {
	*cyberobservableobjectsstix.URLCyberObservableObjectSTIX
}

type FinalyURLObjects struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value string `bson:"value" required:"true"`
}

// WrapperFile тип содержит дополнительный метод MarshalBSON
type WrapperFile struct {
	*cyberobservableobjectsstix.FileCyberObservableObjectSTIX
}

type FinalyFileObjects struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Size               uint64                           `bson:"size"`
	Name               string                           `bson:"name"`
	NameEnc            string                           `bson:"name_enc"`
	MagicNumberHex     string                           `bson:"magic_number_hex"`
	MimeType           string                           `bson:"mime_type"`
	Ctime              time.Time                        `bson:"ctime"`
	Mtime              time.Time                        `bson:"mtime"`
	Atime              time.Time                        `bson:"atime"`
	ParentDirectoryRef stixhelpers.IdentifierTypeSTIX   `bson:"parent_directory_ref"`
	Hashes             stixhelpers.HashesTypeSTIX       `bson:"hashes"`
	ContentRef         stixhelpers.IdentifierTypeSTIX   `bson:"content_ref"`
	ContainsRefs       []stixhelpers.IdentifierTypeSTIX `bson:"contains_refs"`
	Extensions         map[string]interface{}           `bson:"extensions"`
}

// WrapperEmailAddress тип содержит дополнительный метод MarshalBSON
type WrapperEmailAddress struct {
	*cyberobservableobjectsstix.EmailAddressCyberObservableObjectSTIX
}

type FinalyEmailAddressObjects struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value        string                         `bson:"value"`
	DisplayName  string                         `bson:"display_name"`
	BelongsToRef stixhelpers.IdentifierTypeSTIX `bson:"belongs_to_ref"`
}

// WrapperIPv4Address тип содержит дополнительный метод MarshalBSON
type WrapperIPv4Address struct {
	*cyberobservableobjectsstix.IPv4AddressCyberObservableObjectSTIX
}

type FinalyIPv4AddressObjects struct {
	commonproperties.CommonPropertiesObjectSTIX
	commonpropertiesstixco.OptionalCommonPropertiesCyberObservableObjectSTIX
	Value          string                           `bson:"value"`
	ResolvesToRefs []stixhelpers.IdentifierTypeSTIX `bson:"resolves_to_refs"`
	BelongsToRefs  []stixhelpers.IdentifierTypeSTIX `bson:"belongs_to_refs"`
}
