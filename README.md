# AI-calculator

## Overview

AI-calculator is an intelligent, modular calculator web application with a Go backend and a React frontend. It supports:
- Standard and advanced calculations
- Natural language and expression evaluation
- AI-powered calculation and reasoning using Google Gemini (with image+prompt support)

## Features
- **Modular Go backend** (Domain-Driven Design):
  - Calculator domain: arithmetic, expression evaluation
  - ConnectAI domain: AI-powered answers, image+prompt support
- **Frontend in React**: Modern, user-friendly UI
- **Gemini AI integration**: Use Google Gemini 1.5 Flash for intelligent answers
- **Image+Prompt**: Upload an image (e.g., handwritten math) and combine with a prompt
- **Configurable prompts**: Easily manage system prompts in YAML

## Architecture
```
backend/
  calculator/
    entity/      # Request/response types
    service/     # Business logic
    endpoint/    # Endpoint layer
    http/        # HTTP handlers
  connectai/
    entity/
    service/
    endpoint/
    http/
  main.go        # Server wiring
  prompts.yaml   # System prompts for AI
  .env           # Gemini API key
frontend/
  ...            # React app
```

## Setup

### Prerequisites
- Go 1.20+
- Node.js (for frontend)
- Gemini API key (get from Google AI Studio)

### Backend
1. Copy `.env.example` to `.env` and add your Gemini API key:
   ```
   GEMINI_API_KEY=your-gemini-api-key-here
   ```
2. Edit `prompts.yaml` to customize system prompts if needed.
3. Install Go dependencies:
   ```sh
   make install-backend
   ```
4. Start the backend:
   ```sh
   make start-backend
   ```

### Frontend
1. Install dependencies:
   ```sh
   make install-frontend
   ```
2. Start the frontend:
   ```sh
   make start-frontend
   ```

## Usage
- Open the app in your browser (usually at http://localhost:3000)
- Use the calculator for standard operations
- Click "Connect AI" to ask natural language questions or upload an image with a prompt
- The AI will respond using Gemini, following your prompt as the highest priority

## Project Highlights
- **Go kit/DDD inspired structure**: Each domain is self-contained and swappable
- **AI-first**: Easily extend to other LLMs or transports (gRPC, pubsub, etc.)
- **Image+Prompt**: Combines OCR and reasoning in one API call

## License
MIT