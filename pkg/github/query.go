package github

var githubQuery = `query {
	viewer {
	  pinnedRepositories(first: 10){
		totalCount
		nodes {
			id
			name
		  url
		  pushedAt
		  shortDescriptionHTML
		  defaultBranchRef {
			target {
			  ...on Commit {
				history(first:10) {
				  totalCount
				}
			  }
			}
		  }
		  repositoryTopics(first:10) {
			edges {
			  node {
				id
				topic {
				  name
				}
			  }
			}
		  }
		  languages(first: 10) {
			nodes {
			  color
			  id
			  name
			}
		  }
		  primaryLanguage {
			color
			id
			name
		  }
		}
	  }
	}
  }
`
