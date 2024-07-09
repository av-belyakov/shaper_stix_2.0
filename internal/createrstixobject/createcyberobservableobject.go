package createrstixobject

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix/cyberobservableobjects"
)

// CreateDomainNameCyberObservableObjectSTIX формирует объект 'domain-name'
func CreateDomainNameCyberObservableObjectSTIX(observable datamodels.ObservableMessage) *cyberobservableobjects.WrapperDomainName {
	domainName := cyberobservableobjects.NewWrapperDomainNameCyberObservableObjectSTIX()
	domainName.SetValueID(fmt.Sprintf("domain-name--%s", uuid.NewString()))
	domainName.SetValueSpecVersion("2.1")
	domainName.SetValueValue(observable.Data)

	return domainName
}

// CreateURLCyberObservableObjectSTIX формирует объект 'url'
func CreateURLCyberObservableObjectSTIX(observable datamodels.ObservableMessage) *cyberobservableobjects.WrapperURL {
	url := cyberobservableobjects.NewWrapperURLCyberObservableObjectSTIX()
	url.SetValueID(fmt.Sprintf("url--%s", uuid.NewString()))
	url.SetValueSpecVersion("2.1")
	url.SetValueValue(observable.Data)

	return url
}

// CreateFileCyberObservableObjectSTIX формирует объект 'file'
func CreateFileCyberObservableObjectSTIX(observable datamodels.ObservableMessage) *cyberobservableobjects.WrapperFile {
	file := cyberobservableobjects.NewWrapperFileCyberObservableObjectSTIX()
	file.SetValueID(fmt.Sprintf("file--%s", uuid.NewString()))
	file.SetValueSpecVersion("2.1")
	file.SetValueName(observable.Data)

	if observable.Attachment.Name != "" {
		file.SetValueName(observable.Attachment.Name)
	}
	if observable.Attachment.Size > 0 {
		file.SetValueSize(observable.Attachment.Size)
	}
	if len(observable.Attachment.Hashes) > 0 {
		hashes := stixhelpers.HashesTypeSTIX{}

		for k, v := range observable.Attachment.Hashes {
			hashes[fmt.Sprintf("hash_%d", k)] = v
		}

		file.SetValueHashes(hashes)
	}

	return file
}

// CreateEmailAddressCyberObservableObjectSTIX формирует объект 'email-addr'
func CreateEmailAddressCyberObservableObjectSTIX(observable datamodels.ObservableMessage) *cyberobservableobjects.WrapperEmailAddress {
	email := cyberobservableobjects.NewWrapperEmailAddressCyberObservableObjectSTIX()
	email.SetValueID(fmt.Sprintf("email-addr--%s", uuid.NewString()))
	email.SetValueSpecVersion("2.1")
	email.SetValueValue(observable.Data)

	return email
}

// CreateIPv4AddressCyberObservableObjectSTIX формирует объект 'ipv4-addr'
func CreateIPv4AddressCyberObservableObjectSTIX(observable datamodels.ObservableMessage) *cyberobservableobjects.WrapperIPv4Address {
	ipv4 := cyberobservableobjects.NewWrapperIPv4AddressCyberObservableObjectSTIX()
	ipv4.SetValueID(fmt.Sprintf("ipv4-addr--%s", uuid.NewString()))
	ipv4.SetValueSpecVersion("2.1")

	ipv4.SetValueValue(observable.Data)
	if strings.Contains(observable.Data, ":") {
		tmp := strings.Split(observable.Data, ":")

		if len(tmp) == 2 {
			ipv4.SetValueValue(tmp[1])
		}
	}

	return ipv4
}
