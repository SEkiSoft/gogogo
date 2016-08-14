import React, { Component } from 'react';
import crossSection from './images/cross-section.png';
import topLeftCorner from './images/top-left-corner.png';
import topRightCorner from './images/top-right-corner.png';
import topEdge from './images/top-edge.png';
import leftEdge from './images/left-edge.png';
import rightEdge from './images/right-edge.png';
import bottomLeftCorner from './images/bottom-left-corner.png';
import bottomRightCorner from './images/bottom-right-corner.png';
import bottomEdge from './images/bottom-edge.png';
import blackPiece from './images/black-piece.png';
import whitePiece from './images/white-piece.png';

import './CrossSection.css';

class CrossSection extends Component {
	constructor(props) {
		super(props);
		this.setPiece = this.setPiece.bind(this);
	}

	getBoardImage() {
		const { rowIndex, colIndex } = this.props;
		if (colIndex === 0) {
			if (rowIndex === 0) {
				return topLeftCorner;
			}
			if (rowIndex === 18) {
				return bottomLeftCorner;
			}
			return leftEdge;
		}
		if (colIndex === 18) {
			if (rowIndex === 0) {
				return topRightCorner;
			}
			if (rowIndex === 18) {
				return bottomRightCorner;
			}
			return rightEdge;
		}
		if (rowIndex === 0) {
			return topEdge;
		}
		if (rowIndex === 18) {
			return bottomEdge;
		}
		return crossSection;
	}

	setPiece() {
		const { setPiece, pieceColour, userColour, colIndex, rowIndex } = this.props;
		if (!pieceColour) {
			return setPiece(userColour, colIndex, rowIndex);
		}
	}

	render() {
		const { pieceColour } = this.props;
		const hasPiece = pieceColour === 'w' || pieceColour === 'b';
		const crossSectionImage = this.getBoardImage();
		const pieceImage = pieceColour === 'w' ? whitePiece : blackPiece;
		const showPiece = hasPiece ? 'piece' : 'hiddenPiece';

		return (
			<div className='crossSection'>
				<img className='boardSection' src={crossSectionImage} onClick={this.setPiece} />
					<img src={pieceImage} className={showPiece} />
			</div>
		);
	}
}

CrossSection.propTypes = {
	rowIndex: React.PropTypes.number.isRequired,
	colIndex: React.PropTypes.number.isRequired,
	userColour: React.PropTypes.string.isRequired,
	pieceColour: React.PropTypes.string.isRequired
}

export default CrossSection;
