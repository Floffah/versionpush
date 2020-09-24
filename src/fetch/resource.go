package files

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"versionpush/src/util"
)

func GetFinal(paths []string) (string, bool) {
	//if len(paths) == 1 {
	//	return paths[0], false
	//}
	var final string = path.Join(CWD(), ".versionpush", "resources.zip")

	if err := ZipEmUp(paths, final); err != nil {
		util.Fatal("An error occured while trying to compress resources into an archive.")
		os.Exit(1)
	}

	return final, true
}

func ZipEmUp(paths []string, final string) error {
	zipp, err := os.Create(final)
	if err != nil {
		util.Fatal("An error occured while trying to create .versionpush/resources.zip.")
		os.Exit(1)
	}

	defer zipp.Close()
	writer := zip.NewWriter(zipp)
	defer writer.Close()

	for _, file := range paths {
		if err = AddToZip(writer, file); err != nil {
			return err
		}
	}

	return nil
}

func AddToZip(writer *zip.Writer, file string) error {
	zipping, err := os.Open(file)
	if err != nil {
		util.Fatal(fmt.Sprintf("An error occurd while trying to open file %v", file))
		os.Exit(1)
	}
	defer zipping.Close()

	inf, err := zipping.Stat()
	if err != nil {
		util.Fatal(fmt.Sprintf("An error occurd while trying to get statistics of file %v", file))
		os.Exit(1)
	}

	hdr, err := zip.FileInfoHeader(inf)
	if err != nil {
		util.Fatal(fmt.Sprintf("An error occurd while trying to create header for file %v to .versionpush/resources.zip", file))
		os.Exit(1)
	}

	hdr.Name = file
	hdr.Method = zip.Deflate

	write, err := writer.CreateHeader(hdr)
	if err != nil {
		util.Fatal(fmt.Sprintf("An error occurd while trying to create header for the writer of file %v to .versionpush/resources.zip", file))
		os.Exit(1)
	}
	_, err = io.Copy(write, zipping)
	return err
}

func CWD() string {
	dir, err := os.Getwd()
	if err != nil {
		util.Fatal("Could not find current working directory.")
		os.Exit(1)
	}
	return dir
}