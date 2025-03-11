package task4

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckSpam(t *testing.T) {
	inputData := `+78962255222
+7 (612) 122-22-22
8 (812) 222-22-22
88962212222
849622222221`

	expectedOutput := `+78962255222 OK
+7 (612) 122-22-22 SPAM
8 (812) 222-22-22 SPAM
88962212222 OK
849622222221 ERROR
`

	inputFile, err := os.CreateTemp("", "input*.txt")
	require.NoError(t, err)
	defer os.Remove(inputFile.Name())

	outputFile, err := os.CreateTemp("", "output*.txt")
	require.NoError(t, err)
	defer os.Remove(outputFile.Name())

	_, err = inputFile.WriteString(inputData)
	require.NoError(t, err)
	inputFile.Close()

	CheckSpam(inputFile.Name(), outputFile.Name())

	outputBytes, err := os.ReadFile(outputFile.Name())
	require.NoError(t, err)
	outputContent := string(outputBytes)

	require.Equal(t, expectedOutput, outputContent)
}
