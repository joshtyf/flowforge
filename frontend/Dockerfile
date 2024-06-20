FROM node:18-alpine as builder
WORKDIR /frontend

COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM node:18-alpine as runner
WORKDIR /frontend
COPY --from=builder /frontend/package.json .
COPY --from=builder /frontend/package-lock.json .
COPY --from=builder /frontend/next.config.js ./
COPY --from=builder /frontend/public ./public
COPY --from=builder /frontend/.next/standalone ./.next/standalone
COPY --from=builder /frontend/.next/static ./.next/static
COPY --from=builder /frontend/public ./.next/standalone/public
COPY --from=builder /frontend/.next/static ./.next/standalone/.next/static
EXPOSE 3000
ENTRYPOINT ["npm", "run", "start"]