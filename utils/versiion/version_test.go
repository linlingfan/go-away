package versiion

import "testing"

func TestCompareVersion(t *testing.T) {
	result:=CompareVersion("1.3.1","1.2.3",".")
	println(result)
}
