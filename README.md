# Alex's Saga Alien Invasion

## Todos

- [x] CLI
- [x] File parser
- [x] Map
- [x] Ensure repeteability in tests even if we are dealing with "random" stuff
- [x] Ensure pseudo-randomness by seeding rand if not in tests
- [x] Alien spawner
- [x] Alien
  - [x] Movement randomizer
- [x] City
- [x] Turns Loop
- [x] Aliens Battle
- [ ] Documentation
- [ ] UX
- [x] Bonus: alien names randomizer
- [ ] CI/CD and trunk based dev (I am working on this alone... probably overkill for now)

## Thinking 🤔
- [ ] Consider using an EventBus


## Usage

```
invasion [number of aliens invading] [flags]
```

### Options

```
  -h, --help                     help for invasion
      --mapfile string           the file containing the map of the 🌍, [relative path] (default "map.txt")
      --maxTurns int             the number of turns before the mothership calls the aliens back home and the program ends (default 10000)
      --numAliensForBattle int   the number aliens required to start a battle that will end up destroying a city (default 2)
      --verbose                  if True, sets the loglevel to show DEBUG messages
```

#### Help

```sh
go run ./...
```

#### Run

```sh
go run ./... X
```

where **X** is the number of aliens you want to spawn
#### Run with logging for debug/dev purposes

```sh
go run ./... X --verbose
```

where **X** is the number of aliens you want to spawn