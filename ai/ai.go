package ai

//NextMove based on board state, whether the player is a 1 or a 0
//Returns x, y, GG
// GG = true, ai surrenders
//Currently does not look-ahead moves
func NextMove(board [] int, player int) (int, int, bool){
	//Find groups (assumes dead groups have been removed)
	//Classify each group as alive, dead, or unsettled
	//Find number of liberties for each group
	//In this order:
	//1. Defend own groups first
	//2. Kill easy opponent groups
	//3. Start new groups
	//4. Kill difficult opponent groups
}