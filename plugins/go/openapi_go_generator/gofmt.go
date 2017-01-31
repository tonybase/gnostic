// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

// Code templates use "//-" prefixes to comment-out template operators
// that otherwise would add gratuitous blank lines.
// Use this to remove these lines after templates have been expanded.
func stripMarkers(inputBytes []byte) (outputBytes []byte) {
	inputString := string(inputBytes)
	inputLines := strings.Split(inputString, "\n")
	outputLines := make([]string, 0)
	for _, line := range inputLines {
		if strings.Contains(line, "//-") {
			removed := strings.TrimSpace(strings.Replace(line, "//-", "", 1))
			if removed != "" {
				log.Printf("keeping |%s|", removed)
				outputLines = append(outputLines, removed)
			}
		} else {
			outputLines = append(outputLines, line)
		}
	}
	outputString := strings.Join(outputLines, "\n")
	return []byte(outputString)
}

func gofmt(inputBytes []byte) (outputBytes []byte, err error) {
	cmd := exec.Command(runtime.GOROOT() + "/bin/gofmt")
	input, _ := cmd.StdinPipe()
	output, _ := cmd.StdoutPipe()
	cmderr, _ := cmd.StderrPipe()
	err = cmd.Start()
	if err != nil {
		return
	}
	input.Write(stripMarkers(inputBytes))
	input.Close()
	errors, err := ioutil.ReadAll(cmderr)
	if len(errors) > 0 {
		log.Printf("ERRORS %s", errors)
		return inputBytes, nil
	} else {
		outputBytes, err = ioutil.ReadAll(output)
	}
	return
}