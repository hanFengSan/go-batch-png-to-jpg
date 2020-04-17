package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Convert(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return errors.Wrap(err, "Open input file failed")
	}
	defer inputFile.Close()

	imgSrc, err := png.Decode(inputFile)
	if err != nil {
		return errors.Wrap(err, "png.Decode failed")
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return errors.Wrap(err, "os.Create(outputPath) failed")
	}
	defer outputFile.Close()

	var opt jpeg.Options
	opt.Quality = 80
	if err := jpeg.Encode(outputFile, imgSrc, &opt); err != nil {
		return errors.Wrap(err, "jpeg.Encode failed")
	}
	return nil
}

func readFiles() []string {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}
	result := []string{}
	for _, f := range files {
		if strings.Contains(f.Name(), ".png") || strings.Contains(f.Name(), ".PNG") {
			result = append(result, f.Name())
		}
	}
	return result
}

func wait() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please enter any key to exit...")
	reader.ReadString('\n')
}

func mkdir() {
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		os.MkdirAll("./output", 0777)
	}
}

func main() {
	succeedFileNum := 0
	files := readFiles()
	mkdir()
	for index, name := range files {
		i := strconv.Itoa(index + 1)
		outputName := strings.Replace(strings.Replace(name, ".png", ".jpg", 1), ".PNG", ".jpg", 1)
		if err := Convert(name, "./output/" + outputName); err != nil {
			fmt.Println(i + ". " + name + ": Failed!!!")
			fmt.Println(err)
		} else {
			fmt.Println(i + ". " + name + ": Done")
			succeedFileNum++
		}
	}
	fmt.Println(strconv.Itoa(succeedFileNum) + "/" + strconv.Itoa(len(files)) + " done")
	wait()
}
