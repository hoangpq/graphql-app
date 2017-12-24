import React, {Component} from 'react';
import Product from './Product';
import {gql} from '../utils';

export default class ProductList extends Component {

  constructor(props) {
    super(props);
    this.state = {
      products: [],
    }
  }

  componentDidMount() {
    gql(`
      query getProducts {
        products {
          id
          name
          price
          uom {
            id
            name 
          }
        }
      }
    `)
      .then(resp => this.setState({products: resp.products}));
  }

  renderProductList() {
    return this.state.products.map((product, i) => <Product key={`product_${i}`} {...product}/>)
  }

  render() {
    return (
      <div className="right-panel">
        <div className="product-list-title">Top <span>{this.state.products.length}</span> product(s)</div>
        <div className="product-list">
          {this.renderProductList()}
        </div>
      </div>
    )
  }

}
