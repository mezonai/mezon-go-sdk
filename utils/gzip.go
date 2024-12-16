package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io"
)

// Compress compresses a string using Gzip and encodes it to Base64.
func GzipCompress(str string) (string, error) {
	var buf bytes.Buffer
	gzipWriter := gzip.NewWriter(&buf)

	// Write the string to the gzip writer
	_, err := gzipWriter.Write([]byte(str))
	if err != nil {
		return "", err
	}

	// Close the writer to flush the data
	if err := gzipWriter.Close(); err != nil {
		return "", err
	}

	// Encode the compressed data to Base64
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// Decompress decodes a Base64 string and decompresses it using Gzip.
func GzipUncompress(compressedStr string) (string, error) {
	// Decode the Base64 string
	compressedData, err := base64.StdEncoding.DecodeString(compressedStr)
	if err != nil {
		return "", err
	}

	// Create a gzip reader
	buf := bytes.NewReader(compressedData)
	gzipReader, err := gzip.NewReader(buf)
	if err != nil {
		return "", err
	}
	defer gzipReader.Close()

	// Decompress the data
	decompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		return "", err
	}

	return string(decompressedData), nil
}
