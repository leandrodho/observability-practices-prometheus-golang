# Observability Practices: Application Metrics with Prometheus in Golang

This repository provides a concrete implementation demonstrating application monitoring using **Prometheus** metrics inside a Go application.

## Project Structure
- `/blind`: An unmonitored production server example lacking diagnostic metrics.
- `/observed`: A production-ready architecture tracking performance signals using a custom telemetry middleware layer.

## Getting Started

### Prerequisites
Ensure you have Go installed (v1.22+ recommended).

### Dependency Installation
Before running the application, fetch the required Prometheus client packages:
```bash
go mod tidy