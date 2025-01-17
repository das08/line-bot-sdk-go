// Copyright 2018 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package linebot

import (
	"errors"
	"github.com/goccy/go-json"
	"log"
)

// UnmarshalFlexMessageJSON function
func UnmarshalFlexMessageJSON(data []byte) (FlexContainer, error) {
	raw := rawFlexContainer{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	return raw.Container, nil
}

type rawFlexContainer struct {
	Type      FlexContainerType `json:"type"`
	Container FlexContainer     `json:"-"`
}

type RfcAlias struct {
	Alias rawFlexContainer
}

func (c *rawFlexContainer) UnmarshalJSON(data []byte) error {
	//type alias rawFlexContainer
	raw := RfcAlias{}
	if err := json.Unmarshal(data, &raw.Alias); err != nil {
		log.Printf("rawFlexContainer UnmarshalJSON error: %v\n", err)
		return err
	}
	var container FlexContainer
	switch raw.Alias.Type {
	case FlexContainerTypeBubble:
		container = &BubbleContainer{}
	case FlexContainerTypeCarousel:
		container = &CarouselContainer{}
	default:
		return errors.New("invalid container type")
	}
	if err := json.Unmarshal(data, container); err != nil {
		return err
	}
	c.Type = raw.Alias.Type
	c.Container = container
	return nil
}

type rawFlexComponent struct {
	Type      FlexComponentType `json:"type"`
	Component FlexComponent     `json:"-"`
}

type RfcAlias2 struct {
	Alias rawFlexComponent
}

func (c *rawFlexComponent) UnmarshalJSON(data []byte) error {
	//type alias rawFlexComponent
	raw := RfcAlias2{}
	if err := json.Unmarshal(data, &raw.Alias); err != nil {
		log.Printf("rawFlexComponent UnmarshalJSON error: %v\n", err)
		return err
	}
	var component FlexComponent
	switch raw.Alias.Type {
	case FlexComponentTypeBox:
		component = &BoxComponent{}
	case FlexComponentTypeButton:
		component = &ButtonComponent{}
	case FlexComponentTypeFiller:
		component = &FillerComponent{}
	case FlexComponentTypeIcon:
		component = &IconComponent{}
	case FlexComponentTypeImage:
		component = &ImageComponent{}
	case FlexComponentTypeSeparator:
		component = &SeparatorComponent{}
	case FlexComponentTypeSpacer:
		component = &SpacerComponent{}
	case FlexComponentTypeText:
		component = &TextComponent{}
	case FlexComponentTypeVideo:
		component = &VideoComponent{}
	default:
		return errors.New("invalid flex component type")
	}
	if err := json.Unmarshal(data, component); err != nil {
		return err
	}
	c.Type = raw.Alias.Type
	c.Component = component
	return nil
}

type rawAction struct {
	Type   ActionType     `json:"type"`
	Action TemplateAction `json:"-"`
}

func (c *rawAction) UnmarshalJSON(data []byte) error {
	type alias rawAction
	raw := alias{}
	if err := json.Unmarshal(data, &raw); err != nil {
		log.Printf("rawAction UnmarshalJSON error: %v\n", err)
		return err
	}
	var action TemplateAction
	switch raw.Type {
	case ActionTypeURI:
		action = &URIAction{}
	case ActionTypeMessage:
		action = &MessageAction{}
	case ActionTypePostback:
		action = &PostbackAction{}
	case ActionTypeDatetimePicker:
		action = &DatetimePickerAction{}
	default:
		return errors.New("invalid action type")
	}
	if err := json.Unmarshal(data, action); err != nil {
		return err
	}
	c.Type = raw.Type
	c.Action = action
	return nil
}

// UnmarshalJSON method for BoxComponent
func (c *BoxComponent) UnmarshalJSON(data []byte) error {
	type alias BoxComponent
	raw := struct {
		Contents []rawFlexComponent `json:"contents"`
		Action   rawAction          `json:"action"`
		*alias
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		log.Printf("BoxComponent UnmarshalJSON error: %v\n", err)
		return err
	}
	components := make([]FlexComponent, len(raw.Contents))
	for i, content := range raw.Contents {
		components[i] = content.Component
	}
	c.Contents = components
	c.Action = raw.Action.Action
	return nil
}

// UnmarshalJSON method for ButtonComponent
func (c *ButtonComponent) UnmarshalJSON(data []byte) error {
	type alias ButtonComponent
	raw := struct {
		Action rawAction `json:"action"`
		*alias
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		log.Printf("ButtonComponent UnmarshalJSON error: %v\n", err)
		return err
	}
	c.Action = raw.Action.Action
	return nil
}

// UnmarshalJSON method for ImageComponent
func (c *ImageComponent) UnmarshalJSON(data []byte) error {
	type alias ImageComponent
	raw := struct {
		Action rawAction `json:"action"`
		*alias
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		log.Printf("ImageComponent UnmarshalJSON error: %v\n", err)
		return err
	}
	c.Action = raw.Action.Action
	return nil
}

// UnmarshalJSON method for TextComponent
func (c *TextComponent) UnmarshalJSON(data []byte) error {
	type alias TextComponent
	raw := struct {
		Action rawAction `json:"action"`
		*alias
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		log.Printf("TextComponent UnmarshalJSON error: %v", err)
		return err
	}
	c.Action = raw.Action.Action
	return nil
}

// UnmarshalJSON method for VideoComponent
func (c *VideoComponent) UnmarshalJSON(data []byte) error {
	type alias VideoComponent
	raw := struct {
		AltContent rawFlexComponent `json:"altContent"`
		*alias
	}{
		alias: (*alias)(c),
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	c.AltContent = raw.AltContent.Component
	return nil
}
