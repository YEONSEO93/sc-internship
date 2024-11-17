package folder

import "errors"

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...

	// * actual address (not a copy)
	var folderToMove *Folder
	var destination *Folder

	for i := range f.folders {
		if f.folders[i].Name == name {
			folderToMove = &f.folders[i]
		}
		if f.folders[i].Name == dst {
			destination = &f.folders[i]
		}
	}

	// if the folder to move is not found -> error
	if folderToMove == nil {
		return nil, errors.New("source folder not found")
	}

	// if the des folder is not found -> error
	if destination == nil {
		return nil, errors.New("destination folder not found")
	}

	// update the path
	folderToMove.Paths = destination.Paths + "." + folderToMove.Name

	return f.folders, nil
}
