// Copyright (c) 2019 Evan Loughlin
//
// MIT License

package cppcomponents_test

import (
	"testing"

	"github.com/emloughl/CppCodeGenerator/cppcomponents"
)

func Test_ToString(t *testing.T) {
	for _, tt := range Test_ToString_Data {
		newInclude := cppcomponents.NewInclude(tt.dataType)
		actualOutput := newInclude.ToString()
		if tt.expectedOutput != actualOutput {
			t.Errorf("ToString(): EXPECTED: %v, ACTUAL: %v", tt.expectedOutput, actualOutput)
		}
	}
}

var Test_ToString_Data = []struct {
	dataType       string
	expectedOutput string
}{
	// Will use braces for Qt classes
	{"QString", "#include <QString>"},
	{"QMap", "#include <QMap>"},

	// Will use quotes for user-defined types
	{"MyDefinedType", "#include \"MyDefinedType.h\""},

	// Will return nothing for C/C++ standard data types
	{"unsigned int", ""},
	{"double", ""},
	{"char **", ""},
	{"char *", ""},
	{"void", ""},
}

func Test_DeriveDependency_TypeContainsPointers_WillExtractBaseType(t *testing.T) {
	for _, tt := range Test_DeriveDependency_TypeContainsPointers_WillExtractBaseType_Data {
		actualOutput := cppcomponents.DeriveDependency(tt.input)
		if tt.expectedOutput != actualOutput {
			t.Errorf("DerivedDependency(%v): EXPECTED: %v, ACTUAL: %v", tt.input, tt.expectedOutput, actualOutput)
		}
	}
}

var Test_DeriveDependency_TypeContainsPointers_WillExtractBaseType_Data = []struct {
	input          string
	expectedOutput string
}{
	// Will remove trailing spaces and '*' characters
	{"char **", "char"},

	// Will not remove intermediate spaces
	{"unsigned int *", "unsigned int"},

	// Will remove leading '*' characters
	{"*unsigned char", "unsigned char"},

	// Will remove leading ampersands
	{"&int", "int"},
	{"&unsigned int", "unsigned int"},
	{"&&char", "char"},

	// Will remove leading  and trailing spaces
	{"        signed int ", "signed int"},
	{" signed int        ", "signed int"},
}
