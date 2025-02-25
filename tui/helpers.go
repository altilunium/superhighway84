package tui

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"


	"github.com/mrusme/superhighway84/models"
)

func MillisecondsToDate(ms int64) (string) {
  return time.Unix(0, ms * int64(time.Millisecond)).Format("Mon Jan _2 15:04:05 2006")
}

func (t *TUI) OpenArticle(article *models.Article) (models.Article, error) {
  if editor, exist := os.LookupEnv("EDITOR"); exist == false || editor == "" {
    return *article, errors.New("EDITOR environment variable not available, please export!")
  }

  tmpFile, err := ioutil.TempFile(os.TempDir(), "article-*.txt")
  if err != nil {
    return *article, err
  }

  

  tmpContent := []byte(fmt.Sprintf(
    "Subject: %s\nNewsgroup: %s\n= = = = = =\n%s",
    article.Subject, article.Newsgroup, article.Body))
  if _, err = tmpFile.Write(tmpContent); err != nil {
    return *article, err
  }

  if err := tmpFile.Close(); err != nil {
    return *article, err
  }





/*
  wasSuspended := t.App.Suspend()

*/

/*
  if wasSuspended == false {
    return *article, err
  }
*/
  tmpContent, err = os.ReadFile(tmpFile.Name())
  if err != nil {
    return *article, err
  }

  content := strings.SplitAfterN(string(tmpContent), "\n= = = = = =\n", 2)
  if len(content) != 2 {
    return *article, errors.New("Document malformatted")
  }

  newArticle := *article

  headerPart := strings.TrimSpace(content[0])
  headers := strings.Split(headerPart, "\n")

  for _, header := range headers {
    splitHeader := strings.SplitAfterN(header, ":", 2)
    if len(splitHeader) < 2 {
      continue
    }

    headerName := strings.ToLower(strings.TrimSpace(splitHeader[0]))
    headerValue := strings.TrimSpace(splitHeader[1])

    switch(headerName) {
    case "subject:":
      newArticle.Subject = headerValue
    case "newsgroup:":
      newArticle.Newsgroup = headerValue
    }
  }

  newArticle.Body = strings.TrimSpace(content[1])

  if valid, err := newArticle.IsValid(); valid == false {
    return *article, err
  }



go func() {
    //cmd := exec.Command(os.Getenv("EDITOR"), tmpFile.Name())
    cmd := exec.Command("notepad", tmpFile.Name())
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    err := cmd.Run()
    if err != nil {
      log.Println(err)
    }
    return
  }()

  return newArticle, nil
}


func (t *TUI) OpenArticle2(article *models.Article) (models.Article, error) {
  if editor, exist := os.LookupEnv("EDITOR"); exist == false || editor == "" {
    return *article, errors.New("EDITOR environment variable not available, please export!")
  }

  tmpFile, err := ioutil.TempFile(os.TempDir(), "article-*.txt")
  if err != nil {
    return *article, err
  }

  

  tmpContent := []byte(fmt.Sprintf(
    "Subject: %s\nNewsgroup: %s\n= = = = = =\n%s",
    article.Subject, article.Newsgroup, article.Body))
  if _, err = tmpFile.Write(tmpContent); err != nil {
    return *article, err
  }

  if err := tmpFile.Close(); err != nil {
    return *article, err
  }


finished := make(chan bool)
go func(finished chan bool) {
    //cmd := exec.Command(os.Getenv("EDITOR"), tmpFile.Name())
    cmd := exec.Command("notepad", tmpFile.Name())
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    err := cmd.Run()
    if err != nil {
      log.Println(err)
    }
    finished <- true
    return
  }(finished)


<- finished
/*
  wasSuspended := t.App.Suspend()

*/

/*
  if wasSuspended == false {
    return *article, err
  }
*/


  tmpContent, err = os.ReadFile(tmpFile.Name())
  if err != nil {
    return *article, err
  }

  content := strings.SplitAfterN(string(tmpContent), "\n= = = = = =\n", 2)
  if len(content) != 2 {
    return *article, errors.New("Document malformatted")
  }

  newArticle := *article

  headerPart := strings.TrimSpace(content[0])
  headers := strings.Split(headerPart, "\n")

  for _, header := range headers {
    splitHeader := strings.SplitAfterN(header, ":", 2)
    if len(splitHeader) < 2 {
      continue
    }

    headerName := strings.ToLower(strings.TrimSpace(splitHeader[0]))
    headerValue := strings.TrimSpace(splitHeader[1])

    switch(headerName) {
    case "subject:":
      newArticle.Subject = headerValue
    case "newsgroup:":
      newArticle.Newsgroup = headerValue
    }
  }

  newArticle.Body = strings.TrimSpace(content[1])

  if valid, err := newArticle.IsValid(); valid == false {
    return *article, err
  }





  return newArticle, nil
}



