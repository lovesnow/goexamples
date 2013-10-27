package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type File struct {
	Name string `xml:"name,attr"`
	Size string `xml:"size,attr"`
	Time string `xml:"time,attr"`
}

type Directory struct {
	Name string `xml:"name,attr"`
	Size string `xml:"size,attr"`
	Time string `xml:"time,attr"`
	Directory []Directory `xml:"directory"`
	File []File `xml:"file"`
}

type Docment struct {
	XMLName xml.Name `xml:"tree"`
	Directory []Directory `xml:"directory"`
	File []File `xml:"file"`	
}


func getFileInfo(f os.FileInfo) (name, size, time string) {
	name = f.Name()
	size = strconv.FormatInt(f.Size(), 10)
	time = strconv.FormatInt(f.ModTime().Unix(), 10)
	return
}

func processDir(dirpath, rootpath string, dir *Directory) {
	file_list, err := ioutil.ReadDir(dirpath)
	if err == nil {
		for _, f := range file_list {
			name, size, time := getFileInfo(f)
			fullpath := path.Join(dirpath, name)			
			abspath := strings.Replace(fullpath, rootpath, ".", 1)
			
			if f.IsDir() {
				next_dir := Directory{Name: abspath, Size: size, Time: time}
				// 原来如此
				processDir(fullpath, rootpath, &next_dir)
				dir.Directory = append(dir.Directory, next_dir)
												
			} else {
				dir.File = append(dir.File, File{Name: abspath, Size: size, Time: time})
			}
		}
	}
}

func processRootDir(dirpath, rootpath string, dom *Docment) {
	file_list, err := ioutil.ReadDir(dirpath)
	if err == nil {
		for _, f := range file_list {
			name, size, time := getFileInfo(f)
			fullpath := path.Join(dirpath, name)			
			abspath := strings.Replace(fullpath, rootpath, ".", 1)
			
			if f.IsDir() {
				next_dir := Directory{Name: abspath, Size: size, Time: time}
				processDir(fullpath, rootpath, &next_dir)				
				dom.Directory = append(dom.Directory, next_dir)
												
			} else {
				dom.File = append(dom.File, File{Name: abspath, Size: size, Time: time})
			}
		}
	}
}

func parseDomTree(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Dot's ReadFile: ", err)
		return 
	}
	
	var tree Docment
	err = xml.Unmarshal(content, &tree)
	if err != nil {
		fmt.Println("Dot's Unmarshal: ", err)
		return 
	}
	
	fmt.Println(tree)
	
}

func createDomTree(path string) {
	var docment Docment
	
	path = strings.TrimSuffix(path, "/")
	processRootDir(path, path, &docment)
	
	output, err := xml.MarshalIndent(docment, "  ", "    ")
    if err != nil {
        fmt.Printf("error: %v\n", err)
    }
    os.Stdout.Write([]byte(xml.Header))
    os.Stdout.Write(output)
}


func main () {
	if len(os.Args) < 2 {
		fmt.Println(os.Args[0], " Directory")
		return 
		
	} else {
		// processDir(os.Args[1])
		// parseDomTree(os.Args[1])
		createDomTree(os.Args[1])
	}
}
