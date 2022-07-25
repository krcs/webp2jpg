package main

import (
    "fmt"
    "os"
    "strings"
    "path/filepath"
    "image/jpeg"
    "golang.org/x/image/webp"
)

func createOutputFile(inputPath string) string {
    directory := filepath.Dir(inputPath)
    filename := filepath.Base(inputPath)
    result := strings.Replace(filename, filepath.Ext(filename), ".jpg", 1)
    result = filepath.Join(directory, result)
    return result
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Missing or incorrect number of parameters.")
        os.Exit(1)
    }

    filename := os.Args[1]

    outputfilename := createOutputFile(filename)

    input, err := os.Open(filename)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer input.Close()

    image, err := webp.Decode(input)
    if (err != nil) {
        fmt.Println("Error while decoding image:", err)
        os.Exit(1)
    }

    output, err := os.OpenFile(outputfilename, os.O_WRONLY|os.O_CREATE, 0755)
    if (err != nil) {
        fmt.Println("err")
        os.Exit(1)
    }

    if err = jpeg.Encode(output, image, nil); err != nil {
        fmt.Println("Error while encoding image:", err)
        os.Exit(1)
    }

    defer output.Close()
    fmt.Println("File saved in", outputfilename)
}
