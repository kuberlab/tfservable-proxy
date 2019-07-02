package tfhttp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func chunker(body io.ReadCloser, boundary string, pubChan chan []byte, stopChan chan bool) {
	fmt.Print("Chunker: starting\n")

	reader := bufio.NewReader(body)
	defer body.Close()

	defer close(pubChan)
	defer close(stopChan)

	var failure error

ChunkLoop:
	for {
		head, size, err := readChunkHeader(reader, boundary)
		if err != nil {
			failure = err
			break ChunkLoop
		}

		data, err := readChunkData(reader, size)
		if err != nil {
			failure = err
			break ChunkLoop
		}

		select {
		case <-stopChan:
			break ChunkLoop
		case pubChan <- append(head, data...):
		}

		if size == 0 {
			failure = errors.New("received final chunk")
			break ChunkLoop
		}
	}

	if failure != nil {
		fmt.Printf("Chunker: %s\n", failure)
	}

	fmt.Print("Chunker: stopping\n")
}

func readChunkHeader(reader *bufio.Reader, boundary string) (head []byte, size int, err error) {
	head = make([]byte, 0)
	size = -1
	err = nil

	// read boundary
	var line []byte
	line, err = reader.ReadSlice('\n')
	if err != nil {
		return
	}
	if bl := strings.TrimRight(string(line), "\r\n"); bl != boundary {
		err_str := fmt.Sprintf("Invalid boundary received (%s)", bl)
		err = errors.New(err_str)
		return
	}
	head = append(head, line...)

	// read header
	for {
		line, err = reader.ReadSlice('\n')
		if err != nil {
			return
		}
		head = append(head, line...)

		// empty line marks end of header
		line_str := strings.TrimRight(string(line), "\r\n")
		if len(line_str) == 0 {
			break
		}

		// find data size
		parts := strings.SplitN(line_str, ": ", 2)
		if strings.EqualFold(parts[0], "Content-Length") {
			var n int
			n, err = strconv.Atoi(string(parts[1]))
			if err != nil {
				return
			}
			size = n
		}
	}

	if size == -1 {
		err = errors.New("Content-Length chunk header not found")
		return
	}

	return
}

func readChunkData(reader *bufio.Reader, size int) (buf []byte, err error) {
	buf = make([]byte, size)
	err = nil

	pos := 0
	for pos < size {
		var n int
		n, err = reader.Read(buf[pos:])
		if err != nil {
			return
		}

		pos += n
	}

	return
}

func getBoundary(resp http.Response) (string, error) {
	ct := resp.Header.Get("Content-Type")
	prefix := "multipart/x-mixed-replace; boundary="
	if !strings.HasPrefix(ct, prefix) {
		err_str := fmt.Sprintf("Content-Type is invalid (%s)", ct)
		return "", errors.New(err_str)
	}

	boundary := strings.TrimPrefix(ct, prefix)
	return boundary, nil
}

func connectChunker(url, username, password string) (*http.Response, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		err_str := fmt.Sprintf("Request failed (%s)", resp.Status)
		return nil, "", errors.New(err_str)
	}

	boundary, err := getBoundary(*resp)
	if err != nil {
		resp.Body.Close()
		return nil, "", err
	}

	return resp, boundary, nil
}
