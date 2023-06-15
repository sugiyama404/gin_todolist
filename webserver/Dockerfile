# For development phase
FROM node:18.15-buster-slim as development

WORKDIR /usr/src/app/myapp
ENV CHOKIDAR_USEPOLLING=true
ENV NODE_ENV=development
ENV APSERVER_HOST=apserver

COPY ./myapp/package*.json ./
RUN npm install
COPY ./myapp .

# build phase
FROM node:18.15-buster-slim as build
WORKDIR /app
COPY ./myapp/package*.json ./
RUN npm ci
COPY ./myapp .
RUN npm run build

# For staging
FROM gcr.io/distroless/nodejs18-debian11:nonroot as staging
WORKDIR /app

ENV NODE_ENV production
ENV PORT 3000

COPY --from=build /app/next.config.js ./
COPY --from=build /app/public ./public
COPY --from=build /app/.next ./.next
COPY --from=build /app/node_modules ./node_modules

CMD ["./node_modules/next/dist/bin/next", "start"]
