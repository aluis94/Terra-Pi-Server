package templateEngine

import "os"

func CreateCrontTabFile() {

}

func CreateFile(scriptName string) {
	f, err := os.Create(scriptName)
	check(err)
	defer f.Close()

}

func DeleteFile(filename string) {
	err := os.Remove(filename)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
