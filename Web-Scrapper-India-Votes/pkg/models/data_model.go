package models

import "time"

type Basic_PC_Data struct {
	PC_Type       string    // Parliament Consituency Type.
	Poll_Date     time.Time // Polling Date.
	Counting_Date time.Time // Counting Date.
	Result_Date   time.Time // Result Date.
}

type PC_Table_Data struct {
	PC_Position         int    // Position of the Person in the list.
	PC_Candidate_Name   string // Name of the candidate.
	PC_Votes            int    // No of votes received by the candidate.
	PC_Votes_Percentage int    // Percentage of votes received by the candidate.
	PC_Party            string // Party Name of the candidate.
}

type Basic_AC_Data struct {
	Winner_Candidate string // Name of the winner candidate.
	Winner_Party     string // Name of the winner party.
}

type AC_Table_Data struct {
	AC_Name             string // Name of the assembly constituency.
	AC_Candidate_Name   string // Name of the candidate.
	AC_Party            string // Name of the party.
	AC_Votes            int    // Number of votes received.
	AC_Votes_Percentage int    // Percentage of votes received.
	AC_Total_Votes      int    // Total number of votes in the AC.
}

type Data struct {
	PC       bool      // Whether it is a Parliament Cosnsituency Data.
	AC       bool      // Whether it is a Assembly Constituency Data.
	Year     time.Time // Year in which election occured.
	State    string    // State in which election occured.
	PC_Name  string    // Parliament Constituency name.
	Total_AC int       // Total no of assembly constituencies under PC.
	BPC      Basic_PC_Data
	PTD      PC_Table_Data
	BAC      Basic_AC_Data
	ATD      AC_Table_Data
}
