package content

func Add(src map[string]string, id string, content string) {
	src[id] = content
}

func Update(src map[string]string, id string, content string) {
	src[id] = content
}

func Delete(src map[string]string, id string) {
	delete(src, id)
}
