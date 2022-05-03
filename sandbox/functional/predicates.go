package functional

type Workspace struct {
	WorkspaceID  int
	SubscriberID int
	RoleID       int
	Documents    []Document
}

type Document struct {
	CreatedBySubscriber int
	CreatedByRole       int
	DocumentID          int
	DocumentType        string
	DocumentStatus      string
}

func all(workspace Workspace, documentType string, documentStatus string) bool {
	// Find all documents created by subscriber ID that have typue documentType
	var documentsBySubscriber []Document

	for _, doc := range workspace.Documents {
		if doc.CreatedBySubscriber == workspace.SubscriberID && doc.CreatedByRole == workspace.RoleID && doc.DocumentType == documentType {
			documentsBySubscriber = append(documentsBySubscriber, doc)
		}
	}

	if len(documentsBySubscriber) == 0 {
		return false
	}

	for _, doc := range documentsBySubscriber {
		if doc.DocumentStatus != documentStatus {
			return false
		}
	}

	return true
}
