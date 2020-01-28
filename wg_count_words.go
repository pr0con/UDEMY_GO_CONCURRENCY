package main

import(
	"os"
	"fmt"
	"sync"
	"strings"
	"bufio"
	"path/filepath"
)

type words struct {
	sync.Mutex
	found map[string]int
}

func newWords() *words {
	return &words{found: map[string]int{}}
}

func (w *words) add(word string, n int) {
	w.Lock()
	defer w.Unlock()
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}




func isDir(path string) (bool) {
    fi, err := os.Stat(path)
    if err != nil {
	    fmt.Println(err)
        return false
    }

    return fi.Mode().IsDir()
}


func getFileList(files *[]string) filepath.WalkFunc {
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err)
        }
        *files = append(*files, path)
        return nil
    }
}




func countWords(path string, table *words) error {
	if !isDir(path) {	
		if filepath.Ext(strings.TrimSpace(path)) == ".txt" {
		
			fmt.Println("got here");
		
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
		
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)
			
			for scanner.Scan() {
				word := strings.ToLower(scanner.Text())
				table.add(word, 1)
			}
			return scanner.Err()
	    }
	}
	return nil
}





func main() {
	var files []string
	
    root := "/var/www/files_to_count_words"
    err := filepath.Walk(root, getFileList(&files))
    if err != nil {
        panic(err)
    }
    
    var wg sync.WaitGroup
    
    w := newWords()
	for _, f := range files {
		wg.Add(1)
		go func(file string) {
			if err := countWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	
	//display output
	fmt.Println("Words that appear more than once:")
	w.Lock()
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	w.Unlock()		     
    
}
