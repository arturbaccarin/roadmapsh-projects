package json

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateFile(t *testing.T) {
	t.Run("create file without extension", func(t *testing.T) {
		filename := "test_filename"

		err := CreateFile(filename)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}

		defer os.Remove(filename + ".json")

		file, err := os.Open(filename + ".json")
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		defer file.Close()
	})

	t.Run("create file with extension", func(t *testing.T) {
		filename := "test_filename.json"

		err := CreateFile(filename)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}

		defer os.Remove(filename)

		file, err := os.Open(filename)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		defer file.Close()
	})

	t.Run("create file with wrong extension", func(t *testing.T) {
		filename := "<>/?``"

		err := CreateFile(filename)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}

func TestUpdateFile(t *testing.T) {
	tmpDir := t.TempDir()

	filename := filepath.Join(tmpDir, "testfile")
	data := []byte(`{"key":"value"}`)

	err := UpdateFile(filename, data)
	if err != nil {
		t.Fatalf("UpdateFile returned an error: %v", err)
	}

	expectedFile := verifyJSONExtension(filename)

	writtenData, err := os.ReadFile(expectedFile)
	if err != nil {
		t.Fatalf("failed to read written file: %v", err)
	}

	if string(writtenData) != string(data) {
		t.Errorf("file contents do not match\nexpected: %s\ngot: %s", string(data), string(writtenData))
	}
}

func TestFileExists(t *testing.T) {
	tmpFile, err := os.Create("testfile.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up

	if !FileExists(tmpFile.Name()) {
		t.Errorf("Expected file to exist: %s", tmpFile.Name())
	}

	nonExistent := tmpFile.Name() + "_doesnotexist"
	if FileExists(nonExistent) {
		t.Errorf("Expected file NOT to exist: %s", nonExistent)
	}
}

func TestLoadFile(t *testing.T) {
	t.Run("Valid file", func(t *testing.T) {
		tempFile, err := os.Create("testfile.json")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(tempFile.Name())

		// Write some data to the file
		_, err = tempFile.Write([]byte("Hello, World!"))
		if err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
		defer tempFile.Close()

		// Load the file
		data, err := LoadFile(tempFile.Name())
		if err != nil {
			t.Fatalf("Failed to load file: %v", err)
		}

		// Check the loaded data
		expectedData := []byte("Hello, World!")
		if !bytes.Equal(data, expectedData) {
			t.Errorf("Expected data '%s', got '%s'", expectedData, data)
		}

	})

	t.Run("File Not Found", func(t *testing.T) {
		_, err := LoadFile("nonexistent_file.json")
		if err == nil {
			t.Errorf("Expected error for non-existent file, got nil")
		}
	})
}
