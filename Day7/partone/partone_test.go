package partone

import (
	"fmt"
	"testing"
)

var fileTree = map[string]interface{}{ // this would represent "/"
	"someFile": 12345,
	"aDir": map[string]interface{}{
		"someInsideFile": 8675309,
		"aSubDir": map[string]interface{}{
			"anInsideInsideFile": 10,
		},
	},
}

func TestParseCommand(t *testing.T) {

	ParseTree(fileTree, "", 0)

}

func TestReadInput(t *testing.T) {
	fileTree := ReadInput("../test-input.txt")
	fmt.Println(fileTree)
}

func TestBoth(t *testing.T) {
	fileTree := ReadInput("../test-input-2.txt")
	directorySizes := ParseTree(fileTree, "", 0)
	fmt.Println("Directory sizes: ", directorySizes)
}
