package args

import (
	"testing"

	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	type TestCase struct {
		error    string
		expected []string
		data     interface{}
	}

	stringPtr := "test"

	run := func(testCase TestCase) func(t *testing.T) {
		return func(t *testing.T) {
			args, err := MarshalStruct(testCase.data)

			if testCase.error == "" {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, args)
			} else {
				assert.Equal(t, testCase.error, err.Error())
			}
		}
	}

	RegisterMarshalFunc((*height)(nil), marshalHeight)

	t.Run("basic", run(TestCase{
		data: &Basic{
			String:    "test",
			Int:       42,
			Int16:     16,
			Int32:     32,
			Int64:     64,
			UInt16:    16,
			UInt32:    32,
			UInt64:    64,
			Float32:   3.2,
			Float64:   6.4,
			StringPtr: &stringPtr,
			Bool:      true,
		},
		expected: []string{
			"string=test",
			"int=42",
			"int16=16",
			"int32=32",
			"int64=64",
			"u-int16=16",
			"u-int32=32",
			"u-int64=64",
			"float32=3.2",
			"float64=6.4",
			"string-ptr=test",
			"bool=true",
		},
	}))

	t.Run("no-default", run(TestCase{
		data:     &Basic{},
		expected: []string(nil),
	}))

	t.Run("data-must-be-a-pointer", run(TestCase{
		data:  Basic{},
		error: "data must be a pointer to a struct",
	}))

	t.Run("basic-slice", run(TestCase{
		data: &Slice{
			Strings:    []string{"1", "2", "", "3"},
			StringsPtr: []*string{&stringPtr, &stringPtr},
			Basics: []Basic{
				{
					String: "test",
					Int:    42,
				},
				{},
				{
					String: "test",
					Int:    42,
				},
			},
		},
		expected: []string{
			"strings.0=1",
			"strings.1=2",
			"strings.3=3",
			"strings-ptr.0=test",
			"strings-ptr.1=test",
			"basics.0.string=test",
			"basics.0.int=42",
			"basics.2.string=test",
			"basics.2.int=42",
		},
	}))

	t.Run("well-known-types", run(TestCase{
		data: &WellKnownTypes{
			Size: 20 * scw.GB,
		},
		expected: []string{
			"size=20GB",
		},
	}))

	t.Run("nested-basic", run(TestCase{
		data: &Nested{
			Basic: Basic{
				String: "test",
			},
		},
		expected: []string{
			"basic.string=test",
		},
	}))

	t.Run("map-basic", run(TestCase{
		data: &Map{
			Map: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		expected: []string{
			"map.key1=value1",
			"map.key2=value2",
		},
	}))

	t.Run("custom", run(TestCase{
		data: &CustomWrapper{
			CustomStruct: &CustomStruct{
				value: "TEST",
			},
			CustomString: CustomString("TEST"),
		},
		expected: []string{
			"custom-struct=test",
			"custom-string=test",
		},
	}))

	t.Run("insane", run(TestCase{
		data: func() interface{} {
			n1 := &Nested{Basic: Basic{String: "test"}}
			m1 := &map[string]**Nested{"key2": &n1}
			m2 := map[string]**map[string]**Nested{"key1": &m1}
			return &Insane{Map: &m2}
		}(),
		expected: []string{
			"map.key1.key2.basic.string=test",
		},
	}))

	t.Run("data-is-a-map", run(TestCase{
		data: &map[string]string{
			"key1": "v1",
			"key2": "v2",
		},
		expected: []string{
			"key1=v1",
			"key2=v2",
		},
	}))

	t.Run("data-is-an-enum", run(TestCase{
		data: &Enum{Color: ColorBlue, Size: Size1},
		expected: []string{
			"color=blue",
			"size=1",
		},
	}))

	t.Run("data-is-raw-args", run(TestCase{
		data: &RawArgs{
			"pro.access_key",
			"access_key",
		},
		expected: []string{
			"pro.access_key",
			"access_key",
		},
	}))

	h := height(14)
	t.Run("data-is-height-set", run(TestCase{
		expected: []string{
			"height=14cm",
		},
		data: &CustomArgs{
			Height: &h,
		},
	}))

	t.Run("data-is-empty-custom-args", run(TestCase{
		data:     &CustomArgs{},
		expected: []string(nil),
	}))
}

func TestMarshalValue(t *testing.T) {
	type TestCase struct {
		error    string
		expected string
		data     interface{}
	}

	run := func(testCase TestCase) func(t *testing.T) {
		return func(t *testing.T) {
			value, err := MarshalValue(testCase.data)

			if testCase.error == "" {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, value)
			} else {
				assert.Equal(t, testCase.error, err.Error())
			}
		}
	}

	RegisterMarshalFunc((*height)(nil), marshalHeight)

	t.Run("string", run(TestCase{
		data:     "a simple string",
		expected: "a simple string",
	}))

	t.Run("int", run(TestCase{
		data:     42,
		expected: "42",
	}))

	t.Run("custom", run(TestCase{
		data:     CustomString("CUSTOM-STRING"),
		expected: "custom-string",
	}))

	t.Run("nil", run(TestCase{
		data:     nil,
		expected: "",
	}))

	t.Run("typed-nil", run(TestCase{
		data:     []string(nil),
		expected: "",
	}))

	t.Run("custom-func", run(TestCase{
		data:     height(42),
		expected: "42cm",
	}))

	t.Run("a-struct", run(TestCase{
		data:  &Basic{},
		error: "data must be a marshalable value (a scalar type or a Marshaler)",
	}))
}
