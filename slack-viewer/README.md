# Slack Exporter Viewer

Slack Exporter Viewer is a secure tool that allows authorized users to read all the exported history of Slack in JSON format, saves it on a MongoDB database for later search, and search through the saved history. The application provides REST endpoints for all functionalities. The backend of the application is written in Golang, and the frontend is built using React. The application is designed to run on AWS.

## Introduction

Slack Exporter Viewer is an efficient tool that can handle a vast amount of data exported from Slack, including all messages and attachments sent to a particular channel, group or DMs. The tool is designed to provide authorized users with a secure platform to save, search, and export Slack history for their teams. With its intuitive user interface, users can search through all Slack history by user or timeframe.

Slack Exporter Viewer is built with a scalable and distributed architecture, making it suitable for large Slack teams.

## Advantages of the Tool

- Easy import of Slack exported data in JSON format.
- Search through all Slack history by user.
- Simple and intuitive user interface built with React (Slack look & feel).
- Scalable and distributed architecture with Golang and MongoDB.
- Ability to handle vast amounts of data exported from Slack.
- Strong layer of security to ensure only authorized users can view the history.

## Technologies Used

The following technologies are used in this project:

- Golang - Backend language used to build the application's logic.
- MongoDB - A NoSQL document-oriented database used for storing the Slack export data.
- React - A JavaScript library used to build the user interface.
- AWS - The application is designed to run on AWS.

## Running the Application

To run the application, you will need to follow these steps:

1. Clone the repository to your local machine.
2. Configure your MongoDB instance.
3. Configure the application's environment variables.
4. Build the Golang backend and run it.
5. Build the React frontend and run it.

### Configure your MongoDB instance

The application requires a MongoDB database to store the Slack export data. You can create a free MongoDB instance on [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).

Once you have created your MongoDB instance, copy the connection string and replace the `MONGODB_URI` value in the `.env` file with your connection string.

### Configure the application's environment variables

To configure the application's environment variables, copy the `.env.example` file to `.env` and modify the values as necessary. The following variables need to be set:

- `SESSION_SECRET` - A secret key used to encrypt user sessions.
- `AUTHORIZED_USERS` - A comma-separated list of Slack user IDs that are authorized to view the Slack history.

### Build the Golang backend and run it

To build the backend, navigate to the `backend` directory, and run the following command:

```bash
go build && ./slack-exporter-viewer
```

This command will build the backend and start the server. By default, the backend runs on port 8000.

### Build the React frontend and run it

To build the frontend, navigate to the frontend directory and run the following commands:

```bash
npm install
npm start
```

This command will install all required dependencies and start the React development server. By default, the frontend runs on port 3000.

### Usage

#### General

1. Run the `slackctl` command to display the welcome message and a list of available commands:

```
./slackctl
```

2. For help with a specific command, type:

```
./slackctl [command] help
```

#### Generate a Report

Generate a report based on the input data:

```
./slackctl report -u [username] -s [source_directory] -o [output_directory]
```

Replace `[username]`, `[source_directory]`, and `[output_directory]` with the desired values.

#### Search Local

Search for Slack history files in a local directory:

```
./slackctl search-local -u [username] -s [source_directory] -o [output_directory]

```

Replace `[username]`, `[source_directory]`, and `[output_directory]` with the desired values.

#### Search Remote

Search for Slack history files in a remote directory (S3):

```
./slackctl search-remote -u [username] -s [source_directory] -o [output_directory] -b [bucket_name]
```

Replace `[username]`, `[source_directory]`, `[output_directory]`, and `[bucket_name]` with the desired values. Make sure your AWS credentials are set up correctly.

#### Start Server

Start a web server for the Slack APIs:

```
./slackctl server -p [port]
```

Replace `[port]` with the desired port number.

### Contributing

If you would like to contribute to the project, please follow the steps below:

- Fork the repository.
- Create a new branch for your feature or bug fix.
- Make your changes and commit them.
