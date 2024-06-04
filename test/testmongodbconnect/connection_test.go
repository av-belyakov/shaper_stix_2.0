package testmongodbhandler_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/av-belyakov/shaper_stix_2.1/confighandler"
	"github.com/av-belyakov/shaper_stix_2.1/databaseapi/mongodbapi"
	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
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

		ctx, _ := context.WithCancel(context.Background())
		//mongomodule
		_, err = mongodbapi.NewClientMongoDB(ctx, *conf.GetAppMongoDB(), logging, counting)
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
