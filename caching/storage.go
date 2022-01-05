package caching

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Programs struct {
	Programs []Program `json:"programs"`
}

type Program struct {
	ID                     string
	TenantID               string
	Name                   string
	Channel                string
	Currency               string
	DefaultLanguage        string
	DisplayInterestRateBPS string
	Policies               Policy
	CreatedAt              string
	DeletedAt              string
}

type Policy struct {
	Application Application
	Eligibility Eligibility
	Transaction Transaction
}

type Application struct {
	TOF TOF
}

type TOF struct {
	processors []string
}

type Eligibility struct {
	Country       string
	Regions       []string
	BuyerStatuses []string
}

type Transaction struct {
	ExcludeFees                     bool
	PaymentPerFundingEvent          bool
	SupportsMultipleMerchantFunding bool
}

type Storage struct {
	programs     Programs
	progFileName string
}

func (s *Storage) GetCacheName() string {
	return "Storage"
}

func (s *Storage) Initialize() {
	s.progFileName = "./storage_data/program.json"
}

func (s *Storage) GetProgram(name string) Program {
	var prog Program

	programsFile, err := os.Open(s.progFileName)
	if err != nil {
		fmt.Println(err)
		return prog
	}
	defer programsFile.Close()

	byteValue, _ := ioutil.ReadAll(programsFile)
	json.Unmarshal(byteValue, &s.programs)

	for i := 0; i < len(s.programs.Programs); i++ {
		if s.programs.Programs[i].Name == name {
			prog = s.programs.Programs[i]
			break
		}
	}

	return prog
}

func (s *Storage) SetProgram(name string, value Program) {
	// Not implemented
}

func (s *Storage) Flush() {
	// Not implemented
}
