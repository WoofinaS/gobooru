package gel

type PostFilter struct {
	PostID    uint
	PostLimit uint
	PageNum   uint
	Tags      []string
	ChangeID  uint
}

type TagFilter struct {
	TagID       uint
	TagLimit    uint
	AfterID     uint
	Name        string
	Names       []string
	NamePattern string
	OrderBy     string
}

type UserFilter struct {
	UserLimit   uint
	PageNum     uint
	UserName    string
	NamePattern string
}

type CommentFilter struct {
	PostID uint
}

type postSearchResult struct {
	Posts  []post `xml:"post"`
	Limit  uint   `xml:"limit,attr"`
	Offset uint   `xml:"offset,attr"`
	Count  uint   `xml:"count,attr"`
}

type tagSearchResult struct {
	Tags   []tag `xml:"tag"`
	Limit  uint  `xml:"limit,attr"`
	Offset uint  `xml:"offset,attr"`
	Count  uint  `xml:"count,attr"`
}

type userSearchResult struct {
	Users  []user `xml:"user"`
	Limit  uint   `xml:"limit,attr"`
	Offset uint   `xml:"offset,attr"`
	Count  uint   `xml:"count,attr"`
}

type commentSearchResult struct {
	Comments []comment `xml:"comment"`
}

type post struct {
	ID            uint   `xml:"id"`
	CreationDate  string `xml:"created_at"`
	Score         int    `xml:"score"`
	Width         uint   `xml:"width"`
	Height        uint   `xml:"height"`
	MD5Hash       string `xml:"md5"`
	Directory     string `xml:"directory"`
	FileName      string `xml:"image"`
	Rating        string `xml:"rating"`
	SourceURL     string `xml:"source"`
	Change        uint   `xml:"change"`
	Owner         string `xml:"owner"`
	CreatorID     uint   `xml:"creator_id"`
	ParentID      uint   `xml:"parent_id"`
	Sample        uint   `xml:"sample"`
	PreviewHeight uint   `xml:"preview_height"`
	PreviewWidth  uint   `xml:"preview_width"`
	Tags          string `xml:"tags"`
	HasNotes      bool   `xml:"has_notes"`
	HasComments   bool   `xml:"has_comments"`
	FileURL       string `xml:"file_url"`
	PreviewURL    string `xml:"preview_url"`
	Status        string `xml:"status"`
	HasChildren   bool   `xml:"has_children"`
}

type tag struct {
	ID        uint   `xml:"id"`
	Name      string `xml:"name"`
	Count     uint   `xml:"count"`
	Type      uint   `xml:"type"`
	Ambiguous bool   `xml:"ambiguous"`
}

type user struct {
	ID     uint   `xml:"id"`
	Name   string `xml:"username"`
	Active bool   `xml:"active"`
}

type comment struct {
	ID           uint   `xml:"id,attr"`
	PostID       uint   `xml:"post_id,attr"`
	CreationDate string `xml:"created_at,attr"`
	Creator      uint   `xml:"creator,attr"`
	CreatorID    uint   `xml:"creator_id,attr"`
}
