/** @type {import('next').NextConfig} */
const nextConfig = {
  output: "standalone",
  async redirects() {
    return [
      {
        source: "/",
        destination: "/service-catalog",
        permanent: true,
      },
      {
        source: "/settings",
        destination: "/settings/organization",
        permanent: true,
      },
    ]
  },
}

module.exports = nextConfig
