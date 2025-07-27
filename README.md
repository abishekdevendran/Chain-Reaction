# ⛓️ Chain Reaction

A real-time, multiplayer web-based implementation of the classic strategy game Chain Reaction. Built with a Go backend and a Vue/Nuxt frontend.

[![Go Backend Status](https://img.shields.io/badge/backend-Go-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
[![Vue Frontend Status](https://img.shields.io/badge/frontend-Vue.js-4FC08D?style=for-the-badge&logo=vue.js)](https://vuejs.org/)
[![Project Status](https://img.shields.io/badge/status-in%20development-orange?style=for-the-badge)]()

---

## Table of Contents
- [⛓️ Chain Reaction](#️-chain-reaction)
  - [Table of Contents](#table-of-contents)
  - [About The Game](#about-the-game)
  - [Tech Stack](#tech-stack)
  - [Project Roadmap](#project-roadmap)
  - [Getting Started](#getting-started)
  - [License](#license)

## About The Game

Chain Reaction is a strategy game for 2 to 8 players. Players take turns placing "orbs" in the cells of a grid. Once a cell reaches its critical mass, the orbs explode into the surrounding cells, claiming them for the player. A player is eliminated when they have no orbs left on the grid. The last player standing wins.

This project aims to be a fully-featured, real-time multiplayer implementation of the game, playable in any modern web browser.

## Tech Stack

This project is a monorepo containing two main components:

*   **Backend**: Written in **Go**.
    *   Handles all game logic and state management.
    *   Serves a WebSocket API for real-time communication.
    *   Manages player connections and game rooms.
*   **Frontend**: Written in **Vue 3** with the **Nuxt 3** framework.
    *   Provides a clean, responsive, and interactive user interface.
    *   Communicates with the backend via WebSockets.
    *   Uses Pinia for state management and TailwindCSS for styling.

## Project Roadmap

- [x] **Phase 1: Project Setup & Foundation**
  - [x] Initialize GitHub repository with professional structure.
  - [x] Set up monorepo for backend and frontend.
  - [x] Define core project documentation (`README.md`, `LICENSE`).
- [ ] **Phase 2: Backend Core Logic**
  - [ ] Implement core game state (Grid, Players, Orbs).
  - [ ] Develop game turn and explosion logic.
  - [ ] Set up basic HTTP server.
- [ ] **Phase 3: Real-time Communication**
  - [ ] Integrate WebSocket support into the Go server.
  - [ ] Define the WebSocket communication protocol (JSON messages).
- [ ] **Phase 4: Frontend Scaffolding**
  - [ ] Initialize Nuxt 3 project.
  * [ ] Create the basic game board UI.
- [ ] **Phase 5: Full-Stack Integration**
  - [ ] Connect the Vue frontend to the Go WebSocket backend.
  - [ ] Implement turn-based actions from the UI.
  - [ ] Visually represent game state changes received from the server.
- [ ] **Phase 6: Advanced Features**
  - [ ] Implement game lobbies and room codes.
  - [ ] Add player authentication.
  - [ ] Deploy to a cloud platform.

## Getting Started

*(Instructions to be added once the project is runnable).*

## License

Distributed under the MIT License. See `LICENSE` for more information.