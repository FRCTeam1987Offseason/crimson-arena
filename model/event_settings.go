// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model and datastore read/write methods for event-level configuration.

package model

import "github.com/FRCTeam1987/crimson-arena/game"

type EventSettings struct {
	Id                          int `db:"id"`
	Name                        string
	ElimType                    string
	NumElimAlliances            int
	SelectionRound2Order        string
	SelectionRound3Order        string
	TBADownloadEnabled          bool
	TbaPublishingEnabled        bool
	TbaEventCode                string
	TbaSecretId                 string
	TbaSecret                   string
	NetworkSecurityEnabled      bool
	ApAddress                   string
	ApPassword                  string
	ApChannel                   int
	SwitchAddress               string
	SwitchPassword              string
	PlcAddress                  string
	AdminPassword               string
	WarmupDurationSec           int
	AutoDurationSec             int
	PauseDurationSec            int
	TeleopDurationSec           int
	WarningRemainingDurationSec int
}

func (database *Database) GetEventSettings() (*EventSettings, error) {
	allEventSettings, err := database.eventSettingsTable.getAll()
	if err != nil {
		return nil, err
	}
	if len(allEventSettings) == 1 {
		return &allEventSettings[0], nil
	}

	// Database record doesn't exist yet; create it now.
	eventSettings := EventSettings{
		Name:                        "Untitled Event",
		ElimType:                    "single",
		NumElimAlliances:            8,
		SelectionRound2Order:        "L",
		SelectionRound3Order:        "",
		TBADownloadEnabled:          true,
		ApChannel:                   36,
		WarmupDurationSec:           game.MatchTiming.WarmupDurationSec,
		AutoDurationSec:             game.MatchTiming.AutoDurationSec,
		PauseDurationSec:            game.MatchTiming.PauseDurationSec,
		TeleopDurationSec:           game.MatchTiming.TeleopDurationSec,
		WarningRemainingDurationSec: game.MatchTiming.WarningRemainingDurationSec,
	}

	if err := database.eventSettingsTable.create(&eventSettings); err != nil {
		return nil, err
	}
	return &eventSettings, nil
}

func (database *Database) UpdateEventSettings(eventSettings *EventSettings) error {
	return database.eventSettingsTable.update(eventSettings)
}
