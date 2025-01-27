package git

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Links struct {
  Self string `json:"self"`
  Git  string `json:"git"`
  Html string `json:"html"`
}

type File struct {
  Name        string    `json:"name"`
  Path        string    `json:"path"`
  SHA         string    `json:"sha"`
  Size        int       `json:"size"`
  URL         string    `json:"url"`
  HTMLURL     string    `json:"html_url"`
  GitURL      string    `json:"git_url"`
  DownloadURL string    `json:"download_url"`
  Type        string    `json:"type"`
  Content     string    `json:"content"`
  Encoding    string    `json:"encoding"`
  Links       Links     `json:"_links"`
}

func GetFileContentFromGitHub(ctx context.Context, repoOwner string, repoName string, filePath string, fileName string, repoToken string) (string, error) {
  client := &http.Client{}

  url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s/%s", repoOwner, repoName, filePath, fileName)

  req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
  if err != nil {
    return "" , err
  }

  token := fmt.Sprintf("token %s", repoToken)

  req.Header.Add("Authorization", token)
  
  res, err := client.Do(req)
  if err != nil {
    return "", err
  }

  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return "", nil
  }

  var file File
  if err := json.Unmarshal(body, &file); err != nil {
    return "", nil
  }
  
  return file.Content, nil
}