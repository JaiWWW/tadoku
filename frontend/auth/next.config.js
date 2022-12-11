/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  publicRuntimeConfig: {
    kratosPublicEndpoint: process.env.NEXT_PUBLIC_KRATOS_ENDPOINT,
    kratosInternalEndpoint: process.env.NEXT_PUBLIC_KRATOS_INTERNAL_ENDPOINT,
  }
}

const withTM = require('next-transpile-modules')(['tadoku-ui']);

module.exports = withTM(nextConfig)
