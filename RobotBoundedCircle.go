package main

func isRobotBounded(instructions string) bool {
    x, y := 0, 0
    // Direction: 1 = North, 2 = East, 3 = South, 4 = West
    direction := 1
    
    for i := 0; i < len(instructions); i++ {
        switch instructions[i] {
        case 'G':
            switch direction {
            case 1:
                y++
            case 2:
                x++
            case 3:
                y--
            case 4:
                x--
            }
        case 'L':
            direction--
            if direction == 0 {
                direction = 4
            }
        case 'R':
            direction++
            if direction == 5 {
                direction = 1
            }
        }
    }
    if (x == 0 && y == 0) || direction != 1 {
        return true
    }
    
    return false
}

// Runtime:=1ms Beats 76.19%
// Memory:=2.19MB Beats 69.05%
