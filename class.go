package skillshare_dl

type ClassInfo struct {
	ID           int    `json:"id"`
	Sku          int    `json:"sku"`
	Title        string `json:"title"`
	ProjectTitle string `json:"project_title"`
	Embedded     struct {
		Teacher struct {
			Username       int         `json:"username"`
			FullName       string      `json:"full_name"`
			VanityUsername interface{} `json:"vanity_username"`
		} `json:"teacher"`
		Units struct {
			Embedded struct {
				Units []struct {
					Embedded struct {
						Sessions struct {
							Embedded struct {
								Sessions []Session `json:"sessions"`
							} `json:"_embedded"`
						} `json:"sessions"`
					} `json:"_embedded"`
				} `json:"units"`
			} `json:"_embedded"`
		} `json:"units"`
	} `json:"_embedded"`
}

func (ci *ClassInfo) AllSessions() []*Session {
	var sessions []*Session
	for i := range ci.Embedded.Units.Embedded.Units {
		for j := range ci.Embedded.Units.Embedded.Units[i].Embedded.Sessions.Embedded.Sessions {
			sessions = append(sessions, &ci.Embedded.Units.Embedded.Units[i].Embedded.Sessions.Embedded.Sessions[j])
		}
	}
	return sessions
}
