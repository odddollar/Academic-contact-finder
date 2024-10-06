package background

import (
	"testing"
)

func TestFindExactMatch(t *testing.T) {
	var tests = []struct {
		inputOriginal      string
		inputOriginalLower string
		inputToFind        string
		expected           string
	}{
		{"Chris McDonald", "chris mcdonald", "chris", "Chris"},
		{"Chris McDonald", "chris mcdonald", "mcdonald", "McDonald"},
		{"The University of Western Australia", "the university of western australia", "university of western australia", "University of Western Australia"},
		{"Chris McDonald", "Chris McDonald", "Chris", "Chris"},
		{"Chris will be available for Consultation (Office Hours) during the Catch-Up session each week, Wednesdays 3-5pm in CSSE Room 1.24.", "chris will be available for consultation (office hours) during the catch-up session each week, wednesdays 3-5pm in csse room 1.24.", "chris", "Chris"},
		{"<h4>Unit coordinator: Chris McDonald, Rm 2.20 of the CSSE building</h4>", "<h4>unit coordinator: chris mcdonald, rm 2.20 of the csse building</h4>", "chris", "Chris"},
		{"'Computer expert' Chris McDonald is actually Professor Chris McDonald. He holds appointments at the University of Western Australia and at Dartmouth. I was lucky to have him as a lecturer for several course and as my honours supervisor.", "'computer expert' chris mcdonald is actually professor chris mcdonald. he holds appointments at the university of western australia and at dartmouth. i was lucky to have him as a lecturer for several course and as my honours supervisor.", "chris mcdonald", "Chris McDonald"},
		{"<p>'Computer expert' Chris McDonald is actually Professor Chris McDonald.</p>", "<p>'computer expert' chris mcdonald is actually professor chris mcdonald.</p>", "mcdonald", "McDonald"},
		{"Slightly OT: &quot;Computer expert&quot; Chris McDonald is actually Professor Chris McDona...", "slightly ot: &quot;computer expert&quot; chris mcdonald is actually professor chris mcdona...", "mcdonald", "McDonald"},
		{"/profile/Chris-Mcdonald-2?__cf_chl_rt_tk=xK4mUdozmkpp.kTkHGIXDPcRpvlxOlJO0wn8qqg.bgc-1728201658-0.0.1.1-6356", "/profile/Chris-Mcdonald-2?__cf_chl_rt_tk=xK4mUdozmkpp.kTkHGIXDPcRpvlxOlJO0wn8qqg.bgc-1728201658-0.0.1.1-6356", "Chris", "Chris"},
	}

	for _, test := range tests {
		if findExactMatch(test.inputOriginal, test.inputOriginalLower, test.inputToFind) != test.expected {
			t.Errorf("Test failed with input %s, %s, %s and output %s", test.inputOriginal, test.inputOriginalLower, test.inputToFind, test.expected)
		}
	}
}

func TestGetHighestSalutation(t *testing.T) {
	var tests = []struct {
		input    []string
		expected string
	}{
		{[]string{"Doc"}, "Doctor"},
		{[]string{"Dr"}, "Doctor"},
		{[]string{"Dr."}, "Doctor"},
		{[]string{"Asst Prof"}, "Assistant Professor"},
		{[]string{"Asst. Prof."}, "Assistant Professor"},
		{[]string{"Assoc Prof"}, "Associate Professor"},
		{[]string{"Assoc. Prof."}, "Associate Professor"},
		{[]string{"Prof"}, "Professor"},
		{[]string{"Prof."}, "Professor"},
		{[]string{"Doctor"}, "Doctor"},
		{[]string{"Doctor", "Assistant Professor"}, "Assistant Professor"},
		{[]string{"Doctor", "Associate Professor", "Assistant Professor"}, "Associate Professor"},
		{[]string{"Doctor", "Professor", "Associate Professor", "Assistant Professor"}, "Professor"},
		{[]string{"doctor", "professor", "associate professor", "assistant professor"}, "Professor"},
		{[]string{"lecturer", "senior lecturer", "mr", "researcher", "ms", "mrs", "dr"}, "Doctor"},
		{[]string{"doctor", "dr", "assistant prof", "associate"}, "Assistant Professor"},
		{[]string{"doctor", "dr", "associate prof", "assistant prof", "associate", "assistant"}, "Associate Professor"},
		{[]string{"associate professor", "assistant professor", "a. prof", "assoc prof", "asst prof"}, "Professor"},
		{[]string{"associate professor", "assistant professor", "a. professor", "assoc prof", "asst prof"}, "Professor"},
	}

	for _, test := range tests {
		if getHighestSalutation(test.input) != test.expected {
			t.Errorf("Test failed with input %s and output %s", test.input, test.expected)
		}
	}
}
