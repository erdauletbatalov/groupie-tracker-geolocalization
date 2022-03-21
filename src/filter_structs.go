package service

type CreationDate struct {
	From string
	To   string
}

type FirstAlbumDate struct {
	From string
	To   string
}

type NumberOfMembers struct {
	From string
	To   string
}

type Chechboxes struct {
	CDCheck        string
	FADCheck       string
	NOMCheck       string
	LocationsCheck string
}

type Inputs struct {
	CD         CreationDate
	FAD        FirstAlbumDate
	NOM        NumberOfMembers
	Loc        []string
	Chechboxes Chechboxes
}
