package main

import (
	"encoding/json"
	"fmt"
	"os"
	"vatsal-patel/json-task/model"
	"vatsal-patel/json-task/utils"
)

func main() {

	// userData is a []byte data reading from user.json file
	userData := ReadJsonFile("user")
	// techData is a []byte data reading from tech.json file
	techData := ReadJsonFile("tech")
	// contactData is a []byte data reading from contact.json file
	contactData := ReadJsonFile("contact")

	// users is an slice of User type struct
	var users []model.User
	// Unmarshalling from userData to convert that JSON data into User type struct
	err := json.Unmarshal(userData, &users)
	if err != nil {
		panic(err)
	}

	// technologies is an slice of Technology type struct
	var technologies []model.Technology
	// Unmarshalling from techData to convert that JSON data into Technology type struct
	err = json.Unmarshal(techData, &technologies)
	if err != nil {
		panic(err)
	}

	// contacts is an slice of Contact type struct
	var contacts []model.Contact
	// Unmarshalling from contactData to convert that JSON data into Technology type struct
	err = json.Unmarshal(contactData, &contacts)
	if err != nil {
		panic(err)
	}

	// finalUsers is a FinalOutput type slice to store all resultant struct
	var finalUsers []model.FinalOutput = make([]model.FinalOutput, 0, len(users))

	var flag bool = false
	// Range over all users
	for _, user := range users {
		// Range all over technologies
		for _, technology := range technologies {
			// if Id of technology and user matches, then and only then it will continue
			if technology.Id == user.Id {
				// if both Id match then it will range over all contacts
				for _, contact := range contacts {

					// If Id of contact and user match, only then it create the final output data
					if contact.Id == user.Id {

						// NewTechDetails slice to change the tag "techdata" in output
						var newTechDetails []model.NewTechDetail = make([]model.NewTechDetail, 0)
						// newTech that will append in the NewTechDetails to add it in finalUser
						var newTech model.NewTechDetail
						for j := 0; j < len(technologies[technology.Id-1].TechDetails); j++ {
							newTech.Tech = technologies[technology.Id-1].TechDetails[j].Tech
							newTech.Experience = technologies[technology.Id-1].TechDetails[j].Experience
							newTechDetails = append(newTechDetails, newTech)
						}

						// finalUser is the part of output that will add into the final JSON output file
						finalUser := model.FinalOutput{
							Userid:      user.Id,
							Name:        user.Name,
							UserAddress: user.UserAddress,
							TechDetails: newTechDetails,
							Email:       contact.UserContact.Email,
							Phone:       utils.PHONE_CODE[user.UserAddress.Country] + "-" + contact.UserContact.Phone,
						}
						// update the finalUsers by appending the user
						finalUsers = append(finalUsers, finalUser)

						flag = true
						// break the loop
						break
					}
				}
			}
			// if matched tech and contact already found then quit the loop
			if flag {
				break
			}
		}
		// set flag to false for further optimization
		flag = false
	}

	// finalJsonData is data to write into the file, and it will indented because of MarshalIndent function
	finalJsonData, err := json.MarshalIndent(finalUsers, "", "\t")
	if err != nil {
		panic(err)
	}

	// WriteJsonFile function will write the []byte data into the the file
	WriteJsonFile(finalJsonData)

}

// WriteJsonFile function create the file if not exist and write the []byte data to a file
func WriteJsonFile(data []byte) {
	// Create the file
	outputFile, err := os.Create("json-data/final-output.json")
	if err != nil {
		panic(err)
	}
	// Write into the file that created
	outputFile.WriteString(string(data))

	// Close the file after write operation
	outputFile.Close()
}

// ReadJsonFile function reads the file according to passed fileName and return the data in form of byte array or panic if error encounter
func ReadJsonFile(fileName string) []byte {
	// Replace the fileName changed according to passed fileName
	fileName = fmt.Sprintf("json-data/%s.json", fileName)

	data, err := os.ReadFile(fileName)

	// Will panic if error encountered
	if err != nil {
		panic(err)
	}

	// return the data in []byte
	return data
}
