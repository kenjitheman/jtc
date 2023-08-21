package core

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
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

func AsyncConvert(jsonFilePath, outputCsvFileName string) {
	fmt.Println(jsonFilePath, outputCsvFileName) // mock | TODO
}
