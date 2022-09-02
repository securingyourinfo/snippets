func ListFilesInDir(path string) ([]string, error) {
	pfiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	ffiles := make([]string, len(pfiles))
	for _, f := range pfiles {
		ffiles = append(ffiles, f.Name())
	}
	return ffiles, err
}
