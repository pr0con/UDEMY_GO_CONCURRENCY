package main

import(
	"os"
	"fmt"
	"bufio"
	"strings"
	
	"github.com/fatih/color"
)

func showBanner() {
	author := "Zach Lacourse"
	version := "v 0.1"	
	
	name := fmt.Sprintf("pr0T3rm (v.%s)", version)

	banner := `
              _______                                                
_____________ \   _  \   ____  ____   ____       ____  ____   _____  
\____ \_  __ \/  /_\  \_/ ___\/  _ \ /    \    _/ ___\/  _ \ /     \ 
|  |_> >  | \/\  \_/   \  \__(  <_> )   |  \   \  \__(  <_> )  Y Y  \
|   __/|__|    \_____  /\___  >____/|___|  / /\ \___  >____/|__|_|  /
|__|                 \/     \/           \/  \/     \/            \/ 
																		`
																		
	fmt.Println(banner)	
	all_lines := strings.Split(banner, "\n")
	w := len(all_lines[3])
	
	color.Green(fmt.Sprintf("%[1]*s", (w+len(name))/2, name))
	color.Blue(fmt.Sprintf("%[1]*s", (w+len(author))/2, author))
	
	
}

func main() {
	showBanner();
	
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print(">>> ")
	for scanner.Scan() {
		x := scanner.Text()
		fmt.Println(x)
		
	    switch(x) {
		    case "something":
		    	fmt.Println("almost another holiday i forget which one...");
		    	break
	    }		
		
		
		fmt.Print(">>> ")
	}
	
	
	if scanner.Err() != nil {
	    // handle error.
	}	
}