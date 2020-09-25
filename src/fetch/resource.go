package files

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"path/filepath"
	"versionpush/src/util"
)

func GetFinal(paths []string) (string, bool) {
	//if len(paths) == 1 {
	//	return paths[0], false
	//}
	var final string = filepath.Join(CWD(), ".versionpush", "resources.zip")

	if err := ZipEmUp(paths, final); err != nil {
		util.Fatal("An error occured while trying to compress resources into an archive.")
		panic(err)
	}

	return final, true
}

func ZipEmUp(paths []string, final string) error {
	zipp, err := os.Create(final)
	if err != nil {
		util.Fatal("An error occured while trying to create .versionpush/resources.zip.")
		panic(err)
	}

	writer := zip.NewWriter(zipp)

	for _, file := range paths {
		if err = AddToZip(writer, file); err != nil {
			return err
		}
	}

	zserr := zipp.Sync()
	zerr := zipp.Close()

	if zserr != nil {
		return zserr
	}
	if zerr != nil {
		return zerr
	}

	werr := writer.Close()
	if werr != nil {
		return werr
	}

	return nil
}

func AddToZip(writer *zip.Writer, file string) error {
	//zipping, err := os.Open(file)
	//if err != nil {
	//	util.Fatal(fmt.Sprintf("An error occured while trying to open file %v", file))
	//	panic(err)
	//}
	//defer zipping.Close()
	//
	//inf, err := zipping.Stat()
	//if err != nil {
	//	util.Fatal(fmt.Sprintf("An error occured while trying to get statistics of file %v", file))
	//	panic(err)
	//}
	//
	//hdr, err := zip.FileInfoHeader(inf)
	//if err != nil {
	//	util.Fatal(fmt.Sprintf("An error occured while trying to create header for file %v to .versionpush/resources.zip", file))
	//	panic(err)
	//}
	//
	//hdr.Name = file
	//hdr.Method = zip.Deflate
	//
	//write, err := writer.CreateHeader(hdr)
	//if err != nil {
	//	util.Fatal(fmt.Sprintf("An error occured while trying to create header for the writer of file %v to .versionpush/resources.zip", file))
	//	panic(err)
	//}
	//_, err = io.Copy(write, zipping)
	//return err

	f, err := writer.Create(file)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func CWD() string {
	dir, err := os.Getwd()
	if err != nil {
		util.Fatal("Could not find current working directory.")
		os.Exit(1)
	}
	return dir
}
