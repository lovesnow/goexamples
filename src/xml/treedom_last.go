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


func NewFile(name, size, time string) *File {
	return &File{ name, size, time }
}

func NewDirectory(name, size, time string) *Directory {
	return &Directory{ File:NewFile(name, size, time) }
}

type File struct {
	Name string `xml:"name,attr"`
	Size string `xml:"size,attr"`
	Time string `xml:"time,attr"`
}

type Tree struct {
	Dirs []Directory `xml:"directory"`
	Files []File `xml:"file"`
}

type Directory struct {
	*File
	Tree
}

func (d *Directory) AddFile (name, size, time string) {
	d.Files = append(d.Files, *NewFile(name, size, time))
}

func (d *Directory) AddDir (td Directory) {
	d.Dirs = append(d.Dirs, td)
}

type Docment struct {
	XMLName xml.Name `xml:"tree"`
	Tree
}

func (d *Docment) AddFile (name, size, time string) {
	d.Files = append(d.Files, *NewFile(name, size, time))
}

func (d *Docment) AddDir (td Directory) {
	d.Dirs = append(d.Dirs, td)
}

type IDirectory interface {
	AddFile (name, size, time string)
	AddDir (td Directory)
}


func getFileInfo(f os.FileInfo) (name, size, time string) {
	name = f.Name()
	size = strconv.FormatInt(f.Size(), 10)
	time = strconv.FormatInt(f.ModTime().Unix(), 10)
	return
}

func processDir(dirpath, rootpath string, dir IDirectory) {
	file_list, err := ioutil.ReadDir(dirpath)
	if err == nil {
		for _, f := range file_list {
			name, size, time := getFileInfo(f)
			fullpath := path.Join(dirpath, name)			
			abspath := strings.Replace(fullpath, rootpath, ".", 1)
			
			if f.IsDir() {
				
				next_dir := NewDirectory(abspath, size, time)
				// 原来如此
				processDir(fullpath, rootpath, next_dir)
				dir.AddDir(*next_dir)
												
			} else {
				dir.AddFile(abspath, size, time)
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
	processDir(path, path, &docment)
	
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
