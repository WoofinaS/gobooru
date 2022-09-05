package e621

type Client string

type PostFilter struct {
	page  uint
	limit uint
	tags  []string
}

type postSearchResult struct {
	Posts []post `json:"posts,omitempty"`
}

type post struct {
	ID            uint          `json:"id,omitempty"`
	Created_at    string        `json:"created_at,omitempty"`
	Updated_at    string        `json:"Updated_at,omitempty"`
	File          file          `json:"file,omitempty"`
	Preview       preview       `json:"preview,omitempty"`
	Sample        sample        `json:"sample,omitempty"`
	Score         score         `json:"score,omitempty"`
	Tags          tags          `json:"tags,omitempty"`
	Locked_tags   []string      `json:"locked_tags,omitempty"`
	Change_seq    uint          `json:"change_seq,omitempty"`
	Flags         flags         `json:"flags,omitempty"`
	Rating        string        `json:"rating,omitempty"`
	Fav_count     uint          `json:"fav_count,omitempty"`
	Sources       string        `json:"sources,omitempty"`
	Pools         []uint        `json:"pools,omitempty"`
	Relationships relationships `json:"relationships,omitempty"`
	Approver_id   string        `json:"approver_id,omitempty"`
	Uploader_id   string        `json:"uploader_id,omitempty"`
	Description   string        `json:"description,omitempty"`
	Comment_count uint          `json:"comment_count,omitempty"`
	Is_favorited  bool          `json:"is_favorited,omitempty"`
}

type file struct {
	Width  uint   `json:"width,omitempty"`
	Height uint   `json:"height,omitempty"`
	Ext    string `json:"ext,omitempty"`
	Size   uint   `json:"size,omitempty"`
	MD5    string `json:"md5,omitempty"`
	Url    string `json:"url,omitempty"`
}

type preview struct {
	Width  uint   `json:"width,omitempty"`
	Height uint   `json:"height,omitempty"`
	Url    string `json:"url,omitempty"`
}

type sample struct {
	Has    bool   `json:"has,omitempty"`
	Width  uint   `json:"width,omitempty"`
	Height uint   `json:"height,omitempty"`
	Url    string `json:"url,omitempty"`
}

type score struct {
	Up    uint `json:"up,omitempty"`
	Down  uint `json:"down,omitempty"`
	Total int  `json:"total,omitempty"`
}

type tags struct {
	General   []string `json:"general,omitempty"`
	Species   []string `json:"species,omitempty"`
	Character []string `json:"character,omitempty"`
	Artist    []string `json:"artist,omitempty"`
	Invalid   []string `json:"invalid,omitempty"`
	Lore      []string `json:"lore,omitempty"`
	Meta      []string `json:"meta,omitempty"`
}

type flags struct {
	Pending       bool `json:"pending,omitempty"`
	Flagged       bool `json:"flagged,omitempty"`
	Note_locked   bool `json:"note_locked,omitempty"`
	Status_locked bool `json:"status_locked,omitempty"`
	Rating_locked bool `json:"rating_locked,omitempty"`
	Deleted       bool `json:"deleted,omitempty"`
}

type relationships struct {
	Parent_id           uint   `json:"parent_id,omitempty"`
	Has_children        bool   `json:"has_children,omitempty"`
	Has_active_children bool   `json:"has_active_children,omitempty"`
	Children            []uint `json:"children,omitempty"`
}
