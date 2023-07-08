###
# Builder Image
###
FROM node:18-alpine AS builder

# Set the working directory
WORKDIR /app

# Install dependencies
COPY package.json .
COPY package-lock.json .

RUN npm install --omit=dev

# Copy sources (see .dockerignore)
COPY . .

# Build source
RUN npm run build

###
# Production Image
###
FROM nginx:1.21.6-alpine AS app

# Set the working directory
WORKDIR /app

# Copy code
COPY --from=builder /app/build /usr/share/nginx/html

# Inform port that the container listens on
EXPOSE 3000

# Prefix commands and start production
ENTRYPOINT [ "nginx", "-g", "daemon off;" ]