package confighandler

type ConfigApp struct {
	CommonAppConfig
	AppConfigNATS
	AppConfigMongoDB
}

type CommonAppConfig struct {
	LogList []LogSet
	Zabbix  ZabbixOptions
}

type Logs struct {
	Logging []LogSet
}

type LogSet struct {
	WritingStdout bool   `yaml:"writingStdout"`
	WritingFile   bool   `yaml:"writingFile"`
	MaxFileSize   int    `yaml:"maxFileSize"`
	MsgTypeName   string `yaml:"msgTypeName"`
	PathDirectory string `yaml:"pathDirectory"`
}

type ZabbixSet struct {
	Zabbix ZabbixOptions
}

type ZabbixOptions struct {
	NetworkPort int         `yaml:"networkPort"`
	NetworkHost string      `yaml:"networkHost"`
	ZabbixHost  string      `yaml:"zabbixHost"`
	EventTypes  []EventType `yaml:"eventType"`
}

type EventType struct {
	IsTransmit bool      `yaml:"isTransmit"`
	EventType  string    `yaml:"eventType"`
	ZabbixKey  string    `yaml:"zabbixKey"`
	Handshake  Handshake `yaml:"handshake"`
}

type Handshake struct {
	TimeInterval int    `yaml:"timeInterval"`
	Message      string `yaml:"message"`
}

type AppConfigNATS struct {
	Port         int    `yaml:"port"`
	Host         string `yaml:"host"`
	SubjectCase  string `yaml:"subject_case"`
	SubjectAlert string `yaml:"subject_alert"`
}

type AppConfigMongoDB struct {
	Port   int    `yaml:"port"`
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	NameDB string `yaml:"namedb"`
}
