package auth

import (
	//"github.com/tkanos/gonfig"
)

// User defines the information of the user to be added
// ProfilePicture: 
// GET https://api.linkedin.com/v2/me?projection=(id,localizedFirstName,localizedLastName,profilePicture(displayImage~:playableStreams))

// User describes the user info coming back from linkedin
type User struct {
	LocalizedLastName string `json:"localizedLastName"`
	ProfilePicture    struct {
		DisplayImage  string `json:"displayImage"`
		DisplayImages struct {
			Elements []struct {
				Artifact            string `json:"artifact"`
				AuthorizationMethod string `json:"authorizationMethod"`
				Data                struct {
					ComLinkedinDigitalmediaMediaartifactStillImage struct {
						StorageSize struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						} `json:"storageSize"`
						StorageAspectRatio struct {
							WidthAspect  int    `json:"widthAspect"`
							HeightAspect int    `json:"heightAspect"`
							Formatted    string `json:"formatted"`
						} `json:"storageAspectRatio"`
						MediaType    string `json:"mediaType"`
						RawCodecSpec struct {
							Name string `json:"name"`
							Type string `json:"type"`
						} `json:"rawCodecSpec"`
						DisplaySize struct {
							Uom    string `json:"uom"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"displaySize"`
						DisplayAspectRatio struct {
							WidthAspect  int    `json:"widthAspect"`
							HeightAspect int    `json:"heightAspect"`
							Formatted    string `json:"formatted"`
						} `json:"displayAspectRatio"`
					} `json:"com.linkedin.digitalmedia.mediaartifact.StillImage"`
				} `json:"data"`
				Identifiers []struct {
					Identifier                 string `json:"identifier"`
					File                       string `json:"file"`
					Index                      int    `json:"index"`
					MediaType                  string `json:"mediaType"`
					IdentifierType             string `json:"identifierType"`
					IdentifierExpiresInSeconds int    `json:"identifierExpiresInSeconds"`
				} `json:"identifiers"`
			} `json:"elements"`
			Paging struct {
				Count int           `json:"count"`
				Start int           `json:"start"`
				Links []interface{} `json:"links"`
			} `json:"paging"`
		} `json:"displayImage~"`
	} `json:"profilePicture"`
	ID                 string `json:"id"`
	LocalizedFirstName string `json:"localizedFirstName"`
}

// Email struc parses user email
type Email struct {
	Elements []struct {
		Handle string `json:"handle"`
		Handles struct {
			EmailAddress string `json:"emailAddress"`
		} `json:"handle~"`
	} `json:"elements"`
}
