package schema

type Message struct {
	Subject     string
	Body        string
	Attachments interface{}
	From        string
	To          []string

	FileData []byte
	FileName string
}
