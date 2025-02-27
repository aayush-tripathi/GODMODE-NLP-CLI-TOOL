# NLP-to-Bash CLI Tool

## Overview

This repository provides an end-to-end solution for translating natural language instructions into Bash commands. The project leverages a fine-tuned T5-small model to perform the translation and is split into two main components:

- **Backend:** A Python-based Flask API that serves the NLP model.
- **Frontend:** A Go-based Command Line Interface (CLI) tool built with Cobra, which sends queries to the backend and displays the generated Bash commands.

This tool aims to simplify command-line interactions, making it easier for users who may not be comfortable with Bash syntax.

## Features

- **Natural Language to Bash Translation:** Converts plain English queries into executable Bash commands.
- **Data Augmentation:** Enhances training data using paraphrasing (via a pre-trained T5 paraphraser) and noise injection techniques.
- **Flask API Backend:** Serves the fine-tuned T5 model for real-time inference.
- **Go CLI Frontend:** A lightweight CLI tool that communicates with the backend via HTTP.
- **Modular Architecture:** Clean separation of model serving (Python) and user interaction (Go) for easier maintenance and scalability.

## Architecture

- **Backend (Python + Flask):**
  - Loads the fine-tuned T5-small model and its tokenizer.
  - Exposes a `/query` endpoint that accepts JSON payloads containing natural language queries.
  - Returns generated Bash commands as JSON responses.
- **Frontend (Go + Cobra):**
  - Sends HTTP POST requests to the Flask API.
  - Receives and displays the generated Bash command.
- **Data Pipeline:**
  - **Preprocessing:** Normalizes input by lowercasing, trimming, and removing unwanted punctuation.
  - **Augmentation:** Uses paraphrasing and noise injection to create diverse training examples.
  - **Fine-Tuning:** Adapts a pre-trained T5 model to the NLP-to-Bash task using the Hugging Face Trainer API.

## Installation

### Prerequisites

- **Python 3.7+** (with pip)
- **Go 1.16+**
- **Git & Git LFS** (if you plan to track large files)
- **Virtual Environment** (optional but recommended)

### Backend Setup

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/your-username/godModeNLP-CLI-TOOL.git
   cd godModeNLP-CLI-TOOL
2. **Install Dependencies:**

    ```bash
    pip install flask transformers sentencepiece evaluate kagglehub
3. **Run the Flask Server:**

### CLI Setup
1. **Navigate to the Go Project Folder:**
    ```bash
    cd nlpcli
2. **Initialize and Build the CLI Tool:**

    ```bash
    go mod init nlpcli
    go get github.com/spf13/cobra
    go build -o nlpcli
3. **Run the CLI Tool:**

   Make sure the Flask server is running, then execute:

    ```bash
    ./nlpcli suggest "find all jpg files modified last week"

