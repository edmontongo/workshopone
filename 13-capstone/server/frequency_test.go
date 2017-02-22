package main

import (
	"bytes"
	"testing"
)

type testFrequency string

func (tf testFrequency) FetchParagraph() string {
	return string(tf)
}

func TestFrequencyProblemValidatePasses(t *testing.T) {
	fp := FrequencyProblem{"paragraph", [3]string{"one", "two", "three"}}

	if !fp.Validate(bytes.NewBufferString("one,two,three")) {
		t.Error("Expect 'one,two,three' to be correct")
	}
}

func TestFrequencyProblemGenerate(t *testing.T) {
	tests := []struct {
		tf       testFrequency
		solution string
	}{
		{"the fair is right around the corner and tickets have just gone on sale we are selling a limited number of tickets at a discount so move fast and get yours while they are still available this is going to be an event you will not want to miss first off the school fair is a great value when compared with other forms of entertainment also your ticket purchase will help our school and when you help the school it helps the entire community but that’s not all every ticket you purchase enters you in a drawing to win fabulous prizes and don’t forget you will have mountains of fun because there are acres and acres of great rides fun games and entertaining attractions spend time with your family and friends at our school fair buy your tickets now", "the,you,and"},
		{"his began as one paragraph but needed two one for the problem and one for the solution also notice that the second paragraph is a process paragraph it would be very easy to add an introduction and conclusion to these two paragraphs and have a complete essay people often install a kitty door only to discover that they have a problem the problem is their cat will not use the kitty door there are several common reasons why cats won’t use kitty doors first they may not understand how a kitty door works they may not understand that it is a little doorway just for them second many kitty doors are dark and cats cannot see to the other side as such they can’t be sure of what is on the other side of the door so they won’t take the risk one last reason cats won’t use kitty doors is because some cats don’t like the feeling of pushing through and then having the door drag across their back but don’t worry—there are solutions to this problem the first step in solving the problem is to prop the door open with tape this means your cat will now be able to see through to the other side your cat will likely begin using the kitty door immediately once your cat has gotten used to using the kitty door remove the tape sometimes cats will continue to use the kitty door without any more prompting if this does not happen you will want to use food to bribe your cat when it’s feeding time sit on the opposite side of the door from your cat and either click the top of the can or crinkle the cat food bag open the door to show your cat that it is both you and the food waiting on the other side of the door repeat this a couple times and then feed your cat after a couple days of this your kitty door problem will be gone", "door,to,the"},
		{"his began as one paragraph but needed two one for the problem and one for the solution also notice that the second paragraph is a process paragraph it would be very easy to add an introduction and conclusion to these two paragraphs and have a complete essay people often install a kitty door only to discover that they have a problem the problem is their cat will not use the kitty door there are several common reasons why cats won’t use kitty doors first they may not understand how a kitty door works they may not understand that it is a little doorway just for them second many kitty doors are dark and cats cannot see to the other side as such they can’t be sure of what is on the other side of the door so they won’t take the risk one last reason cats won’t use kitty doors is because some cats don’t like the feeling of pushing through and then having the door drag across their back but don’t worry—there are solutions to this problem the first step in solving the problem is to prop the door open with tape this means your cat will now be able to see through to the other side your cat will likely begin using the kitty door immediately once your cat has gotten used to using the kitty door remove the tape sometimes cats will continue to use the kitty door without any more prompting if this does not happen you will want to use food to bribe your cat when it’s feeding time sit on the opposite side of the door from your cat and either click the top of the can or crinkle the cat food bag open the door to show your cat that it is both you and the food waiting on the other side of the door repeat this a couple times and then feed your cat after a couple days of this your kitty door problem will be gone", "to,the,door"},
	}

	for i, test := range tests {
		gen := FrequencyProblemGenerator{test.tf}
		if !gen.Generate().Validate(bytes.NewBufferString(test.solution)) {
			t.Log(gen.Generate())
			t.Errorf("Failed to validate case %d", i)
		}

	}
}

func TestFrequencyProblemValidateFails(t *testing.T) {
	fp := FrequencyProblem{"paragraph", [3]string{"one", "two", "three"}}
	inputs := []string{"an orange", "64", "16.0"}

	for _, test := range inputs {
		if fp.Validate(bytes.NewBufferString(test)) {
			t.Errorf("Did not expect '%s' to be valid input", test)
		}
	}
}
