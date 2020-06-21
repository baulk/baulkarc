package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	fmt.Fprintf(os.Stderr,"%s\n",filepath.Join("C:\\User\\example\\jack","back\\..\\hello"))
}
