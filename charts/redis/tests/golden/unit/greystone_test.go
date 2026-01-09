package golden

// import (
// 	"path/filepath"
// 	"testing"
// 	"tests/golden"

// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// )

// 03 file test to compare new manifests with the old one
// func TestGolden04(t *testing.T) {
// 	t.Parallel()

// 	chartPath, err := filepath.Abs("../../..")
// 	require.NoError(t, err)

// 	suite.Run(t, &golden.TemplateGoldenTest{
// 		ChartPath:      chartPath,
// 		Release:        "golden-file-test",
// 		Namespace:      "test-namespace",
// 		GoldenFileName: "new_03",
// 		Templates: []string{
// 			"templates/configmap.yaml",
// 			"templates/headless-svc.yaml",
// 			"templates/health-configmap.yaml",
// 			// "templates/master/serviceaccount.yaml",
// 			"templates/metrics-svc.yaml",
// 			"templates/networkpolicy.yaml",
// 			"templates/prometheusrule.yaml",
// 			// "templates/replicas/serviceaccount.yaml",
// 			"templates/scripts-configmap.yaml",
// 			"templates/sentinel/service.yaml",
// 			"templates/sentinel/statefulset.yaml",
// 		},
// 		SetValues: defaultSetValues,
// 		ValuesFiles: []string{
// 			"../../values/03.values.yaml",
// 		},
// 		Suite:        suite.Suite{},
// 		IgnoredLines: []string{},
// 	})
// }

// 01 file test to compare new manifests with the old one
// func TestGolden04(t *testing.T) {
// 	t.Parallel()

// 	chartPath, err := filepath.Abs("../../..")
// 	require.NoError(t, err)

// 	suite.Run(t, &golden.TemplateGoldenTest{
// 		ChartPath:      chartPath,
// 		Release:        "golden-file-test",
// 		Namespace:      "test-namespace",
// 		GoldenFileName: "new_01",
// 		Templates: []string{
// 			"templates/configmap.yaml",
// 			"templates/headless-svc.yaml",
// 			"templates/health-configmap.yaml",
// 			// "templates/master/serviceaccount.yaml",
// 			"templates/metrics-svc.yaml",
// 			"templates/networkpolicy.yaml",
// 			// "templates/replicas/serviceaccount.yaml",
// 			"templates/scripts-configmap.yaml",
// 			"templates/sentinel/service.yaml",
// 			"templates/sentinel/statefulset.yaml",
// 			"templates/servicemonitor.yaml",
// 		},
// 		SetValues: defaultSetValues,
// 		ValuesFiles: []string{
// 			"../../values/01.values.yaml",
// 		},
// 		Suite:        suite.Suite{},
// 		IgnoredLines: []string{},
// 	})
// }
