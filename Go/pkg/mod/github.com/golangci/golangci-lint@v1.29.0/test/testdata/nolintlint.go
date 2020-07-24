//args: -Enolintlint
//config: linters-settings.nolintlint.require-explanation=true
//config: linters-settings.nolintlint.require-specific=true
//config: linters-settings.nolintlint.allowing-leading-space=false
package testdata

import "fmt"

func Foo() {
	fmt.Println("not specific")         //nolint // ERROR "directive `.*` should mention specific linter such as `//nolint:my-linter`"
	fmt.Println("not machine readable") // nolint // ERROR "directive `.*`  should be written as `//nolint`"
	fmt.Println("extra spaces")         //  nolint:deadcode // because // ERROR "directive `.*` should not have more than one leading space"
}
