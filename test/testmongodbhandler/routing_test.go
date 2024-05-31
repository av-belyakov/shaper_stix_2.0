package testmongodbhandler_test

import (
	"fmt"
	"shaper_stix/confighandler"
	"shaper_stix/databaseapi/mongodbapi"
	"shaper_stix/datamodels"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Routing", Ordered, func() {
	testDomainObjectList := map[string]bool{
		"AttackPattern":   false,
		"Campaign":        false,
		"CourseOfAction":  false,
		"Grouping":        false,
		"Identity":        false,
		"Indicator":       false,
		"Infrastructure":  false,
		"IntrusionSet":    false,
		"Location":        false,
		"Malware":         false,
		"MalwareAnalysis": false,
		"Note":            false,
		"ObservedData":    false,
		"Opinion":         false,
		"Report":          false,
		"ThreatActor":     false,
		"Tool":            false,
		"Vulnerability":   false,
	}

	testCyberObservableObjectList := map[string]bool{
		"Artifact":           false,
		"AutonomousSystem":   false,
		"Directory":          false,
		"DomainName":         false,
		"EmailAddress":       false,
		"EmailMessage":       false,
		"File":               false,
		"IPv4Address":        false,
		"IPv6Address":        false,
		"MACAddress":         false,
		"Mutex":              false,
		"NetworkTraffic":     false,
		"Process":            false,
		"Software":           false,
		"URL":                false,
		"UserAccount":        false,
		"WindowsRegistryKey": false,
		"X509Certificate":    false,
	}

	var (
		mongomodule *mongodbapi.MongoDBModule
		conf        confighandler.ConfigApp

		done chan struct{}

		logging  = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)

		err, errConf error
	)

	BeforeAll(func() {
		conf, errConf = confighandler.NewConfig(DIR_ROOT)
		mongomodule, err = mongodbapi.NewClientMongoDB(*conf.GetAppMongoDB(), logging, counting)

		/*
			Не могу написать этот тест, надо синхронизировать гроутину с записью
			testDomainObjectList и testCyberObservableObjectList с окончанием
			перебора для этих мап, а то получится что из канала done придет раньше
			чем закончат приходить данные из counting

			И вообще тест почему то блокируется!
		*/

		go func() {
			for {
				select {
				case <-done:
					return

				case v := <-counting:
					fmt.Println(v)

					if v.DataType != "routing_test" {
						continue
					}

					if _, ok := testDomainObjectList[v.DataMsg]; ok {
						testDomainObjectList[v.DataMsg] = true
					}
					if _, ok := testCyberObservableObjectList[v.DataMsg]; ok {
						testCyberObservableObjectList[v.DataMsg] = true
					}
				}
			}
		}()

		for k := range testDomainObjectList {
			mongomodule.ChanInputModule <- mongodbapi.ChanInputMongoDB{
				CommonChanMongoDB: mongodbapi.CommonChanMongoDB{
					ObjectType: k,
				},
			}
		}

		fmt.Println("STOP 1")

		for k := range testCyberObservableObjectList {
			mongomodule.ChanInputModule <- mongodbapi.ChanInputMongoDB{
				CommonChanMongoDB: mongodbapi.CommonChanMongoDB{
					ObjectType: k,
				},
			}
		}

		fmt.Println("STOP 2")

		done <- struct{}{}
	})

	Context("Тест 0. Проверка подключения к СУБД", func() {
		It("Формирование конфигурационного объекта", func() {
			Expect(errConf).ShouldNot(HaveOccurred())
		})

		It("При подключении к СУБД не должно быть ошибок", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 1. Проверка доступности роутинга", func() {
		It("Все роуты из  DomainObject должны быть доступны", func() {
			for k, v := range testDomainObjectList {
				fmt.Println(k, " = ", v)

				if !v {
					fmt.Println(k, "routing error")

					Expect(v).Should(BeTrue())
				}
			}

			Expect(true).Should(BeTrue())
		})
		It("Все роуты из CyberObservable должны быть доступны", func() {
			for k, v := range testCyberObservableObjectList {
				if !v {
					fmt.Println(k, " routing error")

					Expect(v).Should(BeTrue())
				}
			}

			Expect(true).Should(BeTrue())
		})
	})
})
