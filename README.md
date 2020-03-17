# Game of Life

`gol` is Conway's [Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) written in go.

## Purpose
Practice project structure and concurrency in Go

## Progress
### Taks
- [x] Grid Structure
- [x] Cell Rules
- [ ] Main Loop
- [ ] Rendering Grid
- [ ] Controls
### Current Progress
Looking at different ways of using goroutines. Currently using a "worker pool" kind of idea, but its slower than just using a nested loop. 
