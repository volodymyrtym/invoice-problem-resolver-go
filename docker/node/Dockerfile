# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy only package.json and package-lock.json for layer caching of dependencies
COPY ./app/package.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files to the container
COPY ./app .

# Build the assets for production
RUN node esbuild.config.js

# By default, keep container alive unless overridden in docker-compose
CMD ["tail", "-f", "/dev/null"]