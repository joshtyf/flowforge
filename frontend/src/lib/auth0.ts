export const getAuth0AuthorizeLink = () => {
  const url =
    `https://${process.env.AUTH0_DOMAIN}/authorize?` +
    `client_id=${process.env.AUTH0_CLIENT_ID}&` +
    `response_type=token&` +
    `redirect_uri=http://localhost:3000/login/callback` // TODO: update with production URL
  return url
}
