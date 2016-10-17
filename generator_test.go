package multiswagger

import (
	"reflect"
	"sort"
	"testing"
)

func TestAppendKeys(t *testing.T) {
	cases := map[string]struct {
		data     string
		expected map[string]bool
	}{
		"not JSON": {
			`This is not JSON.`,
			map[string]bool{},
		},
		"pattern 1": {
			`{"en":"english","ja":"日本語"}`,
			map[string]bool{
				"en": true,
				"ja": true,
			},
		},
		"pattern 2": {
			`{"zh":"汉语"}`,
			map[string]bool{
				"en": true,
				"ja": true,
				"zh": true,
			},
		},
	}

	var mapKeys []string
	for k := range cases {
		mapKeys = append(mapKeys, k)
	}
	sort.Strings(mapKeys)

	for _, k := range mapKeys {
		tc := cases[k]
		appendKeys(&tc.data)
		if reflect.DeepEqual(keys, tc.expected) != true {
			t.Errorf("%s: appendKeys got %v, want %v", k, keys, tc.expected)
		}
	}
}

func TestExtract(t *testing.T) {
	cases := map[string]struct {
		data     string
		keys     map[string]bool
		key      string
		expected string
	}{
		"not JSON": {
			`This is not JSON.`,
			map[string]bool{
				"en": true,
			},
			"en",
			`This is not JSON.`,
		},
		"pattern 1": {
			`{"en":"english","ja":"日本語"}`,
			map[string]bool{
				"en": true,
				"ja": true,
			},
			"en",
			"english",
		},
		"pattern 2": {
			`{"en":"english","ja":"日本語"}`,
			map[string]bool{
				"en": true,
				"ja": true,
			},
			"ja",
			"日本語",
		},
	}

	var mapKeys []string
	for k := range cases {
		mapKeys = append(mapKeys, k)
	}
	sort.Strings(mapKeys)

	for _, k := range mapKeys {
		tc := cases[k]
		keys = tc.keys
		key = tc.key
		extract(&tc.data)
		if tc.data != tc.expected {
			t.Errorf("%s: extract got %v, want %v", k, tc.data, tc.expected)
		}
	}
}
