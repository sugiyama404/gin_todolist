/** @type {import('next').NextConfig} */

const nextConfig = {
    env: {
        APSERVER_HOST: process.env.APSERVER_HOST
    },
    async rewrites() {
        return [
            {
                source: "/api/:path*",
                destination: `http://${process.env.APSERVER_HOST}:8080/api/:path*`,
            },
        ];
    },

}

module.exports = nextConfig
