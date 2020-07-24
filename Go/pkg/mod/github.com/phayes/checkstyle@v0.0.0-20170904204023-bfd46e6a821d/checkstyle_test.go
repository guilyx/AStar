package checkstyle

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestCheckStyle(t *testing.T) {
	xmlString := `<?xml version="1.0" encoding="UTF-8"?>
	<checkstyle version="1.0.0">
			<file name="/path/to/code/myfile.go">
					<error line="2"  message="msg1" source="Ruleset.RuleName"/>
					<error line="20"  message="msg2" source="Generic.Constant"/>
					<error line="47"  message="msg3" source="ScopeIndent"/>
					<error line="47" message="msg4" source="Format.MultipleAlignment"/>
					<error line="51" message="msg5" source="Comment.FunctionComment"/>
			</file>
	</checkstyle>`

	checkStyleElement := New()
	err := xml.Unmarshal([]byte(xmlString), &checkStyleElement)
	if err != nil {
		t.Error(err)
	}

	// <checkstyle>
	if checkStyleElement.Version != "1.0.0" {
		t.Error("Bad checkstyle version")
	}
	if len(checkStyleElement.File) != 1 {
		t.Error("Wrong number of child <file> elements")
	}

	// <file>
	fileElement := checkStyleElement.File[0]
	if fileElement.Name != "/path/to/code/myfile.go" {
		t.Error("Bad file name")
	}
	if len(fileElement.Error) != 5 {
		t.Error("Wrong number of child <error> elements")
	}

	// <error>
	errorElement := fileElement.Error[0]
	if errorElement.Line != 2 {
		t.Error("Bad line number")
	}
	if errorElement.Message != "msg1" {
		t.Error("Bad error message")
	}
	if errorElement.Source != "Ruleset.RuleName" {
		t.Error("Bad error source")
	}
	if errorElement.Severity != SeverityNone {
		t.Error("Bad error Severity")
	}

	// Test round trip
	roundtripXML := checkStyleElement.String()
	roundtrip := New()
	err = xml.Unmarshal([]byte(roundtripXML), &roundtrip)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(roundtrip, checkStyleElement) {
		t.Error("Round Trip failed")
	}

}

func TestBuildCheckStyle(t *testing.T) {
	checkStyle := New()

	checkfile := checkStyle.EnsureFile("path/to/file")

	checkError := NewError(10, 5, SeverityError, "msg1", "test")
	checkfile.AddError(checkError)

	// Check to make sure they are the same
	checkfileDuplicate := checkStyle.EnsureFile("path/to/file")
	if checkfile != checkfileDuplicate {
		t.Error("checkfile != checkfileDuplicate")
	}

	// Check the output
	if checkStyle.String() != `<checkstyle version="1.0.0"><file name="path/to/file"><error line="10" column="5" severity="error" message="msg1" source="test"></error></file></checkstyle>` {
		t.Error("Wrong output for String()")
	}
}
