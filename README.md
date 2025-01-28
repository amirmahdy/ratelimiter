# Rate Limiter

A simple and efficient rate limiter implementation, designed to control the frequency of actions or requests.

## Overview

This repository provides a rate limiter based on Token Bucket algorithm that helps you enforce limits on how often a particular action or request can be performed. It is useful in scenarios where you need to prevent abuse, manage resource usage, or comply with API rate limits.

## Features

- **Simple API**: Easy-to-use interface for rate limiting.
- **Flexible Configuration**: Set custom limits and time windows.
- **Concurrency-Safe**: Designed to work in concurrent environments using Go's `sync` package.
- **Lightweight**: Minimal dependencies and low overhead.

## Installation

To use this rate limiter in your Go project, run:

```bash
go get github.com/amirmahdy/ratelimiter