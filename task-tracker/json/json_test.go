package json

import (
	"bytes"
	"io"
	"os"
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
	filename := "test_filename.json"
	initialData := []byte("Initial data.\n")

	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("failed to create file: %v", err)
	}

	_, err = file.Write(initialData)
	if err != nil {
		t.Fatalf("failed to write to file: %v", err)
	}

	defer func() {
		file.Close()
		os.Remove(filename)
	}()

	t.Run("Valid append", func(t *testing.T) {
		appendData := []byte("New data added.\n")

		err = UpdateFile(filename, appendData)
		if err != nil {
			t.Fatalf("UpdateFile failed: %v", err)
		}

		file, err = os.Open(filename)
		if err != nil {
			t.Fatalf("failed to open file for reading: %v", err)
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		expectedContent := string(initialData) + "New data added.\n"
		if string(content) != expectedContent {
			t.Errorf("Expected file content '%s', got '%s'", expectedContent, string(content))
		}
	})

	// Error tests
	t.Run("File not found", func(t *testing.T) {
		nonExistentFile := "non_existent_file.json"

		err := UpdateFile(nonExistentFile, []byte("Some data"))
		if err == nil {
			t.Errorf("Expected error for non-existent file, got nil")
		}
	})
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
