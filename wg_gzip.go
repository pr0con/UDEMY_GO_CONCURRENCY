package main

import (
	"io"
	"os"	
	"fmt"
	"sync"
	"compress/gzip"
	"path/filepath"
)

func getFileList(files *[]string) filepath.WalkFunc {
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err)
        }
        *files = append(*files, path)
        return nil
    }
}

func isDir(path string) (bool) {
    fi, err := os.Stat(path)
    if err != nil {
	    fmt.Println(err)
        return false
    }

    return fi.Mode().IsDir()
}


func compressFile(filename string, wg *sync.WaitGroup) {
	fmt.Printf("Compressing %s\n", filename)
	in, err := os.Open(filename)
	
	if err != nil { fmt.Println(err) } else { 
		defer in.Close() 
	
		out, err := os.Create(filename + ".gz")
		if err != nil { fmt.Println(err) }else { 
			defer out.Close()
		
			gzout := gzip.NewWriter(out)
			_, err = io.Copy(gzout, in)
			gzout.Close()			
			
		}=
	
    }
    
    wg.Done()
}


func main() {
    var files []string

	root := "/var/www/files_to_gzip"
	err := filepath.Walk(root, getFileList(&files))
	if err != nil { fmt.Println(err) } else {
		var wg sync.WaitGroup
		
		for _, file := range files {
	    	fmt.Println(file)		
			
			if !isDir(file) {
				wg.Add(1)
				go compressFile(file, &wg)
	    	}			
		}	
		wg.Wait()
		fmt.Println("program done...")
	}
}