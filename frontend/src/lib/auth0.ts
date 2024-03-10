export const getAuth0AuthorizeLink = () => {
  const url =
    `https://${process.env.AUTH0_DOMAIN}/authorize?` +
    `client_id=${process.env.AUTH0_CLIENT_ID}&` +
    `response_type=token&` +
    `redirect_uri=${process.env.APP_BASE_URL}/login/callback` // TODO: update with production URL
  return url
}

export const getAuth0LogoutLink = () => {
  const url =
    `https://${process.env.AUTH0_DOMAIN}/v2/logout?` +
    `client_id=${process.env.AUTH0_CLIENT_ID}&` +
    `returnTo=${process.env.APP_BASE_URL}/login` // TODO: update with production URL
  return url
}
