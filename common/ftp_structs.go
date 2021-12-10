package common

type FtpConnectionConfig struct {
	Ftp      string `yaml:"ftp-ip"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type RecursiveLookForFileInFtpModule struct {
	RecursiveLookForFileInFtp RecursiveLookForFileInFtp `yaml:"recursive-look-for-file-in-ftp"`
}

type RecursiveLookForFileInFtp struct {
	Path string   `yaml:"path-to-look-for"`
	File []string `yaml:"file-to-look-for"`
}

type RetriveFileFromFtpModule struct {
	RetriveFileFromFtp RetriveFileFromFtp `yaml:"retrive-file-from-ftp"`
}

type RetriveFileFromFtp struct {
	Path     string   `yaml:"path-to-look-for"`
	FilesReg []string `yaml:"files-to-retrive"`
	Dest     string   `yaml:"dest-to-save-files"`
}

type LookForFileAndUploadToDestinationModule struct {
	LookForFileAndUploadToDestination LookForFileAndUploadToDestination `yaml:"look-for-file-and-upload-to-destination"`
}

type LookForFileAndUploadToDestination struct {
	Dest string `yaml:"file-dest-root-path"`
	Src  string `yaml:"file-src-root-path"`
}

type DeleteFileOnFtpModule struct {
	DeleteFileOnFtp DeleteFileOnFtp `yaml:"delete-file-on-ftp"`
}

type DeleteFileOnFtp struct {
	Path          string   `yaml:"path-to-look-for"`
	FilesToDelete []string `yaml:"files-to-delete"`
}

type FtpConfig struct {
	Ftp      string `yaml:"ftp-ip"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dest     string `yaml:"file-dest-root-path"`
	Src      string `yaml:"file-src-root-path"`
}
