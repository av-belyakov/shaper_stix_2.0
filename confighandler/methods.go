package confighandler

func (conf *ConfigApp) GetCommonApp() *CommonAppConfig {
	return &conf.CommonAppConfig
}

func (conf *ConfigApp) GetListLogs() []LogSet {
	return conf.LogList
}

func (conf *ConfigApp) GetAppNATS() *AppConfigNATS {
	return &conf.AppConfigNATS
}

func (conf *ConfigApp) GetAppMongoDB() *AppConfigMongoDB {
	return &conf.AppConfigMongoDB
}

func (conf *ConfigApp) Clean() {
	conf = &ConfigApp{}
}
