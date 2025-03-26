package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Send[T any](cli *http.Client, req *http.Request, ret *T) error {
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("code: %d, err: %s", resp.StatusCode, string(respBody))
	}
	// fmt.Println(string(respBody))
	return json.Unmarshal(respBody, ret)
}
