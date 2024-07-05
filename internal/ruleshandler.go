package internal

import (
	"fmt"
	"runtime"

	"github.com/av-belyakov/shaper_stix_2.1/ruleinteraction"
)

// ProcessingRules хранилище правил
type ProcessingRules struct {
	rootDir  string
	rulesDir string
	alerts   *ruleinteraction.ListRule
	cases    *ruleinteraction.ListRule
}

// NewRulesHandler инициализирует хранилище правил
func NewRulesHandler(rootDir, rulesDir string) *ProcessingRules {
	return &ProcessingRules{
		rootDir:  rootDir,
		rulesDir: rulesDir,
	}
}

// GetRootDir возвращает наименование основной директории приложения
func (pr *ProcessingRules) GetRootDir() string {
	return pr.rootDir
}

// GetRulesDir возвращает наименование директории с правилами
func (pr *ProcessingRules) GetRulesDir() string {
	return pr.rulesDir
}

// GetAlertRules возвращает список правил для обработки Alerts
func (pr *ProcessingRules) GetAlertRules() *ruleinteraction.ListRule {
	return pr.alerts
}

// AddAlertRules добавляет новый список правил для обработки Alerts
func (pr *ProcessingRules) AddAlertRules(fileName string) (string, error) {
	var warning string

	lr, lw, err := createRules(pr.rootDir, pr.rulesDir, fileName)
	if err != nil {
		return warning, err
	}

	for _, v := range lw {
		warning += fmt.Sprintf("%s\n", string(v))
	}

	pr.alerts = lr

	return warning, nil
}

// GetCaseRules возвращает список правил для обработки Cases
func (pr *ProcessingRules) GetCaseRules() (*ruleinteraction.ListRule, error) {
	if pr.cases == nil {
		return nil, fmt.Errorf("it is necessary to initialize the case processing rules, to do this, you need to use the method 'AddCaseRules'")
	}

	return pr.cases, nil
}

// GetCaseRules добавляет список правил для обработки Cases
func (pr *ProcessingRules) AddCaseRules(fileName string) (string, error) {
	var warning string

	lr, lw, err := createRules(pr.rootDir, pr.rulesDir, fileName)
	if err != nil {
		return warning, err
	}

	for _, v := range lw {
		warning += fmt.Sprintf("%s\n", string(v))
	}

	pr.cases = lr

	return warning, nil
}

func createRules(rootDir, dirName, fileName string) (*ruleinteraction.ListRule, string, error) {
	var (
		err      error
		warning  string
		warnings []string
		lr       *ruleinteraction.ListRule
	)

	lr, warnings, err = ruleinteraction.NewListRule(rootDir, dirName, fileName)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return lr, warning, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	// проверяем наличие правил Pass или Passany
	if len(lr.GetRulePass()) == 0 && !lr.GetRulePassany() {
		msg := "there are no rules for handling received from NATS or all rules have failed validation"
		_, f, l, _ := runtime.Caller(0)
		return lr, warning, fmt.Errorf(" '%s' %s:%d", msg, f, l-3)
	}

	// если есть какие либо логические ошибки в файле с YAML правилами для обработки сообщений поступающих от NATS
	if len(warnings) > 0 {
		for _, v := range warnings {
			warning += fmt.Sprintln(v)
		}
	}

	return lr, warning, err
}
