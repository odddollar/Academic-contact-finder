package global

import (
	"net/url"
	"testing"
)

// Implement unit testing for String function
func TestFoundContactStructString(t *testing.T) {
	// Create URLs as strings
	var urlstrings = []string{
		"https://teaching.csse.uwa.edu.au/units/CITS3002/",
		"https://research-repository.uwa.edu.au/en/persons/michael-wise",
		"",
		"http://johndoewebsite.com",
		"https://research-repository.uwa.edu.au/en/persons/barry-marshall",
		"https://reporter.anu.edu.au/people/brian-schmidt",
		"https://www.eva.mpg.de/genetics/staff/paabo/",
		"https://charleskkao-memorial.erg.cuhk.edu.hk/",
		"https://ece.princeton.edu/people/shlomo-shamai-shitz",
		"https://ieeexplore.ieee.org/author/37267737500",
		"https://www.u-tokyo.ac.jp/en/",
		"https://www.spacex.com",
		"Σ£þ§¦†É",
	}

	// Provide access to URL objects
	var urls = []*url.URL{}

	// Parse strings into URL objects and fill array
	for _, u := range urlstrings {
		up, _ := url.Parse(u)
		urls = append(urls, up)
	}

	// Create test cases
	var tests = []struct {
		input    FoundContactStruct
		expected string
	}{
		{
			FoundContactStruct{
				FirstName:   "chris",
				LastName:    "mcdonald",
				Salutation:  "professor",
				Email:       "chris@csse.uwa.edu.au",
				Institution: "university of western australia",
				URL:         urls[0]},
			"chris mcdonald, professor\nchris@csse.uwa.edu.au\nuniversity of western australia\nSource: https://teaching.csse.uwa.edu.au/units/CITS3002/\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "Michael",
				LastName:    "Wise",
				Salutation:  "Professor",
				Email:       "michael.wise@uwa.edu.au",
				Institution: "The University of Western Australia",
				URL:         urls[1]},
			"Michael Wise, Professor\nmichael.wise@uwa.edu.au\nThe University of Western Australia\nSource: https://research-repository.uwa.edu.au/en/persons/michael-wise\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "jane",
				LastName:    "doe",
				Salutation:  "dr",
				Email:       "",
				Institution: "",
				URL:         urls[2]},
			"jane doe, dr\n\n\nSource: \n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "John",
				LastName:    "Doe",
				Salutation:  "Associate Professor",
				Email:       "fake@email.com",
				Institution: "WA University",
				URL:         urls[3]},
			"John Doe, Associate Professor\nfake@email.com\nWA University\nSource: http://johndoewebsite.com\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "Barry",
				LastName:    "Marshall",
				Salutation:  "Professor",
				Email:       "barry.marshall@uwa.edu.au",
				Institution: "The University of Western Australia",
				URL:         urls[4]},
			"Barry Marshall, Professor\nbarry.marshall@uwa.edu.au\nThe University of Western Australia\nSource: https://research-repository.uwa.edu.au/en/persons/barry-marshall\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "Brian",
				LastName:    "Schmidt",
				Salutation:  "Professor",
				Email:       "Brian.Schmidt@anu.edu.au",
				Institution: "Australian National University",
				URL:         urls[5]},
			"Brian Schmidt, Professor\nBrian.Schmidt@anu.edu.au\nAustralian National University\nSource: https://reporter.anu.edu.au/people/brian-schmidt\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "Svante",
				LastName:    "Pääbo",
				Salutation:  "Professor",
				Email:       "paabo@eva.mpg.de",
				Institution: "Max Planck Institute for Evolutionary Anthropology",
				URL:         urls[6]},
			"Svante Pääbo, Professor\npaabo@eva.mpg.de\nMax Planck Institute for Evolutionary Anthropology\nSource: https://www.eva.mpg.de/genetics/staff/paabo/\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "Charles Kun (锟)",
				LastName:    "Kao (高)",
				Salutation:  "Professor",
				Email:       "",
				Institution: "The Chinese University of Hong Kong",
				URL:         urls[7]},
			"Charles Kun (锟) Kao (高), Professor\n\nThe Chinese University of Hong Kong\nSource: https://charleskkao-memorial.erg.cuhk.edu.hk/\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "Shlomo [שלמה]",
				LastName:    "Shamai (Shitz) [שמאי]",
				Salutation:  "Professor",
				Email:       "shamai@princeton.edu",
				Institution: "Technion – Israel Institute of Technology",
				URL:         urls[8]},
			"Shlomo [שלמה] Shamai (Shitz) [שמאי], Professor\nshamai@princeton.edu\nTechnion – Israel Institute of Technology\nSource: https://ece.princeton.edu/people/shlomo-shamai-shitz\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "상헌",
				LastName:    "팩",
				Salutation:  "Professor",
				Email:       "shpack@korea.ac.kr",
				Institution: "KAIST - Korea Advanced Institute of Science & Technology",
				URL:         urls[9]},
			"상헌 팩, Professor\nshpack@korea.ac.kr\nKAIST - Korea Advanced Institute of Science & Technology\nSource: https://ieeexplore.ieee.org/author/37267737500\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "ハルト",
				LastName:    "渡辺春虎",
				Salutation:  "Doctor",
				Email:       "",
				Institution: "The University of Tokyo",
				URL:         urls[10]},
			"ハルト 渡辺春虎, Doctor\n\nThe University of Tokyo\nSource: https://www.u-tokyo.ac.jp/en/\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "X Æ A-12",
				LastName:    "M",
				Salutation:  "Dr",
				Email:       "xai@space.x",
				Institution: "Space X University",
				URL:         urls[11]},
			"X Æ A-12 M, Dr\nxai@space.x\nSpace X University\nSource: https://www.spacex.com\n\n",
		},
		{
			FoundContactStruct{
				FirstName:   "ŽŒß",
				LastName:    "ÉÆæôεµñ",
				Salutation:  "Þ®Θ℉εššð®",
				Email:       "£πàî└•çºπ",
				Institution: "ªΦçÅτŠ§ úÑï",
				URL:         urls[12]},
			"ŽŒß ÉÆæôεµñ, Þ®Θ℉εššð®\n£πàî└•çºπ\nªΦçÅτŠ§ úÑï\nSource: %CE%A3%C2%A3%C3%BE%C2%A7%C2%A6%E2%80%A0%C3%89\n\n",
		},
	}

	// Check test cases for expected values
	for _, test := range tests {
		if test.input.String() != test.expected {
			t.Errorf("Test failed with input %s and expected output %s", test.input, test.expected)
		}
	}
}
