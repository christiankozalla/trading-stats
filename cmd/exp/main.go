package main

import "log"

func main() {
	record := make(map[string]any)
	stringSlice := []string{"These", "are", "the", "values"}
	header := []string{"header1", "header2", "header3", "header4"}

	record["Hello"] = "World"

	for i, val := range stringSlice {
		log.Println("appending", val)
		record[header[i]] = val
	}

	log.Println(record)
}
