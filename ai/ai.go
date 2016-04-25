package ai

var board[][] int;
var g[][] int;

//NextMove based on board state, whether the player is a 1 or a -1
//Returns x, y, GG
// GG = true, ai surrenders
//Currently does not look-ahead moves
func NextMove(b [][] int, player int) (int, int, bool){
	//Find groups (assumes dead groups have been removed)
	//Classify each group as alive, dead, or unsettled
	//Find number of liberties for each group
	//In this order:
	//1. Defend own groups first
	//2. Kill easy opponent groups
	//3. Start new groups
	//4. Kill difficult opponent groups
	//TODO: DFS each move up to X moves?
	//This can be really hard

	//Find groups
	//ASSUMES SQUARE BOARD
	board = b;
	for i:=0; i < len(board); i++ {
		for j:=0; j < len(board); j++ {
			g[i][j] = -1;
		}
	}

	//BFS through current group
	var cur = 0;
	for i:=0; i < len(board); i++ {
		for j:=0; j < len(board); j++ {
			if(g[i][j] == -1 && board[i][j] != 0) {
				findGroup(i, j, board[i][j], cur);
				cur++;
			}
		}
	}
}

//DFS time
func findGroup(x int, y int, player int, cur int) {
	if(board[i][j] == player) {
		g[i][j] = cur;

		if(j-1 >= 0 && board[i][j-1] == player) {
			findGroup(i, j-1, player, cur);
		}
		if(j+1 < len(board) && board[i][j+1] == player) {
			findGroup(i, j+1, player, cur);
		}
		if(i-1 >= 0 && board[i-1][j] == player) {
			findGroup(i-1, j, player, cur);
		}
		if(i+1 < len(board) && board[i+1][j] == player) {
			findGroup(i+1, j, player, cur);
		}
	}
} 