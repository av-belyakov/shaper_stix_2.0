package testmongodbhandler_test

import (
	"fmt"
	"shaper_stix/confighandler"
	"shaper_stix/databaseapi/mongodbapi"
	"shaper_stix/datamodels"
	"sync"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const DIR_ROOT = "shaper_stix_2.1"

var _ = Describe("Routing", Ordered, func() {
	testDomainObjectList := map[string]bool{
		"attack-pattern":   false,
		"campaign":         false,
		"course-of-action": false,
		"grouping":         false,
		"identity":         false,
		"indicator":        false,
		"infrastructure":   false,
		"intrusion-set":    false,
		"location":         false,
		"malware":          false,
		"malware-analysis": false,
		"note":             false,
		"observed-data":    false,
		"opinion":          false,
		"report":           false,
		"threat-actor":     false,
		"tool":             false,
		"vulnerability":    false,
	}

	testCyberObservableObjectList := map[string]bool{
		"artifact":             false,
		"autonomous-system":    false,
		"directory":            false,
		"domain-name":          false,
		"email-addr":           false,
		"email-message":        false,
		"file":                 false,
		"ipv4-addr":            false,
		"ipv6-addr":            false,
		"mac-addr":             false,
		"mutex":                false,
		"network-traffic":      false,
		"process":              false,
		"software":             false,
		"url":                  false,
		"user-account":         false,
		"windows-registry-key": false,
		"x509-certificate":     false,
	}

	var (
		mongomodule *mongodbapi.MongoDBModule
		conf        confighandler.ConfigApp

		logging  = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)

		err, errConf error
	)

	BeforeAll(func() {
		conf, errConf = confighandler.NewConfig(DIR_ROOT)
		mongomodule, err = mongodbapi.NewClientMongoDB(*conf.GetAppMongoDB(), logging, counting)
		sumTestElem := len(testDomainObjectList) + len(testCyberObservableObjectList)
		var num int

		var wg sync.WaitGroup
		wg.Add(sumTestElem)

		go func() {
			for v := range counting {
				wg.Done()
				num++
				fmt.Printf("msg: %s, sum: %d, received: %d\n", v.DataMsg, sumTestElem, num)

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
		}()

		for k := range testDomainObjectList {
			mongomodule.ChanInputModule <- mongodbapi.ChanInputMongoDB{
				CommonChanMongoDB: mongodbapi.CommonChanMongoDB{
					ObjectType: k,
				},
			}
		}

		fmt.Println("STOP reading list 'testDomainObjectList'")

		for k := range testCyberObservableObjectList {
			mongomodule.ChanInputModule <- mongodbapi.ChanInputMongoDB{
				CommonChanMongoDB: mongodbapi.CommonChanMongoDB{
					ObjectType: k,
				},
			}
		}

		fmt.Println("STOP reading list 'testCyberObservableObjectList'")

		wg.Wait()

		fmt.Println("STOP BEFOREALL")
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
