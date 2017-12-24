import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {gql} from '../utils';
import {Link} from 'react-router-dom';

export default class ProductDetail extends Component {

  constructor(props) {
    super(props);
    this.state = {
      name: '',
    }
  }

  componentDidMount() {
    // load product detail
    gql(`
      query getProduct($id: Int) {
        product(id: $id) {
          id
          name
          price
        }
      }
    `,
      {id: +this.props.match.params.id}
    )
      .then(resp => {
        this.setState({
          name: resp.product.name,
        })
      });
  }

  render() {
    return (
      <div>
        <span>{this.state.name}</span>
        <br/>
        <Link to="/">Back to home</Link>
      </div>
    )
  }
}

ProductDetail.propTypes = {
  id: PropTypes.string,
};
