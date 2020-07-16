package msteams_webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Send(url string, card *MessageCard) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(card.ToJson()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(req)
}

func NewMessage(text string) *MessageCard {
	return &MessageCard{
		Type:            "MessageCard",
		Text:            text,
		PotentialAction: make([]*Actions, 0),
		Sections:        make([]*Section, 0),
	}
}

func (m *MessageCard) AddSection(section *Section) *MessageCard {
	m.Sections = append(m.Sections, section)
	return m
}

func (m *MessageCard) AddPotentialAction(actions *Actions) *MessageCard {
	m.PotentialAction = append(m.PotentialAction, actions)
	return m
}

func (m *MessageCard) ToJson() []byte {
	parsed, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return parsed
}

func NewHttpPost(name string, target string) *Actions {
	return &Actions{
		Type:   HttpPOST,
		Name:   name,
		Target: target,
	}
}

func NewActionCard(name string) *Actions {
	return &Actions{
		Type:    ActionCard,
		Name:    name,
		Actions: make([]*Actions, 0),
		Inputs:  make([]*Input, 0),
	}
}

func (a *Actions) AddTextInput(input *Input) *Actions {
	if a.Type != ActionCard {
		panic("only action card can add input")
	}
	a.Inputs = append(a.Inputs, input)
	return a
}

func (a *Actions) AddActions(actions *Actions) *Actions {
	a.Actions = append(a.Actions, actions)
	return a
}

func NewTextInput(id string, title string, required bool, value *string) *Input {
	return &Input{
		Type:       TextInput,
		Id:         id,
		Title:      title,
		IsRequired: required,
		Value:      value,
	}
}

func NewSection(text string) *Section {
	return &Section{
		Title:            "",
		StartGroup:       false,
		ActivityImage:    "",
		ActivityTitle:    "",
		ActivitySubtitle: "",
		ActivityText:     "",
		HeroImage:        nil,
		Images:           nil,
		Text:             text,
		Facts:            make([]*Fact, 0),
	}
}

func (s *Section) AddFact(fact *Fact) *Section {
	s.Facts = append(s.Facts, fact)
	return s
}

func NewFact(name string, value string) *Fact {
	return &Fact{
		Name:  name,
		Value: value,
	}
}
