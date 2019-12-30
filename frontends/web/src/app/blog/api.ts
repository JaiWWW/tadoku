import GhostContentAPI from '@tryghost/content-api'
import { Post } from './domain'

const api = GhostContentAPI({
  url: 'https://blog.tadoku.app',
  key: process.env.GHOST_KEY || '',
  version: 'canary',
})

const getPosts = async (): Promise<Post[]> => {
  const response = await api.posts.browse({
    limit: 5,
    include: ['authors', 'tags'],
    formats: ['html'],
  })

  if (!response) {
    return []
  }

  return Object.entries(response)
    .filter(([key]) => key !== 'meta')
    .map(([_, p]) => ({
      uuid: p.uuid,
      title: p.title,
      html: p.html,
    }))
}

const BlogApi = {
  get: getPosts,
}

export default BlogApi
