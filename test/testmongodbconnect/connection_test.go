package testmongodbhandler_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"shaper_stix/confighandler"
	"shaper_stix/databaseapi/mongodbapi"
	"shaper_stix/datamodels"
)

const DIR_ROOT = "shaper_stix_2.1"

var _ = Describe("Connection", Ordered, func() {
	var (
		//mongomodule *mongodbapi.MongoDBModule
		conf confighandler.ConfigApp

		logging  = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)

		err, errConf error
	)

	BeforeAll(func() {
		conf, errConf = confighandler.NewConfig(DIR_ROOT)

		//mongomodule
		_, err = mongodbapi.NewClientMongoDB(*conf.GetAppMongoDB(), logging, counting)
	})

	Context("Тест 1. Проверка подключения к СУБД", func() {
		It("Формирование конфигурационного объекта", func() {
			Expect(errConf).ShouldNot(HaveOccurred())
		})

		It("При подключении к СУБД не должно быть ошибок", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
