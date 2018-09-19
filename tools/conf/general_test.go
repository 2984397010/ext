package conf_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"v2ray.com/core/common"
	. "v2ray.com/ext/tools/conf"
)

func diffMessage(a proto.Message, b proto.Message) error {
	if !proto.Equal(a, b) {
		return fmt.Errorf("Difference in messages\n%v\n%v", a, b)
	}

	return nil
}

func loadJSON(creator func() Buildable) func(string) (proto.Message, error) {
	return func(s string) (proto.Message, error) {
		instance := creator()
		if err := json.Unmarshal([]byte(s), instance); err != nil {
			return nil, err
		}
		return instance.Build()
	}
}

type TestCase struct {
	Input  string
	Parser func(string) (proto.Message, error)
	Output proto.Message
}

func runMultiTestCase(t *testing.T, testCases []TestCase) {
	for _, testCase := range testCases {
		actual, err := testCase.Parser(testCase.Input)
		common.Must(err)
		if !proto.Equal(actual, testCase.Output) {
			t.Fatalf("Failed in test case:\n%s\nActual:\n%v\nExpected:\n%v", testCase.Input, actual, testCase.Output)
		}
	}
}
