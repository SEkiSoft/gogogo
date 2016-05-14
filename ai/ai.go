package ai

type Point struct {
	uint x
	uint y
}

//NextMove based on board state, whether the player is a 1 or a 2
//Returns x, y, GG
// GG = true, ai surrenders
//Currently does not look-ahead moves
func NextMove(board model.Board, player int) (uint, uint, bool){
	//Find groups (assumes dead groups have been removed)
	//Classify each group as alive, dead, or unsettled
	//Find number of liberties for each group
	//In this order:
	//1. Defend own groups first
	//2. Kill easy opponent groups
	//3. Start new groups
	//4. Kill difficult opponent groups

	//Find groups
	for i:=0; i < board.numLines; i++ {
		for j:=0; j < board.numLines; j++ {
			g[i][j] = -1
		}
	}

	//BFS through each potential group
	//Black - 1
	//White - 2
	//Empty - 0
	var cur uint = 0
	for i:=0; i < len(board); i++ {
		for j:=0; j < len(board); j++ {
			if(g[i][j] == -1) {
				findGroup(g, i, j, board[i][j], cur);
				cur++
			}
		}
	}
	//TODO: Make more efficient, currently looping twice to avoid use of append
	//Make slice for number of liberties of each group
	lib := make([]uint, cur)
	//Find number of liberties
	//TODO
	//Store location of each liberty
	loc := map[int][]Point //Map a group to an array of points
	//Store each liberty as a point
	//TODO
	//Find which group belongs to which player
	//Black - 1, White - 2, Empty - 0
	gp := make([]uint, cur)
}

/*
AI Modelling:
Every combination of (playerGroups, playerLiberties, playerStones, opponentGroups, opponentLiberties, opponentStones) represent a situation
Next moves are represented by pG, pL, pS, oG, oL, oS and an operator.
Operators are +, -, * (don't care), < (<= current), > (>= current), and =
Analyze the current situation:
-If encountered before, choose best move if SAMPLE_COUNT > THRESHOLD
-Otherwise, search other situations using constraints: min(pG, pL, pS), vary each variable up to max(pG, pL, pS) - min
-Otherwise, pick a move at random
*/

/*
Machine Learning:
Record every move made, and the preceding situations
Record winning player
Propagate back data
Update winning probabilities of each move at each given situation
Analysis runs both ways
Run analysis in batches (every day?)
*/

//DFS time
func findGroup(g *[][]int, x int, y int, player int, cur int) {
	if(board[i][j] == player) {
		g[i][j] = cur

		if(j-1 >= 0 && board[i][j-1] == player) {
			findGroup(g, i, j-1, player, cur)
		}
		if(j+1 < len(board) && board[i][j+1] == player) {
			findGroup(g, i, j+1, player, cur)
		}
		if(i-1 >= 0 && board[i-1][j] == player) {
			findGroup(g, i-1, j, player, cur)
		}
		if(i+1 < len(board) && board[i+1][j] == player) {
			findGroup(g, i+1, j, player, cur)
		}
	}
} 