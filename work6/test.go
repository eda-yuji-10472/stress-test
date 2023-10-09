package main

import (
    "fmt"
    "os"
)

func main() {
    // Get the file size.
    fileSize, err := os.Stat("random.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Check if the file size is larger than 10GB.
    if fileSize.Size() > 10*1024*1024 {
        // Delete the file.
        err = os.Remove("random.txt")
        if err != nil {
            fmt.Println(err)
            return
        }

        // Recreate the file.
        f, err := os.Create("random.txt")
        if err != nil {
            fmt.Println(err)
            return
        }
		defer f.Close()
    }
}
