package main

import (
    "bettypy/internal/checker"
    "bettypy/internal/log"
    "os"
    "path/filepath"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    logger := log.NewLogger()

    if len(os.Args) != 2 {
        logger.Error("Usage: bettypy <file>")
        os.Exit(1)
    }

    filename := os.Args[1]

    if filepath.Ext(filename) != ".py" {
        logger.Error("Error: Only Python files (.py) are supported.")
        os.Exit(1)
    }

    contentBytes, err := ioutil.ReadFile(filename)
    if err != nil {
        logger.Error("Error reading file: " + err.Error())
        os.Exit(1)
    }

    content := string(contentBytes)

    logger.Info(fmt.Sprintf("%s - Starting analysis...", filename))

    checkers := []checker.Checker{
        checker.NewLineLengthChecker(strings.Split(content, "\n"), filename),
        checker.NewCommentChecker(strings.Split(content, "\n"), filename),
        checker.NewNewLineChecker(content, filename),
        checker.NewDockstringChecker(content, filename),
        checker.NewIdentationChecker(strings.Split(content, "\n"), filename),
    }

    totalErrors := 0
    for _, c := range checkers {
        errors := c.Run()
        for _, e := range errors {
            logger.Warning(e)
            totalErrors++
        }
    }

    if totalErrors == 0 {
        logger.Success("Analysis completed successfully, no errors found.")
    } else {
        logger.Error(fmt.Sprintf("Analysis completed with %d errors.", totalErrors))
    }
}