import React, {Component} from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';

const src = 'https://elements-cover-images-0.imgix.net/f2992836-14d0-47cf-8094-43913e2d08c5?fit=max&fm=jpeg&markalign=center%2Cmiddle&markalpha=18&q=80&w=316&s=1b49538a3d7a4ba9801c26e8e3c9774c';

export default class Product extends Component {

  constructor(props) {
    super(props);
    this.state = {
      withUOM: true,
    };
    this.onCheckboxChange = this.onCheckboxChange.bind(this);
  }

  onCheckboxChange(evt) {
    evt.stopPropagation();
    this.setState(state => {
      console.log(state);
      return {
        withUOM: !state.withUOM
      }
    });
  }

  render() {
    return (
      <div className="product-item">
        <div className="product-item-header">
          <br/>
          <input type="checkbox" checked={this.state.withUOM} onChange={this.onCheckboxChange}/>
          { this.state.withUOM ? 'Checked' : 'Unchecked' }
        </div>
        <div className="product-item-content">
          <span>Price: {this.props.price}</span>
          <span>Uom: {this.props.uom.name}</span>
          <Link to={`/product/detail/${this.props.id}`}>
            <div className="img-holder">
              <span>Go to detail</span>
              {/*<img src={src} />*/}
            </div>
          </Link>
        </div>
      </div>
    )
  }

}

Product.propTypes = {
  id: PropTypes.string,
  name: PropTypes.string,
  price: PropTypes.number,
  uom: PropTypes.shape({
    id: PropTypes.string,
    name: PropTypes.string,
  }),
};
