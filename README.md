# Alex's Saga Alien Invasion

## Todos

- [x] CLI
- [x] File parser
- [x] Map
- [ ] Ensure repeteability in tests even if we are dealing with "random" stuff
- [ ] Ensure pseudo-randomness by seeding rand if not in tests
- [ ] Alien spawner
- [ ] Alien
  - [ ] Movement randomizer
- [ ] City
- [ ] Turns Loop
- [ ] AliensFight
- [ ] Documentation
- [ ] UX
- [ ] Bonus: alien names randomizer

## Thinking 🤔
- [ ] Consider using an EventBus


## Usage

#### Help

```sh
go run ./...
```

#### Run

```sh
go run ./... 15
```
#### Run with logging for debug/dev purposes

```sh
go run ./... 15 --verbose
```