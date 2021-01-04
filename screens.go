package main

import "github.com/racingmars/go3270"

var loginScreen = go3270.Screen {
	{Row: 0, Col: 35, Intense: true, Content: "Route/3270"},
	{Row: 6, Col: 3, Content: "Username: "},
	{Row: 8, Col: 3, Content: "Password: "},
	{Row: 6, Col: 15, Name: "username", Write: true, Highlighting: go3270.Underscore},
	{Row: 6, Col: 30},
	{Row: 8, Col: 15, Name: "password", Write: true, Highlighting: go3270.Underscore, Hidden: true},
	{Row: 8, Col: 30},
	{Row: 12, Col: 3, Content: "Press "},
	{Row: 12, Col: 10, Content: "ENTER", Intense: true},
	{Row: 12, Col: 17, Content: "to log in."},
	{Row: 10, Col: 3, Intense: true, Color: go3270.Red, Name: "errormsg"},
	{Row: 22, Col: 0, Content: "PF3 Exit"},
}

var loginScreenRules = go3270.Rules{
	"username": {Validator: go3270.NonBlank},
	"password": {Validator: go3270.NonBlank, Reset: true},
}

var connectionErrorScreen = go3270.Screen{
	{Row: 0, Col: 35, Intense: true, Content: "Route/3270"},
	{Row: 4, Col: 3, Intense: true, Content: "Error", Color: go3270.Red},
	{Row: 6, Col: 3, Name: "errormsg"},
	{Row: 6, Col: 79},
	{Row: 22, Col: 0, Content: "PF3 Exit"},
	{Row: 20, Col: 1, Content: "Hostname: "},
	{Row: 20, Col: 40, Content: "Port"},
	{Row: 20, Col: 11, Name: "destination", Intense: true},
	{Row: 20, Col: 39},
	{Row: 20, Col: 46, Name: "port", Intense: true},
	{Row: 20, Col: 79},
}