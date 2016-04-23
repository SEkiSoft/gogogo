//Client Side JavaScript code
//Game state is stored server side, client sends moves only

//New Game
function newGame(boardLines, ai) {
	//Tells the server to start a new game
	//Returns game ID
	return 12345;
}

//Send move
function sendMove(x, y, player) {
	//Sends move to server
	//Returns:
	//	-1: communication error
	//	 0: invalid move
	//   1: valid move, accepted
	return -1; 
}

//Get state
function getState() {
	//Gets current game state from server
}