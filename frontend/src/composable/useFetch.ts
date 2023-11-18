import { ofetch } from 'ofetch'
import md5 from 'blueimp-md5'

export function signature(body: string, path: string, timestamp: string) {
  const magic = 'b4afaa99988e236a65c27acb1b1962b2'
  return md5(`${magic}${path}${body}${timestamp}${magic}`)
}

export const tz = String(new Date().getTimezoneOffset());
export function useFetch() {
  return ofetch.create({
    baseURL: import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8090',
    // We set an interceptor for each request to
    // include Bearer token to the request if user is logged in
    onRequest: ({ options, request }) => {
      const timestamp = Date.now().toString()
      options.body = JSON.stringify(options.body)
      const sig = signature(options.body, request.toString(), timestamp);
      if (!options.method) {
        options.method = 'POST'
      }
      options.mode = 'cors'
      options.credentials = 'include'
      options.headers = {
        ...options.headers,
        'Content-Type': 'application/json',
        'X-Signature': `${timestamp},${sig}`,
      }
    },
  })
}
