package processor

type Result struct {
	SearchWord      string
	TargetDirectory string
	Error           error
	FileSearch      []FileSearch
}
