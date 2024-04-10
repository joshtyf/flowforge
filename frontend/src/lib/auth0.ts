import apiClient from "./apiClient"

export const getAuth0AuthorizeLink = () => {
  const url =
    `https://${process.env.NEXT_PUBLIC_AUTH0_DOMAIN}/authorize?` +
    `client_id=${process.env.NEXT_PUBLIC_AUTH0_CLIENT_ID}&` +
    `response_type=token&` +
    `redirect_uri=${process.env.NEXT_PUBLIC_APP_BASE_URL}/login/callback&` +
    `audience=${process.env.NEXT_PUBLIC_AUTH0_AUDIENCE}&` +
    `scope=openid%20profile%20email`

  return url
}

export const getAuth0LogoutLink = () => {
  const url =
    `https://${process.env.NEXT_PUBLIC_AUTH0_DOMAIN}/v2/logout?` +
    `client_id=${process.env.NEXT_PUBLIC_AUTH0_CLIENT_ID}&` +
    `returnTo=${process.env.NEXT_PUBLIC_APP_BASE_URL}/login` // TODO: update with production URL
  return url
}

export async function getUserProfile(accessToken: string) {
  return apiClient
    .get(`https://${process.env.NEXT_PUBLIC_AUTH0_DOMAIN}/userinfo`, {
      headers: { Authorization: `Bearer ${accessToken}` },
    })
    .then((res) => res.data)
}
