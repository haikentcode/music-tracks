# music-tracks

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Overview

This project is a simple API built with Go and Gin that interacts with the Spotify API to store and retrieve album metadata using ISRC codes. ISRC (International Standard Recording Code) is a unique identifier for audio and music video recordings.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#music-tracks-api)
- [Project Structure](#project-structure)

## Installation

Provide instructions on how to install and set up your project. Include any dependencies that need to be installed and steps for configuration.

```bash
# Clone the repository
git clone https://github.com/haikentcode/music-tracks.git

# Navigate to the project directory
cd music-tracks

# Install dependencies
go mod download

# Set your Spotify API credentials as environment variables
export SPOTIFY_CLIENT_ID=your_client_id
export SPOTIFY_CLIENT_SECRET=your_client_secret


# Build and run the project
go run main.go

```

# Music Tracks API

This API provides endpoints to interact with music tracks.

## Endpoints

### Get All Tracks

**Endpoint:** `GET /tracks`

**Description:**
Get a list of all music tracks.

### Create Track by ISRC

**Endpoint:** `PUT /create-track/:isrc`

**Description:**
Create a new music track using its International Standard Recording Code (ISRC).

**Parameters:**

- `isrc`: The International Standard Recording Code for the track.

### Get Track by ISRC

**Endpoint:** `GET /tracks/:isrc`

**Description:**
Get details about a specific music track based on its International Standard Recording Code (ISRC).

**Parameters:**

- `isrc`: The International Standard Recording Code for the track.

### Get Tracks by Artist

**Endpoint:** `GET /tracks/by-artist/:artist`

**Description:**
Get a list of music tracks associated with a specific artist.

**Parameters:**

- `artist`: The name of the artist.

## Usage

To use this API, make HTTP requests to the specified endpoints using your preferred HTTP client.

### Example:

```bash
export API_HOST="http://localhost:8000" # your-api-host
# Get all tracks
curl -X GET ${API_HOST}/v1/tracks

# Create a track by ISRC
curl -X PUT ${API_HOST}/v1/create-track/your-isrc-code

# Get track by ISRC
curl -X GET ${API_HOST}/v1/tracks/your-isrc-code

# Get tracks by artist
curl -X GET ${API_HOST}/v1/tracks/by-artist/your-artist-name
```

## Project-Structure

```
.
├── README.md
├── db
│   └── database.go
├── go.mod
├── go.sum
├── handlers
│   ├── spotify_handler.go
│   └── track_handler.go
├── index.html
├── main.go
├── models
│   └── track.go
├── pkg
│   └── spotify
│       └── spotify.go
├── routes
│   └── routes.go
└── test.db
```
