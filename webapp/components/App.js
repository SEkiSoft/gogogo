import React, { Component } from 'react';
import './App.css';
import CrossSection from './CrossSection';

class App extends Component {
  
  constructor(props) {
    super(props);
    this.state = {
      gameData: [['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', 'w', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', 'w', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', 'w', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', '', ''],
                 ['', 'b', '', '', '', '', '', '', '', '', '', 'b', '', '', '', '', '', '', '']]
    };
  }

  setPiece = (userColour, colIndex, rowIndex) => {
    const gameState = this.state.gameData;
    gameState[rowIndex][colIndex] = userColour;
    this.setState({
      gameData: gameState
    }, this.endTurn())
  }

  endTurn() {
  }

  render() {
    return (
      <div className='App'>
        <h1>Go Go Go</h1>
        {
          this.state.gameData.map((row, rowIndex) => {
            const rowClass = rowIndex >= 1 ? 'row' : '';
            return (
              <div key={rowIndex} className={rowClass}>
              {
                row.map((cell, colIndex) => {
                  return (
                    <CrossSection 
                      key={rowIndex.toString()+ colIndex.toString()} 
                      rowIndex={rowIndex} colIndex={colIndex}
                      userColour={'b'}
                      setPiece={this.setPiece}
                      pieceColour={cell} />
                  )
                })
              }
              </div>
            )
          })
        }
      }
      </div>
    );
  }
}

export default App;
