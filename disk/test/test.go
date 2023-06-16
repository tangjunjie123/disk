package main

type FileUploadPrepareReply struct {
	Identity string `json:"identity"`
	UploadId string `json:"upload_id"`
	Key      string `json:"key"`
}

func fun() (resp *FileUploadPrepareReply) {
	resp.Identity = "1111111111111111111111111111111111"
	return
}
func main() {
	fun()
}
