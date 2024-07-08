package testreportwrap_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/av-belyakov/shaper_stix_2.1/datamodels"
	"github.com/av-belyakov/shaper_stix_2.1/internal/decodejson"
	"github.com/av-belyakov/shaper_stix_2.1/internal/listhandlerjson"
	do "github.com/av-belyakov/shaper_stix_2.1/internal/wrappersobjectstix/domainobjects"
	"github.com/av-belyakov/shaper_stix_2.1/supportingfunctions"
)

func handlerCaseToReport(
	input <-chan datamodels.ChanOutputDecodeJSON,
	listWrapReport map[string][]func(interface{}),
	logging chan<- datamodels.MessageLogging) {

	for data := range input {
		//******************************************************************************
		//******** формирование обьекта относящегося к Report Domain Object STIX *******
		if reports, ok := listWrapReport[data.FieldBranch]; ok {
			for _, f := range reports {
				f(data.Value)
			}
		}
	}

	logging <- datamodels.MessageLogging{
		MsgData: "",
		MsgType: "STOP TEST",
	}
}

var _ = Describe("Reportwrap", Ordered, func() {
	var (
		fileByte []byte
		fileErr  error

		chanStoppedCounting  chan struct{}
		chanOutputDecodeJson chan datamodels.ChanOutputDecodeJSON
		logging              chan datamodels.MessageLogging
		counting             chan datamodels.DataCounterSettings

		//формируем новый объект 'report' в обёртке
		reportWrap *do.WrapperReport = do.NewWrapperReportDomainObjectsSTIX()

		//*************** Обработчик формирующий объект 'report' в обёртке ****************
		listWrapReport = listhandlerjson.NewHandlerReportDomainObjectSTIX(reportWrap)
	)

	BeforeAll(func() {
		chanStoppedCounting = make(chan struct{})
		chanOutputDecodeJson = make(chan datamodels.ChanOutputDecodeJSON)
		logging = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)

		fileByte, fileErr = supportingfunctions.ReadFileJson("test/filestest", "event_1.json")

		go func() {
			fmt.Println("___ Logging START")
			defer fmt.Println("___ Logging STOP")

			for log := range logging {
				if log.MsgType == "error" {
					fmt.Println("----", log, "----")
				}

				if log.MsgType == "STOP TEST" {
					chanStoppedCounting <- struct{}{}

					return
				}
			}
		}()

		//вывод данных счетчика
		go func() {
			for d := range counting {
				fmt.Printf("\tСчетчик %v\n", d.DataType)
			}
		}()

		decodeJson := decodejson.NewDecodeJsonMessageSettings(logging, counting)
		chanOutputDecodeJson = decodeJson.HandlerJsonMessage(fileByte, "test_id_73d8r3", "subject_case")
		go handlerCaseToReport(chanOutputDecodeJson, listWrapReport, logging)
	})

	Context("Тест 1. Проверка успешности чтения Json файла", func() {
		It("При чтении файла не должно быть ошибок", func() {
			Expect(fileErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Проверка успешности заполнения значений", func() {
		It("Должны быть успешно добавлены некоторые значения", func() {
			<-chanStoppedCounting

			Expect(reportWrap.GetName()).Should(Equal("FP (Внутреннее взаимодействие) с 10.67.32.11"))
			Expect(len(reportWrap.GetDescription())).ShouldNot(Equal(0))

			ros := reportWrap.GetReportOutsideSpecification()
			Expect(ros.GetObjectType()).Should(Equal("case"))
			Expect(ros.GetRootId()).Should(Equal("~498274384"))
			Expect(ros.GetObjectId()).Should(Equal("~498274384"))
			Expect(ros.GetCaseId()).Should(Equal("5815"))
			Expect(ros.ImpactStatus).Should(Equal("NoImpact"))
			Expect(ros.ResolutionStatus).Should(Equal("TruePositive"))

			//fmt.Println(reportWrap.ToStringBeautiful())
		})
	})

	/*Context("Тест . ", func() {
		It("", func() {

		})
	})*/
})
