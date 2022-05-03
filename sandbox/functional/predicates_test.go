package functional

import (
	"testing"
)

func TestAllPredicate(t *testing.T) {
	testCases := []struct {
		Description    string
		WorkspaceData  Workspace
		DocumentType   string
		DocumentStatus string
		ExpectedResult bool
	}{
		{
			Description: "All mortgage documents are signed",
			WorkspaceData: Workspace{
				WorkspaceID:  1,
				SubscriberID: 1,
				RoleID:       1,
				Documents: []Document{
					{
						DocumentID:          1,
						DocumentStatus:      "Signed",
						DocumentType:        "National Mortgage",
						CreatedByRole:       1,
						CreatedBySubscriber: 1,
					},
					{
						DocumentID:          2,
						DocumentStatus:      "Signed",
						DocumentType:        "National Mortgage",
						CreatedByRole:       1,
						CreatedBySubscriber: 1,
					},
				},
			},
			DocumentType:   "National Mortgage",
			DocumentStatus: "Signed",
			ExpectedResult: true,
		},
		{
			Description: "Some mortgage documents are not signed",
			WorkspaceData: Workspace{
				WorkspaceID:  1,
				SubscriberID: 1,
				RoleID:       1,
				Documents: []Document{
					{
						DocumentID:          1,
						DocumentStatus:      "Signed",
						DocumentType:        "National Mortgage",
						CreatedByRole:       1,
						CreatedBySubscriber: 1,
					},
					{
						DocumentID:          2,
						DocumentStatus:      "Created",
						DocumentType:        "National Mortgage",
						CreatedByRole:       1,
						CreatedBySubscriber: 1,
					},
				},
			},
			DocumentType:   "National Mortgage",
			DocumentStatus: "Signed",
			ExpectedResult: false,
		},

		{
			Description: "No mortgage documents in workspace",
			WorkspaceData: Workspace{
				WorkspaceID:  1,
				SubscriberID: 1,
				RoleID:       1,
				Documents:    []Document{},
			},
			DocumentType:   "National Mortgage",
			DocumentStatus: "Signed",
			ExpectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			actualResult := all(testCase.WorkspaceData, testCase.DocumentType, testCase.DocumentStatus)
			if actualResult != testCase.ExpectedResult {
				t.Errorf("Test failed: expected [%v], got [%v]", testCase.ExpectedResult, actualResult)
			}
		})
	}

}
