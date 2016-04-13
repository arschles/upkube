package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"github.com/arschles/upkube/maps"
	"gopkg.in/yaml.v2"
)

func main() {
	inFile := flag.String("infile", "infile.yaml", "the input file to parse")
	updatePath := flag.String("path", "a.b", "the dot-delimited path to update in the yaml file")
	updateValue := flag.String("val", "some_value", "the value to set at the given path")
	flag.Parse()

	if *inFile == "" || *updatePath == "" || *updateValue == "" {
		log.Fatalf("Error: 'infile', 'outfile', 'path' and 'val' all need to be set")
	}

	m := make(map[interface{}]interface{})
	fileBytes, err := ioutil.ReadFile(*inFile)
	if err != nil {
		log.Fatalf("Error reading %s (%s)", *inFile, err)
	}
	if err := yaml.Unmarshal(fileBytes, &m); err != nil {
		log.Fatalf("Error unmarshaling JSON at %s (%s)", *inFile, err)
	}

	pathElts := strings.Split(*updatePath, ".")
	newMap, err := traverseAndSet(pathElts, m, *updateValue)
	if err != nil {
		log.Fatalf("Error updating yaml (%s)", err)
	}
	newYamlBytes, err := yaml.Marshal(newMap)
	if err != nil {
		log.Fatalf("Error marshalling new yaml (%s)", err)
	}
	fmt.Println(strings.TrimSpace(string(newYamlBytes)))
}

func traverseAndSet(path []string, m map[interface{}]interface{}, val interface{}) (map[interface{}]interface{}, error) {
	if len(path) < 1 {
		return maps.Empty(), fmt.Errorf("no path elements left")
	} else if len(path) > 1 {
		pathElt := path[0]
		remaining, ok := m[pathElt]
		if !ok {
			return maps.Empty(), fmt.Errorf("Path ends at %s", path)
		}
		remainingMap, ok := remaining.(map[interface{}]interface{})
		if !ok {
			return maps.Empty(), fmt.Errorf("Element at starting path %s is not a map (it's a %s)", path, reflect.TypeOf(remaining))
		}
		subMap, err := traverseAndSet(path[1:], remainingMap, val)
		if err != nil {
			return maps.Empty(), err
		}
		m[path[0]] = subMap
		return m, nil
	}
	_, ok := m[path[0]]
	if !ok {
		return maps.Empty(), fmt.Errorf("Final path element %s doesn't exist", path[0])
	}
	m[path[0]] = val
	return m, nil
}
