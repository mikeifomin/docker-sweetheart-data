package testserver

type RespHealth struct {
	Status string
}
type ParamsHealth struct {
}
type ParamsAddFile struct {
	Filename string
	Contents string
}
type RespAddFile struct {
	IsCanged bool
}

type ParamsListFile struct {
}
type File struct {
	Filename string
	Contents string
}
type RespListFile struct {
	Files []File
}
	
