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

    if len(os.Args) < 2 {
        logger.Error("Usage: bettypy <file1> <file2> ...")
        os.Exit(1)
    }

    totalErrors := 0
    validFilesFound := false

    for _, pattern := range os.Args[1:] {
        files, err := filepath.Glob(pattern)
        if err != nil {
            logger.Error(fmt.Sprintf("Error processing pattern '%s': %s", pattern, err.Error()))
            os.Exit(1)
        }

        if len(files) == 0 {
            logger.Error(fmt.Sprintf("No files matched pattern '%s'", pattern))
            continue
        }

        for _, filename := range files {
            if filepath.Ext(filename) != ".py" {
                logger.Warning(fmt.Sprintf("%s is not a Python file, skipping...", filename))
                continue
            }

            validFilesFound = true

            contentBytes, err := ioutil.ReadFile(filename)
            if err != nil {
                logger.Error(fmt.Sprintf("Error reading file '%s': %s", filename, err.Error()))
                continue
            }

            content := string(contentBytes)
            logger.Info(fmt.Sprintf("%s - Starting analysis...", filename))

            checkers := []checker.Checker{
                checker.NewLineLengthChecker(strings.Split(content, "\n"), filename),
                checker.NewCommentChecker(strings.Split(content, "\n"), filename),
                checker.NewNewLineChecker(content, filename),
                checker.NewDockstringChecker(content, filename),
                checker.NewIdentationChecker(strings.Split(content, "\n"), filename),
                checker.NewCodetagChecker(content, filename),
                checker.NewEmptyFileChecker(content, filename),
                checker.NewUnusedImportChecker(strings.Split(content, "\n"), filename),
                checker.NewFunctionChecker(strings.Split(content, "\n"), filename),
            }

            for _, c := range checkers {
                errors := c.Run()
                for _, e := range errors {
                    logger.Warning(e)
                    totalErrors++
                }
            }
        }
    }

    if !validFilesFound {
        logger.Error("No valid Python files were found. Please check your patterns.")
        os.Exit(1)
    }

    if totalErrors == 0 {
        logger.Success("Analysis completed successfully, no errors found.")
    } else {
        logger.Error(fmt.Sprintf("Analysis completed with %d errors.", totalErrors))
    }
}