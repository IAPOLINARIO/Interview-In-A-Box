# ---- Base Node ----
FROM node:14 AS base
WORKDIR /app
COPY package*.json ./

# ---- Dependencies ----
FROM base AS dependencies
RUN npm ci

# ---- Build ----
FROM base AS build
COPY --from=dependencies /app/node_modules ./node_modules
COPY . .
RUN npm run build

# ---- Release ----
FROM nginx:stable-alpine AS release
COPY --from=build /app/build /usr/share/nginx/html

# Expose the application port
EXPOSE 80

# Start the Nginx server
CMD ["nginx", "-g", "daemon off;"]
