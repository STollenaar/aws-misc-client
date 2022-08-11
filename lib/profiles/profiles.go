package profiles

import (
	"bufio"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	profileNameRe     = regexp.MustCompile(`\[|]`)
	awsAccessRe       = regexp.MustCompile(`aws_access_key_id\s*=\s*`)
	awsSecretAccessRe = regexp.MustCompile(`aws_secret_access_key\s*=\s*`)
	awsSessionRe      = regexp.MustCompile(`aws_session_token\s*=\s*`)
)

// GetProfiles - Returns list of profiles (no auth required)
func (c *ProfileClient) GetProfiles() ([]Profile, error) {
	profiles := []Profile{}
	credentialsFilePath := GetAWSCredentialsFilePath()
	file, err := os.Open(credentialsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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
			currentProfile.Name = profileNameRe.ReplaceAllString(currentLine, "")
		} else if strings.Contains(currentLine, "aws_access_key_id") {
			currentProfile.AWSAccessKeyId = awsAccessRe.ReplaceAllString(currentLine, "")
		} else if strings.Contains(currentLine, "aws_secret_access_key") {
			currentProfile.AWSSecretAccessKey = awsSecretAccessRe.ReplaceAllString(currentLine, "")
		} else if strings.Contains(currentLine, "aws_session_token") {
			currentProfile.AWSSessionToken = awsSessionRe.ReplaceAllString(currentLine, "")
		} else if currentLine == "" {
			profiles = append(profiles, *currentProfile)
			currentProfile = nil
		}
	}
	if currentProfile != nil {
		profiles = append(profiles, *currentProfile)
		currentProfile = nil
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
