package tags

import (
	"h-ayat.github.io/tim/util"
)

func path() string { return util.BaseAddress() + "tags.txt" }

func LoadAllTags() []string {
	return util.ReadAll(path())
}

func AddTag(tag string) {
	util.Append(path(), tag)
}

func DeleteTagFile() {
	util.Delete(path())
}
