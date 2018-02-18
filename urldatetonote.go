/*
Copyright 2018 Joshua Vécsei

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "UrldateToNote"
	app.Usage = "Bibtex urldate to note"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "prefix, p",
			Value: "last visited at",
			Usage: "adds `PREFIX` to urldate",
		},
	}

	app.Action = func(c *cli.Context) error {
		inputFile := ""
		outputFile := "new.bib"
		if c.NArg() > 1 {
			inputFile = c.Args().Get(0)
			outputFile = c.Args().Get(1)
		} else {
			log.Fatal("Error! Usage: urldatetonote input.bib output.bib")
		}
		prefix := c.String("prefix")
		bibFile, err := ioutil.ReadFile(inputFile)

		if err != nil {
			log.Fatal(err)
		}

		urlDateRegexp, _ := regexp.Compile(`(?i)urldate = \{([0-9\.-/a-z]+)\}`)
		updatedBibtex := urlDateRegexp.ReplaceAllString(string(bibFile), "note = {"+prefix+" $1}")
		err = ioutil.WriteFile(outputFile, []byte(updatedBibtex), 777)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Successfully wrote new file %s", outputFile)

		return nil
	}
	app.Run(os.Args)
}
