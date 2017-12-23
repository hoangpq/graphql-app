import React, {Component} from 'react';
import Product from './Product';

export default class ProductList extends Component {

  constructor(props) {
    super(props);
    this.state = {
      products: [],
    }
  }

  componentDidMount() {
    fetch('/api/graphql', {
      method: 'POST',
      /*headers: new Headers({
        'Content-Type': 'application/json',
      }),*/
      body: JSON.stringify({
        query: 'query getProducts { products { id name price uom { id name } } }'
      })
    })
      .then(resp => resp.json())
      .then(resp => this.setState({products: resp.data.products}))
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
