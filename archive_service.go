package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/julienschmidt/httprouter"
)

const fileStorageDir = "test_photos"

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	res.Header().Add("Content-Type", "text/html")
	fmt.Fprint(res, string(body))
}

func Archivate(res http.ResponseWriter, path string) error {
	bashCmd := fmt.Sprintf("zip -r - %s -j", path)
	cmd := exec.Command("bash", "-c", bashCmd)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return err
	}

	buff := make([]byte, 8192)

	for err == nil {
		_, err := stdout.Read(buff)
		if err != nil {
			break
		}

		_, e := res.Write(buff)
		if e != nil {
			cmd.Process.Kill()
			cmd.Wait()
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
		return err
	} else {
		log.Println("Ok")
	}
	return nil
}

func ArchiveHandler(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	archiveDir := ps.ByName("archive_hash")
	archivePath := path.Join(fileStorageDir, archiveDir)

	if _, err := os.Stat(archivePath); err != nil {
		if os.IsNotExist(err) {
			http.NotFound(res, req)
		}
	} else {
		res.Header().Set("Content-Type", "application/zip")
		res.Header().Set("Content-Disposition", "attachment; filename=archive.zip")

		err := Archivate(res, archivePath)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/archive/:archive_hash/", ArchiveHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
