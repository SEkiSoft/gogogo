//Client Side JavaScript code
//Game state is stored server side, client sends moves only

//New Game
function newGame(boardLines, ai) {
	//Tells the server to start a new game
	//Returns game ID
	return 12345;
}

//Existing game
function existingGame(gameID) {
	//Loads existing game from server
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
function getState(gameID) {
	//Gets current game state from server
}

//Draw grid
function drawGrid(boardLines) {
	//Draws grid onto HTML5 canvas
}

//Draw stones
function drawStones(black, white) {
	//Draws black and white stones onto the board
}