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
      relations {
        edges {
          relationType
          id
          node {
            id
            title {
              romaji
            }
          }
        }
      }
    }
  }
`

var AnimeQueryPaged string = `
query ($id: Int, $search: String) {
  Page(page: 1, perPage: 5) {
  media(id: $id, type: ANIME, search: $search) {
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
    relations {
      edges {
        relationType
        id
        node {
          id
          title {
            romaji
          }
        }
      }
    }
  }
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

var CharacterQueryPaged string = `
query ($id: Int, $search: String) {
  Page(page: 1, perPage: 5) {
  characters(id: $id, search: $search) {
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
}
`

var MangaQuery string = `
query ($id: Int, $search: String) {
  Media(id: $id, type: MANGA, search: $search) {
    id
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
    type
    format
    status
    siteUrl
    averageScore
    genres
    volumes
    chapters
    bannerImage
    isAdult
    isFavourite
    relations {
      edges {
        relationType
        id
        node {
          id
          title {
            romaji
          }
        }
      }
    }
  }
}
`

var MangaQueryPaged string = `
query ($id: Int, $search: String = "high school dxd") {
  Page(page: 1, perPage: 5) {
    media(id: $id, type: MANGA, search: $search) {
      id
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
      type
      format
      status
      siteUrl
      averageScore
      genres
      volumes
      chapters
      bannerImage
      isAdult
      isFavourite
      relations {
        edges {
          relationType
          id
          node {
            id
            title {
              romaji
            }
          }
        }
      }
    }
  }
}

`

var AnimeAiringQuery string = `
query ($id: Int, $search: String) {
  Media(id: $id, type: ANIME, search: $search) {
    id
    status
    episodes
    title {
      romaji
      english
      native
    }
    nextAiringEpisode {
      airingAt
      timeUntilAiring
      episode
    }
  }
}

`
