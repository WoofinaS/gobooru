package gel

import (
	"testing"
)

const n = "\n"

var testClient = NewClient("", "") // Modify this to connect as an authenticated user.

func TestClient_SearchPosts(t *testing.T) {
	result, err := testClient.SearchPosts(PostFilter{})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	p0 := result.Posts[0]
	t.Log(n, p0.ID, n, p0.CreationDate, n, p0.Score, n, p0.Width, n,
		p0.Height, n, p0.MD5Hash, n, p0.Directory, n, p0.FileName, n, p0.Rating, n,
		p0.SourceURL, n, p0.Change, n, p0.Owner, n, p0.CreatorID, n, p0.ParentID, n,
		p0.Sample, n, p0.PreviewHeight, n, p0.PreviewWidth, n, p0.Tags, p0.HasNotes, n,
		p0.HasComments, n, p0.FileURL, n, p0.PreviewURL, n, p0.Status, n, p0.HasChildren)
}

func TestClient_SearchTags(t *testing.T) {
	result, err := testClient.SearchTags(TagFilter{})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t0 := result.Tags[0]
	t.Log(n, t0.ID, n, t0.Name, n, t0.Type, n, t0.Ambiguous, n, t0.Count)
}

func TestClient_SearchUsers(t *testing.T) {
	result, err := testClient.SearchUsers(UserFilter{})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	u0 := result.Users[0]
	t.Log(n, u0.ID, n, u0.Name, n, u0.Active)
}

func TestClient_SearchComments(t *testing.T) {
	result, err := testClient.SearchComments(CommentFilter{PostID: 7507608})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	c0 := result.Comments[0]
	t.Log(n, c0.ID, n, c0.PostID, n, c0.CreationDate, n, c0.Creator, n, c0.CreatorID)
}
