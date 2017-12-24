import React, {Component} from 'react';
import {BrowserRouter, Link, Route} from 'react-router-dom';
import ReactDOM from 'react-dom';
import Home from './components/Home';
import ProductDetail from './components/ProductDetail';
// import style
require('./style.scss');

const About = () => (
  <div className="about-screen">
    About
  </div>
);

const App = () => (
  <div className="container">
    <nav>
      <Link to="/">Home</Link>
      <Link to="/about">About</Link>
    </nav>
    <Route exact path="/" component={Home}/>
    <Route path="/about" component={About}/>
    <Route path="/product/detail/:id" component={ProductDetail}/>
  </div>
);

class Container extends Component {

  render() {
    return (
      <BrowserRouter>
        <App/>
      </BrowserRouter>
    )
  }

}

ReactDOM.render(<Container/>, document.querySelector('#app'));

if (module.hot) {
  module.hot.accept();
}
