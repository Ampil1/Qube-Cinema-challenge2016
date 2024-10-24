package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Region struct {
	CityCode     string
	ProvinceCode string
	CountryCode  string
	CityName     string
	ProvinceName string
	CountryName  string
}

type Distributor struct {
	Name    string
	Include []string
	Exclude []string
	Parent  *Distributor
}

var regions = make(map[string]Region)
var distributors = make(map[string]*Distributor)

func loadRegions(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		region := Region{
			CityCode:     record[0],
			ProvinceCode: record[1],
			CountryCode:  record[2],
			CityName:     record[3],
			ProvinceName: record[4],
			CountryName:  record[5],
		}
		regions[region.CityCode] = region
	}

	return nil
}

// Add a distributor with permissions
func addDistributor(name string, include, exclude []string, parent *Distributor) {
	distributors[name] = &Distributor{
		Name:    name,
		Include: include,
		Exclude: exclude,
		Parent:  parent,
	}
	fmt.Printf("Distributor added: %s\n", name)
	fmt.Printf("Included Regions: %v\n", include)
	fmt.Printf("Excluded Regions: %v\n", exclude)
}

func canDistribute(distributor *Distributor, regionCode string) (bool, error) {

	current := distributor
	fmt.Println("Current distributor:", current)
	for current != nil {
		for _, exRegion := range current.Exclude {
			if exRegion == regionCode {
				return false, nil
			}
		}
		for _, inRegion := range current.Include {
			if inRegion == regionCode {
				return true, nil
			}
		}
		current = current.Parent
	}
	return false, nil
}

func main() {

	action := flag.String("action", "check", "Action to perform (add/check)")
	distributorName := flag.String("name", "", "Distributor name")
	includeRegions := flag.String("include", "", "Comma separated region codes to include")
	excludeRegions := flag.String("exclude", "", "Comma separated region codes to exclude")
	parentDistributor := flag.String("parent", "", "Parent distributor name")
	regionCode := flag.String("region", "", "Region code to check distribution")
	csvFile := flag.String("csv", "cities.csv", "CSV file path for regions data")
	flag.Parse()

	err := loadRegions(*csvFile)
	if err != nil {
		fmt.Println("Error loading CSV:", err)
		return
	}
	// fmt.Printf("Loaded regions: %+v\n", regions)

	switch *action {
	case "add":
		if *distributorName == "" {
			fmt.Println("Distributor name is required for 'add' action")
			return
		}

		var parent *Distributor
		if *parentDistributor != "" {
			parent = distributors[*parentDistributor]
			if parent == nil {
				fmt.Println("Parent distributor not found")
				return
			}
		}

		includes := strings.Split(*includeRegions, ",")
		excludes := strings.Split(*excludeRegions, ",")

		addDistributor(*distributorName, includes, excludes, parent)
		fmt.Println("Distributor added successfully")
		fmt.Printf("Current distributors: %+v\n", distributors)

	case "check":
		if *distributorName == "" || *regionCode == "" {
			fmt.Println("Distributor name and region code are required for 'check' action")
			return
		}

		distributor := distributors[*distributorName]
		fmt.Printf("Checking for distributor: %s\n", *distributorName)

		if distributor != nil {
			fmt.Printf("Distributor details: %+v\n", *distributor)
		} else {
			fmt.Println("Distributor not found")
			return
		}

		canDist, err := canDistribute(distributor, *regionCode)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if canDist {
			fmt.Printf("YES, %s can distribute in %s\n", *distributorName, *regionCode)
		} else {
			fmt.Printf("NO, %s cannot distribute in %s\n", *distributorName, *regionCode)
		}

	default:
		fmt.Println("Invalid action. Available actions: add, check")
	}
}
