package partone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var directorySizes map[string]int = map[string]int{}

func ParseTree(tree map[string]interface{}, currentDir string, depth int) map[string]int {
	fmt.Println(currentDir)
	for key, val := range tree {
		for i := 0; i < depth; i++ {
			fmt.Printf("\t")
		}
		switch v := val.(type) {
		case map[string]interface{}:
			fmt.Printf("+")
			newCurrentDir := fmt.Sprintf("%s/%s", currentDir, key)
			depth++
			ParseTree(v, newCurrentDir, depth)
		case int:
		case int32:
		case int64:
			// here v has type S
			fmt.Printf("-")
			subdirs := strings.Split(currentDir, "/")
			temp := subdirs[0]
			for _, subdir := range subdirs {
				if subdir != "" {
					temp = fmt.Sprintf("%s/%s", temp, subdir)
				} else {
					temp = "/"
				}
				directorySizes[temp] += int(v)
			}
			fmt.Printf("file: %s, size: %d", key, v)
			fmt.Println()
		default:
			// no match; here v has the same type as i
			log.Fatal("unknown type")
		}
	}
	return directorySizes
}

func ParseCommand(cmd string) string {
	// cd / means root
	if strings.HasPrefix(cmd, "$ ") {
		return strings.Replace(cmd, "$ ", "", 1)
	}
	return ""
}

func InsertTreeNodes(tree map[string]interface{}, cwd string, files []string) {
	fmt.Printf("Writing files %s to %s\n", files, cwd)
	// TODO navigate the tree based on the CWD string
	currentNode := tree
	dirs := strings.Split(cwd, "/")
	for _, dir := range dirs {
		if dir == "" {
			// this is the root tree so bail out
			continue
		}
		if dirAsNode, ok := currentNode[dir].(map[string]interface{}); ok {
			currentNode = dirAsNode
		}
	}
	for _, file := range files {
		if file[:3] == "dir" {
			dirName := file[4:]
			// make sure we setup the next node
			if _, ok := currentNode[dirName].(map[string]interface{}); !ok {
				currentNode[dirName] = map[string]interface{}{}
			}
		} else {
			fileParts := strings.Split(file, " ")
			fileSize, err := strconv.ParseInt(fileParts[0], 10, 32)
			if err != nil {
				log.Fatal("Tried to parse an invalid file size")
			}
			fileName := fileParts[1]
			currentNode[fileName] = fileSize
		}
	}
}

func ReadInput(filepath string) map[string]interface{} {
	fmt.Println(fmt.Sprintf("Reading file into string: %s", filepath))
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Failed to open file: ", filepath)
		panic(err)
	}
	defer file.Close()

	cwd := ""
	var listFilesLines []string
	fileTree := map[string]interface{}{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		if currentLine[0:2] == "$ " {
			if len(listFilesLines) > 0 {
				// This tells us we've accrued list files
				// we should establish the files listed into the file tree
				// and then reset listFilesLines
				InsertTreeNodes(fileTree, cwd, listFilesLines)
				listFilesLines = make([]string, 0)
			}
			// we only have two commands: cd and ls
			// so the next two characters are the command
			command := currentLine[2:4]
			switch command {
			case "cd":
				destDir := currentLine[5:]
				if destDir == "/" {
					fmt.Println("Resetting current working directory")
					cwd = ""
				} else if destDir == ".." {
					dirs := strings.Split(cwd, "/")
					cwd = strings.Join(dirs[:len(dirs)-1], "/")
				} else {
					cwd = fmt.Sprintf("%s/%s", cwd, destDir)
				}
				fmt.Println("Setting cwd to: ", cwd)
			case "ls":
			default:
				log.Fatal("Unrecognized command")
			}
		} else {
			listFilesLines = append(listFilesLines, currentLine)
		}
	}
	// Flush out the last of the list files lines
	if len(listFilesLines) > 0 {
		InsertTreeNodes(fileTree, cwd, listFilesLines)
	}
	return fileTree
}
