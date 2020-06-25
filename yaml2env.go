package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	src := flag.String("src", "", "source yaml file path")
	flag.Parse()

	if *src == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(*src)
	if err != nil {
		fatal(err)
	}

	contents := make(map[string]interface{})

	err = yaml.Unmarshal(bytes, &contents)
	if err != nil {
		fatal(err)
	}

	data := contents["data"].(map[string]interface{})

	for key, value := range data {
		decoded, err := base64.StdEncoding.DecodeString(value.(string))
		if err != nil {
			fatal(err)
		}

		fmt.Printf("%s=%s\n", key, decoded)
	}
}

func fatal(v interface{}) {
	fmt.Println(v)
	os.Exit(1)
}
