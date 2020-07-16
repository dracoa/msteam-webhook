package msteams_webhook

import (
	"testing"
)

const TestURL = "https://outlook.office.com/webhook/8e8747ff-3b74-4621-8e6e-0647bdc559c5@83eef0ff-65ef-4870-89dc-e1b9c8e4356f/IncomingWebhook/0e9467ffe58f447cbd6ae7b0f6591ca4/8c83e08d-2982-4b26-9ee0-029d133267b0"

func TestSend_Basic(t *testing.T) {
	card := NewMessage("TestSend_Basic")
	resp, err := Send(TestURL, card)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Error(resp.Status)
	}
}

func TestSend_Action(t *testing.T) {
	card := NewMessage("TestSend_Action")
	a := NewHttpPost("Basic Http POST", "https://www.example.com")
	card.AddPotentialAction(a)
	resp, err := Send(TestURL, card)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Error(resp.Status)
	}
}

func TestSend_ActionInput(t *testing.T) {
	card := NewMessage("TestSend_ActionInput")
	card.Title = "Title"

	a := NewActionCard("Basic Action Card")
	a.AddTextInput(NewTextInput("name", "Your Name", false, nil))
	a.AddActions(NewHttpPost("Approve", "https://www.example.com/"))

	card.AddPotentialAction(a)
	card.AddPotentialAction(a)

	resp, err := Send(TestURL, card)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Error(resp.Status)
	}
}

func TestSend_Section(t *testing.T) {
	card := NewMessage("TestSend_ActionInput")
	card.Title = "Title"

	sec := NewSection("<table></table>")
	sec.Title = "Section Title"
	sec.ActivityImage = "https://teams.microsoft.com/api/mt/apac/beta/teams/8e8747ff-3b74-4621-8e6e-0647bdc559c5/profilepicturev2?etag=null&displayName=O365G-SR-IT&voidCache=true"
	sec.ActivitySubtitle = "ActivitySubtitle"
	sec.ActivityText = "ActivityText"
	sec.ActivityTitle = "ActivityTitle"

	sec.AddFact(NewFact("Name", "Value"))

	card.AddSection(sec)

	sec2 := NewSection("Section Text 2")
	sec2.Title = "No start group"
	sec2.StartGroup = true
	card.AddSection(sec2)

	sec3 := NewSection("Section 3 Text 2")
	sec3.Title = "No start group"
	sec3.StartGroup = false
	card.AddSection(sec3)

	resp, err := Send(TestURL, card)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Error(resp.Status)
	}
}

func TestSend_SectionHTML(t *testing.T) {
	card := NewMessage("###Visiting on **17 July** *(1000 - 1200)*")
	sec := NewSection(``)
	sec.StartGroup = true
	sec.AddFact(NewFact("Name", "Chan Tai Man"))
	sec.AddFact(NewFact("SID (Room)", "56781234 (103A)"))
	sec.AddFact(NewFact("Visitors", "1"))
	sec.AddFact(NewFact("Nof. Visit", "5"))
	card.AddSection(sec)
	sep := NewSection(`&nbsp;`)
	card.AddSection(sep)
	card.AddPotentialAction(NewHttpPost("Approve", "https://www.example.com/"))
	reject := NewActionCard("Reject")
	reject.AddTextInput(NewTextInput("message", "Message to student (optional):", false, nil))
	reject.AddActions(NewHttpPost("Send", "https://www.example.com/"))
	card.AddPotentialAction(reject)

	resp, err := Send(TestURL, card)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Error(resp.Status)
	}
}
