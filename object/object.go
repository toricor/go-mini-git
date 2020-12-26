package object

import (
	"strconv"
	"strings"
)

type ObjectType int

const (
	Unknown ObjectType = iota
	Blob
	Commit
	Tree
	Tag
)

var objectTypeLabelMap = map[string]ObjectType{
	"blob":   Blob,
	"commit": Commit,
	"tree":   Tree,
	"tag":    Tag,
}

var objectTypeMap = map[ObjectType]string{
	Blob:   "blob",
	Commit: "commit",
	Tree:   "tree",
	Tag:    "tag",
}

func (t ObjectType) String() string {
	if o, ok := objectTypeMap[t]; ok {
		return o
	}
	return "Unknown ObjectType"
}

type GitObject struct {
	ObjectType ObjectType
	Size       uint64
	Content    string
}

func BuildGitObject(uncompressed string) GitObject {
	splited := strings.SplitN(uncompressed, "\x00", 2)
	header := splited[0]
	content := splited[1]
	objectType := strings.SplitN(header, " ", 2)[0]
	size, _ := strconv.ParseUint(strings.SplitN(header, " ", 2)[1], 10, 64)

	return GitObject{
		ObjectType: objectTypeLabelMap[objectType],
		Size:       size,
		Content:    content,
	}
}
