# Google Sheets Service

This is a Go-based microservice for working with the Google Sheets API. The service provides RESTful endpoints to manage spreadsheets, create tables, and delete data. It is designed to be modular, scalable, and production-ready.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Setup and Installation](#setup-and-installation)
- [Usage](#usage)
- [Commands from Makefile](#commands-from-makefile)
- [API Endpoints](#api-endpoints)
- [Docker](#docker)

## Overview

This service allows you to integrate with Google Sheets via a structured HTTP API. It supports operations like creating new tables, deleting specific data ranges, and retrieving content from your spreadsheets.

## Features

- ✅ Create and manage tables in Google Sheets
- 🔄 Fetch and delete data from specific ranges
- 🚀 Graceful startup/shutdown with context and wait groups
- 🧩 Clean, modular architecture
- 📝 Structured logging with logrus
- ➕ Append new data (rows) to spreadsheets

## Setup and Installation

### Prerequisites

- Go 1.18 or higher installed
- A Google Cloud Platform (GCP) project
- A service account with access to Google Sheets API
- Enabled the Google Sheets API in your GCP project
- Downloaded the service account credentials JSON file
- Shared your target spreadsheet with the service account's email

### Installation Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/mickey-mickser/google-sheets-project.git
   cd .../google-sheets-project

## Usage

Once the service is running, you can interact with it using HTTP clients like curl or Postman. The API endpoints allow you to create new spreadsheets, update data, delete data, and read spreadsheet contents.


## Commands from Makefile
- Run: `make run`
- Build: `make build`
- Test: `make test`
- Build Container: `make build-container`

## API Endpoints
### Depending on the language you select in Google Sheet, the word "Sheet" in the value of the request body depends
**POST /create**  

We need to specify the `sheet_id` of the table. ??????????????

Example URL: `http://localhost:8080/create`

_Description:_ Creates a new table in a Google Sheets document.  
_Request Example:_
```json
{
  "properties": {
    "title": "My New Spreadsheet"
  }
}
```
_Response Example:_
```json
{
  "data": {
    "properties": {
      "autoRecalc": "ON_CHANGE",
      "defaultFormat": {
        "backgroundColor": {
          "blue": 1,
          "green": 1,
          "red": 1
        },
        "backgroundColorStyle": {
          "rgbColor": {
            "blue": 1,
            "green": 1,
            "red": 1
          }
        },
        "padding": {
          "bottom": 2,
          "left": 3,
          "right": 3,
          "top": 2
        },
        "textFormat": {
          "fontFamily": "arial,sans,sans-serif",
          "fontSize": 10,
          "foregroundColor": {},
          "foregroundColorStyle": {
            "rgbColor": {}
          }
        },
        "verticalAlignment": "BOTTOM",
        "wrapStrategy": "OVERFLOW_CELL"
      },
      "locale": "en_US",
      "spreadsheetTheme": {
        "primaryFontFamily": "Arial",
        "themeColors": [
          {
            "color": {
              "rgbColor": {}
            },
            "colorType": "TEXT"
          },
          {
            "color": {
              "rgbColor": {
                "blue": 1,
                "green": 1,
                "red": 1
              }
            },
            "colorType": "BACKGROUND"
          },
          {
            "color": {
              "rgbColor": {
                "blue": 0.95686275,
                "green": 0.52156866,
                "red": 0.25882354
              }
            },
            "colorType": "ACCENT1"
          },
          {
            "color": {
              "rgbColor": {
                "blue": 0.20784314,
                "green": 0.2627451,
                "red": 0.91764706
              }
            },
            "colorType": "ACCENT2"
          },
          {
            "color": {
              "rgbColor": {
                "blue": 0.015686275,
                "green": 0.7372549,
                "red": 0.9843137
              }
            },
            "colorType": "ACCENT3"
          },
          {
            "color": {
              "rgbColor": {
                "blue": 0.3254902,
                "green": 0.65882355,
                "red": 0.20392157
              }
            },
            "colorType": "ACCENT4"
          },
          {
            "color": {
              "rgbColor": {
                "blue": 0.003921569,
                "green": 0.42745098,
                "red": 1
              }
            },
            "colorType": "ACCENT5"
          },
          {
            "color": {
              "rgbColor": {
                "blue": 0.7764706,
                "green": 0.7411765,
                "red": 0.27450982
              }
            },
            "colorType": "ACCENT6"
          },
          {
            "color": {
              "rgbColor": {
                "blue": 0.8,
                "green": 0.33333334,
                "red": 0.06666667
              }
            },
            "colorType": "LINK"
          }
        ]
      },
      "timeZone": "Etc/GMT",
      "title": "My New Spreadsheet"
    },
    "sheets": [
      {
        "properties": {
          "gridProperties": {
            "columnCount": 26,
            "rowCount": 1000
          },
          "sheetType": "GRID",
          "title": "Sheet1"
        }
      }
    ],
    "spreadsheetId": "1a2b3c4d5e6f7g8h9i0j",
    "spreadsheetUrl": "https://docs.google.com/spreadsheets/d/1a2b3c4d5e6f7g8h9i0j/edit"
  },
  "status": "ok"
}
```
### POST /update
We need to specify the `sheet_id` of the table.

Example URL: `http://localhost:8080/update?sheet_id=1a2b3c4d5e6f7g8h9i0j`


_Description:_ Updates the cell or range of the table to the specified values.

**Request Example:**
```json

{
  "updates": [
    {
      "range": "Sheet1!A1:A10",
      "value": "Hello"
    },
    {
      "range": "Sheet1!B1",
      "value": "World"
    }
  ]
}

```
_Response Example:_
```json
{
  "data": {
    "responses": [
      {
        "spreadsheetId": "1a2b3c4d5e6f7g8h9i0j",
        "updatedCells": 1,
        "updatedColumns": 1,
        "updatedRange": "'Sheet1'!A1",
        "updatedRows": 1
      },
      {
        "spreadsheetId": "1a2b3c4d5e6f7g8h9i0j",
        "updatedCells": 1,
        "updatedColumns": 1,
        "updatedRange": "'Sheet1'!B1",
        "updatedRows": 1
      }
    ],
    "spreadsheetId": "1a2b3c4d5e6f7g8h9i0j",
    "totalUpdatedCells": 2,
    "totalUpdatedColumns": 2,
    "totalUpdatedRows": 1,
    "totalUpdatedSheets": 1
  },
  "status": "ok"
}
```
### POST /clearTable

We need to specify the `sheet_id` of the table.


Example URL: `http://localhost:8080/api/v1/sheets/clearTable?sheet_id=1a2b3c4d5e6f7g8h9i0j`

If the range is not specified in the JSON body, the default range will be `A1:Z1000`.


_Description:_ Deletes values in a specified range from a Google Sheet.

**Request Example:**
```json
{
  "range": "A1:C3"
}
```

_Response Example:_
```json
{
    "data": {
        "clearedRange": "'Sheet1'!A1:Z1000",
        "spreadsheetId": "1qPyWOoR5fX..."
    },
    "status": "ok"
}
```

### GET /read

We need to specify the `sheet_id` of the table.

_Description:_ Returns the specified cell range from a Google Sheet.


Example URL: `http://localhost:8080/read?sheet_id=1a2b3c4d5e6f7g8h9i0j`


_Response Example:_
```json
{
  "data": [
    [
      "Example"
    ]
  ],
  "status": "success"
}
```

### POST /share

No sheet_id required. Body only.

_Description:_ Grants read access to a specific user by email.


Example URL: `http://localhost:8080/api/v1/sheets/share`

**Request Example:**
```json
{
  "email": "user@example.com",
  "sendEmail": true,
  "emailMessage": "You now have read access"
}
```

_Response Example:_
```json
{
  "status": "ok",
  "data": { "id": "permission-id" }
}
```

### GET /capabilities

No sheet_id required. Body only.

_Description:_ Return of capabilities and description, requests and response of sheets-service.


Example URL: `http://localhost:8080/api/v1/sheets/capabilities`


_Response Example:_
```json
{
  "capabilities": "json from file"
}
```

## Docker
The provided Dockerfile uses a multi-stage build for efficiency:

Build Stage:
- Uses the official Go image (version defined by GO_VERSION) with support for cross-platform builds.

- Copies the entire project into the container.

- Caches Go modules and downloads dependencies.

- Compiles the project with CGO disabled for a fully static binary.

Final Stage:

- Uses a lightweight Alpine image.

- Copies the compiled binary and configuration files.

- Installs required packages (ca-certificates, tzdata) and updates certificates.

- Creates a non-root user (worker) for security.

- Exposes port 8055 and sets the entry point to run the service.

### Instructions

- 🐳 Running with Docker:

You can containerize and run the Google Sheets service using Docker. Below are the steps to build and run the application in a Docker container.

- 📦 Build the Docker Image
```bash
docker build -t google-sheets-project .
```
OR
```bash
make build-container
```

This command builds the Docker image using the Dockerfile in the project root directory. The resulting image will be named google-sheets-project.

- ▶️ Run the Container
```bash
sudo docker run -d \
--name sheets-service \
-p 8055:8055 \
-v $(pwd)/configs:/configs \
google-sheets-project
```
`-d`: Run container in background (detached mode)

`--name`: Give the container a name

`-p 8055:8055`: Map port 8055 in the container to 8095 on your host

`-v $(pwd)/configs:/configs`: Mount your local configs folder into the container

After this, your service will be available at `http://localhost:8055`.

- 🛑 Stop and Remove the Container

```bash
docker stop sheets-service && docker rm sheets-service
```
This will stop and remove the running container.

- 🧪 Run with Docker Compose (Optional)

If you prefer using docker-compose:
```bash
docker-compose up --build
```
This uses the docker-compose.yml file to build the image and run the container with networking and volume already configured.

- 📜 View Logs

```bash
docker logs -f sheets-service
```
- 🧭 Access the Container
  sudo docker-compose exec server /bin/sh

To enter the running container:
```bash
sudo docker exec -it sheets-service /bin/sh
```
or, if using docker-compose:
```bash
sudo docker exec -it sheets-service /bin/sh
```

