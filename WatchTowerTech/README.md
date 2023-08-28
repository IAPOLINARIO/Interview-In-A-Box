# WatchTowerTech Homework Task

Welcome to the WatchTowerTech Homework Task. This Flask application provides several endpoints that make use of the [Fake Name Generator](https://www.fakenamegenerator.com/) to return details of a fictitious older person, common words (names), and potential blood donors.

## Getting Started

These instructions will help you to set up and run the project on your local machine via Docker.

### Prerequisites

- Docker
- Docker-compose

### Installation

1. Unzip the provided file to your desired directory.

### Running the Application

In the project directory, run the docker-compose file:

```sh
docker-compose up --build
```

The application will be running at [http://localhost:5000](http://localhost:5000)

You can also run Jupyter notebook using the following link: [http://localhost:8888](http://localhost:8888)

## Available Endpoints

1. `/old_person` - Returns details of a fictitious older person.

Example call:

```sh
curl http://localhost:5000/old_person
```

Example result:

```json
{
  "first_name": "John",
  "middle_name": "Doe",
  "last_name": "Smith",
  "date_of_birth": "12/03/1945",
  "tropical_zodiac": "Pisces"
}
```

2. `/top_used_words` - Returns the most common names obtained from the Fake Name Generator.

Example call:

```sh
curl "http://localhost:5000/top_used_words?number_of_names=10"
```

Example result:

```json
{
  "John": 6,
  "Mary": 4,
  "Smith": 3,
  "Doe": 2,
  "Sara": 1
}
```

3. `/get_blood_donors` - Returns a list of potential blood donors for a specific blood type.

Example call:

```sh
curl "http://localhost:5000/get_blood_donors?blood_type=O%2B"
```

Example result:

```json
[
  {
    "blood_type": "O+",
    "donors": [
      {
        "name": "John Doe",
        "age": 25,
        "phone_number": "1234567890"
      },
      {
        "name": "Mary Smith",
        "age": 30,
        "phone_number": "9876543210"
      }
    ]
  }
]
```

## Jupyter Notebooks

Jupyter Notebooks are also provided as a part of this project, which you can use to interact with the application in a more detailed way.

To access the notebooks, navigate to the following URL in your browser: [http://localhost:8888](http://localhost:8888)

When prompted for a password/token, use the toekn printed in the command line when the container starts.

The notebooks are located in the `WatchTowerTech` directory.

## Warning

The application makes numerous requests to the Fake Name Generator website. Please be aware that making too many requests in a short period of time could lead to your IP being temporarily blocked by the website. Make sure to use this application responsibly to avoid disruptions to the service.
