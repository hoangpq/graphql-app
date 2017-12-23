import React, {Component} from 'react';
import PropTypes from 'prop-types';

const src = 'https://elements-cover-images-0.imgix.net/f2992836-14d0-47cf-8094-43913e2d08c5?fit=max&fm=jpeg&markalign=center%2Cmiddle&markalpha=18&q=80&w=316&s=1b49538a3d7a4ba9801c26e8e3c9774c';

export default class Product extends Component {

  render() {
    return (
      <div className="product-item">
        <div className="product-item-header">
          {this.props.name}
        </div>
        <div className="product-item-content">
          <span>Price: {this.props.price}</span>
          <span>Uom: {this.props.uom.name}</span>
          <div className="img-holder">
            <img src={src} />
          </div>
        </div>
      </div>
    )
  }

}

Product.propTypes = {
  name: PropTypes.string,
  price: PropTypes.number,
  uom: PropTypes.shape({
    id: PropTypes.string,
    name: PropTypes.string,
  }),
};
