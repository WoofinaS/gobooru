// The full documentation of the gelbooru api can be seen at
// https://gelbooru.com/index.php?page=wiki&s=view&id=18780
package gobooru

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Client string

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

type postSearchResults struct {
	Posts  []post `xml:"post"`
	Limit  uint   `xml:"limit,attr"`
	Offset uint   `xml:"offset,attr"`
	Count  uint   `xml:"count,attr"`
}

type tagSearchResults struct {
	Tags   []tag `xml:"tag"`
	Limit  uint  `xml:"limit,attr"`
	Offset uint  `xml:"offset,attr"`
	Count  uint  `xml:"count,attr"`
}

type userSearchResults struct {
	Users  []user `xml:"user"`
	Limit  uint   `xml:"limit,attr"`
	Offset uint   `xml:"offset,attr"`
	Count  uint   `xml:"count,attr"`
}

type commentSearchResults struct {
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

const (
	baseUrl           = "https://gelbooru.com/index.php?"
	basePostSearchUrl = baseUrl + "page=dapi&s=post&q=index"
	baseTagsSearchUrl = baseUrl + "page=dapi&s=tag&q=index"
	baseUserSearchUrl = baseUrl + "page=dapi&s=user&q=index"
	baseCommentsUrl   = baseUrl + "page=dapi&s=comment&q=index"
)

// NewClient returns an Client that allows the user to interact with the gelbooru API.
// Authenticating a user account with both an API key and UserID may prevent API rate limiting,
// however, use without either is possible as well.
func NewClient(key, user string) Client {
	if len(key) == 0 || len(user) == 0 {
		return ""
	}
	return Client(fmt.Sprint("&api_key=", key, "&user_id", user))
}

// SearchPosts returns a postSearchResults that contains the results of the gelbooru api
// call using the filters specifed in the passed PostFilter.
func (c Client) SearchPosts(filter PostFilter) (postSearchResults, error) {
	var url strings.Builder
	var results postSearchResults
	if filter.PostLimit > 100 {
		return postSearchResults{}, errors.New("post limit can not be greater then 100")
	}
	url.WriteString(basePostSearchUrl)
	url.WriteString(string(c))
	url.WriteString("&pid=")
	url.WriteString(strconv.Itoa(int(filter.PageNum)))
	url.WriteString("&limit=")
	url.WriteString(strconv.Itoa(int(filter.PostLimit)))
	url.WriteString("&cid=")
	url.WriteString(strconv.Itoa(int(filter.ChangeID)))
	url.WriteString("&id=")
	url.WriteString(strconv.Itoa(int(filter.PostID)))
	url.WriteString("&tags=")
	for i, v := range filter.Tags {
		if i > 0 {
			url.WriteByte('+')
		}
		url.WriteString(v)
	}
	if err := request(url.String(), &results); err != nil {
		return postSearchResults{}, err
	}
	return results, nil
}

// SearchTags returns a tagSearchResults that contains the results of the gelbooru api
// call using the filters specifed in the passed TagFilter.
func (c Client) SearchTags(filter TagFilter) (tagSearchResults, error) {
	var url strings.Builder
	var results tagSearchResults
	if filter.TagLimit > 100 {
		return tagSearchResults{}, errors.New("tag limit can not be greater then 100")
	}
	url.WriteString(basePostSearchUrl)
	url.WriteString(string(c))
	url.WriteString("&id=")
	url.WriteString(strconv.Itoa(int(filter.TagID)))
	url.WriteString("&limit=")
	url.WriteString(strconv.Itoa(int(filter.TagLimit)))
	url.WriteString("&after_id=")
	url.WriteString(strconv.Itoa(int(filter.AfterID)))
	url.WriteString("&name=")
	url.WriteString(filter.Name)
	url.WriteString("&names=")
	for i, v := range filter.Names {
		if i > 0 {
			url.WriteByte('+')
		}
		url.WriteString(v)
	}
	url.WriteString("&name_pattern=")
	url.WriteString(filter.NamePattern)
	url.WriteString("&orderby=")
	url.WriteString(filter.OrderBy)
	if err := request(url.String(), &results); err != nil {
		return tagSearchResults{}, err
	}
	return results, nil
}

// SearchUsers returns a userSearchResults that contains the results of the gelbooru api
// call using the filters specifed in the passed UserFilter.
func (c Client) SearchUsers(filter UserFilter) (userSearchResults, error) {
	var url strings.Builder
	var results userSearchResults
	if filter.UserLimit > 100 {
		return userSearchResults{}, errors.New("user limit can not be greater then 100")
	}
	url.WriteString(basePostSearchUrl)
	url.WriteString(string(c))
	url.WriteString("&limit=%d")
	url.WriteString(strconv.Itoa(int(filter.UserLimit)))
	url.WriteString("&pid=")
	url.WriteString(strconv.Itoa(int(filter.PageNum)))
	url.WriteString("&name=")
	url.WriteString(filter.UserName)
	url.WriteString("&name_pattern=")
	url.WriteString(filter.NamePattern)
	if err := request(url.String(), &results); err != nil {
		return userSearchResults{}, err
	}
	return results, nil
}

// SearchComments returns a commentSearchResults that contains the results of the gelbooru api
// call using the filters specifed in the passed CommentFilter.
func (c Client) SearchComments(filter CommentFilter) (commentSearchResults, error) {
	var url strings.Builder
	var results commentSearchResults
	if filter.PostID == 0 {
		return commentSearchResults{}, errors.New("invalid PostID 0")
	}
	url.WriteString(baseCommentsUrl)
	url.WriteString(string(c))
	url.WriteString("&post_id=")
	url.WriteString(strconv.Itoa(int(filter.PostID)))
	if err := request(url.String(), &results); err != nil {
		return commentSearchResults{}, err
	}
	return results, nil
}

// Requests a XML file from the internet and parses it
func request(url string, v any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(bytes, v)
	if err != nil {
		return err
	}
	return nil
}
