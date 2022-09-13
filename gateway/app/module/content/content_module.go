package content

type ContentModule struct {
	ContentHandler *ContentHandler
	ContentService *ContentService
}

func (m *ContentModule) InitComponents() {
	m.ContentService = new(ContentService)
	m.ContentHandler = NewContentHandler(
		m.ContentService,
	)
}
