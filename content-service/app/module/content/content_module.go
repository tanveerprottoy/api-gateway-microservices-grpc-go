package content

type ContentModule struct {
	ContentRPC        *ContentRPC
	ContentService    *ContentService
	ContentRepository *ContentRepository
}

func (m *ContentModule) InitComponents() {
	m.ContentRepository = new(ContentRepository)
	m.ContentService = NewContentService(
		m.ContentRepository,
	)
	m.ContentRPC = NewContentRPC(
		m.ContentService,
	)
}
