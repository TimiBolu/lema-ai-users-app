# Lema-Ai Users App (Frontend)

## Overview
Lema-Ai Users App is a React-based frontend application for managing users and posts. It uses TypeScript, React Query, and Tailwind CSS to create an efficient, responsive UI. This app connects to a backend API for fetching user data, displaying posts, and managing post creation and deletion. The frontend is designed to be responsive and follows best practices in code organization, state management, and UI design.

## Features
- **Paginated User Table**: Displays a list of users with details such as full name, email, and address.
- **User Posts Page**: Displays the list of posts for each user, including options to delete posts and create new ones.
- **Post Management**: Create new posts, delete existing ones, and display post details like title and body.
- **Error Handling**: Graceful error handling for failed API requests or unexpected backend responses.
- **Responsive Design**: Tailwind CSS is used for building a responsive and flexible layout that works well on various devices.

## Setup Instructions

### Prerequisites
Ensure that you have the following installed on your machine:
- **Node.js** (version 14 or higher)
- **npm** (Node Package Manager)

### Install Dependencies
1. Clone the repository:
   ```bash
   git clone https://github.com/TimiBolu/lema-ai-users-app.git
   cd lema-ai-users-app
   ```

2. Install the project dependencies:
   ```bash
   npm install
   ```

### Run the Development Server
To run the development server, execute the following command:
```bash
npm run dev
```

The app will be available at `http://localhost:5173`.

### Build for Production
To build the app for production, run:
```bash
npm run build
```

To preview the production build, use:
```bash
npm run preview
```

### Run Tests
To run the tests using **Vitest**, use the following command:
```bash
npm run test
```

For coverage reports, use:
```bash
npm run coverage
```

To run tests in watch mode:
```bash
npm run test:watch
```

To run UI tests with Vitest's UI:
```bash
npm run test:ui
```



## Project Structure

### File Structure

```
/src
  /apis            # API services
  /components      # Reusable components (Loader, DirectionBtn, etc.)
  /hooks           # Custom React hooks (useUser, useWindowSize)
  /layout          # The layout of the app
  /mocks           # Setting up the mock server and seeding mock data
  /pages           # Main App Screens
  /router          # Basic App Router
  /utils           # Utility functions and helpers
  /types           # TypeScript type definitions
  main.tsx         # Root of the application
```
