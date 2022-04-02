package appconfig

import (
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/fs"
)

type Example struct {
	One         string
	StringField string
	BoolField   bool
	Int32Field  int32
	Singleword  int
	HostTHING   string
	More        string

	StringFromEnv string
	BoolFromEnv   bool
	UintFromEnv   uint
	NetIPFromEnv  string

	Nest    Nested
	NestPtr *Nested

	Nested

	// TODO: slice, and map
	// TODO: type that defines Unmarshal/Decode method
}

type Nested struct {
	Two     string
	Twine   string
	Numb    int
	Flag    bool
	Ratio   float64
	Another string
}

func TestLoad(t *testing.T) {
	content := `
string_field: from-file
bool_field: true
int_32_field: 2
singleword: 3
host_thing: ok
more: not-this

string_from_env: from-file-2
bool_from_env: false
uint_from_env: 5

nest:
    numb: -2
    another: not-this

nest_ptr:
    two: "the-value"
    ratio: 3.15

two: "from-file-3"
`
	f := fs.NewFile(t, t.Name(), fs.WithContent(content))

	t.Setenv("APPNAME_STRING_FROM_ENV", "from-env-1")
	t.Setenv("APPNAME_BOOL_FROM_ENV", "true")
	t.Setenv("APPNAME_UINT_FROM_ENV", "412")
	t.Setenv("APPNAME_NET_IP_FROM_ENV", "0.0.0.0")
	t.Setenv("APPNAME_NEST_TWINE", "from-env-2")
	t.Setenv("APPNAME_NEST_RATIO", "3.14")
	t.Setenv("APPNAME_NEST_PTR_TWINE", "from-env-3")
	t.Setenv("APPNAME_NEST_PTR_FLAG", "true")
	t.Setenv("APPNAME_TWINE", "from-env-4")
	t.Setenv("APPNAME_MORE", "not-this-either")

	target := Example{
		One:         "left-as-default",
		StringField: "default-1",
		Int32Field:  1,
		More:        "default-2",
		Nest: Nested{
			Two: "left-as-default-2",
		},
	}

	opts := Options{
		Filename:  f.Path(),
		EnvPrefix: "APPNAME",
		Overrides: []string{
			"more=from-override-1",
			"nest.another=from-override-2",
			"nest_ptr.another=from-override-3",
			"another=from-override-4",
		},
	}
	err := Load(&target, opts)
	assert.NilError(t, err)
	expected := Example{
		One:           "left-as-default",
		StringField:   "from-file",
		BoolField:     true,
		Int32Field:    2,
		Singleword:    3,
		HostTHING:     "ok",
		StringFromEnv: "from-env-1",
		BoolFromEnv:   true,
		UintFromEnv:   412,
		NetIPFromEnv:  "0.0.0.0",
		More:          "from-override-1",
		Nest: Nested{
			Two:     "left-as-default-2",
			Twine:   "from-env-2",
			Numb:    -2,
			Ratio:   3.14,
			Another: "from-override-2",
		},
		NestPtr: &Nested{
			Two:     "the-value",
			Twine:   "from-env-3",
			Flag:    true,
			Ratio:   3.15,
			Another: "from-override-3",
		},
		Nested: Nested{
			Two:     "from-file-3",
			Twine:   "from-env-4",
			Another: "from-override-4",
		},
	}
	assert.DeepEqual(t, target, expected)
}
