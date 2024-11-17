package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	// orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// res := folder.GetAllFolders()

	// // example usage
	// folderDriver := folder.NewDriver(res)
	// orgFolder := f#olderDriver.GetFoldersByOrgID(orgID)

	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)

	// Get sample data
	folders := folder.GetAllFolders()

	// Print all the folders
	fmt.Println("🌟 All Folders 🌟")
	printFolders(folders)

	// Get folders by OrgID
	orgID := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
	driver := folder.NewDriver(folders)
	foldersByOrg := driver.GetFoldersByOrgID(orgID)
	fmt.Printf("\n🌟 Folders for OrgID: %s 🌟\n", orgID)
	printFolders(foldersByOrg)

	// Get child folders
	parentFolder := "creative-scalphunter"
	childFolders := driver.GetAllChildFolders(orgID, parentFolder)
	fmt.Printf("\n🌟 Child Folders of '%s' 🌟\n", parentFolder)
	printFolders(childFolders)

	// Move a folder
	srcFolder := "creative-scalphunter"
	destFolder := "steady-insect"
	fmt.Printf("\n🚀 Moving folder '%s' to '%s'...\n", srcFolder, destFolder)
	updatedFolders, err := driver.MoveFolder(srcFolder, destFolder)
	if err != nil {
		fmt.Printf("❌ Error: %s\n", err)
	} else {
		fmt.Println("✅ Move successful. Updated folders:")
		printFolders(updatedFolders)
	}
}

// make it easy to read
func printFolders(folders []folder.Folder) {
	for _, f := range folders {
		fmt.Printf("🗂️ Name: %s, OrgID: %s, Path: %s\n", f.Name, f.OrgId, f.Paths)
	}
}

// go run main.go
