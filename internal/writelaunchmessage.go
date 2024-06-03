package internal

import (
	"fmt"
	"os"
	"runtime"

	"shaper_stix/supportingfunctions"
)

func writeLaunchMessage() (string, error) {
	var (
		appName   = ""
		appStatus = "production"
	)

	if an, err := supportingfunctions.GetAppName("README.md", 1); err != nil {
		_, f, l, _ := runtime.Caller(0)
		return "", fmt.Errorf(" '%s' %s:%d", err, f, l-2)
	} else {
		appName = an
	}

	envValue, ok := os.LookupEnv("GO_PHELASTIC_MAIN")
	if ok && envValue == "development" {
		appStatus = envValue
	}

	appVersion := supportingfunctions.GetAppVersion(appName)
	msg := fmt.Sprintf("Shaper_stix_2.1 application, version %s is running. Application status is '%s'\n", appVersion, appStatus)

	return msg, nil
}
