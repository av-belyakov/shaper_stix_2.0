package createrstixobject

import (
	"fmt"

	methodstixobjects "github.com/av-belyakov/methodstixobjects/cmd"
	"github.com/av-belyakov/methodstixobjects/datamodels/cyberobservableobjectsstix"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/google/uuid"
)

// CreateDomainNameCyberObservableObjectSTIX формирует объект 'domain-name'
func CreateDomainNameCyberObservableObjectSTIX(observable datamodels.ObservableMessage) *cyberobservableobjectsstix.DomainNameCyberObservableObjectSTIX {
	domainName := methodstixobjects.NewDomainNameCyberObservableObjectSTIX()
	domainName.SetValueID(fmt.Sprintf("domain-name-%s", uuid.NewString()))
	domainName.SetValueValue(observable.Data)

	return domainName
}

// CreateURLCyberObservableObjectSTIX формирует объект 'url'
func CreateURLCyberObservableObjectSTIX(observable datamodels.ObservableMessage) *cyberobservableobjectsstix.URLCyberObservableObjectSTIX {
	url := methodstixobjects.NewURLCyberObservableObjectSTIX()
	url.SetValueID(fmt.Sprintf("url-%s", uuid.NewString()))
	url.SetValueValue(observable.Data)

	return url
}

// CreateFileCyberObservableObjectSTIX формирует объект 'file'
func CreateFileCyberObservableObjectSTIX(observable datamodels.ObservableMessage) *cyberobservableobjectsstix.FileCyberObservableObjectSTIX {
	file := methodstixobjects.NewFileCyberObservableObjectSTIX()
	file.SetValueID(fmt.Sprintf("file-%s", uuid.NewString()))
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
func CreateEmailAddressCyberObservableObjectSTIX(observable datamodels.ObservableMessage) *cyberobservableobjectsstix.EmailAddressCyberObservableObjectSTIX {
	email := methodstixobjects.NewEmailAddressCyberObservableObjectSTIX()
	email.SetValueID(fmt.Sprintf("email-addr-%s", uuid.NewString()))
	email.SetValueValue(observable.Data)

	return email
}

// CreateIPv4AddressCyberObservableObjectSTIX формирует объект 'ipv4-addr'
func CreateIPv4AddressCyberObservableObjectSTIX(observable datamodels.ObservableMessage) *cyberobservableobjectsstix.IPv4AddressCyberObservableObjectSTIX {
	ipv4 := methodstixobjects.NewIPv4AddressCyberObservableObjectSTIX()
	ipv4.SetValueID(fmt.Sprintf("ipv4-addr-%s", uuid.NewString()))
	ipv4.SetValueValue(observable.Data)

	return ipv4
}
