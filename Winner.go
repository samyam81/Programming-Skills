package main

func tictactoe(moves [][]int) string {
    board := make([][]string, 3)
    for i := range board {
        board[i] = make([]string, 3)
    }
    
    for idx, move := range moves {
        row, col := move[0], move[1]
        if idx % 2 == 0 {
            board[row][col] = "X"
        } else {
            board[row][col] = "O"
        }
    }
    
    winner := checkWinner(board)
    
    if winner == "X" {
        return "A"
    } else if winner == "O" {
        return "B"
    }
    
    if len(moves) < 9 {
        return "Pending"
    } else {
        return "Draw"
    }
}

func checkWinner(board [][]string) string {
    for i := 0; i < 3; i++ {
        if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
            return board[i][0]
        }
    }
    
    for j := 0; j < 3; j++ {
        if board[0][j] != "" && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
            return board[0][j]
        }
    }
    
    if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
        return board[0][0]
    }
    if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
        return board[0][2]
    }
    
    return ""
}

// Runtime:= 2ms Beats 20.31%
// Memory:= 2.28MB Beats 67.19%