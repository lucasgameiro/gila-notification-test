/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  async rewrites() {
    return [
      {
        source: '/api/message',
        destination: process.env.BACKEND_ADDRESS + '/message',
      },
    ]
  },
  output: 'standalone',
}

module.exports = nextConfig
