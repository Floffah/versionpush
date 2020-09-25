package files

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"versionpush/src/util"
	"golang.org/x/sys/unix"
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
	if !Writable(final) {
		util.Fatal(fmt.Sprintf("File %v is not writable. Aborting...", final))
		os.Exit(1)
	}

	zipp, err := os.Create(final)
	if err != nil {
		util.Fatal("An error occured while trying to create .versionpush/resources.zip.")
		panic(err)
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
		util.Fatal(fmt.Sprintf("An error occured while trying to open file %v", file))
		panic(err)
	}
	defer zipping.Close()

	inf, err := zipping.Stat()
	if err != nil {
		util.Fatal(fmt.Sprintf("An error occured while trying to get statistics of file %v", file))
		panic(err)
	}

	hdr, err := zip.FileInfoHeader(inf)
	if err != nil {
		util.Fatal(fmt.Sprintf("An error occured while trying to create header for file %v to .versionpush/resources.zip", file))
		panic(err)
	}

	hdr.Name = file
	hdr.Method = zip.Deflate

	write, err := writer.CreateHeader(hdr)
	if err != nil {
		util.Fatal(fmt.Sprintf("An error occured while trying to create header for the writer of file %v to .versionpush/resources.zip", file))
		panic(err)
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

func Writable(path string) bool {
    return unix.Access(path, unix.W_OK) == nil
}