# EPI's GoData Tools
This project tries to augment GoData with customized tools for EPI.

The project is split into the backend(api) and the frontend.

## Backend
A Go http server that provides the endpoints for retrieving the outbreak cases:
1. `/outbreaks`: returns a list of all outbreaks
2. `/casesByOutbreak`:  returns all cases for a specific outbreak.

## Frontend
This is a React project that communicates with the `Backend` to fetch data.
