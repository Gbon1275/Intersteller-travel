package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge (fuel int) {
  fmt.Println("Fuel left:", fuel)
}


// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var requiredFuel int
  
  switch planet {
    case "Venus":
      requiredFuel = 300000
    case "Mercury":
      requiredFuel = 500000
    case "Mars":
      requiredFuel = 700000
    default:
      requiredFuel = 0
  }
  
  return requiredFuel
}

// Create the function greetPlanet() here
func greetPlanet (planet string) {
  fmt.Println("Next stop:", planet)
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fule to fly there")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining int = fuel
  
  fuleCost := calculateFuel(planet)
  
  if fuleCost >= fuelRemaining {
    cantFly()
  }
  
  return fuelRemaining
  
}

func main() {
  // Test your functions!
 
  // Create `planetChoice` and `fuel`
  fuel := 1000000
  planetChoice := "Venus"
  // And then liftoff!
  fuel = flyToPlanet(planetChoice,fuel)
  fuelGauge(fuel)
}
