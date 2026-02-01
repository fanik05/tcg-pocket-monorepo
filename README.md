# Pokémon TCG Pocket - Tournament Organizer

A high-performance monorepo for organizing and tracking custom Pokémon TCG Pocket tournaments. Built with a Go-powered pairing engine and a real-time Next.js frontend.

## 🏗️ Architecture Overview
- **/apps/api**: The Backend. Written in **Go**, handling the Swiss pairing algorithm, tie-breaker logic, and tournament state.
- **/apps/web**: The Frontend. A **Next.js** application focused on a mobile-first "Match Hub" experience for players and organizers.

## 🚀 Getting Started

### Prerequisites
- [Go](https://go.dev/doc/install) (1.21+)
- [Node.js](https://nodejs.org/) (20+)
- [Make](https://www.gnu.org/software/make/)

### Installation & Development
This project uses a `Makefile` to simplify monorepo management.

1. **Install all dependencies:**
   ```bash
   make install