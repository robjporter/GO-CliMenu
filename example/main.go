package main

import (
	"../"
	"fmt"
	"os"
)

func callback(id string) {
	fmt.Println("Chose item:", id)
}

func introTitle() {
	fmt.Println(`   _   ___ ___
  /_\ / __|_ _|
 / _ \ (__ | |
/_/ \_\___|___|  Simple Configuration Tool
==========================================`)
}

func newLine() {
	fmt.Println("")
}

func topMenu() string {
	climenu.Clear()
	introTitle()
	newLine()
	menu := climenu.NewButtonMenu("", "Choose an action")
	menu.AddMenuItem("Setup ACI System Connection", "connection")
	menu.AddMenuItem("Manage Tenants", "tenants")
	menu.AddMenuItem("Quit", "quit")
	action, escaped := menu.Run()
	if escaped || action == "quit" {
		os.Exit(0)
	}
	return action
}

func aciConnectionManagement() string {
	climenu.Clear()
	introTitle()
	newLine()
	menu := climenu.NewButtonMenu("", "Choose an action")
	menu.AddMenuItem("Set Destination Name/IP", "createTenant")
	menu.AddMenuItem("Set Username", "DeleteTenant")
	menu.AddMenuItem("Set Password", "Find Tenant")
	menu.AddMenuItem("Back", "back")
	action, escaped := menu.Run()
	if escaped {
		os.Exit(0)
	}
	return action
}

func connectionManagement(action string) {
	fmt.Println(action)
}

func aciTenants() string {
	climenu.Clear()
	introTitle()
	newLine()
	menu := climenu.NewButtonMenu("", "Choose an action")
	menu.AddMenuItem("Create Tenant", "createTenant")
	menu.AddMenuItem("Delete Tenant", "DeleteTenant")
	menu.AddMenuItem("Find Tenant", "Find Tenant")
	menu.AddMenuItem("Quit", "quit")
	action, escaped := menu.Run()
	if escaped {
		os.Exit(0)
	}
	return action
}

func tenantAction(action string) {
	fmt.Println(action)
}

func menu2() string {
	climenu.Clear()
	introTitle()
	newLine()
	menu := climenu.NewButtonMenu("", "Choose an action")
	menu.AddMenuItem("Create entry", "create")
	menu.AddMenuItem("Edit entry", "edit")
	menu.AddMenuItem("Back", "back")
	action, escaped := menu.Run()
	if escaped {
		os.Exit(0)
	}
	return action
}

func menu3() []string {
	climenu.Clear()
	introTitle()
	newLine()
	checkbox := climenu.NewCheckboxMenu("Let's try some checkboxes", "Select options", "OK", "Cancel")
	checkbox.AddMenuItem("Apples", "apples")
	checkbox.AddMenuItem("Oranges", "oranges")
	checkbox.AddMenuItem("Bananas", "bananas")
	checkbox.AddMenuItem("Back", "back")

	selection, escaped := checkbox.Run()
	if escaped {
		os.Exit(0)
	}
	return selection
}

func menus() {
	result := topMenu()

	if result == "connection" {
		result = aciConnectionManagement()
		if result == "back" {
			menus()
		} else {
			connectionManagement(result)
		}
	} else if result == "tenants" {
		result := aciTenants()
		if result == "back" {
			menus()
		} else {
			tenantAction(result)
		}
	}
}

func main() {
	for {
		menus()
	}
}
