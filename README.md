# Game of Life - Go Implementation

A terminal-based implementation of Conway's Game of Life written in Go.

## Description

This project is an implementation of Conway's Game of Life, a cellular automaton devised by mathematician John Conway. The "game" is a zero-player game, meaning that its evolution is determined by its initial state, with no further input from humans.

The universe is represented as a grid of cells, each of which can be in one of two states: alive or dead. The state of each cell in the next generation is determined by its current state and the number of live neighbors it has, according to the following rules:

1. Any live cell with fewer than two live neighbors dies (underpopulation)
2. Any live cell with two or three live neighbors lives on to the next generation
3. Any live cell with more than three live neighbors dies (overpopulation)
4. Any dead cell with exactly three live neighbors becomes a live cell (reproduction)

## Features

- Customizable universe size
- Configurable seed for reproducible randomization
- Variable number of generations to run
- Adjustable pause time between generations
- Terminal-based visualization
