//go:build unit
// +build unit

package yamltohtml_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Dmitrylolo/golang-tests/yamltohtml"
)

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestMain(m *testing.M) {
	fmt.Println("Testing YAML to HTML")
	code := m.Run()
	fmt.Println("All tests have been run")
	os.Exit(code)
}

func TestYamltoHTML(t *testing.T) {
	testCases := []TestCase{
		{
			desc: "Test Case 1",
			path: "testdata/test_01.yaml",
			expected: `<!DOCTYPE html>
		<html>
			<head>
				<title>Awesome Test Page</title>
			</head>
			<body>The content of the awesome test page</body>
		</html>`,
		},
		{
			desc: "Test Case 2",
			path: "testdata/test_02.yaml",
			expected: `<!DOCTYPE html>
		<html>
			<head>
				<title>Second Test Page</title>
			</head>
			<body>Content of the second test page</body>
		</html>`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			html, err := yamltohtml.YamltoHTML(testCase.path)
			if err != nil {
				t.Fatal(err)
			}

			t.Log(html)

			if html != testCase.expected {
				t.Fatalf("Expected %s, got %s", testCase.expected, html)
			}
		})
	}
}
