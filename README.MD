# MPEG-DASH Video Streaming Server in Go

A scalable MPEG-DASH video streaming server implemented in Go, following Clean Architecture and Service-Oriented Architecture (SOA) principles. Designed for maintainability, testability, and extensibility, this server delivers adaptive streaming video content using MPEG-DASH protocol.

## Features

- ✅ Serves MPEG-DASH manifests (.mpd) and segmented video chunks (.m4s)
- ✅ Adaptive bitrate streaming support
- ✅ Modular design with separation of concerns (domain, use cases, interfaces)
- ✅ RESTful API for video streaming requests
- ✅ Clean Architecture layers for easy maintenance and testing
- ✅ Support for video storage abstraction (local filesystem, cloud storage, etc.)
- ✅ Ready for integration with DRM and analytics services (extendable)

## Architecture Overview

```
┌────────────────────┐
│     Delivery       │ ← HTTP Handlers / API Layer (Interface adapters)
├────────────────────┤
│   Application      │ ← Use cases / Interactors (Business logic)
├────────────────────┤
│      Domain        │ ← Entities, business rules
├────────────────────┤
│  Infrastructure    │ ← DB, file storage, external services
└────────────────────┘
```

### Layer Responsibilities

- **Domain**: Core video entities (Video, Segment, Manifest) and interfaces
- **Application**: Use cases such as GenerateManifest, FetchSegment
- **Infrastructure**: Implementations for file storage, HTTP serving, logging
- **Delivery**: HTTP handlers that expose endpoints like `/videos/{id}/manifest.mpd`, `/videos/{id}/segment_{n}.m4s`

## Getting Started

### Prerequisites

- Go 1.20+
- FFmpeg (for video encoding and segment generation)
- Sample video files encoded and packaged as MPEG-DASH segments

### Installation

```bash
git clone https://github.com/yourusername/go-mpeg-dash-streaming.git
cd go-mpeg-dash-streaming
go mod download
```

### Generating MPEG-DASH Content (Example)

Use FFmpeg to encode and package your video:

```bash
ffmpeg -i input.mp4 \
  -map 0 \
  -codec:v libx264 -codec:a aac \
  -f dash \
  -seg_duration 4 \
  -dash_segment_filename "segment_$Number$.m4s" \
  manifest.mpd
```

Place the generated `manifest.mpd` and `.m4s` segments in your configured video storage folder.

### Configuration

Configure video storage path and server port in `config.yaml` or environment variables.

Example `config.yaml`:

```yaml
server:
  port: 8080

storage:
  path: "./videos"
```

### Running the Server

```bash
go run ./cmd/server
```

Server starts and listens on `http://localhost:8080`

## API Endpoints

| Endpoint                       | Description                    |
| ------------------------------ | ------------------------------ |
| `/videos/{id}/manifest.mpd`    | Serves MPEG-DASH manifest file |
| `/videos/{id}/segment_{n}.m4s` | Serves MPEG-DASH video segment |

### Example Requests

```http
GET http://localhost:8080/videos/sample/manifest.mpd
GET http://localhost:8080/videos/sample/segment_1.m4s
```

## Project Structure

```
├── cmd/                    # Entry point, main.go
├── internal/
│   ├── domain/            # Core entities and interfaces
│   ├── application/       # Use cases / business logic
│   ├── infra/             # File storage, HTTP server implementations
│   └── delivery/          # HTTP handlers / controllers
├── configs/               # Configuration files
├── scripts/               # Video packaging scripts, helpers
├── docs/                  # Documentation
└── test/                  # Unit and integration tests
```

## Extensibility

- **DRM Support**: Add by integrating with application layer use cases
- **Cloud Storage**: Support S3, GCS by implementing storage interfaces
- **Analytics**: Add metrics and analytics by injecting logging and monitoring services

## Testing

Run tests with:

```bash
go test ./...
```

## License

MIT License © [Timothy Kuo]
