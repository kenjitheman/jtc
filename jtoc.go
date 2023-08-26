package jtoc

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

func Convert(jsonFilePath, outputCsvFileName string) {
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println("[ERROR] error reading JSON file:", err)
		return
	}

	var data []map[string]interface{}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("[ERROR] error unmarshaling JSON:", err)
		return
	}

	csvFile, err := os.Create(outputCsvFileName)
	if err != nil {
		fmt.Println("[ERROR] error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	if len(data) > 0 {
		header := make([]string, 0)
		for key := range data[0] {
			header = append(header, key)
		}
		csvWriter.Write(header)
	}

	for _, item := range data {
		row := make([]string, 0)
		for _, value := range item {
			row = append(row, fmt.Sprintf("%v", value))
		}
		csvWriter.Write(row)
	}

	fmt.Println("[SUCCESS] conversion successful!")
}

func AsyncConvert(jsonFilePath, outputCsvFileName string, numWorkers int) {
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println("[ERROR] error reading JSON file:", err)
		return
	}

	var data []map[string]interface{}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("[ERROR] error unmarshaling JSON:", err)
		return
	}

	jsonChannel := make(chan map[string]interface{})

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go processJSON(jsonChannel, &wg, outputCsvFileName)
	}

	for _, item := range data {
		jsonChannel <- item
	}
	close(jsonChannel)

	wg.Wait()

	fmt.Println("[SUCCESS] asynchronous conversion successful !")
}

func processJSON(
	jsonChannel <-chan map[string]interface{},
	wg *sync.WaitGroup,
	outputCsvFileName string,
) {
	defer wg.Done()

	csvFile, err := os.Create(outputCsvFileName)
	if err != nil {
		fmt.Println("[ERROR] error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	for item := range jsonChannel {
		row := make([]string, 0)
		for _, value := range item {
			row = append(row, fmt.Sprintf("%v", value))
		}
		csvWriter.Write(row)
	}
}
