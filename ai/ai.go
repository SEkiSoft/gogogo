// Copyright (c) 2016 SEkiSoft
// See License.txt

package ai

import (
	"github.com/sekisoft/gogogo/api"
	"github.com/sekisoft/gogogo/model"
	"github.com/sekisoft/gogogo/store"
)

//NextMove based on board state, whether the player is a 1 or a 2
//Returns x, y, GG
// GG = true, ai surrenders
//Currently does not look-ahead moves
func NextMove(game *model.Game, player *model.Player) *model.Move {
	//Find groups (assumes dead groups have been removed)
	//Classify each group as alive, dead, or unsettled
	//Find number of liberties for each group
	//In this order:
	//1. Defend own groups first
	//2. Kill easy opponent groups
	//3. Start new groups
	//4. Kill difficult opponent groups

	return nil
}

//Takes all moves from game from DB, and analyzes them
func AnalyzeData(moves []*model.Move) *model.Ai {

}

//Get stats
func GetStats() {

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
