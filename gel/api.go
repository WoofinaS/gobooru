// The full documentation of the gelbooru API can be seen at
// https://gelbooru.com/index.php?page=wiki&s=view&id=18780
package gel

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	postSearchPrefix    = "https://gelbooru.com/index.php?page=dapi&q=index&s=post"
	tagSearchPrefix     = "https://gelbooru.com/index.php?page=dapi&q=index&s=tag"
	userSearchPrefix    = "https://gelbooru.com/index.php?page=dapi&q=index&s=user"
	commentSearchPrefix = "https://gelbooru.com/index.php?page=dapi&q=index&s=comment"
)

// NewClient returns an Client that allows the user to interact with the gelbooru API.
// Authenticating a user account with both an API key and UserID may prevent API rate limiting,
// however, use without either is possible as well.
func NewClient(key, user string) Client {
	if len(key) == 0 || len(user) == 0 {
		return ""
	}

	return Client("&api_key=" + key + "&user_id=" + user)
}

// SearchPosts returns a postSearchResult that contains the parsed response to the gelbooru API call
// using the filters specifed in the passed PostFilter. Post searches have a hard limit of 100
// results.
func (c Client) SearchPosts(f PostFilter) (result *postSearchResult, err error) {
	if f.PostLimit > 100 {
		return result, errors.New("PostLimit can not be greater than 100")
	}

	var url strings.Builder
	url.WriteString(postSearchPrefix)
	url.WriteString(string(c))
	if f.PageNum > 0 {
		url.WriteString("&pid=")
		url.WriteString(strconv.Itoa(int(f.PageNum)))
	}
	if f.PostLimit > 0 {
		url.WriteString("&limit=")
		url.WriteString(strconv.Itoa(int(f.PostLimit)))
	}
	if f.ChangeID > 0 {
		url.WriteString("&cid=")
		url.WriteString(strconv.Itoa(int(f.ChangeID)))
	}
	if f.PostID > 0 {
		url.WriteString("&id=")
		url.WriteString(strconv.Itoa(int(f.PostID)))
	}
	for i, v := range f.Tags {
		if i > 0 {
			url.WriteByte('+')
		} else {
			url.WriteString("&tags=")
		}
		url.WriteString(v)
	}

	err = request(url.String(), &result)
	return
}

// SearchTags returns a tagSearchResult that contains the parsed response to the gelbooru API call
// using the filters specifed in the passed TagFilter. Tag searches have a hard limit of 1000
// results. NOTE: This may be a gelbooru API bug, as its docs suggest the limit is 100.
func (c Client) SearchTags(f TagFilter) (result *tagSearchResult, err error) {
	if f.TagLimit > 1000 {
		return result, errors.New("TagLimit can not be greater than 1000")
	}

	var url strings.Builder
	url.WriteString(tagSearchPrefix)
	url.WriteString(string(c))
	if f.TagID > 0 {
		url.WriteString("&id=")
		url.WriteString(strconv.Itoa(int(f.TagID)))
	}
	if f.TagLimit > 0 {
		url.WriteString("&limit=")
		url.WriteString(strconv.Itoa(int(f.TagLimit)))
	}
	if f.AfterID > 0 {
		url.WriteString("&after_id=")
		url.WriteString(strconv.Itoa(int(f.AfterID)))
	}
	if len(f.Name) > 0 {
		url.WriteString("&name=")
		url.WriteString(f.Name)
	}
	for i, v := range f.Names {
		if i > 0 {
			url.WriteByte('+')
		} else {
			url.WriteString("&names=")
		}
		url.WriteString(v)
	}
	if len(f.NamePattern) > 0 {
		url.WriteString("&name_pattern=")
		url.WriteString(f.NamePattern)
	}
	switch f.OrderBy {
	case "":
	case "date", "count", "name":
		url.WriteString("&orderby=")
		url.WriteString(f.OrderBy)
	default:
		return result, errors.New("invalid OrderBy " + f.OrderBy)
	}

	err = request(url.String(), &result)
	return
}

// SearchUsers returns a userSearchResult that contains the parsed response to the gelbooru API call
// using the filters specifed in the passed UserFilter. User searches have a hard limit of 100
// results.
func (c Client) SearchUsers(f UserFilter) (result *userSearchResult, err error) {
	if f.UserLimit > 100 {
		return result, errors.New("UserLimit can not be greater than 100")
	}

	var url strings.Builder
	url.WriteString(userSearchPrefix)
	url.WriteString(string(c))
	if f.UserLimit > 0 {
		url.WriteString("&limit=")
		url.WriteString(strconv.Itoa(int(f.UserLimit)))
	}
	if f.PageNum > 0 {
		url.WriteString("&pid=")
		url.WriteString(strconv.Itoa(int(f.PageNum)))
	}
	if len(f.UserName) > 0 {
		url.WriteString("&name=")
		url.WriteString(f.UserName)
	}
	if len(f.NamePattern) > 0 {
		url.WriteString("&name_pattern=")
		url.WriteString(f.NamePattern)
	}

	err = request(url.String(), &result)
	return
}

// SearchComments returns a commentSearchResult that contains the parsed response to the gelbooru
// API call using the filters specifed in the passed CommentFilter.
func (c Client) SearchComments(f CommentFilter) (result *commentSearchResult, err error) {
	if f.PostID == 0 {
		return result, errors.New("invalid PostID 0")
	}

	err = request(commentSearchPrefix+string(c)+"&post_id="+strconv.Itoa(int(f.PostID)), &result)
	return
}

// Request is an internal function.
func request(url string, v interface{}) error {
	// TODO: is there a better way to write this?
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return xml.Unmarshal(bytes, v)
}
