// Package hashablemaps offers tools for handling and manipulating
// map data structures in Go with hashable keys.
package hashablemaps_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablemaps"
	"github.com/GrewalAS/godash/utils"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]],
		[]string,
	]{
		{
			Name: "Empty HashableMap",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{},
			},
			Want: []string{},
		},
		{
			Name: "Single element HashableMap",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}},
			},
			Want: []string{"a"},
		},
		{
			Name: "Multiple elements HashableMap",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{
					"a": {Value: "a"},
					"b": {Value: "b"},
					"c": {Value: "c"},
				},
			},
			Want: []string{"a", "b", "c"},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Keys()", hashablemaps.Keys[TestHashable], testCases)
}
