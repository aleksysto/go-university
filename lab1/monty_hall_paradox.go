package main

import (
	"fmt"
	"math/rand/v2"
)

const number int = 3

func removeFromSliceOrdered[T any](s []T, at int) []T {
    s1 := s[:at]
    s2 := s[(at+1):]
    return append(s1, s2...)
}

func main() {
    var losses int = 0
    var wins int = 0
    n_games := 100
    strategy := true
    
    for i := 0; i < n_games; i++ {
        boxes := make([]bool, number)
        rGoodBox := rand.IntN(len(boxes))
        boxes[rGoodBox] = true

        playerChoice := rand.IntN(len(boxes))
        
        if !strategy {
            if boxes[playerChoice] {
                wins += 1
            } else {
                losses += 1
            }
            continue
        }


        for {
            x := rand.IntN(len(boxes))
            if !boxes[x] && x != playerChoice {
                boxes = removeFromSliceOrdered(boxes, x)
                if playerChoice != 0 {
                playerChoice -= 1
                }
                break
            }
        }
       // j := 0
       // for {
       //     x := rand.IntN(len(boxes))
       //     if !boxes[x] && x != playerChoice {
       //         j++
       //         boxes = removeFromSliceOrdered(boxes, x)
       //         if playerChoice != 0 {
       //             playerChoice -= 1
       //         }
       //     }
       //     if j > 96 {
       //         break
       //     }
       // }

        for {
            newChoice := rand.IntN(len(boxes))
                if newChoice != playerChoice {
                    playerChoice = newChoice
                    break
                }
        }
        fmt.Println(boxes)
        if boxes[playerChoice] {
            wins += 1
        } else {
            losses += 1
        }
    }
    percentWin := float64(wins) / float64(100)
    fmt.Println("Wins: ", percentWin, wins, losses)
}
