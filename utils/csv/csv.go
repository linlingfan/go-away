package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

type CsvTable struct {
	FileName string
	Records  []CsvRecord
}

type CsvRecord struct {
	Record map[string]string
}

func (c *CsvRecord) GetInt(field string) int {
	r, err := strconv.Atoi(strings.TrimSpace(c.Record[field]))
	if err != nil {
		panic(err)
	}
	return r
}

func (c *CsvRecord) GetString(field string) string {
	r, ok := c.Record[field]
	if !ok {
		panic("GetString error! field :" + field)
	}
	return r
}

// 读取行数 row=0 表示全部行数(从哪一行读起)
func LoadCsvFile(fileName string, row int) *CsvTable {

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("%+v", err)
		return nil
	}
	defer file.Close()
	reader := csv.NewReader(file)
	if reader == nil {
		log.Printf("NewReader return nil and file :%+v", file)
		return nil
	}
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("%s+v", err)
		return nil
	}

	if len(records) == 1 && strings.Contains(records[0][0], "\r") {
		log.Printf("不支持的格式")
		return nil
	}

	if len(records) < row {
		log.Printf("csv row less : %d", row)
		return nil
	}
	colNum := len(records[0])
	recordNum := len(records)
	var allRecords []CsvRecord
	for i := row; i < recordNum; i++ {
		record := &CsvRecord{make(map[string]string)}
		for k := 0; k < colNum; k++ {
			record.Record[strings.TrimSpace(records[0][k])] = strings.TrimSpace(records[i][k])
		}
		allRecords = append(allRecords, *record)
	}
	var result = &CsvTable{
		fileName,
		allRecords,
	}
	return result
}
