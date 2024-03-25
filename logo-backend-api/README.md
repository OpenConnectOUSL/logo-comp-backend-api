# Logo Design Submission Web Interface

This project is a simple web interface for logo design submissions. It's built with a Go-based backend API and a Next.js frontend.

## Features

- **Image Storage**: Images are stored using [edgestore.dev](https://edgestore.dev).
- **Drag and Drop**: Users can upload images using a drag and drop interface.
- **User Details**: User details are stored in the backend API.

## Deployment

The backend API is deployed on [DigitalOcean](https://www.digitalocean.com). The frontend is deployed on [Netlify](https://www.netlify.com) and is connected to the backend API.

API URL :- https://206.189.138.245/v1/healthcheck

## Usage

To use the web interface, visit the deployed site [https://logocomp.netlify.app/](https://logocomp.netlify.app/).

## Running the Backend API

To run the Go backend API locally, you'll need to have Go installed. Follow these steps:

1. Navigate to the directory containing the Go code:
    ```bash
    cd <go-directory>
    ```

2. Build the Go application:
    ```bash
    go build
    ```

3. Run the Go application:
    ```bash
    ./<go-application>
    ```

Replace `<go-directory>` with the path to the directory containing your Go code, and `<go-application>` with the name of your Go application.

The API will start running on its configured port (default is usually `:4000`). You can then send requests to `localhost:<port>` to interact with the API.

## Running the Frontend Application



To run the Next.js frontend application locally, you'll need to have Node.js installed. Follow these steps:

1. Navigate to the directory containing the Next.js code:

    ```
    git clone https://github.com/mayura-andrew/logo-dropbox.git
    ```
    ```bash
    cd <nextjs-directory>
    ```

2. Install the dependencies:
    ```bash
    npm install
    ```

3. Start the development server:
    ```bash
    npm run dev
    ```

Replace `<nextjs-directory>` with the path to the directory containing your Next.js code.

Visit `localhost:4000` in your browser to see the application running.
