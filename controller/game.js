// Copyright (c) 2016 David Lu
// See License.txt

//Client Side JavaScript code
//Game state is stored server side, client sends moves only

//Server URL
const serverRoot = "http://localhost";
var id = "";

//New Game
function newGame(boardLines, ai) {
	//Tells the server to start a new game
	var xhr = new XMLHttpRequest();
	xhr.open('GET', "$serverRoot/newGame", false);
	xhr.send();

	//Gets game ID
	id = JSON.parse(xhr.responseText).id;

	//Loads game state
	drawGame(id);
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
	//Query AI move from server
	var xhr = new XMLHttpRequest();
	xhr.open('GET', "$serverRoot/ai/$id", false);
	xhr.send();

	//Get response
	var response = JSON.parse(xhr.responseText);
	var x = response.x;
	var y = response.y;
	var gg = response.gg;

	//Process response
}

//Draw game
function drawGame(gameID) {

}

//Draw grid
function drawGrid(boardLines) {
	//Draws grid onto HTML5 canvas
}

//Draw stones
function drawStones(black, white) {
	//Draws black and white stones onto the board
}