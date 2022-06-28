# Alex's Saga Alien Invasion

## Todos

- [x] CLI
- [x] File parser
- [x] Map
- [x] Ensure repeteability in tests even if we are dealing with "random" stuff
- [x] Ensure pseudo-randomness by seeding rand if not in tests
- [x] Alien spawner
- [x] Alien
  - [ ] Movement randomizer
- [ ] City
- [ ] Turns Loop
- [ ] AliensFight
- [ ] Documentation
- [ ] UX
- [ ] Bonus: alien names randomizer
- [ ] CI/CD and trunk based dev (I am working on this alone... probably overkill for now)

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