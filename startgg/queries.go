package startgg

var EventSets = `
query EventSets($eventId: ID!, $page: Int!, $perPage: Int!) {
	event(id: $eventId) {
	  id
	  name
	  sets(page: $page, perPage: $perPage, sortType: STANDARD) {
		pageInfo {
		  total
		}
		nodes {
		  id
		  slots {
			id
			entrant {
			  id
			  name
			}
		  }
		}
	  }
	}
  }
  
`

var EventName = `
query EventName($eventId: ID) {
	event(id: $eventId) {
	  id
	  name
	}
  }
  
`

var TournamentIdFromSlug = `
query TournamentIdFromSlug($slug: String) {
	tournament(slug: $slug){
		id
		name
	}
}`
