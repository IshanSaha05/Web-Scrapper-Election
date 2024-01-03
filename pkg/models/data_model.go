package models

type CommonACData struct {
	ParliamentConstituency string `bson:"ParliamentConstituency"` // Parliament Constituency Name -- Done
	State                  string `bson:"State"`                  // State Name -- Done
	Year                   int    `bson:"Year"`                   // Year of election -- Done
	WinningCandidate       string `bson:"WinningCandidate"`       // Winning Candidate Name -- Done
	WinningCandidateParty  string `bson:"WinningCandidateParty"`  // Winning Candidate Party Name -- Done
	TotalACs               int    `bson:"TotalACs"`               // Total Number of ACs present in the PC -- Done
}

type CommonACTableData struct {
	AssemblyConstituency string `bson:"AssemblyConstituency"` // Assembly Constituency Name
	TotalVotes           int    `bson:"TotalVotes"`           // Total number of votes given in that PC or AC
}

type DifferentACData struct {
	District        string  `bson:"District"`        // District Name
	Type            string  `bson:"Type"`            // Whether the details are of a particular PC or AC. If PC value = "PC", otherwise "AC"
	Candidate       string  `bson:"Candidate"`       // Candidate Name
	CandidateParty  string  `bson:"CandidateParty"`  // Candidate Party Name
	Votes           int     `bson:"Votes"`           // Votes received by the Candidate
	VotesPercentage float64 `bson:"VotesPercentage"` // Votes Percentage
}

type ACData struct {
	CACData      CommonACData      `bson:",inline"` // Storing the commonACData
	CACTableData CommonACTableData `bson:",inline"` // Storing the commonACTableData
	DACData      DifferentACData   `bson:",inline"` // Storing the DifferentACData
}
