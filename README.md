<div align="center"><i>CLI tool to post tweets on x (twitter) from terminal</i></div>

<br>

![preview](./assets/preview.gif)

## Table of Contents

- [Usage](#usage)
- [Installation](#installation)
- [Authentication & Setup](#authentication--setup)
- [Config File](#configuration)
- [Docs](#docs)

## Usage

### Quick Start
Post a tweet directly from your terminal:
```bash
x "first tweet from terminal!"
```

### Commands Overview

```bash
USAGE
  x <command>

COMMANDS
  init         Manage your Twitter Developer Apps (Add, Use, Delete)
  auth         Authenticate the active app with your Twitter account
  tweet (t)    Post a tweet
  version (v)  Show CLI version
```

## Installation

### Linux / macOS
1. Download the latest release binary.
2. Move it to your path:
   ```bash
   chmod +x x
   sudo mv x /usr/local/bin/
   ```

### Windows
1. Download the `x.exe` binary.
2. Add its location to your System PATH.

## Authentication & Setup

The CLI is now **local-first**. You need to bring your own Twitter Developer App credentials.

### 1. Create a Twitter App & Get Credentials
**Important:** Do NOT use the "Consumer Keys" or "Bearer Token". You need the **OAuth 2.0 Client ID and Secret**.

1. Go to the [Twitter Developer Portal](https://developer.twitter.com/en/portal/dashboard) and create a Free project/app.
2. Navigate to your App's **Settings**.
3. Under **User authentication settings**, click **Edit** (or Set up).
4. Configure the following:
   - **App Permissions**: Select "Read and write".
   - **Type of App**: Select "Native App".
   - **Redirect URL**: Enter `http://localhost:3000/callback`. **This is critical.**
   - **Website URL**: Enter any valid URL (e.g., `https://twitter.com`).
5. Click **Save**.
6. **IMMEDIATELY COPY** the **Client ID** and **Client Secret** shown on the screen.
   - *Note: These are different from the API Key and Secret found in the "Keys and Tokens" tab.*

### 2. Register App in CLI
Run the interactive init wizard:
```bash
x init add
```
Follow the prompts to enter your App Name, Client ID, and Client Secret.

### 3. Authenticate
Once added, authenticate your user account:
```bash
x auth
```
This will open your browser to authorize the app.

### 4. Manage Multiple Apps
You can add multiple apps and switch between them:
```bash
x init use     # Switch active app
x init delete  # Delete an app
```

## Configuration

Your data is stored locally in `~/.x-cli/config.json`. 

### Config Structure
```json
{
  "active_app": "MyBot",
  "apps": [
    {
      "name": "MyBot",
      "client_id": "YOUR_CLIENT_ID",
      "client_secret": "YOUR_CLIENT_SECRET",
      "user": {
        "access_token": "EncryptedOAuthToken...",
        "refresh_token": "EncryptedRefreshToken...",
        "expiry": "2024-05-20T15:04:05Z"
      }
    }
  ]
}
```

- **active_app**: The name of the app currently being used for tweets.
- **apps**: List of all registered apps.
  - **user**: Contains the OAuth tokens for the authenticated user. This is automatically populated/updated when you run `x auth` or when the CLI refreshes your token.

**Security Note:** This file contains sensitive credentials. Ensure your home directory is secure.

## Docs
[Docs](https://deepwiki.com/devhindo/x)
