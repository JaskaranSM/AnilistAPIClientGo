package AnilistAPIClientGo

var AnimeQuery string = `
	query ($id: Int, $search: String) {
    Media(id: $id, type: ANIME, search: $search) {
      id
      isAdult
      title {
        romaji
        english
        native
      }
      description
      startDate {
        year
        month
        day
      }
      episodes
      season
      type
      format
      status
      duration
      isFavourite
      siteUrl
      studios {
        nodes {
          name
        }
      }
      tags {
        name
      }
      trailer {
        id
        site
        thumbnail
      }
      averageScore
      genres
      bannerImage
    }
  }
`

var CharacterQuery string = `
	query ($id: Int, $search: String) {
    Character(id: $id, search: $search) {
      id
      name {
        first
        last
        full
        native
      }
      siteUrl
      image {
        large
      }
      media {
        edges {
          characterRole
          node {
            id
            title {
              romaji
            }
          }
        }
      }
      description
    }
  }
`
