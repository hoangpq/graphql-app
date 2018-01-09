import React, { Component } from 'react';
import ProductList from './ProductList';
import SideBar from './SideBar';

export default class Home extends Component {

  render() {
    return (
      <div className="home-screen">
        <SideBar/>
        <ProductList/>
      </div>
    )
  }

}