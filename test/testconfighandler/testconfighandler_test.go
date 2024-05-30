package testconfighandler_test

import (
	"fmt"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"shaper_stix/confighandler"
)

var _ = Describe("Testconfighandler", func() {
	var (
		err  error
		conf confighandler.ConfigApp
	)

	Context("Тест 1. Чтение конфигурационного файла (по умолчанию config_prod.yaml)", func() {
		conf, err = confighandler.NewConfig("shaper_stix_2.1")

		It("При чтении конфигурационного файла ошибок быть не должно", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Все пораметры конфигрурационного файла для NATS должны быть успешно получены", func() {
			cn := conf.GetAppNATS()

			Expect(cn.Host).Should(Equal("nats.cloud.gcm"))
			Expect(cn.Port).Should(Equal(4222))
			Expect(cn.SubjectCase).Should(Equal("main_caseupdate"))
			Expect(cn.SubjectAlert).Should(Equal("main_alertupdate"))
		})
		It("Все пораметры конфигрурационного файла для MongoDB должны быть успешно получены", func() {
			cmdb := conf.GetAppMongoDB()

			Expect(cmdb.Host).Should(Equal("192.168.9.208"))
			Expect(cmdb.Port).Should(Equal(37017))
			Expect(cmdb.User).Should(Equal("module-isems-mrsict"))
			Expect(cmdb.Passwd).Should(Equal("vkL6Znj$Pmt1e1"))
			Expect(cmdb.NameDB).Should(Equal("isems-mrsict"))
		})
	})

	Context("Тест 2. Проверяем установленные значения переменных окружения", func() {
		const (
			NATS_HOST         = "nats.cloud.gcm.test.test"
			NATS_PORT         = 4545
			NATS_SUBJECTCASE  = "main_CASEupdate"
			NATS_SUBJECTALERT = "main_ALERTupdate"

			MDB_HOST   = "199.166.199.166"
			MDB_PORT   = 11111
			MDB_USER   = "module_placeholder_elasticsearch.test.test"
			MDB_PASSWD = "gDbv5cf7*F2.test.test"
			MDB_NAMEDB = "placeholder_elasticsearch.test.test"
		)

		os.Setenv("GO_PHELASTIC_NHOST", NATS_HOST)
		os.Setenv("GO_PHELASTIC_NPORT", strconv.Itoa(NATS_PORT))
		os.Setenv("GO_PHELASTIC_SUBJECTCASE", NATS_SUBJECTCASE)
		os.Setenv("GO_PHELASTIC_SUBJECTALERT", NATS_SUBJECTALERT)

		os.Setenv("GO_PHELASTIC_MONGOHOST", MDB_HOST)
		os.Setenv("GO_PHELASTIC_MONGOPORT", strconv.Itoa(MDB_PORT))
		os.Setenv("GO_PHELASTIC_MONGOUSER", MDB_USER)
		os.Setenv("GO_PHELASTIC_MONGOPASSWD", MDB_PASSWD)
		os.Setenv("GO_PHELASTIC_MONGONAMEDB", MDB_NAMEDB)

		confEnv, err := confighandler.NewConfig("shaper_stix_2.1")

		It("При чтении конфигурационного файла ошибок быть не должно", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("Все параметры конфигурации для NATS должны быть успешно установлены через соответствующие переменные окружения", func() {
			cn := confEnv.GetAppNATS()

			Expect(cn.Host).Should(Equal(NATS_HOST))
			Expect(cn.Port).Should(Equal(NATS_PORT))
			Expect(cn.SubjectCase).Should(Equal("main_CASEupdate"))
			Expect(cn.SubjectAlert).Should(Equal("main_ALERTupdate"))
		})
		It("Все параметры конфигурации для MONGODB должны быть успешно установлены через соответствующие переменные окружения", func() {
			cmdb := confEnv.GetAppMongoDB()

			Expect(cmdb.Host).Should(Equal(MDB_HOST))
			Expect(cmdb.Port).Should(Equal(MDB_PORT))
			Expect(cmdb.User).Should(Equal(MDB_USER))
			Expect(cmdb.Passwd).Should(Equal(MDB_PASSWD))
			Expect(cmdb.NameDB).Should(Equal(MDB_NAMEDB))
		})
	})

	Context("Тест 3. Проверяем работу функции NewConfig с разными значениями переменной окружения GO_PHMISP_MAIN", func() {
		It("Должно быть получено содержимое общего файла 'config.yaml'", func() {
			conf, err := confighandler.NewConfig("shaper_stix_2.1")

			//fmt.Println("conf = ", conf)

			//for k, v := range conf.GetListOrganization() {
			//	fmt.Printf("%d. OrgName: %s, SourceName: %s\n", k, v.OrgName, v.SourceName)
			//}

			commonApp := conf.GetCommonApp()

			//fmt.Println("------------------------ ZABBIX -------------------------")
			//fmt.Println("NetworkHost:", conf.GetCommonApp().Zabbix.NetworkHost)
			//fmt.Println("NetworkPort:", conf.GetCommonApp().Zabbix.NetworkPort)

			Expect(commonApp.Zabbix.NetworkHost).Should(Equal("192.168.9.45"))
			Expect(commonApp.Zabbix.NetworkPort).Should(Equal(10051))
			Expect(commonApp.Zabbix.ZabbixHost).Should(Equal("test-uchet-db.cloud.gcm"))
			Expect(len(commonApp.Zabbix.EventTypes)).Should(Equal(3))

			fmt.Println("===========================")
			fmt.Println(commonApp.Zabbix.EventTypes)
			fmt.Println("===========================")

			Expect(commonApp.Zabbix.EventTypes[0].EventType).Should(Equal("error"))
			Expect(commonApp.Zabbix.EventTypes[0].ZabbixKey).Should(Equal("shaper_stix.error"))
			Expect(commonApp.Zabbix.EventTypes[0].IsTransmit).Should(BeTrue())
			Expect(commonApp.Zabbix.EventTypes[0].Handshake.TimeInterval).Should(Equal(0))
			Expect(commonApp.Zabbix.EventTypes[0].Handshake.Message).Should(Equal(""))
			Expect(commonApp.Zabbix.EventTypes[1].EventType).Should(Equal("info"))
			Expect(commonApp.Zabbix.EventTypes[1].ZabbixKey).Should(Equal("shaper_stix.info"))
			Expect(commonApp.Zabbix.EventTypes[1].IsTransmit).Should(BeTrue())
			Expect(commonApp.Zabbix.EventTypes[2].EventType).Should(Equal("handshake"))
			Expect(commonApp.Zabbix.EventTypes[2].ZabbixKey).Should(Equal("shaper_stix.handshake"))
			Expect(commonApp.Zabbix.EventTypes[2].IsTransmit).Should(BeTrue())

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(conf.GetListLogs())).Should(Equal(5))
		})
	})
})
