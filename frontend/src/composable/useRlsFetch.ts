import { ofetch } from 'ofetch'

export function useRlsFetch() {
  return ofetch.create({
    baseURL: process.env.NODE_ENV === 'production' ? 'http://172.16.0.200:8080' : '',
    // We set an interceptor for each request to
    // include Bearer token to the request if user is logged in
    onRequest: ({ options }) => {
      options.headers = {
        ...options.headers,
        'Content-Type': 'application/json'
      }
    },
  })
}
