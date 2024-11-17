package folder

import (
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...

	childFolders := []Folder{}

	// for loof -> find all folders that match the OrgID & child folders
	for _, folder := range f.folders {
		if folder.OrgId == orgID && strings.HasPrefix(folder.Paths, name+".") {
			childFolders = append(childFolders, folder)
		}
	}

	return childFolders
}
