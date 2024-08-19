# Race Condition Example in Go

The `race.go` file demonstrates a simple example of a race condition in Go. A race condition occurs when multiple goroutines access shared resources concurrently, leading to unpredictable behavior.

## Code Overview

The [main](https://github.com/Sgudkov/race/blob/main/race.go#L10:L34) function in `race.go` launches 5 goroutines, each of which increments a shared `counters` map and a `totalCount` integer. The goroutines use a mutex (`mu`) to ensure exclusive access to the shared resources.

## Key Components

* `counters` map: a shared map that stores integer values, incremented by each goroutine.
* `totalCount` integer: a shared integer that keeps track of the total count, incremented atomically using `atomic.AddInt32`.
* `mu` mutex: a mutex that ensures exclusive access to the shared resources.

## Example Output

The program prints the final state of the `counters` map and the `totalCount` integer after all goroutines have completed.

## Note

This example demonstrates a simple race condition scenario. In a real-world scenario, you would want to handle errors and edge cases more robustly.

## Description of race.go

The `race.go` file contains a simple example of a race condition in Go. The [main](https://github.com/Sgudkov/race/blob/main/race.go#L10:L34) function launches 5 goroutines, each of which increments a shared `counters` map and a `totalCount` integer. The goroutines use a mutex (`mu`) to ensure exclusive access to the shared resources. The program prints the final state of the `counters` map and the `totalCount` integer after all goroutines have completed.