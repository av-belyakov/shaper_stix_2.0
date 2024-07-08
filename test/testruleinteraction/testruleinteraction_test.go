package testruleinteraction_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/av-belyakov/shaper_stix_2.1/internal"
	"github.com/av-belyakov/shaper_stix_2.1/ruleinteraction"
)

var _ = Describe("Testruleinteraction", Ordered, func() {
	const (
		ROOT_DIR   = "shaper_stix_2.1"
		RULE_ALERT = "msgrule_alert.yaml"
		RULE_CASE  = "msgrule_case.yaml"
	)

	var (
		alertRules, caseRules *ruleinteraction.ListRule
		law, lcw              []string
		procRules             *internal.ProcessingRules

		errAlertRules, errCaseRules error
	)

	printRuleResult := func(r *ruleinteraction.ListRule) string {
		resultPrint := fmt.Sprintln("RULES:")

		resultPrint += fmt.Sprintln("  REPLACE:")
		for k, v := range r.Rules.Replace {
			resultPrint += fmt.Sprintln("  ", k+1, ".")
			resultPrint += fmt.Sprintf("    searchField: '%s'\n", v.SearchField)
			resultPrint += fmt.Sprintf("    searchValue: '%s'\n", v.SearchValue)
			resultPrint += fmt.Sprintf("    replaceValue: '%s'\n", v.ReplaceValue)
		}

		resultPrint += fmt.Sprintln("  PASS:")
		for key, value := range r.Rules.Pass {
			resultPrint += fmt.Sprintln("  ", key+1, ".")
			for k, v := range value.ListAnd {
				resultPrint += fmt.Sprintln("    ", k+1, ".")
				resultPrint += fmt.Sprintf("      searchField: '%s'\n", v.SearchField)
				resultPrint += fmt.Sprintf("      searchValue: '%s'\n", v.SearchValue)
			}
		}

		resultPrint += fmt.Sprintln("  EXCLUDE:")
		for k, v := range r.Rules.Exclude {
			resultPrint += fmt.Sprintln("    ", k+1, ".")
			resultPrint += fmt.Sprintf("      searchField: '%s'\n", v.SearchField)
			resultPrint += fmt.Sprintf("      searchValue: '%s'\n", v.SearchValue)
			resultPrint += fmt.Sprintf("      accurateComparison: '%v'\n", v.AccurateComparison)
		}

		resultPrint += fmt.Sprintf("  PASSANY: '%v'\n", r.Rules.Passany)

		return resultPrint
	}

	BeforeAll(func() {
		//инициализация списка правил Alert
		alertRules, law, errAlertRules = ruleinteraction.NewListRule(ROOT_DIR, "configs", RULE_ALERT)

		//инициализация списка правил Case
		caseRules, lcw, errCaseRules = ruleinteraction.NewListRule(ROOT_DIR, "configs", RULE_CASE)

		//инициализация хранилища правил
		procRules = internal.NewRulesHandler(ROOT_DIR, "configs")
	})

	Context("Тест 1. Чтение файлов правил Alert и Case по отдельности", func() {
		It("Файл с правилами Alert должен быть успешно прочитан", func() {
			//инициализация списка правил
			fmt.Println("NEW RULES FILE", RULE_ALERT, ":")
			for k, v := range law {
				fmt.Printf("%d. %s\n", k, v)
			}
			fmt.Println("new rule result:")
			fmt.Println(printRuleResult(alertRules))

			_, _, err := alertRules.ReplacementRuleHandler("string", "objectType", "dfff")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(errAlertRules).ShouldNot(HaveOccurred())
		})

		It("Файл с правилами Case должен быть успешно прочитан", func() {
			//инициализация списка правил
			fmt.Println("NEW RULES FILE", RULE_CASE, ":")
			for k, v := range lcw {
				fmt.Printf("%d. %s\n", k, v)
			}
			fmt.Println("new rule result:")
			fmt.Println(printRuleResult(caseRules))

			_, _, err := alertRules.ReplacementRuleHandler("string", "objectType", "dfff")
			Expect(err).ShouldNot(HaveOccurred())

			Expect(errCaseRules).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 2. Инициализация функции ProcessingRules", func() {
		It("Должна быть добавлена основная директория приложения", func() {
			Expect(procRules.GetRootDir()).Should(Equal(ROOT_DIR))
		})
		It("Должна быть добавлена директория с правилами приложения", func() {
			Expect(procRules.GetRulesDir()).Should(Equal("configs"))
		})
		It("Должен быть добавлен список правил Alert", func() {
			warning, err := procRules.AddAlertRules(RULE_ALERT)

			fmt.Println("Alert Warning:", warning)
			fmt.Printf("Alert list:\n%v\n", procRules.GetAlertRules())

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(warning)).Should(Equal(0))
		})
		It("Должен быть добавлен список правил Case", func() {
			warning, err := procRules.AddCaseRules(RULE_CASE)

			listRule, err := procRules.GetCaseRules()
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("Case Warning:", warning)
			fmt.Printf("Case list:\n%v\n", listRule)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(warning)).Should(Equal(0))
		})
	})
})
