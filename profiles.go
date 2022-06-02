package awsprofilerclient

import (
	"bufio"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

// GetCoffees - Returns list of coffees (no auth required)
func (c *Client) GetProfiles() ([]Profile, error) {
	profiles := []Profile{}
	credentialsFilePath := GetAWSCredentialsFilePath()
	file, err := os.Open(credentialsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var re = regexp.MustCompile(`\[]`)

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example
	var currentProfile *Profile
	for scanner.Scan() {
		currentLine := scanner.Text()
		if strings.Contains(currentLine, "[") {
			if currentProfile == nil {
				currentProfile = new(Profile)
			} else {
				profiles = append(profiles, *currentProfile)
				currentProfile = new(Profile)
			}
			currentProfile.Name = re.ReplaceAllString(currentLine, "")
		} else if strings.Contains(currentLine, "aws_access_key_id") {
			currentProfile.AWSAccessKeyId = strings.ReplaceAll(currentLine, "aws_access_key_id=", "")
		} else if strings.Contains(currentLine, "aws_secret_access_key") {
			currentProfile.AWSSecretAccessKey = strings.ReplaceAll(currentLine, "aws_secret_access_key=", "")
		} else if strings.Contains(currentLine, "aws_session_token") {
			currentProfile.AWSSessionToken = strings.ReplaceAll(currentLine, "aws_session_token=", "")
		} else if currentLine == "" {
			profiles = append(profiles, *currentProfile)
			currentProfile = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return profiles, nil
}

func GetAWSCredentialsFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return path.Join(homeDir, ".aws", "credentials")
}
