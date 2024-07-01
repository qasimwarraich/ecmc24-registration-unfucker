package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type ParticipantInfo struct {
	name          string
	email         string
	category      string
	pronouns      string
	message       string
	cityOrCompany string
	raceNumber    string
}

var OUTDIR = "./out/"

func main() {
	fmt.Println("Welcome to the ECMC24 Registration Unfucker")

	createOutDir()

	file, err := os.Open("./ecmc24-registration.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err.Error())
	}
	lines = lines[1:]

	participants := generateParticipantList(lines)
	volunteers := generateVolunteerList(lines)

	generateParticipantsCSV(participants)
	generateVolunteersCSV(volunteers)
	generateEmailLists(participants, volunteers)

	fmt.Println("Unfucking sucessfully completed!")
	fmt.Println("Check the created out directory for generated files")
	fmt.Println("chistole ❤️")

}

func createOutDir() {
	err := os.MkdirAll(OUTDIR, 0755)
	if err != nil {
		log.Fatal("creating output directory: ", err.Error())
	} else {
		fmt.Println("Output directory './out' created")
	}
}

func generateParticipantList(lines [][]string) []ParticipantInfo {
	registeredParticipants := []ParticipantInfo{}
	for _, line := range lines {
		if line[len(line)-1] != "" {
			registeredParticipants = append(registeredParticipants, ParticipantInfo{
				name:          line[len(line)-1],
				email:         line[len(line)-2],
				category:      line[3],
				pronouns:      line[4],
				message:       line[5],
				cityOrCompany: line[10],
				raceNumber:    line[9],
			})
		}
	}

	return registeredParticipants
}

func generateVolunteerList(lines [][]string) []ParticipantInfo {
	registeredVolunteers := []ParticipantInfo{}
	for _, line := range lines {
		if line[11] != "" {
			registeredVolunteers = append(registeredVolunteers, ParticipantInfo{
				name:          line[11],
				email:         line[2],
				category:      line[3],
				pronouns:      line[4],
				message:       line[5],
				cityOrCompany: line[10],
				raceNumber:    line[9],
			})
		}
	}

	return registeredVolunteers
}

func generateParticipantsCSV(participants []ParticipantInfo) {
	fmt.Println("Generating Participant CSV")
	header := []string{"name", "email", "category", "pronouns", "racenumber", "city/team", "message"}
	file, err := os.Create(OUTDIR + "participants-ecmc24.csv")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	err = w.Write(header)
	if err != nil {
		log.Fatalln("Couldn't write header to file", err)
	}

	var data [][]string
	for _, v := range participants {
		row := []string{
			v.name,
			v.email,
			v.category,
			v.pronouns,
			v.raceNumber,
			v.cityOrCompany,
			v.message,
		}
		if err := w.Write(row); err != nil {
			log.Fatalln("Couldn't write row to file", err)
		}

		if err := w.WriteAll(data); err != nil {
			log.Fatalln("Couldn't write rows to file", err)
		}
	}
}

func generateVolunteersCSV(volunteers []ParticipantInfo) {
	fmt.Println("Generating volunteers CSV")
	header := []string{"name", "email", "category", "pronouns", "racenumber", "city/team", "message"}
	file, err := os.Create(OUTDIR + "volunteers-ecmc24.csv")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	err = w.Write(header)
	if err != nil {
		log.Fatalln("Couldn't write header to file", err)
	}

	var data [][]string
	for _, v := range volunteers {
		row := []string{
			v.name,
			v.email,
			v.category,
			v.pronouns,
			v.raceNumber,
			v.cityOrCompany,
			v.message,
		}
		if err := w.Write(row); err != nil {
			log.Fatalln("Couldn't write row to file", err)
		}

		if err := w.WriteAll(data); err != nil {
			log.Fatalln("Couldn't write rows to file", err)
		}
	}
}

func generateEmailLists(participants []ParticipantInfo, volunteers []ParticipantInfo) {
	fmt.Println("Generating email lists")

	volunteerEmails := []string{}
	for _, volunteer := range volunteers {
		volunteerEmails = append(volunteerEmails, volunteer.email)
	}

	participantEmails := []string{}
	for _, participant := range participants {
		participantEmails = append(participantEmails, participant.email)
	}

	allEmails := append(volunteerEmails, participantEmails...)

	volunteerEmailFile, err := os.Create(OUTDIR + "volunteer-emails-ecmc24.txt")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer volunteerEmailFile.Close()

	participantEmailFile, err := os.Create(OUTDIR + "participant-emails-ecmc24.txt")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer participantEmailFile.Close()

	allEmailFile, err := os.Create(OUTDIR + "all-emails-ecmc24.txt")
	if err != nil {
		log.Fatalln("Couldn't create file", err)
	}
	defer allEmailFile.Close()

	volunteerEmailString := strings.Join(volunteerEmails, ",")
	_, err = volunteerEmailFile.WriteString(volunteerEmailString)
	if err != nil {
		log.Fatal(err)
	}

	participantEmailString := strings.Join(participantEmails, ",")
	_, err = participantEmailFile.WriteString(participantEmailString)
	if err != nil {
		log.Fatal(err)
	}

	allEmailString := strings.Join(allEmails, ",")
	_, err = allEmailFile.WriteString(allEmailString)
	if err != nil {
		log.Fatal(err)
	}

}
