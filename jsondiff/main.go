package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jsondiff/content"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	srcDir = "src"
	dstDir = "dst"
)

func main() {
	dirs, err := ListDir(fmt.Sprintf("./%s", srcDir))
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, dir := range dirs {
		if err := os.Mkdir(fmt.Sprintf("./%s/%s", dstDir, dir), os.ModePerm); err != nil {
			log.Fatalln(err.Error())
		}
	}

	fmt.Println(dirs)

	files, err := ListFile(fmt.Sprintf("./%s", srcDir))
	if err != nil {
		log.Fatalln(err.Error())
	}

	addId := "0"
	addContent := "add"

	updateId := "3"
	updateContent := "updated"

	deleteId := "5"

	for _, s := range files {
		ct := Import(fmt.Sprintf("./%s", s))
		fmt.Printf("src: %+v\n", ct)
		content.Add(ct, addId, addContent)
		content.Update(ct, updateId, updateContent)
		content.Delete(ct, deleteId)
		fmt.Printf("dst: %+v\n", ct)
		d := ReplaceDir(s)
		if err := Export(fmt.Sprintf("./%s", d), ct); err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func Import(path string) map[string]string {
	ba, _ := ioutil.ReadFile(path)
	var res map[string]string
	_ = json.Unmarshal(ba, &res)
	return res
}

func ListFile(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filenames := make([]string, 0)

	for _, f := range files {
		if f.IsDir() {
			fs, err := ListFile(filepath.Join(dir, f.Name()))
			if err != nil {
				return nil, err
			}
			filenames = append(filenames, fs...)
			continue
		}
		filenames = append(filenames, filepath.Join(dir, f.Name()))
	}

	return filenames, nil
}

func ListDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	dirs := make([]string, 0)

	for _, f := range files {
		if f.Name() == srcDir {
			continue
		}
		if f.IsDir() {
			dirs = append(dirs, f.Name())
			fs, err := ListDir(filepath.Join(dir, f.Name()))
			if err != nil {
				return nil, err
			}
			dirs = append(dirs, fs...)
			continue
		}
	}

	return dirs, nil
}

func Export(path string, content map[string]string) error {
	// bt, err := json.Marshal(content)
	// if err != nil {
	// 	return err
	// }

	// var buf bytes.Buffer

	// if err := json.Indent(&buf, bt, "", " "); err != nil {
	// 	return err
	// }

	// file, err := os.Create(path)
	// if err != nil {
	// 	return err
	// }

	// defer file.Close()

	// if err := json.NewEncoder(file).Encode(buf.String()); err != nil {
	// 	return err
	// }

	// return nil

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	if err := json.NewEncoder(file).Encode(content); err != nil {
		return err
	}

	return nil
}

func ReplaceDir(path string) string {
	return strings.Replace(path, srcDir, dstDir, 1)
}
