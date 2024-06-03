package mongodbapi

import (
	"shaper_stix/datamodels"
)

// ArtifactCyberObservableObjectSTIX объект "Artifact", по терминалогии STIX, позволяет
// захватывать массив байтов (8 бит) в виде строки в кодировке base64 или связывать его
// с полезной нагрузкой, подобной файлу. Обязательно должен быть заполнено одно из полей
// PayloadBin или URL
func (w *Wrappers) AddArtifactCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Artifact
	//****************************************************************
	//fmt.Println("Additing Artifact Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "artifact",
	}
}

// AutonomousSystemCyberObservableObjectSTIX объект "Autonomous System", по терминалогии STIX,
// содержит параметры Автономной системы
func (w *Wrappers) AddAutonomousSystemCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа AutonomousSystem
	//****************************************************************
	//fmt.Println("Additing AutonomousSystem Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "autonomous-system",
	}
}

// DirectoryCyberObservableObjectSTIX объект "Directory", по терминалогии STIX, содержит
// свойства, общие для каталога файловой системы
func (w *Wrappers) AddDirectoryCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Directory
	//****************************************************************
	//fmt.Println("Additing Directory Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "directory",
	}
}

// DomainNameCyberObservableObjectSTIX объект "Domain Name", по терминалогии STIX,
// содержит сетевое доменное имя
func (w *Wrappers) AddDomainNameCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа DomainName
	//****************************************************************
	//fmt.Println("Additing DomainName Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "domain-name",
	}
}

// EmailAddressCyberObservableObjectSTIX объект "Email Address", по терминалогии STIX,
// содержит представление единственного email адреса
func (w *Wrappers) AddEmailAddressCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа EmailAddress
	//****************************************************************
	//fmt.Println("Additing EmailAddress Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "email-addr",
	}
}

// EmailMessageCyberObservableObjectSTIX объект "Email Message", по терминалогии STIX,
// содержит экземпляр email сообщения
func (w *Wrappers) AddEmailMessageCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа EmailMessage
	//****************************************************************
	//fmt.Println("Additing  Cyber EmailMessage Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "email-message",
	}
}

// FileCyberObservableObjectSTIX объект "File Object", по терминалогии STIX,
// последекодирования из JSON (основной, рабочий объект)
func (w *Wrappers) AddFileCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа File
	//****************************************************************
	//fmt.Println("Additing File Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "file",
	}
}

// IPv4AddressCyberObservableObjectSTIX объект "IPv4 Address Object", по
// терминалогии STIX, содержит один или более IPv4 адресов, выраженных с
// помощью нотации CIDR.
func (w *Wrappers) AddIPv4AddressCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа IPv4Address
	//****************************************************************
	//fmt.Println("Additing IPv4Address Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "ipv4-addr",
	}
}

// IPv6AddressCyberObservableObjectSTIX объект "IPv6 Address Object", по
// терминалогии STIX, содержит один или более IPv6 адресов, выраженных с
// помощью нотации CIDR.
func (w *Wrappers) AddIPv6AddressCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа IPv6Address
	//****************************************************************
	//fmt.Println("Additing IPv6Address Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "ipv6-addr",
	}
}

// MACAddressCyberObservableObjectSTIX объект "MAC Address Object", по
// терминалогии STIX, содержит объект MAC-адрес, представляющий собой
// один адрес управления доступом к среде (MAC).
func (w *Wrappers) AddMACAddressCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа MACAddress
	//****************************************************************
	//fmt.Println("Additing MACAddress Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "mac-addr",
	}
}

// MutexCyberObservableObjectSTIX объект "Mutex Object", по терминалогии STIX,
// содержит свойства объекта взаимного исключения (mutex).
func (w *Wrappers) AddMutexCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Mutex
	//****************************************************************
	//fmt.Println("Additing Mutex Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "mutex",
	}
}

// NetworkTrafficCyberObservableObjectSTIX объект "Network Traffic Object", по
// терминалогии STIX, содержит объект. Сетевого трафика представляющий собой
// произвольный сетевой трафик, который исходит из источника и адресуется адресату.
func (w *Wrappers) AddNetworkTrafficCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа NetworkTraffic
	//****************************************************************
	//fmt.Println("Additing NetworkTraffic Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "network-traffic",
	}
}

// ProcessCyberObservableObjectSTIX объект "Process Object", по терминологии STIX,
// содержит общие свойства экземпляра компьютерной программы, выполняемой в
// операционной системе. Объект процесса ДОЛЖЕН содержать хотя бы одно
// свойство (отличное от типа) этого объекта (или одного из его расширений).
func (w *Wrappers) AddProcessCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Process
	//****************************************************************
	//fmt.Println("Additing Process Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "process",
	}
}

// SoftwareCyberObservableObjectSTIX объект "Software Object", по терминологии STIX,
// содержит свойства, связанные с программным обеспечением, включая программные продукты.
func (w *Wrappers) AddSoftwareCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа Software
	//****************************************************************
	//fmt.Println("Additing Software Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "software",
	}
}

// URLCyberObservableObjectSTIX объект "URL Object", по терминологии STIX, содержит
// унифицированный указатель информационного ресурса (URL).
func (w *Wrappers) AddURLCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа URL
	//****************************************************************
	//fmt.Println("Additing URL Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "url",
	}
}

// UserAccountCyberObservableObjectSTIX объект "User Account Object", по терминалогии
// STIX, содержит экземпляр любого типа учетной записи пользователя, включая, учетные
// записи операционной системы, устройства, службы обмена сообщениями и платформы
// социальных сетей и других прочих учетных записей
func (w *Wrappers) AddUserAccountCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа UserAccount
	//****************************************************************
	//fmt.Println("Additing UserAccount Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "user-account",
	}
}

// WindowsRegistryKeyCyberObservableObjectSTIX объект "Windows Registry Key Object", по
// терминалогии STIX. Содержит описание значений полей раздела реестра Windows.
func (w *Wrappers) AddWindowsRegistryKeyCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа WindowsRegistryKey
	//****************************************************************
	//fmt.Println("Additing WindowsRegistryKey Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "windows-registry-key",
	}
}

// X509CertificateCyberObservableObjectSTIX объект "X.509 Certificate Object", по
// терминологии STIX, представлет свойства сертификата X.509, определенные в
// рекомендациях ITU X.509 [X.509]. X.509  Certificate объект должен содержать
// по крайней мере одно cвойство специфичное для этого объекта (помимо type).
func (w *Wrappers) AddX509CertificateCO(data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {

	//****************************************************************
	//**** Здесь обработчик для добавления объекта типа X509Certificate
	//****************************************************************
	//fmt.Println("Additing X509Certificate Cyber Observable Object STIX")

	//************************************
	//*** выполняется только для тестов
	//************************************
	counting <- datamodels.DataCounterSettings{
		DataType: "routing_test",
		DataMsg:  "x509-certificate",
	}
}
