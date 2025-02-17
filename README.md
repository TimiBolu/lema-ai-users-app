# Lema-AI Users App

## Overview
Lema-AI Users App is a full-stack application designed to manage users and their posts efficiently. The project consists of two main components:

1. **Frontend**: A React-based web application that provides an intuitive user interface for managing users and posts.
2. **Backend**: A Go-based REST API that handles user and post management with SQLite as the database.

Each component has its own dedicated README file with setup instructions and additional details.

## Project Structure
```bash
lema-ai-users-app/
├── backend/          # Backend service (Go)
│   ├── README.md    # Backend documentation
│   ├── main.go      # Entry point for the backend
│   ├── handlers/    # API handlers
│   ├── models/      # Database models
│   ├── router/      # API router
│   └── database/    # Database setup
│
├── frontend/         # Frontend application (React)
│   ├── README.md    # Frontend documentation
│   ├── src/         # React source code
│   ├── public/      # Static assets
│   └── package.json # Project dependencies
│
├── .gitignore       # Git ignore file
├── LICENSE          # MIT License
└── README.md        # This project documentation
```

## Getting Started

### Prerequisites
Ensure you have the following installed:
- **Node.js** (v14 or higher) for the frontend
- **Go** (v1.23.0 or higher) for the backend
- **SQLite** for database management

### Installation and Setup
#### Clone the repository
```bash
git clone https://github.com/TimiBolu/lema-ai-users-app.git
cd lema-ai-users-app
```

### Backend Setup
Navigate to the backend directory and follow the [backend README](backend/README.md) for setup instructions.
```bash
cd backend
cat README.md
```

### Frontend Setup
Navigate to the frontend directory and follow the [frontend README](frontend/README.md) for setup instructions.
```bash
cd frontend
cat README.md
```

## Contributing
If you would like to contribute to the project, please follow these steps:
1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Submit a pull request.

## License
MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact
For questions or feedback, reach out to **Timi Adesina** at [adesinatim@gmail.com](mailto:adesinatim@gmail.com).
