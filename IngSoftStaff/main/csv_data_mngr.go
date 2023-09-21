
package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
	 "regexp"
)


//https://articles.wesionary.team/easy-working-with-csv-in-golang-using-gocsv-package-9c8424728bbe

//"Persona,Nombre,Apellido,En Funcionarios,Materias,Email ORT,Email personal,Celular,Graduado hasta 2023.1,Posgrado"
type NotUsed struct {
	Name string
}

type Person struct {
	Id            string `csv:"Persona"`
	FirstName     string `csv:"Nombre"`
	Surname       string `csv:"Apellido"`
	WorkEmail     string `csv:"Email_ORT"`
	PersonalEmail string `csv:"Email_Personal"`
	CellPhone     string `csv:"Celular"`
	Graduation    string `csv:"Graduado"`
	Degree        string `csv:"Graduado"`

}



func main() {

	 delimiter, err := ReadCSVHeader();
	 if err != nil {
		panic(err)
	}

	// set the pipe as the delimiter for reading
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = delimiter
		return r
	})

	clientsFile, err := os.Open("../../../data/profesIngSoft.csv")
//	clientsFile, err := os.OpenFile("../pkg/db/profesIngSoftOk.csv", os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()


	clients := []*Person{}
	
	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}
	for _, client := range clients {
		fmt.Println("Hello", client.FirstName)
	}


}

func ReadCSVHeader() (rune, error) {
	file, err := os.Open("../../../data/profesIngSoft.csv")
    if err != nil {
        return 0, err
    }

    // remember to close the file at the end of the program
    defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ','
	csvReader.FieldsPerRecord = -1

	// read csv header
	var header []string
	header, err = csvReader.Read()
	if err != nil {
        return 0, err
    }

	var r rune
	var ok bool

	if len(header) == 1 {
		r, ok = FindSeparator(header[0])
		if !ok {
			return 0, nil
		}
	}

	return r, nil

}



func FindSeparator(s string) (rune, bool) {
    re := regexp.MustCompile(`[;,]`)
    match := re.FindString(s)
    if match == "" {
        return 0, false
    }
    return rune(match[0]), true
}