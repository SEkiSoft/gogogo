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
	//Using POST
}

//Send move
function sendMove(x, y, player, gameID) {
	//Sends move to server using POST
	//Returns:
	//	-1: communication error
	//	 0: invalid move
	//   1: valid move, accepted
	return -1; 
}

//Get AI move
function getAI() {
	//Query AI move from server using POST
	return [0, 0];
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